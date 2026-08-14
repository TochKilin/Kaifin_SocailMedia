package likes_mobile

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type LikesMobileService interface {
	ToggleLike(req *CreateLikeMobileRequest) (bool, *error_responses.ErrorResponse) // returns true=liked, false=unliked
	Show(req ShowLikesMobileRequest) (*LikesMobileResponse, *error_responses.ErrorResponse)
}

type LikesMobileServiceImpl struct {
	dbpool  *sqlx.DB
	Repo    *LikesMobileRepoImpl
	UserCtx *share.UserContext
	Redis   *redis.Client
}

func NewLikesMobileServiceImpl(db *sqlx.DB, rdb *redis.Client) *LikesMobileServiceImpl {
	return &LikesMobileServiceImpl{
		dbpool: db,
		Repo:   NewLikesMobileRepoImpl(db),
		Redis:  rdb,
	}
}

func (s *LikesMobileServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *LikesMobileServiceImpl) Show(req ShowLikesMobileRequest, uctx *share.UserContext) (*LikesMobileResponse, *error_responses.ErrorResponse) {
	var currentUserID int64
	if uctx != nil {
		currentUserID = uctx.UserID
	}
	return s.Repo.Show(req, currentUserID)
}

func (s *LikesMobileServiceImpl) Create(req *CreateLikeMobileRequest, uctx *share.UserContext) (bool, *error_responses.ErrorResponse) {
	return s.Repo.ToggleLike(req, uctx.UserID)
}

func (s *LikesMobileServiceImpl) Delete(postID int64, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.Delete(postID, uctx.UserID)
}
