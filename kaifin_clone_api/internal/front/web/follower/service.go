package follower

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type FollowersService interface {
	SetUserCtx(ctx *share.UserContext) bool
	Toggle(req ToggleFollowRequest) (*FollowStatusResponse, *error_responses.ErrorResponse)
	Show(req FollowShowRequest) (*FollowStatusResponse, *error_responses.ErrorResponse)
}

type FollowersServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Repo    *FollowersRepoImpl
}

func NewFollowersServiceImpl(dbpool *sqlx.DB) *FollowersServiceImpl {
	return &FollowersServiceImpl{
		dbpool: dbpool,
		Repo:   NewFollowersRepoImpl(dbpool),
	}
}

func (s *FollowersServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *FollowersServiceImpl) currentUserID() int64 {
	if s.UserCtx != nil {
		return s.UserCtx.UserID
	}
	return 0
}

func (s *FollowersServiceImpl) Toggle(req ToggleFollowRequest) (*FollowStatusResponse, *error_responses.ErrorResponse) {
	followerID := s.currentUserID()

	isFollowing, err := s.Repo.ToggleFollow(followerID, req.UserID)
	if err != nil {
		return nil, err
	}

	return s.new(req.UserID, isFollowing)
}

func (s *FollowersServiceImpl) Show(req FollowShowRequest) (*FollowStatusResponse, *error_responses.ErrorResponse) {
	followerID := s.currentUserID()

	isFollowing, err := s.Repo.IsFollowing(followerID, req.UserID)
	if err != nil {
		return nil, err
	}

	return s.new(req.UserID, isFollowing)
}
