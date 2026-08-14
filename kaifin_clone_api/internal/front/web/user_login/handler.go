package user_login

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/logs"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type AuthUserHandlerImpl struct {
	Se *AuthUserServiceImpl
	db *sqlx.DB
	ws *websocket.WebSocketManager
}

func NewAuthUserHandlerImpl(db *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *AuthUserHandlerImpl {
	s := NewAuthUserServiceImpl(db, rdb)
	return &AuthUserHandlerImpl{
		Se: s,
		ws: ws,
	}
}

func (h *AuthUserHandlerImpl) UserLogin(c fiber.Ctx) error {
	req := AuthLoginRequest{}
	v := utls.NewValidator()
	log := logs.NewCustomLog("Login", "request received", "info")
	_ = log
	if err := req.bind(c, v); err != nil {
		return err
	}
	rs, errRes := h.Se.UserLogin(req)

	if errRes != nil {
		log := logs.NewCustomLog("Login", "login failed", "error")
		_ = log
		smg, err_smg := translate.TranslateWithError(c, errRes.MessageID)
		if err_smg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.NewResponseError(
				err_smg.ErrorString(),
				constants.Translate_Failed,
				err_smg.Err,
			))
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.NewResponseError(
			smg,
			constants.Login_failed,
			errRes,
		))

	}
	log = logs.NewCustomLog("Login", "login success", "info")
	_ = log
	successMsg, err_msg := translate.TranslateWithError(c, "login_success")
	if err_msg != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.NewResponseError(
			err_msg.ErrorString(),
			constants.Translate_Failed,
			err_msg.Err,
		))
	}

	return c.Status(fiber.StatusOK).JSON(response.NewResponse(
		successMsg,
		constants.Login_success,
		rs,
	))

}

func (h *AuthUserHandlerImpl) Profile(c fiber.Ctx) error {
	userCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponseError(
			"invalid user context",
			constants.Login_failed,
			fmt.Errorf("user context not found"),
		))
	}

	profile, errRes := h.Se.Profile(userCtx.UserID)
	if errRes != nil {
		smg, err_smg := translate.TranslateWithError(c, errRes.MessageID)
		if err_smg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.NewResponseError(
				err_smg.ErrorString(),
				constants.Translate_Failed,
				err_smg.Err,
			))
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.NewResponseError(
			smg,
			constants.Login_failed,
			errRes,
		))
	}

	return c.Status(fiber.StatusOK).JSON(response.NewResponse(
		"profile fetched successfully",
		constants.Login_success,
		profile,
	))
}
