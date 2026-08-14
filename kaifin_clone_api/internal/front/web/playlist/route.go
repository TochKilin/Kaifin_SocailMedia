package playlist

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type PlaylistRouteImpl struct {
	Handler *PlaylistHandlerImpl
	ws      *websocket.WebSocketManager
}

func NewPlaylistRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *PlaylistRouteImpl {
	handler := NewPlaylistHandlerImpl(dbpool, rdb)
	route := &PlaylistRouteImpl{
		Handler: handler,
		ws:      ws,
	}

	playlists := app.Group("/api/v1/front/playlists")
	playlists.Get("/show", handler.Show)
	playlists.Get("/top", handler.Top)
	playlists.Post("/create", handler.Create)
	playlists.Put("/:id", handler.Update)
	playlists.Delete("/:id", handler.Delete)
	playlists.Post("/:id/songs", handler.AddSong)
	playlists.Delete("/:id/songs/:songId", handler.RemoveSong)
	playlists.Get("/:id", handler.ShowDetail)

	return route
}
