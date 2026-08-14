package profile_mobile

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type ProfileMobileRouteImpl struct {
	Handler *ProfileMobileHandlerImpl
}

func NewProfileMobileRouteImpl(
	app *fiber.App,
	db *sqlx.DB,
) *ProfileMobileRouteImpl {

	handler := NewProfileMobileHandlerImpl(db)

	api := app.Group(
		"/api/v1/mobile/profile",
	)

	api.Get("/show", handler.Profile)
	api.Put("/update/:id", handler.Update)

	return &ProfileMobileRouteImpl{
		Handler: handler,
	}
}
