package template

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type TemplateRouteImpl struct {
	th *TemplateHandlerImpl
}

func NewTemplateRouteImpl(app *fiber.App, dbpool *sqlx.DB) *TemplateRouteImpl {
	h := NewTemplateHandlerImpl(dbpool)

	group := app.Group("/api/v1/front/templates")
	group.Get("/show", h.List)
	group.Get("/image/:id", h.ServeImage)
	group.Post("/create", h.Create)

	return &TemplateRouteImpl{
		th: h,
	}
}
