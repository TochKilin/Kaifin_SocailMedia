package bookmark

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type BookMarkRouteImpl struct {
	hb *BookMarkHandlerImpl
	ws *websocket.WebSocketManager
}

func NewBookMarkRouteImpl(app *fiber.App, db *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *BookMarkRouteImpl {
	h := NewBookMarkHandlerImpl(db, rdb)
	bookmark := app.Group("/api/v1/front/bookmarks")
	bookmark.Post("/create", h.Toggle)
	bookmark.Get("/show", h.Show)
	return &BookMarkRouteImpl{
		hb: h,
	}
}
