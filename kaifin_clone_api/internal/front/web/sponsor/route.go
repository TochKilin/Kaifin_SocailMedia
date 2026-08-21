package sponsor

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type SponsorsRouteImpl struct {
	sh *SponsorsHandlerImpl
}

func NewSponsorsRouteImpl(app *fiber.App, dbpool *sqlx.DB) *SponsorsRouteImpl {
	h := NewSponsorsHandlerImpl(dbpool)

	publicGroup := app.Group("/api/v1/front/sponsors")
	publicGroup.Get("/show", h.Show)

	adminGroup := app.Group("/api/v1/admin/sponsors")
	adminGroup.Post("/create", h.Create)
	adminGroup.Put("/update/:id", h.Update)
	adminGroup.Delete("/delete/:id", h.Delete)

	return &SponsorsRouteImpl{sh: h}
}
