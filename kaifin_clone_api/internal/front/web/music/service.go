package music

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type PostMusicService interface {
	Create(req CreatePostRequest) (*PostResponse, *error_responses.ErrorResponse)
	Show(req ShowPostRequest) (*PostListResponse, *error_responses.ErrorResponse)
	Update(id int64, req UpdatePostRequest) (*PostResponse, *error_responses.ErrorResponse)
	Delete(id int64) *error_responses.ErrorResponse
}

type PostMusicServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	Repo    *PostMusicRepoImpl
}

func NewPostServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *PostMusicServiceImpl {
	return &PostMusicServiceImpl{
		dbpool: dbpool,
		Redis:  rdb,
		Repo:   NewPostRepoImpl(dbpool),
	}
}

func (s *PostMusicServiceImpl) Create(req CreatePostRequest) (*PostResponse, *error_responses.ErrorResponse) {
	if e := req.validateBusinessRules(); e != nil {
		return nil, e
	}
	return s.Repo.Create(s.UserCtx.UserID, req)
}

func (s *PostMusicServiceImpl) Show(req ShowPostRequest) (*PostListResponse, *error_responses.ErrorResponse) {
	// Show is allowed without a logged-in user (public feed), so guard
	// against a nil UserCtx instead of assuming it's always set.
	var requesterID int64
	if s.UserCtx != nil {
		requesterID = s.UserCtx.UserID
	}
	return s.Repo.Show(requesterID, req)
}

func (s *PostMusicServiceImpl) Update(id int64, req UpdatePostRequest) (*PostResponse, *error_responses.ErrorResponse) {
	if e := req.validateBusinessRules(); e != nil {
		return nil, e
	}
	return s.Repo.Update(id, s.UserCtx.UserID, req)
}

func (s *PostMusicServiceImpl) Delete(id int64) *error_responses.ErrorResponse {
	return s.Repo.Delete(id, s.UserCtx.UserID)
}
