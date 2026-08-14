package notification

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	"kaifin_clone_api/internal/admin/websocket" // ហៅយក WebSocket Manager
)

type NotificationsRouteImpl struct {
	Handler *NotificationsHandlerImpl
}

// បន្ថែម ws *websocket.WebSocketManager ចូលក្នុង Parameter
func NewNotificationsRouteImpl(app *fiber.App, db *sqlx.DB, ws *websocket.WebSocketManager) *NotificationsRouteImpl {
	// បញ្ជូន ws ទៅឱ្យ Handler វិញ
	h := NewNotificationsHandlerImpl(db, ws)

	group := app.Group("/api/v1/front/notifications")
	group.Get("/show", h.GetList)
	group.Post("/read", h.Read)

	// បន្ថែម WebSocket Route សម្រាប់ให้ Client Connect (ຖ້າຕ້ອງການ)
	// group.Get("/ws", websocket.New(...))

	return &NotificationsRouteImpl{
		Handler: h,
	}
}
