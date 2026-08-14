package communities

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type CommunitiesHandlerImpl struct {
	cs CommunitiesServiceImpl
}

func NewCommunitiesHandlerImpl(app *fiber.App, dbpool *sqlx.DB) *CommunitiesHandlerImpl {
	return &CommunitiesHandlerImpl{
		cs: *NewCommunitiesServiceImpl(dbpool),
	}
}

func (h *CommunitiesHandlerImpl) Show(c fiber.Ctx) error {
	var req ShowCommunitiesRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, err),
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

	msg, e_msg := translate.TranslateWithError(c, "communities_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, data),
	)
}
