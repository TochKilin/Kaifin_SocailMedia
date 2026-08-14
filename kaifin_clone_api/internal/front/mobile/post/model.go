package post_mobile

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

type CreatePostMobileRequest struct {
	CommunityID *int64   `form:"community_id"`
	Caption     string   `form:"caption"`
	PostType    string   `form:"post_type"`
	CodeContent *string  `form:"code_content"`
	LinkURL     *string  `form:"link_url"`
	ImageURLs   []string `form:"-"`
	Hashtags    []string `form:"tag_name"`
	StickerIDs  []int64  `form:"sticker_ids[]"`
	VideoURL    *string  `form:"-"`
}

type ShowPostMobileRequest struct {
	PageOption share.Paging   `query:"page_option"`
	Filters    []share.Filter `query:"filters"`
	Sorts      []share.Sort   `query:"sorts"`
	Search     string         `query:"search"`
	CurrencyID int            `query:"currency_id"`
}

type Post struct {
	ID            int64     `json:"id" db:"id"`
	Username      string    `json:"user_name" db:"user_name"`
	UserID        int64     `json:"user_id" db:"user_id"`
	ProfileImages *string   `json:"profile_images" db:"profile_images"`
	CommunityID   *int64    `json:"community_id" db:"community_id"`
	Caption       string    `json:"caption" db:"caption"`
	PostType      string    `json:"post_type" db:"post_type"`
	CodeContent   *string   `json:"code_content" db:"code_content"`
	LinkURL       *string   `json:"link_url" db:"link_url"`
	ViewsCount    int       `json:"views_count" db:"views_count"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Images        *string   `json:"images" db:"images"`
	Hashtags      *string   `json:"tag_name" db:"tag_name"`
	CommentCount  int       `json:"comment_count" db:"comment_count"`
	StickerIDs    *string   `json:"sticker_ids" db:"sticker_ids"`
	VideoURL      *string   `json:"video_path" db:"video_path"`
	ThumbnailURL  *string   `json:"thumbnail_path" db:"thumbnail_path"`
	VideoDuration *int      `db:"duration" json:"duration"`
}

type PostResponse struct {
	Posts []Post `json:"posts"`
	Total int    `json:"total"`
}

func (p *Post) new(req *CreatePostMobileRequest, uctx *share.UserContext) error {
	msg := error_responses.ErrorResponse{}
	switch req.PostType {
	case "text", "image", "link", "code", "video":

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

func (u *CreatePostMobileRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(u); err != nil {
		return err
	}
	if err := v.Validate(u); err != nil {
		return err
	}

	return nil
}

func (u *ShowPostMobileRequest) bind(c fiber.Ctx, v *utls.Validator) error {
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
