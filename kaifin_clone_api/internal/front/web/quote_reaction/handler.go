package quotereaction

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type QuoteReactionHandler struct {
	Service QuoteReactionService
	ws      *websocket.WebSocketManager
}

func NewQuoteReactionHandler(dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *QuoteReactionHandler {
	return &QuoteReactionHandler{
		Service: NewQuoteReactionServiceImpl(dbpool, rdb),
		ws:      ws,
	}
}

func (h *QuoteReactionHandler) React(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		msg, _ := translate.TranslateWithError(c, "unauthorized")
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, fmt.Errorf("missing user context")),
		)
	}
	h.Service.SetUserCtx(&uCtx)

	req := &ReactRequest{}
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

	reaction, e := h.Service.React(req, &uCtx)
	if e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "reaction_added")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, reaction),
	)
}

func (h *QuoteReactionHandler) Unreact(c fiber.Ctx) error {
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

	req := &UnreactRequest{QuoteID: quoteID}

	if e := h.Service.Unreact(req, &uCtx); e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "reaction_removed")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}

func (h *QuoteReactionHandler) ListTypes(c fiber.Ctx) error {
	types, e := h.Service.ListTypes()
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

	msg, e_msg := translate.TranslateWithError(c, "reaction_types_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, types),
	)
}

func (h *QuoteReactionHandler) Show(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, errors.New("user context not found")),
		)
	}

	quoteID, err := strconv.ParseInt(c.Params("quote_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid quote id", constants.Generic_invalid, err),
		)
	}

	reaction, e := h.Service.GetByUser(quoteID, &uCtx) // ➕ service method takes *uctx now, matching React/Unreact pattern
	if e != nil {
		return c.Status(fiber.StatusOK).JSON(
			response.NewResponse("ok", constants.Generic_success, fiber.Map{
				"reaction_type_id": nil,
			}),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("ok", constants.Generic_success, fiber.Map{
			"reaction_type_id": reaction.ReactionTypeID,
		}),
	)
}
