package follower

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type FollowersRepo interface {
	ToggleFollow(followerID, followingID int64) (bool, *error_responses.ErrorResponse)
	IsFollowing(followerID, followingID int64) (bool, *error_responses.ErrorResponse)
	CountFollowers(userID int64) (int64, *error_responses.ErrorResponse)
	CountFollowing(userID int64) (int64, *error_responses.ErrorResponse)
}

type FollowersRepoImpl struct {
	dbpool *sqlx.DB
}

func NewFollowersRepoImpl(db *sqlx.DB) *FollowersRepoImpl {
	return &FollowersRepoImpl{
		dbpool: db,
	}
}

func (r *FollowersRepoImpl) ToggleFollow(followerID, followingID int64) (bool, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	if followerID == followingID {
		return false, msg.NewErrorResponse("cannot_follow_self", fmt.Errorf("user %d cannot follow themselves", followerID))
	}

	tx, err := r.dbpool.Beginx()
	if err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()
	var exists bool
	queryCheck := `SELECT EXISTS(SELECT 1 FROM tbl_follows WHERE follower_id = $1 AND following_id = $2)`
	if err := tx.Get(&exists, queryCheck, followerID, followingID); err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}

	var isFollowing bool
	if exists {
		queryDel := `DELETE FROM tbl_follows WHERE follower_id = $1 AND following_id = $2`
		if _, err := tx.Exec(queryDel, followerID, followingID); err != nil {
			return false, msg.NewErrorResponse("database_error", err)
		}
		isFollowing = false
	} else {
		queryIns := `INSERT INTO tbl_follows (follower_id, following_id, created_at) VALUES ($1, $2, NOW())`
		if _, err := tx.Exec(queryIns, followerID, followingID); err != nil {
			return false, msg.NewErrorResponse("database_error", err)
		}
		isFollowing = true
	}

	if err := tx.Commit(); err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}

	return isFollowing, nil
}

func (r *FollowersRepoImpl) IsFollowing(followerID, followingID int64) (bool, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM tbl_follows WHERE follower_id = $1 AND following_id = $2)`
	if err := r.dbpool.Get(&exists, query, followerID, followingID); err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}

	return exists, nil
}

func (r *FollowersRepoImpl) CountFollowers(userID int64) (int64, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var count int64
	query := `SELECT COUNT(*) FROM tbl_follows WHERE following_id = $1`
	if err := r.dbpool.Get(&count, query, userID); err != nil {
		return 0, msg.NewErrorResponse("database_error", err)
	}

	return count, nil
}

func (r *FollowersRepoImpl) CountFollowing(userID int64) (int64, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var count int64

	query := `SELECT COUNT(*) FROM tbl_follows WHERE follower_id = $1`
	if err := r.dbpool.Get(&count, query, userID); err != nil {
		return 0, msg.NewErrorResponse("database_error", err)
	}

	return count, nil
}
