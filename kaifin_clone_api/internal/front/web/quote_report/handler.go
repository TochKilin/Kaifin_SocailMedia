package quotereport

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

type QuoteReportHandler struct {
	Service QuoteReportService
}

func NewQuoteReportHandler(dbpool *sqlx.DB) *QuoteReportHandler {
	return &QuoteReportHandler{Service: NewQuoteReportServiceImpl(dbpool)}
}

func (h *QuoteReportHandler) Create(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		msg, _ := translate.TranslateWithError(c, "unauthorized")
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, fmt.Errorf("missing user context")),
		)
	}
	h.Service.SetUserCtx(&uCtx)

	req := &CreateReportRequest{}
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

	if e := h.Service.Create(req, &uCtx); e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "quote_reported")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}

func (h *QuoteReportHandler) Show(c fiber.Ctx) error {
	status := c.Query("status")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}

	reports, total, e := h.Service.Show(status, page, perPage)
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

	msg, e_msg := translate.TranslateWithError(c, "reports_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponseWithPaing(msg, constants.Generic_success, reports, page, perPage, total),
	)
}

func (h *QuoteReportHandler) UpdateStatus(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_report_id")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}

	req := &UpdateReportStatusRequest{}
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

	if e := h.Service.UpdateStatus(id, req); e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "report_status_updated")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}
