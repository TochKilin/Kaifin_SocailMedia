package menu

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type MenuRepo interface {
	Show(req ShowMenuRequest) (*MenuResponse, *error_responses.ErrorResponse)
}

type MenuRepoImpl struct {
	dbpool *sqlx.DB
}

func NewMenuRepoImpl(db *sqlx.DB) *MenuRepoImpl {
	return &MenuRepoImpl{
		dbpool: db,
	}
}

func (r *MenuRepoImpl) Show(req ShowMenuRequest) (*MenuResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var menus []Menu
	query := fmt.Sprintf(
		`SELECT
			m.id,
			m.title,
			m.sort_order,
			m.is_active,
			m.created_at
		FROM tbl_menus m
		WHERE 1=1
		ORDER BY m.sort_order ASC`,
	)

	err := r.dbpool.Select(&menus, query)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &MenuResponse{
		Menus: menus,
		Total: len(menus),
	}, nil
}
