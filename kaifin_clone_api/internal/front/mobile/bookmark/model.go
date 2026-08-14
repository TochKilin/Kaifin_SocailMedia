package bookmark_mobile

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/utls"
)

type BookMark struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	PostID    int64     `json:"post_id" db:"post_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type BookmarStatus struct {
	BookMarked bool  `json:"bookmarked"`
	Total      int64 `json:"total"`
}

type ToggleBookmarkRequest struct {
	PostID int64 `json:"post_id" validate:"required"`
}

type ShowBookmarkRequest struct {
	Search string `query:"search"`
	PostID int64  `query:"post_id" validate:"required"`
}

func (u *ToggleBookmarkRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(u); err != nil {
		return err
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

type BookmarkResponse struct {
	Bookmarks []BookMark `json:"bookmarks"`
	Total     int        `json:"total"`
}

func (u *ShowBookmarkRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(u); err != nil {
		return err
	}
	if u.Search == "" {
		u.Search = c.Query("q")
	}
	return v.Validate(u)
}
