package post_mobile

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type PostMobileRouteImpl struct {
	ph *PostMobileHandlerImpl
	ws *websocket.WebSocketManager
}

func NewPostMobilrRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *PostMobileRouteImpl {
	p := NewPostMobileHandlerImpl(dbpool, rdb, ws)
	posts := app.Group("/api/v1/mobile/posts")
	posts.Post("/create", p.Create)
	posts.Get("/show", p.Show)
	posts.Delete("/delete/:id", p.Delete)
	posts.Post("/view/:id", p.View)
	return &PostMobileRouteImpl{
		ph: p,
		ws: ws,
	}
}
