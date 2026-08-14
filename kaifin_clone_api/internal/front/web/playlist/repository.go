package playlist

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type PlaylistRepo interface {
	Create(userID int64, req CreatePlaylistRequest) (*PlaylistResponse, *error_responses.ErrorResponse)
	Show(requesterID int64, req ShowPlaylistRequest) (*PlaylistListResponse, *error_responses.ErrorResponse)
	Top(limit int) ([]PlaylistResponse, *error_responses.ErrorResponse)
	GetByID(id int64) (*Playlist, *error_responses.ErrorResponse)
	Update(id int64, userID int64, req UpdatePlaylistRequest) (*PlaylistResponse, *error_responses.ErrorResponse)
	Delete(id int64, userID int64) *error_responses.ErrorResponse
	AddSong(playlistID int64, userID int64, songID int64) *error_responses.ErrorResponse
	RemoveSong(playlistID int64, userID int64, songID int64) *error_responses.ErrorResponse
	GetDetailByID(id int64) (*PlaylistDetailResponse, *error_responses.ErrorResponse)
}

type PlaylistRepoImpl struct {
	dbpool *sqlx.DB
}

func NewPlaylistRepoImpl(db *sqlx.DB) *PlaylistRepoImpl {
	return &PlaylistRepoImpl{dbpool: db}
}

// ---------- Create ----------

func (r *PlaylistRepoImpl) Create(userID int64, req CreatePlaylistRequest) (*PlaylistResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	isPublic := true
	if req.IsPublic != nil {
		isPublic = *req.IsPublic
	}

	tx, err := r.dbpool.Beginx()
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	var newPlaylist Playlist
	insertQuery := `
		INSERT INTO playlists (user_id, name, cover_url, is_public)
		VALUES ($1, $2, $3, $4)
		RETURNING id, user_id, name, cover_url, is_public, is_featured, created_at`

	err = tx.QueryRowx(insertQuery, userID, req.Name, req.CoverURL, isPublic).StructScan(&newPlaylist)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	songsCount, err := insertPlaylistSongs(tx, newPlaylist.ID, req.SongIDs)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	resp := toPlaylistResponse(newPlaylist, songsCount)
	return &resp, nil
}

// ---------- Show / List ----------

func (r *PlaylistRepoImpl) Show(requesterID int64, req ShowPlaylistRequest) (*PlaylistListResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	offset := (req.Page - 1) * req.Limit

	// Visible if it's public, or it's the requester's own playlist.
	baseQuery := `
		FROM playlists p
		WHERE (p.is_public = TRUE OR p.user_id = $1)
		AND ($2 = '' OR p.name ILIKE '%' || $2 || '%')`

	var total int
	countQuery := `SELECT COUNT(*) ` + baseQuery
	if err := r.dbpool.Get(&total, countQuery, requesterID, req.Search); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var playlists []Playlist
	listQuery := `
		SELECT p.id, p.user_id, p.name, p.cover_url, p.is_public, p.is_featured, p.created_at
		` + baseQuery + `
		ORDER BY p.created_at DESC
		LIMIT $3 OFFSET $4`
	if err := r.dbpool.Select(&playlists, listQuery, requesterID, req.Search, req.Limit, offset); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	responses, err := r.attachSongsCount(playlists)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &PlaylistListResponse{
		Playlists: responses,
		Total:     total,
		Page:      req.Page,
		Limit:     req.Limit,
	}, nil
}

// ---------- Top ----------
// Ranking: is_featured DESC, then songs_count DESC — agreed with the
// product owner as the "Top Playlist" ordering (see conversation).

func (r *PlaylistRepoImpl) Top(limit int) ([]PlaylistResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	query := `
		SELECT p.id, p.user_id, p.name, p.cover_url, p.is_public, p.is_featured, p.created_at,
			COUNT(ps.id) AS songs_count
		FROM playlists p
		LEFT JOIN playlist_songs ps ON ps.playlist_id = p.id
		WHERE p.is_public = TRUE
		GROUP BY p.id
		ORDER BY p.is_featured DESC, songs_count DESC, p.created_at DESC
		LIMIT $1`

	rows, err := r.dbpool.Queryx(query, limit)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	defer rows.Close()

	responses := make([]PlaylistResponse, 0, limit)
	for rows.Next() {
		var p Playlist
		var songsCount int
		if err := rows.Scan(
			&p.ID, &p.UserID, &p.Name, &p.CoverURL, &p.IsPublic, &p.IsFeatured, &p.CreatedAt, &songsCount,
		); err != nil {
			return nil, msg.NewErrorResponse("database_error", err)
		}
		responses = append(responses, toPlaylistResponse(p, songsCount))
	}
	if err := rows.Err(); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return responses, nil
}

