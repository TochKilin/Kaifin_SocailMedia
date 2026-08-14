package playlist

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/utls"
)

// Playlist maps to the `playlists` table.
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

// ---------- Create ----------
// NOTE: cover_url is a real file upload (multipart file part), not a
// text URL — same pattern as songs.file_url/cover_url. It's populated
// by the handler after saving the uploaded file, not bound from the
// form directly.

type CreatePlaylistRequest struct {
	Name     string  `json:"name" form:"name" validate:"required"`
	CoverURL string  `json:"cover_url"` // set by the handler, not bound from form
	IsPublic *bool   `json:"is_public" form:"is_public"`
	SongIDs  []int64 `json:"song_ids"` // optional — songs to add right away
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
	CoverURL *string  `json:"cover_url"` // set by the handler only if a new file was uploaded
	IsPublic *bool    `json:"is_public" form:"is_public"`
	SongIDs  *[]int64 `json:"song_ids"` // nil = leave songs untouched, [] = clear, [...] = replace
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

// ---------- Show / List ----------

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

// ---------- Top ----------
// Ranking agreed on: is_featured DESC, then songs_count DESC.

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

// ---------- Delete ----------
// Delete only needs the :id path param plus the authenticated user —
// both are resolved in the handler, so there's no request body to bind.

// ---------- Add / Remove a single song ----------
// Lighter-weight alternative to Update's full song_ids replace — lets
// the frontend's "Add to playlist" button add one song without
// re-sending the whole list.

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
