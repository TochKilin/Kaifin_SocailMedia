package template

import (
	"errors"
	"io"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/translate"
)

type TemplateHandlerImpl struct {
	ts TemplateService
}

func NewTemplateHandlerImpl(dbpool *sqlx.DB) *TemplateHandlerImpl {
	return &TemplateHandlerImpl{
		ts: NewTemplateServiceImpl(dbpool),
	}
}

func (h *TemplateHandlerImpl) List(c fiber.Ctx) error {
	result, e := h.ts.List()
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

	msg, e_msg := translate.TranslateWithError(c, "templates_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, result),
	)
}

func (h *TemplateHandlerImpl) ServeImage(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid id")
	}

	fileData, fileType, err := h.ts.GetImage(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("not found")
	}

	c.Set("Content-Type", fileType)
	return c.Send(fileData)
}

func (h *TemplateHandlerImpl) Create(c fiber.Ctx) error {
	name := c.FormValue("name")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Name required", constants.Generic_invalid, errors.New("empty name")),
		)
	}

	fh, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("File required", constants.Generic_invalid, err),
		)
	}
	f, err := fh.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Cannot read file", constants.Generic_error, err),
		)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Cannot read file", constants.Generic_error, err),
		)
	}

	if e := h.ts.Create(name, data, fh.Header.Get("Content-Type")); e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "template_created")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}
