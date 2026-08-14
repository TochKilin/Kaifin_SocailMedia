package auth

import (
	//community package

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
	constants "kaifin_clone_api/pkg/constants"
	contants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/logs"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

// import "kaifin_clone_api/internal/admin/auth"

type AuthHandlerImpl struct {
	Se *AuthServiceImpl
	db *sqlx.DB
	ws *websocket.WebSocketManager
}

func NewAuthHandlerImpl(db *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *AuthHandlerImpl {
	s := NewAuthServiceImpl(db, rdb)
	return &AuthHandlerImpl{
		Se: s,
		ws: ws,
	}
}

func (h *AuthHandlerImpl) Login(c fiber.Ctx) error {
	req := AuthLoginRequest{}
	v := utls.NewValidator()
	log := logs.NewCustomLog("Login", "request received", "info")
	_ = log
	if err := req.bind(c, v); err != nil {
		return err
	}
	// internal library use for perform bussinis logic or deal with database
	rs, errRes := h.Se.Login(req)

	if errRes != nil {
		// log when error
		log := logs.NewCustomLog("Login", "login failed", "error")
		_ = log
		smg, err_smg := translate.TranslateWithError(c, errRes.MessageID)
		if err_smg != nil { // translate not found
			return c.Status(fiber.StatusBadRequest).JSON(response.NewResponseError(
				err_smg.ErrorString(),
				contants.Translate_Failed,
				err_smg.Err,
			))
		}
		// error from repos
		return c.Status(fiber.StatusBadRequest).JSON(response.NewResponseError(
			smg,
			constants.Login_failed,
			errRes,
		))

	}

	// when success
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
