package menu

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/utls"
)

type Menu struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	SortOrder int       `json:"sort_order" db:"sort_order"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type MenuResponse struct {
	Menus []Menu `json:"menus"`
	Total int    `json:"total"`
}

type CreateMenuRequest struct {
	Title     string `json:"title" validate:"required"`
	SortOrder int    `json:"sort_order"`
	IsActive  *bool  `json:"is_active"`
}

type UpdateMenuRequest struct {
	Title     *string `json:"title"`
	SortOrder *int    `json:"sort_order"`
	IsActive  *bool   `json:"is_active"`
}

type ShowMenuRequest struct {
	Search string `query:"search"`
}

func (u *CreateMenuRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(u); err != nil {
		return err
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

func (u *ShowMenuRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(u); err != nil {
		return err
	}
	if u.Search == "" {
		u.Search = c.Query("q")
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

func (m *Menu) new(req *CreateMenuRequest) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	if req.Title == "" {
		return msg.NewErrorResponse("invalid_menu_title", fmt.Errorf("menu title is required"))
	}

	m.Title = req.Title
	m.SortOrder = req.SortOrder
	if req.IsActive != nil {
		m.IsActive = *req.IsActive
	} else {
		m.IsActive = true
	}

	return nil
}
