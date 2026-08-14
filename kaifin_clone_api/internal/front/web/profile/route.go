package profile

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type ProfileRouteImpl struct {
	Handler *ProfileHandlerImpl
}

func NewProfileRouteImpl(
	app *fiber.App,
	db *sqlx.DB,
) *ProfileRouteImpl {

	handler := NewProfileHandlerImpl(db)

	api := app.Group(
		"/api/v1/front/profile",
	)

	api.Get("/show", handler.Profile)
	api.Put("/update/:id", handler.Update)
	api.Put("/update-cover", handler.UpdateCover)
	api.Put("/update-info", handler.UpdateInfo)

	return &ProfileRouteImpl{
		Handler: handler,
	}
}