// ---------- GetByID ----------

func (r *PlaylistRepoImpl) GetByID(id int64) (*Playlist, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var p Playlist
	query := `
		SELECT id, user_id, name, cover_url, is_public, is_featured, created_at
		FROM playlists WHERE id = $1`
	if err := r.dbpool.Get(&p, query, id); err != nil {
		return nil, msg.NewErrorResponse("playlist_not_found", err)
	}
	return &p, nil
}

// ---------- Update ----------

func (r *PlaylistRepoImpl) Update(id int64, userID int64, req UpdatePlaylistRequest) (*PlaylistResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	tx, err := r.dbpool.Beginx()
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	var updated Playlist
	updateQuery := `
		UPDATE playlists SET
			name      = COALESCE($1, name),
			cover_url = COALESCE($2, cover_url),
			is_public = COALESCE($3, is_public)
		WHERE id = $4 AND user_id = $5
		RETURNING id, user_id, name, cover_url, is_public, is_featured, created_at`

	err = tx.QueryRowx(updateQuery, req.Name, req.CoverURL, req.IsPublic, id, userID).StructScan(&updated)
	if err != nil {
		return nil, msg.NewErrorResponse("playlist_not_found_or_forbidden", err)
	}

	var songsCount int
	if req.SongIDs != nil {
		if _, err := tx.Exec(`DELETE FROM playlist_songs WHERE playlist_id = $1`, id); err != nil {
			return nil, msg.NewErrorResponse("database_error", err)
		}
		songsCount, err = insertPlaylistSongs(tx, id, *req.SongIDs)
		if err != nil {
			return nil, msg.NewErrorResponse("database_error", err)
		}
	} else {
		if err := tx.Get(&songsCount, `SELECT COUNT(*) FROM playlist_songs WHERE playlist_id = $1`, id); err != nil {
			return nil, msg.NewErrorResponse("database_error", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	resp := toPlaylistResponse(updated, songsCount)
	return &resp, nil
}

// ---------- Delete ----------

func (r *PlaylistRepoImpl) Delete(id int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	tx, err := r.dbpool.Beginx()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	if _, err := tx.Exec(`DELETE FROM playlist_songs WHERE playlist_id = $1`, id); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	result, err := tx.Exec(`DELETE FROM playlists WHERE id = $1 AND user_id = $2`, id, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	if rows == 0 {
		return msg.NewErrorResponse("playlist_not_found_or_forbidden", fmt.Errorf("playlist %d not found for user %d", id, userID))
	}

	if err := tx.Commit(); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

// ---------- Add / Remove a single song ----------

func (r *PlaylistRepoImpl) AddSong(playlistID int64, userID int64, songID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	// Ownership check first — only the playlist's owner can add songs to it.
	var owns bool
	if err := r.dbpool.Get(&owns, `SELECT EXISTS(SELECT 1 FROM playlists WHERE id = $1 AND user_id = $2)`, playlistID, userID); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	if !owns {
		return msg.NewErrorResponse("playlist_not_found_or_forbidden", fmt.Errorf("playlist %d not found for user %d", playlistID, userID))
	}

	var nextOrder int
	if err := r.dbpool.Get(&nextOrder, `SELECT COALESCE(MAX(sort_order), -1) + 1 FROM playlist_songs WHERE playlist_id = $1`, playlistID); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	_, err := r.dbpool.Exec(
		`INSERT INTO playlist_songs (playlist_id, song_id, sort_order)
		 VALUES ($1, $2, $3)
		 ON CONFLICT (playlist_id, song_id) DO NOTHING`,
		playlistID, songID, nextOrder,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *PlaylistRepoImpl) RemoveSong(playlistID int64, userID int64, songID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	var owns bool
	if err := r.dbpool.Get(&owns, `SELECT EXISTS(SELECT 1 FROM playlists WHERE id = $1 AND user_id = $2)`, playlistID, userID); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	if !owns {
		return msg.NewErrorResponse("playlist_not_found_or_forbidden", fmt.Errorf("playlist %d not found for user %d", playlistID, userID))
	}

	if _, err := r.dbpool.Exec(`DELETE FROM playlist_songs WHERE playlist_id = $1 AND song_id = $2`, playlistID, songID); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

// ---------- shared helpers ----------

func insertPlaylistSongs(tx *sqlx.Tx, playlistID int64, songIDs []int64) (int, error) {
	if len(songIDs) == 0 {
		return 0, nil
	}
	query := `INSERT INTO playlist_songs (playlist_id, song_id, sort_order) VALUES ($1, $2, $3)`
	count := 0
	for i, songID := range songIDs {
		if _, err := tx.Exec(query, playlistID, songID, i); err != nil {
			return 0, err
		}
		count++
	}
	return count, nil
}

func (r *PlaylistRepoImpl) attachSongsCount(playlists []Playlist) ([]PlaylistResponse, error) {
	responses := make([]PlaylistResponse, 0, len(playlists))
	for _, p := range playlists {
		var count int
		if err := r.dbpool.Get(&count, `SELECT COUNT(*) FROM playlist_songs WHERE playlist_id = $1`, p.ID); err != nil {
			return nil, err
		}
		responses = append(responses, toPlaylistResponse(p, count))
	}
	return responses, nil
}

// func (r *PlaylistRepoImpl) GetDetailByID(id int64) (*PlaylistDetailResponse, *error_responses.ErrorResponse) {
// 	msg := error_responses.ErrorResponse{}

// 	var p Playlist
// 	if err := r.dbpool.Get(&p, `SELECT id, user_id, name, cover_url, is_public, is_featured, created_at FROM playlists WHERE id = $1`, id); err != nil {
// 		return nil, msg.NewErrorResponse("playlist_not_found", err)
// 	}

// 	var songs []PlaylistSongItem
// 	query := `
// 		SELECT s.id, s.title, s.singer_name, s.cover_url, s.file_url, s.duration, ps.sort_order
// 		FROM playlist_songs ps
// 		JOIN songs s ON s.id = ps.song_id
// 		WHERE ps.playlist_id = $1
// 		ORDER BY ps.sort_order ASC`
// 	if err := r.dbpool.Select(&songs, query, id); err != nil {
// 		return nil, msg.NewErrorResponse("database_error", err)
// 	}

// 	return &PlaylistDetailResponse{
// 		ID: p.ID, UserID: p.UserID, Name: p.Name, CoverURL: p.CoverURL,
// 		IsPublic: p.IsPublic, IsFeatured: p.IsFeatured,
// 		SongsCount: len(songs), Songs: songs, CreatedAt: p.CreatedAt,
// 	}, nil
// }

func (r *PlaylistRepoImpl) GetDetailByID(id int64) (*PlaylistDetailResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var p Playlist
	query := `
		SELECT id, user_id, name, cover_url, is_public, is_featured, created_at
		FROM playlists WHERE id = $1`
	if err := r.dbpool.Get(&p, query, id); err != nil {
		return nil, msg.NewErrorResponse("playlist_not_found", err)
	}

	var songs []PlaylistSongItem
	songsQuery := `
		SELECT
			s.id, s.title, s.artist_id, s.cover_url, s.file_url, s.duration,
			ps.sort_order
		FROM playlist_songs ps
		JOIN songs s ON s.id = ps.song_id
		WHERE ps.playlist_id = $1
		ORDER BY ps.sort_order ASC`
	if err := r.dbpool.Select(&songs, songsQuery, id); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &PlaylistDetailResponse{
		ID:         p.ID,
		UserID:     p.UserID,
		Name:       p.Name,
		CoverURL:   p.CoverURL,
		IsPublic:   p.IsPublic,
		IsFeatured: p.IsFeatured,
		SongsCount: len(songs),
		Songs:      songs,
		CreatedAt:  p.CreatedAt,
	}, nil
}
