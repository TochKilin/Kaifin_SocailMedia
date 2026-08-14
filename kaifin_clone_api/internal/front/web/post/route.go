package post

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type PostRouteImpl struct {
	ph *PostHandlerImpl
	ws *websocket.WebSocketManager
}

func NewPostRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *PostRouteImpl {
	p := NewPostHandlerImpl(dbpool, rdb, ws)

	posts := app.Group("/api/v1/front/posts")
	posts.Post("/create", p.Create)
	posts.Get("/show", p.Show)
	posts.Delete("/delete/:id", p.Delete)
	posts.Post("/view/:id", p.View)
	posts.Post("/shares/create", p.CreateShare)
	posts.Get("/shares/show", p.ShowShares)

	return &PostRouteImpl{
		ph: p,
		ws: ws,
	}
}
