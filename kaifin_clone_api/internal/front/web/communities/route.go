package communities

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type CommunitiesRouteImpl struct {
	ch *CommunitiesHandlerImpl
}

func NewCommunitiesRoute(app *fiber.App, dbpool *sqlx.DB) *CommunitiesRouteImpl {
	h := NewCommunitiesHandlerImpl(app, dbpool)
	comm := app.Group("/api/v1/front/communities")
	comm.Get("/show", h.Show)
	comm.Get("/:id", h.ShowDetail)
	comm.Get("/:id/members", h.ShowMembers)
	comm.Post("/create", h.Create)
	comm.Post("/:id/join", h.ToggleJoin)
	comm.Put("/:id/avatar", h.UpdateAvatar)
	comm.Put("/:id/cover", h.UpdateCover)
	return &CommunitiesRouteImpl{
		ch: h,
	}
}
