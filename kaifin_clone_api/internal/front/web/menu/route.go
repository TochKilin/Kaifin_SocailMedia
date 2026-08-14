package menu

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type MenuRouteImpl struct {
	Handler *MenuHandlerImpl
	ws      *websocket.WebSocketManager
}

func NewMenuRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *MenuRouteImpl {
	handler := NewMenuHandlerImpl(dbpool, rdb)
	route := &MenuRouteImpl{
		Handler: handler,
		ws:      ws,
	}
	menus := app.Group("/api/v1/front/menus")
	menus.Get("/show", handler.Show)

	return route
}
