package songs

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type SongRouteImpl struct {
	Handler *SongHandlerImpl
	ws      *websocket.WebSocketManager
}

func NewSongRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *SongRouteImpl {
	handler := NewSongHandlerImpl(dbpool, rdb)
	route := &SongRouteImpl{
		Handler: handler,
		ws:      ws,
	}

	songs := app.Group("/api/v1/front/songs")
	songs.Get("/show", handler.Show)
	songs.Post("/create", handler.Create)
	songs.Put("/update:id", handler.Update)
	songs.Delete("/delete/:id", handler.Delete)

	return route
}
