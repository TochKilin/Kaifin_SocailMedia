package quote

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type QuoteRoute struct {
	handler *QuoteHandler
	ws      *websocket.WebSocketManager
}

func NewQuoteRoute(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *QuoteRoute {
	h := NewQuoteHandler(dbpool, rdb, ws)
	v1 := app.Group("/api/v1/front")
	quotes := v1.Group("/quotes")

	quotes.Get("/show", h.Show)
	quotes.Get("/show/:id", h.ShowOne)
	quotes.Post("/create", h.Create)
	quotes.Put("/update/:id", h.Update)
	quotes.Delete("/delete/:id", h.Delete)

	return &QuoteRoute{
		handler: h,
		ws:      ws,
	}
}
