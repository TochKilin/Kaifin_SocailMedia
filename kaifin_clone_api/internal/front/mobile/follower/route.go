package follower_mobile

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type FollowersMobileRouteImpl struct {
	FollowersHandler *FollowersMobileHandlerImpl
}

func NewFollowersMobileRouteImpl(app *fiber.App, db *sqlx.DB) *FollowersMobileRouteImpl {
	h := NewFollowersMobileHandlerImpl(db)
	follower := app.Group("/api/v1/mobile/followers")
	follower.Post("/create", h.Toggle)
	follower.Get("/show", h.Show)
	return &FollowersMobileRouteImpl{
		FollowersHandler: h,
	}
}
