package quotehidden

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type QuoteHiddenHandler struct {
	Service QuoteHiddenService
}

func NewQuoteHiddenHandler(dbpool *sqlx.DB) *QuoteHiddenHandler {
	return &QuoteHiddenHandler{Service: NewQuoteHiddenServiceImpl(dbpool)}
}

func (h *QuoteHiddenHandler) Hide(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		msg, _ := translate.TranslateWithError(c, "unauthorized")
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, fmt.Errorf("missing user context")),
		)
	}
	h.Service.SetUserCtx(&uCtx)

	req := &HideRequest{}
	v := utls.NewValidator()
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request body", constants.Generic_invalid, err),
		)
	}
	if err := v.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Validation failed", constants.Generic_invalid, err),
		)
	}

	if e := h.Service.Hide(req, &uCtx); e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "quote_hidden")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}

func (h *QuoteHiddenHandler) Unhide(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		msg, _ := translate.TranslateWithError(c, "unauthorized")
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, fmt.Errorf("missing user context")),
		)
	}
	h.Service.SetUserCtx(&uCtx)

	quoteID, err := strconv.ParseInt(c.Params("quote_id"), 10, 64)
	if err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_quote_id")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}

	if e := h.Service.Unhide(quoteID, &uCtx); e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "quote_unhidden")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}
