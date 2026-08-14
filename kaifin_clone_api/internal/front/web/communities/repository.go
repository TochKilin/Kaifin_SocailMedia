package communities

import (
	"strconv"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CommunitiesRepo interface {
	Show(req ShowCommunitiesRequest) (*CommunitiesResponse, *error_responses.ErrorResponse)
}

type CommunitiesRepoImpl struct {
	dbpool *sqlx.DB
}

func NewCommunitiesRepoImpl(db *sqlx.DB) *CommunitiesRepoImpl {
	return &CommunitiesRepoImpl{
		dbpool: db,
	}
}

func (r *CommunitiesRepoImpl) Show(req ShowCommunitiesRequest) (*CommunitiesResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var communities []Communities
	query := `
		SELECT 
			id,
			name,
			description,
			created_at,
			updated_at,
			created_by
		FROM tbl_communities
		WHERE 1=1
	`
	args := []interface{}{}
	index := 1

	if req.Search != "" {
		query += " AND name ILIKE $" + strconv.Itoa(index)
		args = append(args, "%"+req.Search+"%")
		index++
	}

	query += " ORDER BY created_at DESC"

	if err := r.dbpool.Select(&communities, query, args...); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &CommunitiesResponse{
		Communities: communities,
		Total:       len(communities),
	}, nil
}
