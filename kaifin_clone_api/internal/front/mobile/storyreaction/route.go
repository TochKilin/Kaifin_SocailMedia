package storyreaction_mobile

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type StoryReactionMobileRouteImpl struct {
	h *StoryReactionMobileHandlerImpl
}

func NewStoryReactionMoblieRouteImpl(app *fiber.App, dbpool *sqlx.DB) *StoryReactionMobileRouteImpl {
	h := NewStoryReactionMobileHandlerImpl(dbpool)
	group := app.Group("/api/v1/mobile/story-reactions")
	group.Post("/react", h.React)
	group.Get("/show", h.Show)
	group.Delete("/delete/:story_id", h.Delete)
	return &StoryReactionMobileRouteImpl{h: h}
}
