package bookmark_mobile

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type BookMarkMobileRepo interface {
	ToggleBookmark(userID, postID int64) (bool, *error_responses.ErrorResponse)
	CountByPost(postID int64) (int64, *error_responses.ErrorResponse)
}

type BookMarkMobileRepoImpl struct {
	dbpool *sqlx.DB
}

func NewBookMarkRepoImpl(db *sqlx.DB) *BookMarkMobileRepoImpl {
	return &BookMarkMobileRepoImpl{
		dbpool: db,
	}
}

func (r *BookMarkMobileRepoImpl) ToggleBookmark(userID, postID int64) (bool, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	tx, err := r.dbpool.Beginx()
	if err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	var exists bool
	queryCheck := `SELECT EXISTS(SELECT 1 FROM tbl_post_bookmarks WHERE user_id = $1 AND post_id = $2)`
	if err := tx.Get(&exists, queryCheck, userID, postID); err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}

	var isBookmarked bool
	if exists {
		queryDel := `DELETE FROM tbl_post_bookmarks WHERE user_id = $1 AND post_id = $2`
		if _, err := tx.Exec(queryDel, userID, postID); err != nil {
			return false, msg.NewErrorResponse("database_error", err)
		}
		isBookmarked = false
	} else {
		queryIns := `INSERT INTO tbl_post_bookmarks (user_id, post_id, created_at) VALUES ($1, $2, NOW())`
		if _, err := tx.Exec(queryIns, userID, postID); err != nil {
			return false, msg.NewErrorResponse("database_error", err)
		}
		isBookmarked = true
	}

	if err := tx.Commit(); err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}

	return isBookmarked, nil
}

func (r *BookMarkMobileRepoImpl) CountByPost(postID int64) (int64, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var count int64
	query := `SELECT COUNT(*) FROM tbl_post_bookmarks WHERE post_id = $1`
	err := r.dbpool.Get(&count, query, postID)
	if err != nil {
		return 0, msg.NewErrorResponse("database_error", err)
	}

	return count, nil
}

func (r *BookMarkMobileRepoImpl) Show(userID int64, req ShowBookmarkRequest) (*BookmarkResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var bookmarks []BookMark
	query := `SELECT id, user_id, post_id, created_at 
              FROM tbl_post_bookmarks 
              WHERE user_id = $1 AND post_id = $2 
              ORDER BY created_at DESC`

	err := r.dbpool.Select(&bookmarks, query, userID, req.PostID)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &BookmarkResponse{
		Bookmarks: bookmarks,
		Total:     len(bookmarks),
	}, nil
}
