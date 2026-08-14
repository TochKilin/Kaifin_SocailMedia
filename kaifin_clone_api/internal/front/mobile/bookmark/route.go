package bookmark_mobile

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type BookMarkMobileRouteImpl struct {
	hb *BookMarkMobileHandlerImpl
	ws *websocket.WebSocketManager
}

func NewBookMarkMobileRouteImpl(app *fiber.App, db *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *BookMarkMobileRouteImpl {
	h := NewBookMarkHandlerImpl(db, rdb)
	bookmark := app.Group("/api/v1/mobile/bookmarks")
	bookmark.Post("/create", h.Toggle)
	bookmark.Get("/show", h.Show)
	return &BookMarkMobileRouteImpl{
		hb: h,
	}
}
