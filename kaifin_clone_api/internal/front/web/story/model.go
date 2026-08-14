package story

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type CreateStoryRequest struct {
	MediaType string `form:"media_type"`
	MediaURL  string `form:"-"`
}

type ShowStoryRequest struct {
	PageOption share.Paging `query:"page_option"`
	UserID     *int64       `query:"user_id"`
}

type StoryResponse struct {
	Stories []Story `json:"stories"`
	Total   int     `json:"total"`
}

type Story struct {
	ID            int64     `json:"id" db:"id"`
	UserID        int64     `json:"user_id" db:"user_id"`
	Username      string    `json:"user_name" db:"user_name"`
	ProfileImages *string   `json:"profile_images" db:"profile_images"`
	MediaURL      string    `json:"media_url" db:"media_url"`
	MediaType     string    `json:"media_type" db:"media_type"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	ExpiresAt     time.Time `json:"expires_at" db:"expires_at"`
}

func (u *CreateStoryRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(u); err != nil {
		return err
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

func (s *Story) new(req *CreateStoryRequest, uctx *share.UserContext) error {
	msg := error_responses.ErrorResponse{}

	if req.MediaURL == "" {
		return msg.NewErrorResponse("media_required", fmt.Errorf("media file is required"))
	}

	switch req.MediaType {
	case "image", "video":
	default:
		req.MediaType = "image"
	}

	s.UserID = uctx.UserID
	s.MediaURL = req.MediaURL
	s.MediaType = req.MediaType

	return nil
}

func (u *ShowStoryRequest) bind(c fiber.Ctx, v *utls.Validator) error {
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
		} else {
			u.PageOption.Perpage = 20
		}
	}

	if uid := c.Query("user_id"); uid != "" {
		if n, err := strconv.ParseInt(uid, 10, 64); err == nil {
			u.UserID = &n
		}
	}

	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}
