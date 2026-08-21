package quotereaction

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type QuoteReactionRoute struct {
	handler *QuoteReactionHandler
	ws      *websocket.WebSocketManager
}

func NewQuoteReactionRoute(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *QuoteReactionRoute {
	h := NewQuoteReactionHandler(dbpool, rdb, ws)
	v1 := app.Group("/api/v1/front")

	v1.Get("/reaction-types", h.ListTypes)

	reactions := v1.Group("/quote-reactions")
	reactions.Post("/create", h.React)
	reactions.Delete("/:quote_id", h.Unreact)

	reactions.Get("/:quote_id", h.Show)

	return &QuoteReactionRoute{
		handler: h,
		ws:      ws,
	}
}
