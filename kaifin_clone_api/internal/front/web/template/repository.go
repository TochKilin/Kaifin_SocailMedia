package template

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type TemplateRepo interface {
	List() (*TemplatesResponse, *error_responses.ErrorResponse)
	GetImage(id int64) ([]byte, string, error)
	Create(name string, fileData []byte, fileType string) *error_responses.ErrorResponse
}

type TemplateRepoImpl struct {
	dbpool *sqlx.DB
}

func NewTemplateRepoImpl(db *sqlx.DB) *TemplateRepoImpl {
	return &TemplateRepoImpl{
		dbpool: db,
	}
}

func (r *TemplateRepoImpl) List() (*TemplatesResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	type row struct {
		ID        int64       `db:"id"`
		Name      string      `db:"name"`
		CreatedAt interface{} `db:"created_at"`
	}
	var rows []row

	err := r.dbpool.Select(&rows, `
		SELECT id, name, created_at FROM tbl_post_template
		WHERE is_active = TRUE
		ORDER BY sort_order ASC, id ASC
	`)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	templates := make([]Template, 0, len(rows))
	for _, r2 := range rows {
		templates = append(templates, Template{
			ID:           r2.ID,
			Name:         r2.Name,
			ThumbnailURL: fmt.Sprintf("/api/v1/front/templates/image/%d", r2.ID),
		})
	}

	return &TemplatesResponse{Templates: templates}, nil
}

func (r *TemplateRepoImpl) GetImage(id int64) ([]byte, string, error) {
	var fileData []byte
	var fileType string
	err := r.dbpool.QueryRow(`
		SELECT file_data, file_type FROM tbl_post_template
		WHERE id = $1 AND is_active = TRUE
	`, id).Scan(&fileData, &fileType)
	return fileData, fileType, err
}

func (r *TemplateRepoImpl) Create(name string, fileData []byte, fileType string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	_, err := r.dbpool.Exec(`
		INSERT INTO tbl_post_template (name, file_data, file_type)
		VALUES ($1, $2, $3)
	`, name, fileData, fileType)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}
