package mobile_login

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type AuthUserMobileRouteImpl struct {
	AuthUserHandler *AuthUserMobileHandlerImpl
}

func NewAuthUserMobileRouteImpl(app *fiber.App, db *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *AuthUserMobileRouteImpl {
	uh := NewAuthUserHandlerImpl(db, rdb, ws)
	mo := app.Group("/api/v1/mobile/auth")
	mo.Post("/login-mobile", uh.UserLogin)

	return &AuthUserMobileRouteImpl{
		AuthUserHandler: uh,
	}
}
