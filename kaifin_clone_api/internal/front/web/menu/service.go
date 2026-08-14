package menu

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type MenuService interface {
	Show(req ShowMenuRequest) (*MenuResponse, *error_responses.ErrorResponse)
}

type MenuServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	Repo    *MenuRepoImpl
}

func NewMenuServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *MenuServiceImpl {
	return &MenuServiceImpl{
		dbpool: dbpool,
		Redis:  rdb,
		Repo:   NewMenuRepoImpl(dbpool),
	}
}

func (s *MenuServiceImpl) Show(req ShowMenuRequest) (*MenuResponse, *error_responses.ErrorResponse) {
	return s.Repo.Show(req)
}
