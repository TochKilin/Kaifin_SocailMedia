package quotereport

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

func NewQuoteReportRoute(app *fiber.App, dbpool *sqlx.DB) {
	h := NewQuoteReportHandler(dbpool)

	// user-facing: report a quote
	v1 := app.Group("/api/v1")
	v1.Post("/quote-reports", h.Create)

	// admin-facing: review reports
	admin := app.Group("/api/v1/admin")
	reports := admin.Group("/quote-reports")
	reports.Get("/", h.Show)
	reports.Put("/status/:id", h.UpdateStatus)
}
