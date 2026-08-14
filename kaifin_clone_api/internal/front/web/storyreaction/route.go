package storyreaction

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type StoryReactionRouteImpl struct {
	h *StoryReactionHandlerImpl
}

func NewStoryReactionRouteImpl(app *fiber.App, dbpool *sqlx.DB) *StoryReactionRouteImpl {
	h := NewStoryReactionHandlerImpl(dbpool)
	group := app.Group("/api/v1/front/story-reactions")
	group.Post("/react", h.React)
	group.Get("/show", h.Show)
	group.Delete("/delete/:story_id", h.Delete)
	return &StoryReactionRouteImpl{h: h}
}
