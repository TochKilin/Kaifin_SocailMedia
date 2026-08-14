package communities

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CommunitiesService interface {
	Show(req ShowCommunitiesRequest) (*CommunitiesResponse, *error_responses.ErrorResponse)
}

type CommunitiesServiceImpl struct {
	dbpool *sqlx.DB
	Repo   *CommunitiesRepoImpl
}

func NewCommunitiesServiceImpl(dbpool *sqlx.DB) *CommunitiesServiceImpl {
	return &CommunitiesServiceImpl{
		dbpool: dbpool,
		Repo:   NewCommunitiesRepoImpl(dbpool),
	}
}

func (s *CommunitiesServiceImpl) Show(req ShowCommunitiesRequest) (*CommunitiesResponse, *error_responses.ErrorResponse) {
	return s.Repo.Show(req)
}
