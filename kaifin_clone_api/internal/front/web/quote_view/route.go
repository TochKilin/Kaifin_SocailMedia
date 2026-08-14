package quoteview

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	"kaifin_clone_api/internal/admin/websocket"
)

type QuoteViewRoutel struct {
	handler *QuoteViewHandler
	ws      *websocket.WebSocketManager
}

func NewQuoteViewRoute(app *fiber.App, dbpool *sqlx.DB) *QuoteViewRoutel {
	h := NewQuoteViewHandler(dbpool)
	v1 := app.Group("/api/v1")
	v1.Post("/quote-views", h.Track)
	return &QuoteViewRoutel{
		handler: h,
	}
}
