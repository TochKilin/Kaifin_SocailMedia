package songs

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/utls"
)

type Song struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	ArtistID  int64     `json:"artist_id" db:"artist_id"`
	Duration  int       `json:"duration" db:"duration"` // seconds
	FileURL   string    `json:"file_url" db:"file_url"`
	CoverURL  string    `json:"cover_url" db:"cover_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type SongResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	ArtistID  int64     `json:"artist_id"`
	Duration  int       `json:"duration"`
	FileURL   string    `json:"file_url"`
	CoverURL  string    `json:"cover_url"`
	CreatedAt time.Time `json:"created_at"`
}

type SongListResponse struct {
	Songs []SongResponse `json:"songs"`
	Total int            `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

type CreateSongRequest struct {
	Title    string `json:"title" form:"title" validate:"required"`
	Duration int    `json:"duration" form:"duration"`
	FileURL  string `json:"file_url"`
	CoverURL string `json:"cover_url"`
}

func (r *CreateSongRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type UpdateSongRequest struct {
	Title    *string `json:"title" form:"title"`
	Duration *int    `json:"duration" form:"duration"`
	FileURL  *string `json:"file_url"`
	CoverURL *string `json:"cover_url"`
}

func (r *UpdateSongRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type ShowSongRequest struct {
	Search string `query:"search"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
}

func (r *ShowSongRequest) bind(c fiber.Ctx, v *utls.Validator) error {
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

func toSongResponse(s Song) SongResponse {
	return SongResponse{
		ID:        s.ID,
		Title:     s.Title,
		ArtistID:  s.ArtistID,
		Duration:  s.Duration,
		FileURL:   s.FileURL,
		CoverURL:  s.CoverURL,
		CreatedAt: s.CreatedAt,
	}
}
