package story

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type StoryHandlerImpl struct {
	ss StoryService
	ws *websocket.WebSocketManager
}

func NewStoryHandlerImpl(dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *StoryHandlerImpl {
	return &StoryHandlerImpl{
		ss: NewStoryServiceImpl(dbpool, rdb),
		ws: ws,
	}
}

func (h *StoryHandlerImpl) Create(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, errors.New("user context not found")),
		)
	}

	req := &CreateStoryRequest{}
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request body", constants.Generic_invalid, err),
		)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("invalid_form", constants.Generic_invalid, err),
		)
	}

	files := form.File["media_url"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("media_required", constants.Generic_invalid, fmt.Errorf("media file is required")),
		)
	}

	file := files[0]
	filePath := fmt.Sprintf("uploads/%s", file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(500).JSON(
			response.NewResponseError("upload_failed", constants.Generic_error, err),
		)
	}
	req.MediaURL = "/" + filePath

	if e := h.ss.Create(req, &uCtx); e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	msg, e_msg := translate.TranslateWithError(c, "story_created")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, true),
	)
}

func (h *StoryHandlerImpl) Show(c fiber.Ctx) error {
	var req ShowStoryRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		msg, err_msg := translate.TranslateWithError(c, "invalid_request")
		if err_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(err_msg.ErrorString(), constants.Translate_Failed, err_msg.Err),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, err),
		)
	}

	stories, e := h.ss.Show(req)
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

	msg, e_msg := translate.TranslateWithError(c, "stories_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponseWithPaing(msg, constants.Generic_success, stories, req.PageOption.Page, req.PageOption.Perpage, stories.Total),
	)
}

func (h *StoryHandlerImpl) Delete(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, errors.New("user context not found")),
		)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_story_id")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}

	if e := h.ss.Delete(id, uCtx.UserID); e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "story_deleted")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}
