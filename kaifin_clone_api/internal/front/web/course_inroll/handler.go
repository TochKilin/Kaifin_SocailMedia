package courseinroll

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
)

type CourseEnrollmentHandler struct {
	Service CourseEnrollmentService
}

func NewCourseEnrollmentHandler(dbpool *sqlx.DB) *CourseEnrollmentHandler {
	return &CourseEnrollmentHandler{
		Service: NewCourseEnrollmentServiceImpl(dbpool),
	}
}

func (h *CourseEnrollmentHandler) ShowEnrolled(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("unauthorized", constants.Generic_invalid, fmt.Errorf("missing user context")),
		)
	}

	req := ShowEnrolledRequest{}
	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("invalid request", constants.Generic_invalid, err),
		)
	}

	result, e := h.Service.ShowEnrolled(uCtx.UserID, req)
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e.MessageID, constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("enrolled courses retrieved", constants.Generic_success, result),
	)
}

func (h *CourseEnrollmentHandler) CreateEnrollment(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("unauthorized", constants.Generic_invalid, fmt.Errorf("missing user context")),
		)
	}

	var req CreateEnrollmentRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("invalid request", constants.Generic_invalid, err),
		)
	}
	if len(req.CourseIDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("course_ids_required", constants.Generic_invalid, fmt.Errorf("course_ids is required")),
		)
	}

	result, e := h.Service.CreateEnrollments(uCtx.UserID, req)
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e.MessageID, constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("enrolled successfully", constants.Generic_success, result),
	)
}
