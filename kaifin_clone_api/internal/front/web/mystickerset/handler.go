package mystickerset

import (
	"errors"
	"io"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
)

type MysetStickerHandlerImpl struct {
	ss MysetStickerServiceImpl
}

func NewMysetStickerHanslerImpl(dbpool *sqlx.DB, rdb *redis.Client) *MysetStickerHandlerImpl {
	return &MysetStickerHandlerImpl{
		ss: *NewMysetStickerServiceImpl(dbpool, rdb),
	}
}

func (h *MysetStickerHandlerImpl) Show(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, errors.New("user context not found")),
		)
	}

	result, e := h.ss.ShowMySets(uCtx.UserID)
	if e != nil {
		msg, _ := translate.TranslateWithError(c, e.MessageID)
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	msg, _ := translate.TranslateWithError(c, "my_sticker_sets_retrieved")
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, result),
	)
}

func (h *MysetStickerHandlerImpl) Delete(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, errors.New("user context not found")),
		)
	}

	packID, err := strconv.ParseInt(c.Params("pack_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid pack id", constants.Generic_invalid, err),
		)
	}

	if e := h.ss.RemoveMySet(packID, uCtx.UserID); e != nil {
		msg, _ := translate.TranslateWithError(c, e.MessageID)
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	msg, _ := translate.TranslateWithError(c, "sticker_set_removed")
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}

func (h *MysetStickerHandlerImpl) Create(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, errors.New("user context not found")),
		)
	}

	packID, err := strconv.ParseInt(c.Params("pack_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid pack id", constants.Generic_invalid, err),
		)
	}

	if e := h.ss.AddPackToMySets(packID, uCtx.UserID); e != nil {
		msg, _ := translate.TranslateWithError(c, e.MessageID)
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	msg, _ := translate.TranslateWithError(c, "pack_added")
	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}

func (h *MysetStickerHandlerImpl) ServeImage(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid id")
	}

	fileData, fileType, err := h.ss.GetStickerImage(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("not found")
	}

	c.Set("Content-Type", fileType)
	return c.Send(fileData)
}

func (h *MysetStickerHandlerImpl) ListPacks(c fiber.Ctx) error {
	result, e := h.ss.ListPacks()
	if e != nil {
		msg, _ := translate.TranslateWithError(c, e.MessageID)
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err))
	}
	msg, _ := translate.TranslateWithError(c, "sticker_packs_retrieved")
	return c.Status(fiber.StatusOK).JSON(response.NewResponse(msg, constants.Generic_success, result))
}

func (h *MysetStickerHandlerImpl) ShowPackStickers(c fiber.Ctx) error {
	packID, err := strconv.ParseInt(c.Query("pack_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid pack id", constants.Generic_invalid, err))
	}
	result, e := h.ss.ListStickersByPack(packID)
	if e != nil {
		msg, _ := translate.TranslateWithError(c, e.MessageID)
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err))
	}
	msg, _ := translate.TranslateWithError(c, "stickers_retrieved")
	return c.Status(fiber.StatusOK).JSON(response.NewResponse(msg, constants.Generic_success, result))
}

func (h *MysetStickerHandlerImpl) CreateSticker(c fiber.Ctx) error {
	packID, err := strconv.ParseInt(c.FormValue("pack_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid pack id", constants.Generic_invalid, err))
	}
	fh, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("File required", constants.Generic_invalid, err))
	}
	f, err := fh.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Cannot read file", constants.Generic_error, err))
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Cannot read file", constants.Generic_error, err))
	}

	if e := h.ss.CreateSticker(packID, data, fh.Header.Get("Content-Type")); e != nil {
		msg, _ := translate.TranslateWithError(c, e.MessageID)
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err))
	}
	msg, _ := translate.TranslateWithError(c, "sticker_created")
	return c.Status(fiber.StatusCreated).JSON(response.NewResponse(msg, constants.Generic_success, nil))
}

func (h *MysetStickerHandlerImpl) CreatePack(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, errors.New("user context not found")))
	}

	name := c.FormValue("name")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Pack name required", constants.Generic_invalid, errors.New("empty name")))
	}

	packID, e := h.ss.CreatePack(name, uCtx.UserID)
	if e != nil {
		msg, _ := translate.TranslateWithError(c, e.MessageID)
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err))
	}

	msg, _ := translate.TranslateWithError(c, "pack_created")
	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, fiber.Map{"pack_id": packID}))
}
