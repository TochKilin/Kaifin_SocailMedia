package comments_mobile

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type CommentsMobileHandlerImpl struct {
	cs CommentsMobileServiceImpl
}

func NewCommentsMobileHanslerImpl(dbpool *sqlx.DB) *CommentsMobileHandlerImpl {
	return &CommentsMobileHandlerImpl{
		cs: *NewCommentsServiceImpl(dbpool),
	}
}

func (h *CommentsMobileHandlerImpl) Create(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	req := &CreateCommentMobileRequest{}
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request body", constants.Generic_invalid, err),
		)
	}

	if e := h.cs.Create(req, &uCtx); e != nil {
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

func (h *CommentsMobileHandlerImpl) Show(c fiber.Ctx) error {
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		h.cs.SetUserCtx(&uCtx)
	}

	var req ShowCommentsMobileRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	data, e := h.cs.Show(req)
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
		response.NewResponse("comments_retrieved", constants.Generic_success, data),
	)
}

func (h *CommentsMobileHandlerImpl) Delete(c fiber.Ctx) error {
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
