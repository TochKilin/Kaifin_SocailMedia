package quote

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type QuoteShowRequest struct {
	PageOption share.Paging   `json:"paging_options" query:"paging_options" validate:"required"`
	Sorts      []share.Sort   `json:"sorts,omitempty" query:"sorts"`
	Filters    []share.Filter `json:"filters,omitempty" query:"filters"`
	Search     string         `json:"q,omitempty" query:"q"`
	Tab        string         `json:"tab,omitempty" query:"tab"` // popular | latest
}

type CreateQuoteRequest struct {
	Title      string `json:"title" validate:"required,max=150"`
	Content    string `json:"content" validate:"required,max=100"`
	Visibility string `json:"visibility" validate:"omitempty,oneof=public private"`
	Status     string `json:"status" validate:"omitempty,oneof=draft published"`
}

func (r *CreateQuoteRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type UpdateQuoteRequest struct {
	Title      *string `json:"title"`
	Content    *string `json:"content"`
	Visibility *string `json:"visibility"`
	Status     *string `json:"status"`
}

// mirror to db table column (quotes)
type Quote struct {
	ID            int64     `json:"id" db:"id"`
	UserID        int64     `json:"user_id" db:"user_id"`
	Title         string    `json:"title" db:"title"`
	Content       string    `json:"content" db:"content"`
	Visibility    string    `json:"visibility" db:"visibility"`
	Status        string    `json:"status" db:"status"`
	ViewsCount    int       `json:"views_count" db:"views_count"`
	LikesCount    int       `json:"likes_count" db:"likes_count"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	ProfileImages string    `json:"profile_images,omitempty" db:"profile_images"`

	// joined field (មិនមែនក្នុង table quotes ទេ តែយកតាម query join)
	Username         string `json:"username,omitempty" db:"username"`
	MyReactionTypeID *int16 `json:"my_reaction_type_id,omitempty" db:"my_reaction_type_id"`
}

func (r *Quote) new(req *CreateQuoteRequest, uctx *share.UserContext) {
	r.UserID = uctx.UserID
	r.Title = req.Title
	r.Content = req.Content

	if req.Visibility == "" {
		r.Visibility = "public"
	} else {
		r.Visibility = req.Visibility
	}

	if req.Status == "" {
		r.Status = "published"
	} else {
		r.Status = req.Status
	}
}

type QuoteResponse struct {
	Quotes []Quote `json:"quotes"`
	Total  int
}

func (u *QuoteShowRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(u); err != nil {
		return err
	}

	if u.PageOption.Page == 0 {
		u.PageOption.Page = 1
	}
	if u.PageOption.Perpage == 0 {
		u.PageOption.Perpage = 10
	}
	if u.Search == "" {
		u.Search = c.Query("q")
	}
	if u.Tab == "" {
		u.Tab = c.Query("tab") // popular | latest
	}

	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}
