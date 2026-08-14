package music

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type PostMusicRouteImpl struct {
	Handler *PostMusicHandlerImpl
	ws      *websocket.WebSocketManager
}

func NewPostMusicRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *PostMusicRouteImpl {
	handler := NewPostHandlerImpl(dbpool, rdb)
	route := &PostMusicRouteImpl{
		Handler: handler,
		ws:      ws,
	}

	posts := app.Group("/api/v1/front/music")
	posts.Get("/show", handler.Show)
	posts.Post("/create", handler.Create)
	posts.Put("/:id", handler.Update)
	posts.Delete("/:id", handler.Delete)

	return route
}
