package chat

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	"kaifin_clone_api/internal/admin/websocket"
)

type ChatsRouteImpl struct {
	ch *ChatsHandlerImpl
	ws *websocket.WebSocketManager
}

func NewChatsRouteImpl(app *fiber.App, dbpool *sqlx.DB, ws *websocket.WebSocketManager) *ChatsRouteImpl {
	h := NewChatsHandlerImpl(dbpool, ws)
	chats := app.Group("/api/v1/front/chats")
	chats.Get("/show", h.ShowConversations)
	chats.Get("/:conversation_id/messages", h.ShowMessages)
	chats.Post("/send", h.SendMessage)
	chats.Post("/:message_id/react", h.ToggleReaction)
	chats.Post("/:conversation_id/read", h.MarkAsRead)
	chats.Post("/start", h.StartConversation)
	chats.Get("/users/search", h.SearchUsers)
	return &ChatsRouteImpl{ch: h}
}
