package article

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type Article struct {
	ID              int64     `json:"id" db:"id"`
	UserID          int64     `json:"user_id" db:"user_id"`
	Username        string    `json:"user_name" db:"user_name"`
	ProfileImage    *string   `json:"profile_images" db:"profile_images"`
	Title           string    `json:"title" db:"title"`
	Summary         *string   `json:"summary" db:"summary"`
	CoverImage      *string   `json:"cover_image" db:"cover_image"`
	Category        string    `json:"category" db:"category"`
	CodeSubcategory *string   `json:"code_subcategory" db:"code_subcategory"`
	Visibility      string    `json:"visibility" db:"visibility"`
	Status          string    `json:"status" db:"status"`
	ViewsCount      int64     `json:"views_count" db:"views_count"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`

	LikeCount    int  `json:"like_count" db:"like_count"`
	Liked        bool `json:"liked" db:"liked"`
	CommentCount int  `json:"comment_count" db:"comment_count"`
	SaveCount    int  `json:"save_count" db:"save_count"`
	Saved        bool `json:"saved" db:"saved"`

	Tags   []string       `json:"tags" db:"-"`
	Blocks []ArticleBlock `json:"blocks,omitempty" db:"-"`
}

type ArticleBlock struct {
	ID        int64     `json:"id" db:"id"`
	ArticleID int64     `json:"article_id" db:"article_id"`
	BlockType string    `json:"block_type" db:"block_type"`
	Title     *string   `json:"title" db:"title"`
	Content   *string   `json:"content" db:"content"`
	Position  int       `json:"position" db:"position"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ArticleBlockInput struct {
	BlockType string `json:"block_type" validate:"required,oneof=text image"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type CreateArticleRequest struct {
	Title           string              `json:"title" validate:"required"`
	Summary         string              `json:"summary"`
	CoverImage      string              `json:"cover_image"`
	Category        string              `json:"category" validate:"required,oneof=general following code read"`
	CodeSubcategory string              `json:"code_subcategory" validate:"omitempty,oneof=backend frontend ai tools"`
	Visibility      string              `json:"visibility" validate:"omitempty,oneof=public private"`
	Tags            []string            `json:"tags"`
	Blocks          []ArticleBlockInput `json:"blocks"`
}

type UpdateArticleRequest struct {
	Title           string              `json:"title" validate:"required"`
	Summary         string              `json:"summary"`
	CoverImage      string              `json:"cover_image"`
	Category        string              `json:"category" validate:"required,oneof=general following code read"`
	CodeSubcategory string              `json:"code_subcategory" validate:"omitempty,oneof=backend frontend ai tools"`
	Visibility      string              `json:"visibility" validate:"omitempty,oneof=public private"`
	Tags            []string            `json:"tags"`
	Blocks          []ArticleBlockInput `json:"blocks"`
}

type ShowArticlesRequest struct {
	Category        string       `query:"category"`
	CodeSubcategory string       `query:"code_subcategory"`
	Search          string       `query:"search"`
	Page            int          `query:"page"`
	PerPage         int          `query:"per_page"`
	Sorts           []share.Sort `query:"sorts"`
}

func (u *ShowArticlesRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(u); err != nil {
		return err
	}
	if u.Page <= 0 {
		u.Page = 1
	}
	if u.PerPage <= 0 || u.PerPage > 50 {
		u.PerPage = 10
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

type ArticlesResponse struct {
	Articles []Article `json:"articles"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	PerPage  int       `json:"per_page"`
}

type ReportArticleRequest struct {
	ArticleID  int64  `json:"article_id" validate:"required"`
	ReportType string `json:"report_type" validate:"required,oneof=bug feedback feature other"`
	Text       string `json:"text"`
}

func (r *ReportArticleRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

func (a *Article) new(req *CreateArticleRequest, uctx *share.UserContext) error {
	a.UserID = uctx.UserID
	a.Title = req.Title
	if req.Summary != "" {
		s := req.Summary
		a.Summary = &s
	}
	if req.CoverImage != "" {
		ci := req.CoverImage
		a.CoverImage = &ci
	}
	a.Category = req.Category
	if req.CodeSubcategory != "" {
		cs := req.CodeSubcategory
		a.CodeSubcategory = &cs
	}
	a.Visibility = req.Visibility
	if a.Visibility == "" {
		a.Visibility = "public"
	}
	a.Status = "published"
	return nil
}

func (a *Article) applyUpdate(req *UpdateArticleRequest) {
	a.Title = req.Title
	if req.Summary != "" {
		s := req.Summary
		a.Summary = &s
	} else {
		a.Summary = nil
	}
	if req.CoverImage != "" {
		ci := req.CoverImage
		a.CoverImage = &ci
	} else {
		a.CoverImage = nil
	}
	a.Category = req.Category
	if req.CodeSubcategory != "" {
		cs := req.CodeSubcategory
		a.CodeSubcategory = &cs
	} else {
		a.CodeSubcategory = nil
	}
	a.Visibility = req.Visibility
	if a.Visibility == "" {
		a.Visibility = "public"
	}
}

func blocksFromInput(input []ArticleBlockInput) []ArticleBlock {
	blocks := make([]ArticleBlock, 0, len(input))
	for i, b := range input {
		blk := ArticleBlock{BlockType: b.BlockType, Position: i}
		if b.Title != "" {
			t := b.Title
			blk.Title = &t
		}
		if b.Content != "" {
			ct := b.Content
			blk.Content = &ct
		}
		blocks = append(blocks, blk)
	}
	return blocks
}
