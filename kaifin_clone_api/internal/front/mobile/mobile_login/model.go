package mobile_login

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/utls"
)

// Login requests
type AuthLoginMobileRequest struct {
	UserName string `json:"user_name" validate:"required,min=4,max=30"`
	Password string `json:"password" validate:"required,min=8"`
	RoleID   int32  `json:"role_id" validate:"omitempty"`
}

// Login response
type AuthUserLoginMobileResponse struct {
	Auth struct {
		Token     string `json:"token"`
		TokenType string `json:"token_type"`
		UserName  string `json:"user_name"`
	} `json:"auth"`
}

// Bind login request
func (r *AuthLoginMobileRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

// User model
type AuthUser struct {
	ID       int64  `db:"id"`
	UserName string `db:"user_name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	RoleID   int32  `db:"role_id"`
}

// Auth response
type AuthProfileResponse struct {
	ID            int64      `json:"id" db:"id"`
	UserName      string     `json:"user_name" db:"user_name"`
	FirstName     string     `json:"first_name" db:"first_name"`
	LastName      string     `json:"last_name" db:"last_name"`
	Email         string     `json:"email" db:"email"`
	ProfileImages *string    `json:"profile_images" db:"profile_images"`
	RoleID        int        `json:"role_id" db:"role_id"`
	RoleName      string     `json:"role_name" db:"role_name"`
	LastLogin     *time.Time `json:"last_login" db:"last_login"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
}
