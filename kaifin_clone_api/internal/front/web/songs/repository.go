package songs

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type SongRepo interface {
	Create(artistID int64, req CreateSongRequest) (*SongResponse, *error_responses.ErrorResponse)
	Show(req ShowSongRequest) (*SongListResponse, *error_responses.ErrorResponse)
	GetByID(id int64) (*Song, *error_responses.ErrorResponse)
	Update(id int64, artistID int64, req UpdateSongRequest) (*SongResponse, *error_responses.ErrorResponse)
	Delete(id int64, artistID int64) *error_responses.ErrorResponse
}

type SongRepoImpl struct {
	dbpool *sqlx.DB
}

func NewSongRepoImpl(db *sqlx.DB) *SongRepoImpl {
	return &SongRepoImpl{
		dbpool: db,
	}
}

// ---------- Create ----------

func (r *SongRepoImpl) Create(artistID int64, req CreateSongRequest) (*SongResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var newSong Song
	query := `
		INSERT INTO songs (title, artist_id, duration, file_url, cover_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, artist_id, duration, file_url, cover_url, created_at`

	err := r.dbpool.QueryRowx(
		query,
		req.Title, artistID, req.Duration, req.FileURL, req.CoverURL,
	).StructScan(&newSong)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	resp := toSongResponse(newSong)
	return &resp, nil
}

// ---------- Show / List ----------

func (r *SongRepoImpl) Show(req ShowSongRequest) (*SongListResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	offset := (req.Page - 1) * req.Limit

	baseQuery := `
		FROM songs s
		WHERE ($1 = '' OR s.title ILIKE '%' || $1 || '%')`

	var total int
	countQuery := `SELECT COUNT(*) ` + baseQuery
	if err := r.dbpool.Get(&total, countQuery, req.Search); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var songs []Song
	listQuery := `
		SELECT s.id, s.title, s.artist_id, s.duration, s.file_url, s.cover_url, s.created_at
		` + baseQuery + `
		ORDER BY s.created_at DESC
		LIMIT $2 OFFSET $3`
	if err := r.dbpool.Select(&songs, listQuery, req.Search, req.Limit, offset); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	responses := make([]SongResponse, 0, len(songs))
	for _, s := range songs {
		responses = append(responses, toSongResponse(s))
	}

	return &SongListResponse{
		Songs: responses,
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}

// ---------- GetByID ----------

func (r *SongRepoImpl) GetByID(id int64) (*Song, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var s Song
	query := `
		SELECT id, title, artist_id, duration, file_url, cover_url, created_at
		FROM songs WHERE id = $1`
	if err := r.dbpool.Get(&s, query, id); err != nil {
		return nil, msg.NewErrorResponse("song_not_found", err)
	}
	return &s, nil
}

// ---------- Update ----------

func (r *SongRepoImpl) Update(id int64, artistID int64, req UpdateSongRequest) (*SongResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var updated Song
	query := `
		UPDATE songs SET
			title    = COALESCE($1, title),
			duration = COALESCE($2, duration),
			file_url = COALESCE($3, file_url),
			cover_url = COALESCE($4, cover_url)
		WHERE id = $5 AND artist_id = $6
		RETURNING id, title, artist_id, duration, file_url, cover_url, created_at`

	err := r.dbpool.QueryRowx(
		query,
		req.Title, req.Duration, req.FileURL, req.CoverURL, id, artistID,
	).StructScan(&updated)
	if err != nil {
		// Either the row doesn't exist, or it exists but belongs to a
		// different artist — both surface as the same "not yours" error.
		return nil, msg.NewErrorResponse("song_not_found_or_forbidden", err)
	}

	resp := toSongResponse(updated)
	return &resp, nil
}

// ---------- Delete ----------

func (r *SongRepoImpl) Delete(id int64, artistID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	result, err := r.dbpool.Exec(`DELETE FROM songs WHERE id = $1 AND artist_id = $2`, id, artistID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	if rows == 0 {
		return msg.NewErrorResponse("song_not_found_or_forbidden", fmt.Errorf("song %d not found for artist %d", id, artistID))
	}
	return nil
}
