package streaklevel

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type LevelRouteImpl struct {
	lh *LevelHandlerImpl
}

func NewLevelRouteImpl(app *fiber.App, dbpool *sqlx.DB) *LevelRouteImpl {
	h := NewLevelHandlerImpl(dbpool)
	levels := app.Group("/api/v1/front/levels")
	levels.Get("/list", h.ListLevels)
	levels.Get("/status", h.GetStatus)
	levels.Post("/checkin", h.CheckIn)
	levels.Get("/leaderboard", h.GetLeaderboard)

	return &LevelRouteImpl{
		lh: h,
	}
}
