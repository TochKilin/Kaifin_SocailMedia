package bookmark_mobile

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	types "kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type BookMarkMobileHandlerImpl struct {
	bs *BookMarkMobileServiceImpl
}

func NewBookMarkHandlerImpl(db *sqlx.DB, rdb *redis.Client) *BookMarkMobileHandlerImpl {
	return &BookMarkMobileHandlerImpl{
		bs: NewBookMarkMobileServiceImpl(db, rdb),
	}
}

func (h *BookMarkMobileHandlerImpl) Toggle(c fiber.Ctx) error {
	var toggleReq ToggleBookmarkRequest
	v := utls.NewValidator()
	if err := toggleReq.bind(c, v); err != nil {
		msg, err_msg := translate.TranslateWithError(c, "invalid_request")
		if err_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(err_msg.ErrorString(), constants.Translate_Failed, err_msg.Err),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, err),
		)
	}

	userCtx, ok := c.Locals("UserContext").(types.UserContext)
	if !ok {
		msg, err_msg := translate.TranslateWithError(c, "unauthorized")
		if err_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(err_msg.ErrorString(), constants.Translate_Failed, err_msg.Err),
			)
		}
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError(msg, constants.Invalid_request, fmt.Errorf("user context missing or invalid")),
		)
	}

	userID := userCtx.UserID
	status, e := h.bs.Toggle(userID, toggleReq)
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

	msg, e_msg := translate.TranslateWithError(c, "bookmark_toggled")
	if e_msg != nil {
		msg = "Bookmark updated successfully"
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, status),
	)
}

func (h *BookMarkMobileHandlerImpl) Show(c fiber.Ctx) error {
	var showReq ShowBookmarkRequest
	v := utls.NewValidator()
	if err := showReq.bind(c, v); err != nil {
		msg, err_msg := translate.TranslateWithError(c, "invalid_request")
		if err_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(err_msg.ErrorString(), constants.Translate_Failed, err_msg.Err),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, err),
		)
	}

	userCtx, ok := c.Locals("UserContext").(types.UserContext)
	if !ok {
		msg, err_msg := translate.TranslateWithError(c, "unauthorized")
		if err_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(err_msg.ErrorString(), constants.Translate_Failed, err_msg.Err),
			)
		}
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError(msg, constants.Invalid_request, fmt.Errorf("user context missing")),
		)
	}

	userID := userCtx.UserID
	bookmarks, e := h.bs.Show(userID, showReq)
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

	msg, e_msg := translate.TranslateWithError(c, "posts_retrieved")
	if e_msg != nil {
		msg = "Bookmarks retrieved successfully"
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, bookmarks),
	)
}
