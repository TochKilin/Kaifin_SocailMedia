package likes_mobile

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/utls"
)

type LikeMobile struct {
	ID           int64     `json:"id" db:"id"`
	PostID       int64     `json:"post_id" db:"post_id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	ProfileImage *string   `json:"profile_images" db:"profile_images"`
	Type         string    `json:"reaction_type" db:"reaction_type"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UserName     *string   `json:"user_name" db:"user_name"`
}

type CreateLikeMobileRequest struct {
	PostID int64  `json:"post_id" validate:"required"`
	Type   string `json:"reaction_type"`
}

func (r *CreateLikeMobileRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if r.Type == "" {
		r.Type = "like"
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type ShowLikesMobileRequest struct {
	PostID int64 `query:"post_id" validate:"required"`
}

func (u *ShowLikesMobileRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(u); err != nil {
		return err
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

type LikeMobileSummary struct {
	Type  string `json:"reaaction_type" db:"reaction_type"`
	Count int    `json:"count" db:"count"`
}

type LikesMobileResponse struct {
	Likes      []LikeMobile        `json:"likes"`
	Summary    []LikeMobileSummary `json:"summary"`
	Total      int                 `json:"total"`
	LikedByMe  bool                `json:"liked_by_me"`
	MyReaction string              `json:"my_reaction"`
}
