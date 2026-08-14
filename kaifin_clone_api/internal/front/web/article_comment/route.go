package articlecomment

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type CommentsRouteImpl struct {
	ch *CommentsHandlerImpl
}

func NewCommentsRouteImpl(app *fiber.App, dbpool *sqlx.DB) *CommentsRouteImpl {
	h := NewCommentsHandlerImpl(dbpool)

	articlesGroup := app.Group("/api/v1/front/articles")
	articlesGroup.Get("/:id/comments", h.Show)
	articlesGroup.Post("/:id/comments", h.Create)

	commentsGroup := app.Group("/api/v1/front/comments")
	commentsGroup.Put("/:id", h.Update)
	commentsGroup.Delete("/:id", h.Delete)

	return &CommentsRouteImpl{
		ch: h,
	}
}
