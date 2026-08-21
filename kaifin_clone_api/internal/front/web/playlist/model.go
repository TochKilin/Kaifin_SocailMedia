package playlist

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/utls"
)

type Playlist struct {
	ID         int64     `json:"id" db:"id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	Name       string    `json:"name" db:"name"`
	CoverURL   string    `json:"cover_url" db:"cover_url"`
	IsPublic   bool      `json:"is_public" db:"is_public"`
	IsFeatured bool      `json:"is_featured" db:"is_featured"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type PlaylistResponse struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	Name       string    `json:"name"`
	CoverURL   string    `json:"cover_url"`
	IsPublic   bool      `json:"is_public"`
	IsFeatured bool      `json:"is_featured"`
	SongsCount int       `json:"songs_count"`
	CreatedAt  time.Time `json:"created_at"`
}

type PlaylistListResponse struct {
	Playlists []PlaylistResponse `json:"playlists"`
	Total     int                `json:"total"`
	Page      int                `json:"page"`
	Limit     int                `json:"limit"`
}

type CreatePlaylistRequest struct {
	Name     string  `json:"name" form:"name" validate:"required"`
	CoverURL string  `json:"cover_url"`
	IsPublic *bool   `json:"is_public" form:"is_public"`
	SongIDs  []int64 `json:"song_ids"`
}

func (r *CreatePlaylistRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

// ---------- Update ----------

type UpdatePlaylistRequest struct {
	Name     *string  `json:"name" form:"name"`
	CoverURL *string  `json:"cover_url"`
	IsPublic *bool    `json:"is_public" form:"is_public"`
	SongIDs  *[]int64 `json:"song_ids"`
}

func (r *UpdatePlaylistRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type ShowPlaylistRequest struct {
	Search string `query:"search"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
}

func (r *ShowPlaylistRequest) bind(c fiber.Ctx, v *utls.Validator) error {
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

type TopPlaylistRequest struct {
	Limit int `query:"limit"`
}

func (r *TopPlaylistRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(r); err != nil {
		return err
	}
	if r.Limit <= 0 || r.Limit > 50 {
		r.Limit = 4
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type AddPlaylistSongRequest struct {
	SongID int64 `json:"song_id" form:"song_id" validate:"required"`
}

func (r *AddPlaylistSongRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

func toPlaylistResponse(p Playlist, songsCount int) PlaylistResponse {
	return PlaylistResponse{
		ID:         p.ID,
		UserID:     p.UserID,
		Name:       p.Name,
		CoverURL:   p.CoverURL,
		IsPublic:   p.IsPublic,
		IsFeatured: p.IsFeatured,
		SongsCount: songsCount,
		CreatedAt:  p.CreatedAt,
	}
}

type ShowPlaylistDetailRequest struct {
	ID int64 `params:"id"`
}

type PlaylistSongItem struct {
	ID        int64  `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	ArtistID  int64  `json:"artist_id" db:"artist_id"`
	CoverURL  string `json:"cover_url" db:"cover_url"`
	FileURL   string `json:"file_url" db:"file_url"`
	Duration  int    `json:"duration" db:"duration"`
	SortOrder int    `json:"sort_order" db:"sort_order"`
}

type PlaylistDetailResponse struct {
	ID         int64              `json:"id"`
	UserID     int64              `json:"user_id"`
	Name       string             `json:"name"`
	CoverURL   string             `json:"cover_url"`
	IsPublic   bool               `json:"is_public"`
	IsFeatured bool               `json:"is_featured"`
	SongsCount int                `json:"songs_count"`
	Songs      []PlaylistSongItem `json:"songs"`
	CreatedAt  time.Time          `json:"created_at"`
}
