// request.go
package userregister

import (
	"fmt"
	"mime/multipart"
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

type AuthRegisterRequest struct {
	FirstName     string                `form:"first_name" validate:"required,min=2,max=50"`
	LastName      string                `form:"last_name" validate:"required,min=2,max=50"`
	UserName      string                `form:"user_name" validate:"required,min=4,max=30"`
	Email         string                `form:"email" validate:"required,email"`
	Password      string                `form:"password" validate:"required,min=8"`
	RoleName      string                `form:"-"`
	RoleID        int                   `form:"-"`
	ProfileImages *multipart.FileHeader `form:"profile_images"`
}

func (r *AuthRegisterRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	fmt.Printf("REQUEST DATA: %+v\n", r)
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type UserShowRequest struct {
	PageOption share.Paging   `json:"paging_options" query:"paging_options" validate:"required"`
	Sorts      []share.Sort   `json:"sorts,omitempty" query:"sorts"`
	Filters    []share.Filter `json:"filters.omitempty" query:"filters"`
	Search     string         `json:"q,omitempty" query:"q"`
	CurrencyID int            `json:"currency_id,omitempty" query:"currency_id"`
}

type UpdateUserRequest struct {
	FirstName     *string               `form:"first_name"`
	LastName      *string               `form:"last_name"`
	Email         *string               `form:"email"`
	RoleID        *int                  `form:"role_id"`
	ProfileImages *multipart.FileHeader `form:"profile_images"`
	RoleName      *string               `form:"role_name"`
	StatusID      *int                  `form:"status_id"`
}

type PagingRequest struct {
	Page    int `query:"page" validate:"min=1"`
	PerPage int `query:"per_page" validate:"min=1,max=100"`
}

type UserProfileResponse struct {
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

func (u *User) new(req *AuthRegisterRequest, c fiber.Ctx, db *sqlx.DB, rdb *redis.Client, uctx *share.UserContext) error {
	msg := error_responses.ErrorResponse{}
	newrepo := NewUserRegisterRepoImpl(db, rdb)
	existing, _ := newrepo.GetByUserName(req.UserName)
	if existing != nil {
		return msg.NewErrorResponse("user_name_taken", fmt.Errorf("username exists"))
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return msg.NewErrorResponse("password_hash_failed", err)
	}

	var createdBy int64 = 0
	if uctx != nil {
		createdBy = uctx.UserID
	}

	u.UserName = req.UserName
	u.Password = string(hashed)
	u.FirstName = req.FirstName
	u.LastName = req.LastName
	u.Email = req.Email
	u.RoleID = 4
	u.RoleName = "user"
	u.IsAdmin = false
	if req.ProfileImages != nil {
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), req.ProfileImages.Filename)
		savePath := "./uploads/" + filename

		if err := c.SaveFile(req.ProfileImages, savePath); err != nil {
			return msg.NewErrorResponse("upload_failed", err)
		}

		u.ProfileImages = &filename
	}

	u.CreatedBy = &createdBy

	return nil
}

type User struct {
	ID        int64  `db:"id"`
	UserName  string `db:"user_name"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Password  string `db:"password"`

	LoginSession *string    `db:"login_session" json:"-"`
	LastLogin    *time.Time `db:"last_login" json:"last_login"`

	CurrencyID *int `db:"currency_id" json:"currency_id"`
	LanguageID *int `db:"language_id" json:"language_id"`
	StatusID   *int `db:"status_id" json:"status_id"`
	Order      *int `db:"order" json:"order"`

	RoleName string `db:"role_name"`
	RoleID   int    `db:"role_id"`
	IsAdmin  bool   `db:"is_admin"`

	UpdatedBy *int64     `db:"updated_by"`
	UpdatedAt *time.Time `db:"updated_at"`

	DeletedBy *int64     `db:"deleted_by"`
	DeletedAt *time.Time `db:"deleted_at"`

	CreatedBy     *int64    `db:"created_by"`
	ProfileImages *string   `db:"profile_images"`
	CreatedAt     time.Time `db:"created_at"`
}

type UserResponse struct {
	Users []User `json:"users"`
	Total int
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
