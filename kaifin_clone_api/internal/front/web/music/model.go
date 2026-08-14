package music

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/utls"
)

// Post maps to the `music_posts` table (renamed so it never collides
// with your existing generic `posts` / post-feed table).
type Post struct {
	ID              int64     `json:"id" db:"id"`
	UserID          int64     `json:"user_id" db:"user_id"`
	SongID          *int64    `json:"song_id" db:"song_id"`
	Content         string    `json:"content" db:"content"`
	Type            string    `json:"type" db:"type"`
	Audience        string    `json:"audience" db:"audience"`
	DisableComments bool      `json:"disable_comments" db:"disable_comments"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

// PostTag maps to `music_post_tags` — "Tag friend name" on a post.
type PostTag struct {
	ID           int64     `json:"id" db:"id"`
	PostID       int64     `json:"post_id" db:"music_post_id"`
	TaggedUserID int64     `json:"tagged_user_id" db:"tagged_user_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type PostResponse struct {
	ID              int64     `json:"id"`
	UserID          int64     `json:"user_id"`
	SongID          *int64    `json:"song_id"`
	Content         string    `json:"content"`
	Type            string    `json:"type"`
	Audience        string    `json:"audience"`
	DisableComments bool      `json:"disable_comments"`
	TaggedUserIDs   []int64   `json:"tagged_user_ids"`
	CreatedAt       time.Time `json:"created_at"`
}

type PostListResponse struct {
	Posts []PostResponse `json:"posts"`
	Total int            `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

// ---------- Create ----------
// Covers: Share to News Feed, Tag friend, Audience/Privacy, Disable
// comments, Standard/Status post. "Send as message" and "Copy link" are
// separate actions (see the message/share modules), not part of create.

type CreatePostRequest struct {
	SongID          *int64  `json:"song_id"`
	Content         string  `json:"content"`
	Type            string  `json:"type" validate:"omitempty,oneof=standard status"`
	Audience        string  `json:"audience" validate:"omitempty,oneof=friend everyone"`
	DisableComments bool    `json:"disable_comments"`
	TaggedUserIDs   []int64 `json:"tagged_user_ids"`
}

func (r *CreatePostRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if r.Type == "" {
		r.Type = "standard"
	}
	if r.Audience == "" {
		r.Audience = "everyone"
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

// validateBusinessRules holds rules a plain struct-tag `validate` can't
// express (e.g. "song_id required only when type=standard").
func (r *CreatePostRequest) validateBusinessRules() *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	if r.Type == "standard" && r.SongID == nil {
		return msg.NewErrorResponse("song_required_for_standard_post", fmt.Errorf("song_id is required for a standard post"))
	}
	if r.Content == "" && r.SongID == nil {
		return msg.NewErrorResponse("empty_post", fmt.Errorf("post must have either content or a song"))
	}
	return nil
}

// ---------- Update ----------
// All fields are pointers so a field left out of the request body is
// left untouched (partial update). TaggedUserIDs: nil = leave tags as
// they are, [] = clear all tags, [...] = replace tags entirely.

type UpdatePostRequest struct {
	Content         *string  `json:"content"`
	Type            *string  `json:"type" validate:"omitempty,oneof=standard status"`
	Audience        *string  `json:"audience" validate:"omitempty,oneof=friend everyone"`
	DisableComments *bool    `json:"disable_comments"`
	TaggedUserIDs   *[]int64 `json:"tagged_user_ids"`
}

func (r *UpdatePostRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *UpdatePostRequest) validateBusinessRules() *error_responses.ErrorResponse {
	// No extra cross-field rule needed yet for updates — add here if
	// your product rules grow (mirrors CreatePostRequest's method).
	return nil
}

// ---------- Show / List ----------

type ShowPostRequest struct {
	Search string `query:"search"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
}

func (r *ShowPostRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(r); err != nil {
		return err
	}
	if r.Search == "" {
		r.Search = c.Query("q")
	}
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.Limit <= 0 || r.Limit > 100 {
		r.Limit = 20
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}
