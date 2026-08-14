package likes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type LikesRouteImpl struct {
	lh *LikesHandlerImpl
}

func NewLikesRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client) *LikesRouteImpl {
	h := NewLikesHandlerImpl(dbpool, rdb)
	likes := app.Group("/api/v1/front/likes")
	likes.Post("/create", h.Create)
	likes.Get("/show", h.Show)
	return &LikesRouteImpl{lh: h}
}
