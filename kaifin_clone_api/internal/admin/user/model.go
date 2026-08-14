package user

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type UserShowRequest struct {
	PageOption share.Paging   `json:"paging_options" query:"paging_options" validate:"required"`
	Sorts      []share.Sort   `json:"sorts,omitempty" query:"sorts"`
	Filters    []share.Filter `json:"filters.omitempty" query:"filters"`
	Search     string         `json:"q,omitempty" query:"q"`
	CurrencyID int            `json:"currency_id,omitempty" query:"currency_id"`
}

type CreateUserRequest struct {
	UserName  string  `json:"user_name" validate:"required,min=4"`
	Password  string  `json:"password" validate:"required,min=6"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	RoleID    int     `json:"role_id" validate:"required"`
	RoleName  string  `json:"role_name" validate:"required"`
}

func (r *CreateUserRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(&r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type UpdateUserRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	RoleID    *int    `json:"role_id"`
	RoleName  *string `json:"role_name"`
	StatusID  *int    `json:"status_id"`
}

// mirror to db table column
type User struct {
	ID           int64      `json:"id" db:"id"`
	FirstName    *string    `json:"first_name" db:"first_name"`
	LastName     *string    `json:"last_name" db:"last_name"`
	UserName     string     `json:"user_name" db:"user_name"`
	Email        *string    `json:"email" db:"email"`
	Password     string     `json:"-" db:"password"`
	RoleName     string     `json:"role_name" db:"role_name"`
	RoleID       int        `json:"role_id" db:"role_id"`
	IsAdmin      bool       `json:"is_admin" db:"is_admin"`
	LoginSession *string    `json:"-" db:"login_session"`
	LastLogin    *time.Time `json:"last_login" db:"last_login"`
	CurrencyID   *int       `json:"currency_id" db:"currency_id"`
	LanguageID   *int       `json:"language_id" db:"language_id"`
	StatusID     *int       `json:"status_id" db:"status_id"`
	Order        *int       `json:"order" db:"order"`
	CreatedBy    *int64     `json:"-" db:"created_by"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy    *int64     `json:"-" db:"updated_by"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
	DeletedBy    *int64     `json:"-" db:"deleted_by"`
	DeletedAt    *time.Time `json:"-" db:"deleted_at"`
}

func (r *User) new(req *CreateUserRequest, db *sqlx.DB, _ *redis.Client, uctx *share.UserContext) error {
	msg := error_responses.ErrorResponse{}
	newrepo := NewUserRepoImpl(db)
	existing, _ := newrepo.GetByUserName(req.UserName)
	if existing != nil {
		return msg.NewErrorResponse("user_name_taken", fmt.Errorf("username exists"))
	}
	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return msg.NewErrorResponse("password_hash_failed", err)
	}
	created_by := uctx.UserID
	r.UserName = req.UserName
	r.Password = string(hashed)
	r.FirstName = req.FirstName
	r.LastName = req.LastName
	r.Email = req.Email
	r.RoleID = req.RoleID
	r.RoleName = req.RoleName
	r.CreatedBy = &created_by

	return nil
}

type UserResponse struct {
	Users []User `json:"users"`
	Total int
}

type PagingRequest struct {
	Page    int `query:"page" validate:"min=1"`
	PerPage int `query:"per_page" validate:"min=1,max=100"`
}

func (u *UserShowRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(u); err != nil { // call url query string
		return err
	}

	if u.PageOption.Page == 0 {
		if p, err := strconv.Atoi(c.Query("page")); err == nil { // Atio = convert to integer
			u.PageOption.Page = p
		} else {
			u.PageOption.Page = 1
		}
	}

	if u.PageOption.Perpage == 0 {
		if pp, err := strconv.Atoi(c.Query("perpage")); err == nil {
			u.PageOption.Perpage = pp
		} else if pp, err := strconv.Atoi(c.Query("per_page")); err == nil {
			u.PageOption.Perpage = pp
		} else {
			u.PageOption.Perpage = 10
		}
	}

	for i := range u.Filters {
		value := c.Query(fmt.Sprintf("filters[%d][value]", i))
		if intValue, err := strconv.Atoi(value); err == nil {
			u.Filters[i].Value = intValue
		} else if boolValue, err := strconv.ParseBool(value); err == nil {
			u.Filters[i].Value = boolValue
		} else {
			u.Filters[i].Value = value
		}
	}

	if u.Search == "" {
		u.Search = c.Query("q")
	}

	if u.CurrencyID == 0 {
		if v := strings.TrimSpace(c.Query("currency_id")); v != "" {
			if n, err := strconv.Atoi(v); err == nil {
				u.CurrencyID = n
			}
		}
	}

	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}
