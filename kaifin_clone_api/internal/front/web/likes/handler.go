package likes

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type LikesHandlerImpl struct {
	ls *LikesServiceImpl
}

func NewLikesHandlerImpl(dbpool *sqlx.DB, rdb *redis.Client) *LikesHandlerImpl {
	return &LikesHandlerImpl{
		ls: NewLikesServiceImpl(dbpool, rdb),
	}
}

func (h *LikesHandlerImpl) Create(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, errors.New("user context not found")),
		)
	}

	req := &CreateLikeRequest{}
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request body", constants.Generic_invalid, err),
		)
	}

	liked, e := h.ls.Create(req, &uCtx)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			errStr := "translation error"
			var rawErr error = e_msg.Err
			if rawErr != nil {
				errStr = rawErr.Error()
			} else {
				rawErr = errors.New(errStr)
			}
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(errStr, constants.Translate_Failed, rawErr),
			)
		}
		serviceErr := e.Err
		if serviceErr == nil {
			serviceErr = errors.New("unknown service error")
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_error, serviceErr),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("success", constants.Generic_success, fiber.Map{"liked": liked}),
	)
}

func (h *LikesHandlerImpl) Show(c fiber.Ctx) error {
	var uCtx *share.UserContext
	if ctx, ok := c.Locals("UserContext").(share.UserContext); ok {
		uCtx = &ctx
	}

	var req ShowLikesRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	data, e := h.ls.Show(req, uCtx)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			errStr := "translation error"
			var rawErr error = e_msg.Err
			if rawErr != nil {
				errStr = rawErr.Error()
			} else {
				rawErr = errors.New(errStr)
			}
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(errStr, constants.Translate_Failed, rawErr),
			)
		}
		serviceErr := e.Err
		if serviceErr == nil {
			serviceErr = errors.New("unknown service error")
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, serviceErr),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("success", constants.Generic_success, data),
	)
}
