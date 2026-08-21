package addcard

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type CartRouteImpl struct {
	ch *CartHandlerImpl
}

func NewCartRouteImpl(app *fiber.App, dbpool *sqlx.DB) *CartRouteImpl {
	h := NewCartHandlerImpl(dbpool)

	cartGroup := app.Group("/api/v1/front/cart")
	cartGroup.Post("/create", h.Add)
	cartGroup.Delete("/remove/:course_id", h.Remove)
	cartGroup.Get("/show", h.Show)

	return &CartRouteImpl{
		ch: h,
	}
}
