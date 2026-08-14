package likes_mobile

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type LikesMobileRouteImpl struct {
	lh *LikesMobileHandlerImpl
}

func NewLikesMobileRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client) *LikesMobileRouteImpl {
	h := NewLikesHandlerImpl(dbpool, rdb)
	likes := app.Group("/api/v1/mobile/likes")
	likes.Post("/create", h.Create)
	likes.Get("/show", h.Show)
	return &LikesMobileRouteImpl{lh: h}
}
