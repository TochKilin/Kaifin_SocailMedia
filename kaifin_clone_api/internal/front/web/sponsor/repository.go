package sponsor

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type SponsorsRepo interface {
	Create(s *Sponsor) *error_responses.ErrorResponse
	Update(s *Sponsor) *error_responses.ErrorResponse
	Delete(id int64) *error_responses.ErrorResponse
	Show(req ShowSponsorsRequest) (*SponsorsResponse, *error_responses.ErrorResponse)
}

type SponsorsRepoImpl struct {
	dbpool *sqlx.DB
}

func NewSponsorsRepoImpl(db *sqlx.DB) *SponsorsRepoImpl {
	return &SponsorsRepoImpl{dbpool: db}
}

func (r *SponsorsRepoImpl) Create(s *Sponsor) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	err := r.dbpool.QueryRow(`
		INSERT INTO tbl_sponsors
			(name, logo_image, website_url, is_verified, is_active, sort_order, created_at, updated_at)
		VALUES ($1, $2, $3, $4, true, $5, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`, s.Name, s.LogoImage, s.WebsiteURL, s.IsVerified, s.SortOrder,
	).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *SponsorsRepoImpl) Update(s *Sponsor) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	result, err := r.dbpool.Exec(`
		UPDATE tbl_sponsors
		SET name = $1, logo_image = COALESCE($2, logo_image), website_url = $3,
			is_verified = $4, is_active = $5, sort_order = $6, updated_at = NOW()
		WHERE id = $7
	`, s.Name, s.LogoImage, s.WebsiteURL, s.IsVerified, s.IsActive, s.SortOrder, s.ID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("sponsor_not_found", nil)
	}
	return nil
}

func (r *SponsorsRepoImpl) Delete(id int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	result, err := r.dbpool.Exec(`DELETE FROM tbl_sponsors WHERE id = $1`, id)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("sponsor_not_found", nil)
	}
	return nil
}

func (r *SponsorsRepoImpl) Show(req ShowSponsorsRequest) (*SponsorsResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var total int
	if err := r.dbpool.Get(&total, `SELECT COUNT(*) FROM tbl_sponsors WHERE is_active = true`); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	offset := (req.Page - 1) * req.PerPage
	var sponsors []Sponsor
	if err := r.dbpool.Select(&sponsors, `
		SELECT id, name, logo_image, website_url, is_verified, is_active, sort_order, created_at, updated_at
		FROM tbl_sponsors
		WHERE is_active = true
		ORDER BY sort_order ASC, id ASC
		LIMIT $1 OFFSET $2
	`, req.PerPage, offset); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &SponsorsResponse{Sponsors: sponsors, Total: total}, nil
}
