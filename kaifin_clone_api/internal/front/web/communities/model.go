package communities

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type Communities struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	CreatedBy   *int64    `json:"created_by" db:"created_by"`
}

type CommunitiesResponse struct {
	Communities []Communities `json:"communities"`
	Total       int           `json:"total"`
}

type CreatCommunitiesRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

func (r *CreatCommunitiesRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(&r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type ShowCommunitiesRequest struct {
	PageOption share.Paging   `query:"page_option"`
	Filters    []share.Filter `query:"filters"`
	Sorts      []share.Sort   `query:"sorts"`
	Search     string         `query:"search"`
}

func (u *ShowCommunitiesRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(u); err != nil {
		return err
	}

	if u.PageOption.Page == 0 {
		if p, err := strconv.Atoi(c.Query("page")); err == nil {
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

	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}
