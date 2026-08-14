package follower_mobile

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type FollowersMobileService interface {
	SetUserCtx(ctx *share.UserContext) bool
	Toggle(req ToggleFollowRequest) (*FollowStatusResponse, *error_responses.ErrorResponse)
	Show(req FollowShowRequest) (*FollowStatusResponse, *error_responses.ErrorResponse)
}

type FollowersMobileServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Repo    *FollowersMobileRepoImpl
}

func NewFollowersMobileServiceImpl(dbpool *sqlx.DB) *FollowersMobileServiceImpl {
	return &FollowersMobileServiceImpl{
		dbpool: dbpool,
		Repo:   NewFollowersMobileRepoImpl(dbpool),
	}
}

func (s *FollowersMobileServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *FollowersMobileServiceImpl) currentUserID() int64 {
	if s.UserCtx != nil {
		return s.UserCtx.UserID
	}
	return 0
}

func (s *FollowersMobileServiceImpl) Toggle(req ToggleFollowRequest) (*FollowStatusResponse, *error_responses.ErrorResponse) {
	followerID := s.currentUserID()

	isFollowing, err := s.Repo.ToggleFollow(followerID, req.UserID)
	if err != nil {
		return nil, err
	}

	return s.new(req.UserID, isFollowing)
}

func (s *FollowersMobileServiceImpl) Show(req FollowShowRequest) (*FollowStatusResponse, *error_responses.ErrorResponse) {
	followerID := s.currentUserID()

	isFollowing, err := s.Repo.IsFollowing(followerID, req.UserID)
	if err != nil {
		return nil, err
	}

	return s.new(req.UserID, isFollowing)
}
