package courseinroll

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type EnrollmentRouteImpl struct {
	Handler *CourseEnrollmentHandler
}

func NewCourseEnrollmentRoute(app *fiber.App, dbpool *sqlx.DB) *EnrollmentRouteImpl {
	h := NewCourseEnrollmentHandler(dbpool)
	group := app.Group("/api/v1/front/course-enrollments")
	group.Get("/show", h.ShowEnrolled)
	group.Post("/create", h.CreateEnrollment)

	return &EnrollmentRouteImpl{
		Handler: h,
	}
}
