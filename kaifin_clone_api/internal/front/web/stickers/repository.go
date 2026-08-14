package stickers

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type StickersRepo interface {
	Show(req ShowStickersRequest) (*StickersResponse, *error_responses.ErrorResponse)
	Create(sticker *Sticker) *error_responses.ErrorResponse
	Update(sticker *Sticker, updateFile bool) *error_responses.ErrorResponse
}

type StickersRepoImpl struct {
	dbpool *sqlx.DB
}

func NewStickersRepoImp(db *sqlx.DB) *StickersRepoImpl {
	return &StickersRepoImpl{
		dbpool: db,
	}
}

func (r *StickersRepoImpl) Show(req ShowStickersRequest) (*StickersResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var list []Sticker
	var err error

	if len(req.IDs) > 0 {
		query, args, inErr := sqlx.In(`
			SELECT id, pack_id, file_name, file_type, file_data, trigger_code, created_at
			FROM stickers
			WHERE id IN (?)
			ORDER BY id ASC
		`, req.IDs)
		if inErr != nil {
			return nil, msg.NewErrorResponse("database_error", inErr)
		}
		query = r.dbpool.Rebind(query)
		err = r.dbpool.Select(&list, query, args...)
	} else {
		err = r.dbpool.Select(&list, `
			SELECT id, pack_id, file_name, file_type, file_data, trigger_code, created_at
			FROM stickers
			WHERE pack_id = $1
			ORDER BY id ASC
		`, req.PackID)
	}

	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	for i := range list {
		list[i].FormatBase64URL()
	}

	return &StickersResponse{
		Stickers: list,
		Total:    len(list),
	}, nil
}

func (r *StickersRepoImpl) Create(sticker *Sticker) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	err := r.dbpool.QueryRow(`
		INSERT INTO stickers (pack_id, file_name, file_type, file_data, trigger_code, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
		RETURNING id, created_at
	`,
		sticker.PackID,
		sticker.FileName,
		sticker.FileType,
		sticker.FileData,
		sticker.TriggerCode,
	).Scan(&sticker.ID, &sticker.CreatedAt)

	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *StickersRepoImpl) Update(sticker *Sticker, updateFile bool) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	var err error
	if updateFile {
		_, err = r.dbpool.Exec(`
			UPDATE stickers 
			SET pack_id = $1, file_name = $2, file_type = $3, file_data = $4, trigger_code = $5
			WHERE id = $6
		`, sticker.PackID, sticker.FileName, sticker.FileType, sticker.FileData, sticker.TriggerCode, sticker.ID)
	} else {
		_, err = r.dbpool.Exec(`
			UPDATE stickers 
			SET pack_id = $1, trigger_code = $2
			WHERE id = $3
		`, sticker.PackID, sticker.TriggerCode, sticker.ID)
	}

	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}
