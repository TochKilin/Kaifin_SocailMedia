package userregister

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type RegisterRouteImpl struct {
	AuthUserHandler *UserRegisterHandler
}

// Route register
func NewRegisterRouteImpl(app *fiber.App, db *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *RegisterRouteImpl {
	re := NewUserRegisterHandler(db, rdb, ws)
	clients := app.Group("/api/v1/front/register")
	clients.Post("/create", re.Create)
	clients.Get("/", re.Show)
	clients.Put("/update/:id", re.Update)
	clients.Delete("/delete/:id", re.Delete)

	return &RegisterRouteImpl{
		AuthUserHandler: re,
	}
}
