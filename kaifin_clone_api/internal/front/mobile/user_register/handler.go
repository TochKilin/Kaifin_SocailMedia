package userregister

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-playground/validator"
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

type UserRegisterHandler struct {
	Service UserRegisterServiceImpl
	ws      *websocket.WebSocketManager
}

func NewUserRegisterHandler(dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *UserRegisterHandler {
	return &UserRegisterHandler{
		Service: *NewUserRegisterServiceImpl(dbpool, rdb),
		ws:      ws,
	}
}

// Create user
func (h *UserRegisterHandler) Create(c fiber.Ctx) error {
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		b, _ := json.MarshalIndent(uCtx, "", "  ")
		h.Service.SetUserCtx(&uCtx) //
		fmt.Println("jwt_data:", string(b))
	}
	req := &AuthRegisterRequest{}
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request body", constants.Generic_invalid, err),
		)
	}
	if err := v.Validate(req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			fe := ve[0]
			msg, _ := translate.TranslateWithError(c, "validation_"+fe.Tag(),
				map[string]any{
					"Field": fe.Field(),
					"Param": fe.Param(),
				})
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(msg, constants.Generic_invalid, err),
			)
		}
		return err
	}

	// var createdBy int64 = 1
	e := h.Service.Create(c, req)
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
	msg, e_msg := translate.TranslateWithError(c, "user_created")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, true),
	)
}

// Show users
func (h *UserRegisterHandler) Show(c fiber.Ctx) error {
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		b, _ := json.MarshalIndent(uCtx, "", "  ")
		h.Service.SetUserCtx(&uCtx) //
		fmt.Println("jwt_data:", string(b))
	}

	var usersShowRequest UserShowRequest
	v := utls.NewValidator()
	if err := usersShowRequest.bind(c, v); err != nil {
		msg, err_msg := translate.TranslateWithError(c, "invalid_request")
		if err_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(
					err_msg.ErrorString(),
					constants.Translate_Failed,
					err_msg.Err,
				),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(
				msg,
				constants.Invalid_request,
				err,
			),
		)
	}
	users, e := h.Service.Show(usersShowRequest)
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
	msg, e_msg := translate.TranslateWithError(c, "users_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponseWithPaing(msg, constants.Generic_success, users, usersShowRequest.PageOption.Page, usersShowRequest.PageOption.Perpage, users.Total),
	)
}

// Update users
func (h *UserRegisterHandler) Update(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_user_id")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}
	req := &UpdateUserRequest{}
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request body", constants.Generic_invalid, err),
		)
	}
	var updatedBy int64 = 1

	user, e := h.Service.Update(c, id, req, updatedBy)
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
	msg, e_msg := translate.TranslateWithError(c, "user_updated")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, user),
	)
}

// Delete user
func (h *UserRegisterHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_user_id")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}
	var deletedBy int64 = 1

	if e := h.Service.Delete(id, deletedBy); e != nil {
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
	msg, e_msg := translate.TranslateWithError(c, "user_deleted")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}

func (h *UserRegisterHandler) Profile(c fiber.Ctx) error {

	fmt.Println("UserContext:", c.Locals("UserContext"))
	fmt.Println("user_id:", c.Locals("user_id"))

	uCtx, ok := c.Locals("UserContext").(share.UserContext)

	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"message": "invalid user context",
			"success": false,
		})
	}

	res, err := h.Service.Profile(uCtx.UserID)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    res,
	})
}
