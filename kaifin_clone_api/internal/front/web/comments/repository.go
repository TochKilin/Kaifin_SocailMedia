package comments

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CommentsRepo interface {
	Create(comment *Comment, imagePaths []string) *error_responses.ErrorResponse
	Delete(id int64, userID int64) *error_responses.ErrorResponse
	Show(req ShowCommentsRequest, userID int64) (*CommentsResponse, *error_responses.ErrorResponse)
	ToggleLike(commentID, userID int64) (bool, *error_responses.ErrorResponse)
}

type CommentsRepoImpl struct {
	dbpool *sqlx.DB
}

func NewCommentsRepoImpl(db *sqlx.DB) *CommentsRepoImpl {
	return &CommentsRepoImpl{dbpool: db}
}

func (r *CommentsRepoImpl) Create(comment *Comment, imagePaths []string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	tx, err := r.dbpool.Beginx()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	err = tx.QueryRow(`
		INSERT INTO tbl_comments (post_id, user_id, parent_comment_id, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`, comment.PostID, comment.UserID, comment.ParentID, comment.Content,
	).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	for _, path := range imagePaths {
		if _, err = tx.Exec(`
			INSERT INTO tbl_comment_images (comment_id, image_path, created_at)
			VALUES ($1, $2, NOW())
		`, comment.ID, path); err != nil {
			return msg.NewErrorResponse("database_error", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *CommentsRepoImpl) Delete(id int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	result, err := r.dbpool.Exec(`
		DELETE FROM tbl_comments WHERE id = $1 AND user_id = $2
	`, id, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("comment_not_found", nil)
	}
	return nil
}

func (r *CommentsRepoImpl) Show(req ShowCommentsRequest, userID int64) (*CommentsResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var comments []Comment
	err := r.dbpool.Select(&comments, `
		SELECT
			c.id, c.post_id, c.user_id, u.user_name AS user_name,
			u.profile_images AS profile_images, c.parent_comment_id,
			c.content, c.created_at, c.updated_at,
			COALESCE(lc.cnt, 0) AS like_count,
			COALESCE(ul.liked, false) AS liked
		FROM tbl_comments c
		LEFT JOIN tbl_users u ON u.id = c.user_id
		LEFT JOIN (
			SELECT comment_id, COUNT(*) AS cnt
			FROM tbl_comment_likes
			GROUP BY comment_id
		) lc ON lc.comment_id = c.id
		LEFT JOIN (
			SELECT comment_id, true AS liked
			FROM tbl_comment_likes
			WHERE user_id = $2
		) ul ON ul.comment_id = c.id
		WHERE c.post_id = $1
		ORDER BY c.created_at ASC
	`, req.PostID, userID)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	if len(comments) > 0 {
		ids := make([]int64, len(comments))
		idx := map[int64]int{}
		for i, c := range comments {
			ids[i] = c.ID
			idx[c.ID] = i
		}

		var images []CommentImage
		q, args, _ := sqlx.In(`
			SELECT id, comment_id, image_path, created_at
			FROM tbl_comment_images
			WHERE comment_id IN (?)
			ORDER BY id ASC
		`, ids)
		q = r.dbpool.Rebind(q)
		if err := r.dbpool.Select(&images, q, args...); err != nil {
			return nil, msg.NewErrorResponse("database_error", err)
		}
		for _, img := range images {
			comments[idx[img.CommentID]].Images = append(comments[idx[img.CommentID]].Images, img)
		}
	}

	return &CommentsResponse{Comments: comments, Total: len(comments)}, nil
}

func (r *CommentsRepoImpl) ToggleLike(commentID, userID int64) (bool, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var exists bool
	err := r.dbpool.Get(&exists, `
		SELECT EXISTS(SELECT 1 FROM tbl_comment_likes WHERE comment_id = $1 AND user_id = $2)
	`, commentID, userID)
	if err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}

	if exists {
		_, err = r.dbpool.Exec(`DELETE FROM tbl_comment_likes WHERE comment_id = $1 AND user_id = $2`, commentID, userID)
		if err != nil {
			return false, msg.NewErrorResponse("database_error", err)
		}
		return false, nil
	}

	_, err = r.dbpool.Exec(`INSERT INTO tbl_comment_likes (comment_id, user_id) VALUES ($1, $2)`, commentID, userID)
	if err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}
	return true, nil
}
