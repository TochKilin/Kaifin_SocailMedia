package user_login

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type AuthUserRouteImpl struct {
	AuthUserHandler *AuthUserHandlerImpl
}

func NewAuthUserRouteImpl(app *fiber.App, db *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *AuthUserRouteImpl {
	uh := NewAuthUserHandlerImpl(db, rdb, ws)
	f := app.Group("/api/v1/front/auth")
	f.Post("/login-user", uh.UserLogin)
	f.Get("/profile", uh.Profile)

	return &AuthUserRouteImpl{
		AuthUserHandler: uh,
	}
}
