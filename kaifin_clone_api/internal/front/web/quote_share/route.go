package quoteshare

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	"kaifin_clone_api/internal/admin/websocket"
)

type QuoteShareRoutel struct {
	handler *QuoteShareHandler
	ws      *websocket.WebSocketManager
}

func NewQuoteShareRoute(app *fiber.App, dbpool *sqlx.DB) *QuoteShareRoutel{
	h := NewQuoteShareHandler(dbpool)
	v1 := app.Group("/api/v1")
	v1.Post("/quote-shares", h.Track)
	return &QuoteShareRoutel{
		handler: h,
	}
}
