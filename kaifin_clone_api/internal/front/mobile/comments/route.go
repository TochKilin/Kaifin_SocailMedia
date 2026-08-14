package comments_mobile

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type CommentsMobileRouteImpl struct {
	ch *CommentsMobileHandlerImpl
}

func NewCommentsMobileRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *CommentsMobileRouteImpl {
	h := NewCommentsMobileHanslerImpl(dbpool)
	comments := app.Group("/api/v1/mobile/comments")
	comments.Post("/create", h.Create)
	comments.Get("/show", h.Show)
	comments.Delete("/delete/:id", h.Delete)
	return &CommentsMobileRouteImpl{
		ch: h,
	}
}
