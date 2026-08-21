package addcard

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type CartHandlerImpl struct {
	cs *CartServiceImpl
}

func NewCartHandlerImpl(dbpool *sqlx.DB) *CartHandlerImpl {
	return &CartHandlerImpl{cs: NewCartServiceImpl(dbpool)}
}

func (h *CartHandlerImpl) Add(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	var req AddToCartRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Generic_invalid, err),
		)
	}
	v := utls.NewValidator()
	if verr := v.Validate(&req); verr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, verr),
		)
	}

	if e := h.cs.AddItem(uCtx.UserID, req.CourseID); e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Failed to add to cart", constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse("added_to_cart", constants.Generic_success, true),
	)
}

func (h *CartHandlerImpl) Remove(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	courseID, err := strconv.ParseInt(c.Params("course_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid course id", constants.Generic_invalid, err),
		)
	}

	if e := h.cs.RemoveItem(uCtx.UserID, courseID); e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Failed to remove from cart", constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("removed_from_cart", constants.Generic_success, nil),
	)
}

func (h *CartHandlerImpl) Show(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	data, e := h.cs.Show(uCtx.UserID)
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Failed to fetch cart", constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("cart_retrieved", constants.Generic_success, data),
	)
}
