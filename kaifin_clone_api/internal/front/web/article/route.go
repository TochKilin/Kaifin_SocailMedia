package article

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type ArticlesRouteImpl struct {
	ah *ArticlesHandlerImpl
}

func NewArticlesRouteImpl(app *fiber.App, dbpool *sqlx.DB) *ArticlesRouteImpl {
	h := NewArticlesHandlerImpl(dbpool)
	articlesGroup := app.Group("/api/v1/front/articles")
	articlesGroup.Post("/create", h.Create)
	articlesGroup.Put("/update/:id", h.Update)
	articlesGroup.Get("/show", h.Show)
	articlesGroup.Get("/:id", h.Detail)
	articlesGroup.Delete("/delete/:id", h.Delete)
	articlesGroup.Post("/:id/like", h.ToggleLike)
	articlesGroup.Post("/:id/save", h.ToggleSave)
	articlesGroup.Post("/report", h.Report)
	articlesGroup.Post("/upload-image", h.UploadImage)
	return &ArticlesRouteImpl{
		ah: h,
	}
}
