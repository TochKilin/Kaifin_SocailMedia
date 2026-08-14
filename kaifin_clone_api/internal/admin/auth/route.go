package auth

import (
	//community package

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type AuthRoute struct {
	Authhandler *AuthHandlerImpl
}

func NewAuthRoute(app *fiber.App, db *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *AuthRoute {
	h := NewAuthHandlerImpl(db, rdb, ws)
	f := app.Group("/api/v1/admin/auth")
	f.Post("/login", h.Login)
	return &AuthRoute{
		Authhandler: h,
	}
}
