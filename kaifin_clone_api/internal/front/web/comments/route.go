package comments

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type CommentsRouteImpl struct {
	ch *CommentsHandlerImpl
}

func NewCommentsRouteImpl(app *fiber.App, dbpool *sqlx.DB) *CommentsRouteImpl {
	h := NewCommentsHanslerImpl(dbpool)
	comments := app.Group("/api/v1/front/comments")
	comments.Post("/create", h.Create)
	comments.Get("/show", h.Show)
	comments.Delete("/delete/:id", h.Delete)
	comments.Post("/:id/like", h.ToggleLike)
	return &CommentsRouteImpl{
		ch: h,
	}
}
