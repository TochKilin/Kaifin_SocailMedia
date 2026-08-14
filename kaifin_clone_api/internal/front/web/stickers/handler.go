package stickers

import (
	"io"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type StickersHandlerImpl struct {
	ss *StickersServiceImpl
}

func NewStickersHandlerImpl(dbpool *sqlx.DB) *StickersHandlerImpl {
	return &StickersHandlerImpl{
		ss: NewStickerServiceImpl(dbpool),
	}
}

func (h *StickersHandlerImpl) Show(c fiber.Ctx) error {
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		h.ss.SetUserCtx(&uCtx)
	}

	var req ShowStickersRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request params", constants.Invalid_request, err),
		)
	}

	data, e := h.ss.Show(req)
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
		response.NewResponse("stickers_retrieved", constants.Generic_success, data),
	)
}

func (h *StickersHandlerImpl) Create(c fiber.Ctx) error {
	_, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	req := &CreateStickerRequest{}
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request body", constants.Generic_invalid, err),
		)
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("File is required", constants.Generic_invalid, err),
		)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Failed to open file", constants.Generic_error, err),
		)
	}
	defer file.Close()
	fileData, _ := io.ReadAll(file)
	sticker := &Sticker{
		FileName: fileHeader.Filename,
		FileType: fileHeader.Header.Get("Content-Type"),
		FileData: fileData,
	}

	if e := h.ss.Create(req, sticker); e != nil {
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
		response.NewResponse("sticker_created", constants.Generic_success, true),
	)
}

func (h *StickersHandlerImpl) Update(c fiber.Ctx) error {
	_, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	req := &UpdateStickerRequest{}
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request body", constants.Generic_invalid, err),
		)
	}

	sticker := &Sticker{}
	updateFile := false
	fileHeader, err := c.FormFile("file")
	if err == nil {
		updateFile = true
		file, _ := fileHeader.Open()
		defer file.Close()
		fileData, _ := io.ReadAll(file)

		sticker.FileName = fileHeader.Filename
		sticker.FileType = fileHeader.Header.Get("Content-Type")
		sticker.FileData = fileData
	}

	if e := h.ss.Update(req, sticker, updateFile); e != nil {
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
		response.NewResponse("sticker_updated", constants.Generic_success, true),
	)
}
