package chat

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	"kaifin_clone_api/internal/admin/websocket"
	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type ChatsHandlerImpl struct {
	dbpool *sqlx.DB
	ws     *websocket.WebSocketManager
}

func NewChatsHandlerImpl(dbpool *sqlx.DB, ws *websocket.WebSocketManager) *ChatsHandlerImpl {
	return &ChatsHandlerImpl{
		dbpool: dbpool,
	}
}

func (h *ChatsHandlerImpl) newService(uCtx *share.UserContext) *ChatsServiceImpl {
	cs := NewChatsServiceImpl(h.dbpool)
	cs.UserCtx = uCtx
	return cs
}

func (h *ChatsHandlerImpl) ShowConversations(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	var req ShowConversationsRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	cs := h.newService(&uCtx)
	data, e := cs.ShowConversations(req)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("conversations_retrieved", constants.Generic_success, data),
	)
}

func (h *ChatsHandlerImpl) ShowMessages(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	var req ShowMessagesRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	cs := h.newService(&uCtx)
	data, e := cs.ShowMessages(req)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		status := fiber.StatusInternalServerError
		if e.MessageID == "forbidden" {
			status = fiber.StatusForbidden
		}
		return c.Status(status).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("messages_retrieved", constants.Generic_success, data),
	)
}

func (h *ChatsHandlerImpl) SendMessage(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	convID, err := strconv.ParseInt(c.FormValue("conversation_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid conversation_id", constants.Generic_invalid, err),
		)
	}

	req := &SendMessageRequest{
		ConversationID: convID,
		Content:        c.FormValue("content"),
		Type:           c.FormValue("type"),
	}

	var attachments []MessageAttachment
	if form, ferr := c.MultipartForm(); ferr == nil && form != nil {
		files := form.File["attachments"]
		if req.Content == "" && len(files) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError("Content or attachment required", constants.Generic_invalid, nil),
			)
		}
		for _, fh := range files {
			path, atype, serr := saveMessageAttachment(c, fh)
			if serr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					response.NewResponseError("Failed to save attachment", constants.Generic_error, serr),
				)
			}
			size := fmt.Sprintf("%.1f KB", float64(fh.Size)/1024)
			fname := fh.Filename
			attachments = append(attachments, MessageAttachment{
				Type:     atype,
				URL:      path,
				FileName: &fname,
				FileSize: &size,
			})
		}
		if req.Type == "" && len(attachments) > 0 {
			req.Type = attachments[0].Type
		}
	} else if req.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Content is required", constants.Generic_invalid, nil),
		)
	}

	cs := h.newService(&uCtx)
	msgData, e := cs.SendMessage(req, attachments)
	if e != nil {
		log.Printf("SendMessage error: %v", e.Err)
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		status := fiber.StatusInternalServerError
		if e.MessageID == "forbidden" {
			status = fiber.StatusForbidden
		}
		return c.Status(status).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse("message_sent", constants.Generic_success, msgData),
	)
}

func saveMessageAttachment(c fiber.Ctx, fh *multipart.FileHeader) (string, string, error) {
	ext := filepath.Ext(fh.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), randomHex(8), ext)
	dst := filepath.Join("uploads", "messages", filename)
	if err := c.SaveFile(fh, dst); err != nil {
		log.Printf("SaveFile error: %v", err)
		return "", "", err
	}
	return filepath.Join("messages", filename), attachmentType(ext), nil
}

func attachmentType(ext string) string {
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif", ".webp":
		return "image"
	case ".mp4", ".mov", ".webm":
		return "video"
	case ".mp3", ".wav", ".m4a", ".ogg":
		return "voice"
	default:
		return "file"
	}
}

func randomHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return hex.EncodeToString(b)
}

func (h *ChatsHandlerImpl) ToggleReaction(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	var req ToggleReactionRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	cs := h.newService(&uCtx)
	reacted, e := cs.ToggleReaction(req)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("reaction_toggled", constants.Generic_success, fiber.Map{"reacted": reacted}),
	)
}

func (h *ChatsHandlerImpl) MarkAsRead(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	convID, err := strconv.ParseInt(c.Params("conversation_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid conversation_id", constants.Generic_invalid, err),
		)
	}

	cs := h.newService(&uCtx)
	if e := cs.MarkAsRead(convID); e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("marked_as_read", constants.Generic_success, nil),
	)
}

func (h *ChatsHandlerImpl) StartConversation(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	targetUserID, err := strconv.ParseInt(c.FormValue("target_user_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid target_user_id", constants.Generic_invalid, err),
		)
	}

	cs := h.newService(&uCtx)
	convID, e := cs.StartConversation(targetUserID)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("conversation_ready", constants.Generic_success, fiber.Map{"conversation_id": convID}),
	)
}

// បន្ថែមក្នុង chats_handler.go
func (h *ChatsHandlerImpl) SearchUsers(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	var req SearchUsersRequest
	if err := req.bind(c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	cs := h.newService(&uCtx)
	users, e := cs.SearchUsers(req)
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e.Err.Error(), constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("users_found", constants.Generic_success, fiber.Map{"users": users}),
	)
}
