package auth

import (
	//internal package
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/utls" //community package
)

type AuthLoginRequest struct {
	UserName string `json:"user_name" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=4"`
}

type AuthLoginResponse struct {
	Auth struct {
		Token     string `json:"token"`
		TokenType string `json:"token_type"`
		UserName  string `json:"user_name"`
	} `json:"auth"`
}

func (r *AuthLoginRequest) bind(c fiber.Ctx, v *utls.Validator) error {
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

type Auth struct {
	// type Users []User `json:"users"`,
	ID           int       `json:"id" db:"id"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	UserName     string    `json:"user_name" db:"user_name"`
	Email        string    `json:"email" db:"email"`
	Password     string    `json:"password" db:"password"`
	RoleName     string    `json:"role_name" db:"role_name"`
	RoleID       int32     `json:"role_id" db:"role_id"`
	IsAdmin      bool      `json:"is_admin" db:"is_admin"`
	LoginSession string    `json:"login_session" db:"login_session"`
	LastLogin    time.Time `json:"last_login" db:"last_login"`
	CurrencyID   int       `json:"currency_id" db:"currency_id"`
	LanguageID   int       `json:"language_id" db:"language_id"`
	StatusID     int16     `json:"status_id" db:"status_id"`
	Order        int       `json:"order" db:"order"`
	CreatedBy    int       `json:"created_by" db:"created_by"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedBy    int       `json:"updated_by" db:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	DeletedBy    int       `json:"deleted_by" db:"deleted_by"`
}
