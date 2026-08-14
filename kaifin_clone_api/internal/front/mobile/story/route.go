package story_mobile

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type StoryMobileRouteImgl struct {
	sh *StoryMobileHandlerImpl
	ws *websocket.WebSocketManager
}

func NewStoryMobileRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *StoryMobileRouteImgl {
	h := NewStoryMobileHandlerImpl(dbpool, rdb, ws)
	storyGroup := app.Group("/api/v1/mobile/stories")
	storyGroup.Post("/create", h.Create)
	storyGroup.Get("/show", h.Show)
	storyGroup.Delete("/delete/:id", h.Delete)
	return &StoryMobileRouteImgl{
		sh: h,
		ws: ws,
	}

}
