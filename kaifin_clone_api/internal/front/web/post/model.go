package post

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type HashtagInput struct {
	Name      string
	StickerID *int64
}

type CreatePostRequest struct {
	CommunityID *int64   `form:"community_id"`
	Caption     string   `form:"caption"`
	PostType    string   `form:"post_type"`
	CodeContent *string  `form:"code_content"`
	LinkURL     *string  `form:"link_url"`
	ImageURLs   []string `form:"-"`
	Hashtags    []HashtagInput
	// Hashtags    []string `form:"tag_name"`
	StickerIDs []int64 `form:"sticker_ids[]"`
	VideoURL   *string `form:"-"`
}

type ShowPostRequest struct {
	PageOption share.Paging   `query:"page_option"`
	Filters    []share.Filter `query:"filters"`
	Sorts      []share.Sort   `query:"sorts"`
	Search     string         `query:"search"`
	CurrencyID int            `query:"currency_id"`
	UserID     string         `query:"user_id"`
	FeedOnly   bool           `query:"feed_only"`
}

type Post struct {
	ID                      int64      `json:"id" db:"id"`
	Username                string     `json:"user_name" db:"user_name"`
	UserID                  int64      `json:"user_id" db:"user_id"`
	ProfileImages           *string    `json:"profile_images" db:"profile_images"`
	CommunityID             *int64     `json:"community_id" db:"community_id"`
	Caption                 string     `json:"caption" db:"caption"`
	PostType                string     `json:"post_type" db:"post_type"`
	CodeContent             *string    `json:"code_content" db:"code_content"`
	LinkURL                 *string    `json:"link_url" db:"link_url"`
	ViewsCount              int        `json:"views_count" db:"views_count"`
	CreatedAt               time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt               time.Time  `json:"updated_at" db:"updated_at"`
	Images                  *string    `json:"images" db:"images"`
	Hashtags                *string    `json:"tag_name" db:"tag_name"`
	TagData                 *string    `json:"tag_data" db:"tag_data"`
	CommentCount            int        `json:"comment_count" db:"comment_count"`
	StickerIDs              *string    `json:"sticker_ids" db:"sticker_ids"`
	VideoURL                *string    `json:"video_path" db:"video_path"`
	ThumbnailURL            *string    `json:"thumbnail_path" db:"thumbnail_path"`
	VideoDuration           *int       `db:"duration" json:"duration"`
	RepostID                *int64     `json:"repost_id" db:"repost_id"`
	RepostedByUserID        *int64     `json:"reposted_by_user_id" db:"reposted_by_user_id"`
	RepostedByUsername      *string    `json:"reposted_by_username" db:"reposted_by_username"`
	RepostedAt              *time.Time `json:"reposted_at" db:"reposted_at"`
	SortAt                  time.Time  `json:"-" db:"sort_at"`
	RepostedByProfileImages *string    `json:"reposted_by_profile_images" db:"reposted_by_profile_images"`
	TagStickerIDs           *string    `json:"tag_sticker_ids" db:"tag_sticker_ids"`
}

type CreateShareRequest struct {
	PostID int64 `json:"post_id" form:"post_id" validate:"required"`
}

func (u *CreateShareRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(u); err != nil {
		return err
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

type PostResponse struct {
	Posts []Post `json:"posts"`
	Total int    `json:"total"`
}

func (p *Post) new(req *CreatePostRequest, uctx *share.UserContext) error {
	msg := error_responses.ErrorResponse{}
	switch req.PostType {
	case "text", "image", "link", "code", "video", "quote":

	default:
		return msg.NewErrorResponse(
			"invalid_post_type",
			fmt.Errorf("invalid post type"),
		)
	}

	p.UserID = uctx.UserID
	p.CommunityID = req.CommunityID
	p.Caption = req.Caption
	p.PostType = req.PostType
	p.CodeContent = req.CodeContent
	p.LinkURL = req.LinkURL

	return nil
}

func (u *CreatePostRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(u); err != nil {
		return err
	}
	if err := v.Validate(u); err != nil {
		return err
	}

	return nil
}

func (u *ShowPostRequest) bind(c fiber.Ctx, v *utls.Validator) error {
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

	if v := c.Query("feed_only"); v == "true" || v == "1" {
		u.FeedOnly = true
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

	if u.UserID != "" {
		if uid, err := strconv.ParseInt(u.UserID, 10, 64); err == nil {
			u.Filters = append(u.Filters, share.Filter{
				Property: "p.user_id",
				Value:    int(uid),
			})
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
