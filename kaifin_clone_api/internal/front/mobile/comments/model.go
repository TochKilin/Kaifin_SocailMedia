package comments_mobile

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type CommentMobile struct {
	ID           int64     `json:"id" db:"id"`
	PostID       int64     `json:"post_id" db:"post_id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	Username     string    `json:"user_name" db:"user_name"`
	ParentID     *int64    `json:"parent_comment_id" db:"parent_comment_id"`
	Content      string    `json:"content" db:"content"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	ProfileImage *string   `json:"profile_images" db:"profile_images"`
}

type CreateCommentMobileRequest struct {
	PostID   int64  `json:"post_id" validate:"required"`
	ParentID *int64 `json:"parent_comment_id"`
	Content  string `json:"content" validate:"required"`
}

func (r *CreateCommentMobileRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type ShowCommentsMobileRequest struct {
	PostID int64        `query:"post_id" validate:"required"`
	Sorts  []share.Sort `query:"sorts"`
}

func (u *ShowCommentsMobileRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(u); err != nil {
		return err
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

type CommentsMobileResponse struct {
	Comments []CommentMobile `json:"comments"`
	Total    int             `json:"total"`
}

func (c *CommentMobile) new(req *CreateCommentMobileRequest, uctx *share.UserContext) error {
	c.PostID = req.PostID
	c.UserID = uctx.UserID
	c.ParentID = req.ParentID
	c.Content = req.Content
	return nil
}
