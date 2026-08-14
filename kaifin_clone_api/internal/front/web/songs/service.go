package songs

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type SongService interface {
	Create(req CreateSongRequest) (*SongResponse, *error_responses.ErrorResponse)
	Show(req ShowSongRequest) (*SongListResponse, *error_responses.ErrorResponse)
	Update(id int64, req UpdateSongRequest) (*SongResponse, *error_responses.ErrorResponse)
	Delete(id int64) *error_responses.ErrorResponse
}

type SongServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	Repo    *SongRepoImpl
}

func NewSongServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *SongServiceImpl {
	return &SongServiceImpl{
		dbpool: dbpool,
		Redis:  rdb,
		Repo:   NewSongRepoImpl(dbpool),
	}
}

func (s *SongServiceImpl) Create(req CreateSongRequest) (*SongResponse, *error_responses.ErrorResponse) {
	return s.Repo.Create(s.UserCtx.UserID, req)
}

func (s *SongServiceImpl) Show(req ShowSongRequest) (*SongListResponse, *error_responses.ErrorResponse) {
	return s.Repo.Show(req)
}

func (s *SongServiceImpl) Update(id int64, req UpdateSongRequest) (*SongResponse, *error_responses.ErrorResponse) {
	return s.Repo.Update(id, s.UserCtx.UserID, req)
}

func (s *SongServiceImpl) Delete(id int64) *error_responses.ErrorResponse {
	return s.Repo.Delete(id, s.UserCtx.UserID)
}
