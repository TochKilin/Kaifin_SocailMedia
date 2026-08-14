package mystickerset

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type MysetStickerRepo interface {
}

type MysetStickerRepoImpl struct {
	dbpool *sqlx.DB
}

func NewMysetStickerRepoImpl(db *sqlx.DB) *MysetStickerRepoImpl {
	return &MysetStickerRepoImpl{
		dbpool: db,
	}
}

func (r *MysetStickerRepoImpl) Show(userID int64) (*MyStickerSetsResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	type setRow struct {
		PackID    int64     `db:"pack_id"`
		PackName  string    `db:"pack_name"`
		AddedAt   time.Time `db:"added_at"`
		StickerID *int64    `db:"sticker_id"`
	}
	var setRows []setRow

	err := r.dbpool.Select(&setRows, `
		SELECT us.pack_id, sp.pack_name, us.added_at
		FROM user_stickers us
		JOIN sticker_packs sp ON sp.id = us.pack_id
		WHERE us.user_id = $1
		ORDER BY us.added_at DESC
	`, userID)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	result := make([]MyStickerSet, 0, len(setRows))
	for _, sr := range setRows {
		type stickerIDRow struct {
			ID int64 `db:"id"`
		}
		var stickerIDs []stickerIDRow

		err := r.dbpool.Select(&stickerIDs, `
			SELECT id FROM stickers WHERE pack_id = $1 ORDER BY id ASC
		`, sr.PackID)
		if err != nil {
			return nil, msg.NewErrorResponse("database_error", err)
		}

		stickers := make([]StickerInPack, 0, len(stickerIDs))
		for _, s := range stickerIDs {
			stickers = append(stickers, StickerInPack{
				ID:  s.ID,
				URL: fmt.Sprintf("/api/v1/front/stickers/image/%d", s.ID),
			})
		}

		result = append(result, MyStickerSet{
			PackID:   sr.PackID,
			PackName: sr.PackName,
			AddedAt:  sr.AddedAt,
			Stickers: stickers,
		})
	}

	return &MyStickerSetsResponse{Sets: result}, nil
}

// Delete = ដក pack ចេញពី collection របស់ user (មិនលុប pack/stickers ដើម)
func (r *MysetStickerRepoImpl) Delete(packID int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	result, err := r.dbpool.Exec(`
		DELETE FROM user_stickers WHERE pack_id = $1 AND user_id = $2
	`, packID, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("set_not_found", fmt.Errorf("pack %d not in user %d's collection", packID, userID))
	}
	return nil
}

// Add pack ចូល collection (ប្រើពេលចុច "ADD" លើ pack-card)
func (r *MysetStickerRepoImpl) Create(packID int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	_, err := r.dbpool.Exec(`
		INSERT INTO user_stickers (user_id, pack_id, added_at)
		VALUES ($1, $2, NOW())
		ON CONFLICT (user_id, pack_id) DO NOTHING
	`, userID, packID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

// Serve raw image bytes ពី bytea
func (r *MysetStickerRepoImpl) GetStickerImage(id int64) ([]byte, string, error) {
	var fileData []byte
	var fileType string
	err := r.dbpool.QueryRow(`SELECT file_data, file_type FROM stickers WHERE id = $1`, id).Scan(&fileData, &fileType)
	return fileData, fileType, err
}

func (r *MysetStickerRepoImpl) ListPacks() (*StickerPacksResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	type packRow struct {
		ID           int64  `db:"id"`
		Name         string `db:"pack_name"`
		StickerCount int    `db:"sticker_count"`
	}
	var rows []packRow

	err := r.dbpool.Select(&rows, `
		SELECT sp.id, sp.pack_name,
		       COUNT(s.id) AS sticker_count
		FROM sticker_packs sp
		LEFT JOIN stickers s ON s.pack_id = sp.id
		GROUP BY sp.id, sp.pack_name
		ORDER BY sp.id
	`)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	packs := make([]StickerPack, 0, len(rows))
	for _, r2 := range rows {
		var thumb string
		if r2.StickerCount > 0 {
			var firstID int64
			err := r.dbpool.Get(&firstID, `SELECT MIN(id) FROM stickers WHERE pack_id = $1`, r2.ID)
			if err == nil {
				thumb = fmt.Sprintf("/api/v1/front/stickers/image/%d", firstID)
			}
		}
		packs = append(packs, StickerPack{
			ID:           r2.ID,
			Name:         r2.Name,
			StickerCount: r2.StickerCount,
			ThumbnailURL: thumb,
		})
	}

	return &StickerPacksResponse{Packs: packs}, nil
}

func (r *MysetStickerRepoImpl) ListStickersByPack(packID int64) (*StickerListResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	type row struct {
		ID int64 `db:"id"`
	}
	var rows []row
	err := r.dbpool.Select(&rows, `SELECT id FROM stickers WHERE pack_id = $1 ORDER BY id ASC`, packID)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	stickers := make([]StickerInPack, 0, len(rows))
	for _, rr := range rows {
		stickers = append(stickers, StickerInPack{ID: rr.ID, URL: fmt.Sprintf("/api/v1/front/stickers/image/%d", rr.ID)})
	}
	return &StickerListResponse{PackID: packID, Stickers: stickers}, nil
}

func (r *MysetStickerRepoImpl) CreateSticker(packID int64, fileData []byte, fileType string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	_, err := r.dbpool.Exec(`
		INSERT INTO stickers (pack_id, file_data, file_type) VALUES ($1, $2, $3)
	`, packID, fileData, fileType)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *MysetStickerRepoImpl) CreatePack(name string, userID int64) (int64, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var packID int64
	err := r.dbpool.Get(&packID, `
		INSERT INTO sticker_packs (pack_name, created_by, created_at)
		VALUES ($1, $2, NOW())
		RETURNING id
	`, name, userID)
	if err != nil {
		return 0, msg.NewErrorResponse("database_error", err)
	}
	return packID, nil
}
