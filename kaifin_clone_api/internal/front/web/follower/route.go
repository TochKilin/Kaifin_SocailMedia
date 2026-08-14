package follower

import (
	wsManager "kaifin_clone_api/internal/admin/websocket"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type FollowersRouteImpl struct {
	FollowersHandler *FollowersHandlerImpl
}

func NewFollowersRouteImpl(app *fiber.App, db *sqlx.DB, ws *wsManager.WebSocketManager) *FollowersRouteImpl {
	h := NewFollowersHandlerImpl(db, ws)
	follower := app.Group("/api/v1/front/followers")
	follower.Post("/create", h.Toggle)
	follower.Get("/show", h.Show)
	return &FollowersRouteImpl{
		FollowersHandler: h,
	}
}
