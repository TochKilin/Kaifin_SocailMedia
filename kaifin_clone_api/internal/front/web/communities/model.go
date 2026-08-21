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
	AvatarURL   *string   `json:"avatar_url" db:"avatar_url"`
	CoverURL    *string   `json:"cover_url" db:"cover_url"`
	CategoryID  *int64    `json:"category_id" db:"category_id"`
	IsVerified  bool      `json:"is_verified" db:"is_verified"`
	Privacy     string    `json:"privacy" db:"privacy"`
	MemberCount int64     `json:"member_count" db:"member_count"`
	HotScore    float64   `json:"hot_score" db:"hot_score"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	CreatedBy   *int64    `json:"created_by" db:"created_by"`

	IsJoined bool `json:"is_joined" db:"is_joined"`
}

type CommunitiesResponse struct {
	Communities []Communities `json:"communities"`
	Total       int           `json:"total"`
}

type CreatCommunitiesRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	AvatarURL   *string `json:"avatar_url"`
	CoverURL    *string `json:"cover_url"`
	CategoryID  *int64  `json:"category_id"`
	Privacy     string  `json:"privacy" validate:"omitempty,oneof=public private"`
}

func (r *CreatCommunitiesRequest) validateOnly(v *utls.Validator) error {
	if r.Privacy == "" {
		r.Privacy = "public"
	}
	return v.Validate(r)
}

func (r *CreatCommunitiesRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(&r); err != nil {
		return err
	}
	if r.Privacy == "" {
		r.Privacy = "public"
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
	CategoryID *int64         `query:"category_id"`
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

	if catID, err := strconv.ParseInt(c.Query("category_id"), 10, 64); err == nil {
		u.CategoryID = &catID
	}

	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

type ShowCommunityDetailRequest struct {
	ID int64 `params:"id" validate:"required"`
}

func (r *ShowCommunityDetailRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().URI(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type CommunityMember struct {
	ID           int64     `json:"id" db:"id"`
	CommunityID  int64     `json:"community_id" db:"community_id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	Username     string    `json:"user_name" db:"user_name"`
	ProfileImage *string   `json:"profile_images" db:"profile_images"`
	Role         string    `json:"role" db:"role"`
	Status       string    `json:"status" db:"status"`
	JoinedAt     time.Time `json:"joined_at" db:"joined_at"`
}

type ShowMembersRequest struct {
	CommunityID int64        `params:"id" validate:"required"`
	PageOption  share.Paging `query:"page_option"`
}

func (r *ShowMembersRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().URI(r); err != nil {
		return err
	}
	communityID := r.CommunityID

	if err := c.Bind().Query(r); err != nil {
		return err
	}
	r.CommunityID = communityID

	if r.PageOption.Page == 0 {
		if p, err := strconv.Atoi(c.Query("page")); err == nil {
			r.PageOption.Page = p
		} else {
			r.PageOption.Page = 1
		}
	}
	if r.PageOption.Perpage == 0 {
		if pp, err := strconv.Atoi(c.Query("perpage")); err == nil {
			r.PageOption.Perpage = pp
		} else {
			r.PageOption.Perpage = 30
		}
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type MembersResponse struct {
	Members []CommunityMember `json:"members"`
	Total   int               `json:"total"`
}

type UpdateCommunityImageRequest struct {
	ID int64 `params:"id" validate:"required"`
}

func (r *UpdateCommunityImageRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().URI(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}
