package likes

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type LikesService interface {
	ToggleLike(req *CreateLikeRequest) (bool, *error_responses.ErrorResponse) // returns true=liked, false=unliked
	Show(req ShowLikesRequest) (*LikesResponse, *error_responses.ErrorResponse)
}

type LikesServiceImpl struct {
	dbpool  *sqlx.DB
	Repo    *LikesRepoImpl
	UserCtx *share.UserContext
	Redis   *redis.Client
}

func NewLikesServiceImpl(db *sqlx.DB, rdb *redis.Client) *LikesServiceImpl {
	return &LikesServiceImpl{
		dbpool: db,
		Repo:   NewLikesRepoImpl(db),
		Redis:  rdb,
	}
}

func (s *LikesServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *LikesServiceImpl) Show(req ShowLikesRequest, uctx *share.UserContext) (*LikesResponse, *error_responses.ErrorResponse) {
	var currentUserID int64
	if uctx != nil {
		currentUserID = uctx.UserID
	}
	return s.Repo.Show(req, currentUserID)
}

func (s *LikesServiceImpl) Create(req *CreateLikeRequest, uctx *share.UserContext) (bool, *error_responses.ErrorResponse) {
	return s.Repo.ToggleLike(req, uctx.UserID)
}

func (s *LikesServiceImpl) Delete(postID int64, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.Delete(postID, uctx.UserID)
}
