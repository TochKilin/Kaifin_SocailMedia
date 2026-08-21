package notification

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	"kaifin_clone_api/internal/admin/websocket"
)

type NotificationsRouteImpl struct {
	Handler *NotificationsHandlerImpl
}

func NewNotificationsRouteImpl(app *fiber.App, db *sqlx.DB, ws *websocket.WebSocketManager) *NotificationsRouteImpl {
	h := NewNotificationsHandlerImpl(db, ws)

	group := app.Group("/api/v1/front/notifications")
	group.Get("/show", h.GetList)
	group.Post("/read", h.Read)

	return &NotificationsRouteImpl{
		Handler: h,
	}
}
