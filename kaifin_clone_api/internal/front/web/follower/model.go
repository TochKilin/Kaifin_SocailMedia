package follower

import (
	"time"

	"github.com/gofiber/fiber/v3"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/utls"
)

type Follow struct {
	ID          int64     `json:"id" db:"id"`
	FollowerID  int64     `json:"follower_id" db:"follower_id"`
	FollowingID int64     `json:"following_id" db:"following_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type ToggleFollowRequest struct {
	UserID int64 `json:"user_id" validate:"required"`
}

func (r *ToggleFollowRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type FollowShowRequest struct {
	UserID int64 `query:"user_id" validate:"required"`
}

func (r *FollowShowRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type FollowStatusResponse struct {
	IsFollowing    bool  `json:"is_following"`
	FollowerCount  int64 `json:"follower_count"`
	FollowingCount int64 `json:"following_count"`
}

func (s *FollowersServiceImpl) new(userID int64, isFollowing bool) (*FollowStatusResponse, *error_responses.ErrorResponse) {
	followerCount, err := s.Repo.CountFollowers(userID)
	if err != nil {
		return nil, err
	}

	followingCount, err := s.Repo.CountFollowing(userID)
	if err != nil {
		return nil, err
	}

	return &FollowStatusResponse{
		IsFollowing:    isFollowing,
		FollowerCount:  followerCount,
		FollowingCount: followingCount,
	}, nil
}
