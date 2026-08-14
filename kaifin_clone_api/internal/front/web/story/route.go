package story

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type StoryRouteImgl struct {
	sh *StoryHandlerImpl
	ws *websocket.WebSocketManager
}

func NewStoryRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *StoryRouteImgl {
	h := NewStoryHandlerImpl(dbpool, rdb, ws)
	storyGroup := app.Group("/api/v1/front/stories")
	storyGroup.Post("/create", h.Create)
	storyGroup.Get("/show", h.Show)
	storyGroup.Delete("/delete/:id", h.Delete)
	return &StoryRouteImgl{
		sh: h,
		ws: ws,
	}

}
