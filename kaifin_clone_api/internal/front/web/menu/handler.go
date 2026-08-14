package menu

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type MenuHandlerImpl struct {
	ms MenuService
}

func NewMenuHandlerImpl(dbpool *sqlx.DB, rdb *redis.Client) *MenuHandlerImpl {
	return &MenuHandlerImpl{
		ms: NewMenuServiceImpl(dbpool, rdb),
	}
}

func (h *MenuHandlerImpl) Show(c fiber.Ctx) error {
	var showMenuRequest ShowMenuRequest
	v := utls.NewValidator()
	if err := showMenuRequest.bind(c, v); err != nil {
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

	menus, e := h.ms.Show(showMenuRequest)
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

	msg, e_msg := translate.TranslateWithError(c, "menus_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, menus),
	)
}
