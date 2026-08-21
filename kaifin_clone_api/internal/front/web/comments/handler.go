package comments

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type CommentsHandlerImpl struct {
	cs CommentsServiceImpl
}

func NewCommentsHanslerImpl(dbpool *sqlx.DB) *CommentsHandlerImpl {
	return &CommentsHandlerImpl{
		cs: *NewCommentsServiceImpl(dbpool),
	}
}

func (h *CommentsHandlerImpl) Create(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	postID, err := strconv.ParseInt(c.FormValue("post_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid post_id", constants.Generic_invalid, err),
		)
	}

	var parentID *int64
	if pc := c.FormValue("parent_comment_id"); pc != "" {
		if v, err := strconv.ParseInt(pc, 10, 64); err == nil {
			parentID = &v
		}
	}

	content := strings.TrimSpace(c.FormValue("content"))

	req := &CreateCommentRequest{PostID: postID, ParentID: parentID, Content: content}

	var savedPaths []string
	if form, ferr := c.MultipartForm(); ferr == nil && form != nil {
		files := form.File["images"]
		if content == "" && len(files) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError("Content or image required", constants.Generic_invalid, nil),
			)
		}
		for _, fh := range files {
			path, serr := saveCommentImage(c, fh)
			if serr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					response.NewResponseError("Failed to save image", constants.Generic_error, serr),
				)
			}
			savedPaths = append(savedPaths, path)
		}
	} else if content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Content is required", constants.Generic_invalid, nil),
		)
	}

	if e := h.cs.Create(req, &uCtx, savedPaths); e != nil {
		log.Printf("Create comment error: %v", e.Err)
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse("comment_created", constants.Generic_success, true),
	)
}

func saveCommentImage(c fiber.Ctx, fh *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(fh.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), randomHex(8), ext)
	dst := filepath.Join("uploads", "comments", filename)
	if err := c.SaveFile(fh, dst); err != nil {
		log.Printf("SaveFile error: %v", err)
		return "", err
	}
	return filepath.Join("comments", filename), nil
}
func randomHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {

		return strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return hex.EncodeToString(b)
}

func (h *CommentsHandlerImpl) Show(c fiber.Ctx) error {
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		h.cs.SetUserCtx(&uCtx)
	}

	var req ShowCommentsRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	data, e := h.cs.Show(req)
	if e != nil {
		log.Printf("Show comments error: %v", e.Err)
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
		response.NewResponse("comments_retrieved", constants.Generic_success, data),
	)
}

func (h *CommentsHandlerImpl) Delete(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid comment id", constants.Generic_invalid, err),
		)
	}

	if e := h.cs.Delete(id, &uCtx); e != nil {
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
		response.NewResponse("comment_deleted", constants.Generic_success, nil),
	)
}

func (h *CommentsHandlerImpl) ToggleLike(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid comment id", constants.Generic_invalid, err),
		)
	}

	liked, e := h.cs.ToggleLike(id, &uCtx)
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
		response.NewResponse("like_toggled", constants.Generic_success, fiber.Map{"liked": liked}),
	)
}
