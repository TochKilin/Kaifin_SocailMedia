package stickers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type StickersRouteImpl struct {
	sh *StickersHandlerImpl
}

func NewStickersRouteImpl(app *fiber.App, dbpool *sqlx.DB) *StickersRouteImpl {
	h := NewStickersHandlerImpl(dbpool)
	stickers := app.Group("/api/v1/front/stickers")
	stickers.Post("/create", h.Create)
	stickers.Get("/show", h.Show)
	stickers.Put("/update", h.Update)

	return &StickersRouteImpl{
		sh: h,
	}
}
