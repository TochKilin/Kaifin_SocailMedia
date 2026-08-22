package notificationbell

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	wsManager "kaifin_clone_api/internal/admin/websocket"
)

type NotiRouteImpl struct {
	lh *NotificationsHandlerImpl
}

func NewNotiRoute(app *fiber.App, db *sqlx.DB, ws *wsManager.WebSocketManager) *NotiRouteImpl {
	lh := NewNotificationsServiceImpl(db)
	h := NewNotificationsHandlerImpl(lh, ws)

	grp := app.Group("/api/v1/front/notifications")
	grp.Get("/show", h.Show)
	grp.Post("/read", h.MarkAsRead)

	// grp.Get("/ws", websocket.New(h.HandleWS))

	return &NotiRouteImpl{
		lh: h,
	}
}
