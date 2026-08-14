package comments_mobile

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CommentsMobileRepo interface {
	Create(comment *CommentMobile) *error_responses.ErrorResponse
	Delete(id int64, userID int64) *error_responses.ErrorResponse
	Show(req ShowCommentsMobileRequest) (*CommentsMobileResponse, *error_responses.ErrorResponse)
}

type CommentsMobileRepoImpl struct {
	dbpool *sqlx.DB
}

func NewCommentsMobileRepoImpl(db *sqlx.DB) *CommentsMobileRepoImpl {
	return &CommentsMobileRepoImpl{dbpool: db}
}

func (r *CommentsMobileRepoImpl) Create(comment *CommentMobile) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	err := r.dbpool.QueryRow(`
		INSERT INTO tbl_comments (post_id, user_id, parent_comment_id, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`,
		comment.PostID,
		comment.UserID,
		comment.ParentID,
		comment.Content,
	).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)

	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *CommentsMobileRepoImpl) Delete(id int64, userID int64) *error_responses.ErrorResponse {
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

func (r *CommentsMobileRepoImpl) Show(req ShowCommentsMobileRequest) (*CommentsMobileResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var comments []CommentMobile
	err := r.dbpool.Select(&comments, `
		SELECT
			c.id,
			c.post_id,
			c.user_id,
			u.user_name AS user_name,
			u.profile_images AS profile_images,
			c.parent_comment_id,
			c.content,
			c.created_at,
			c.updated_at
		FROM tbl_comments c
		LEFT JOIN tbl_users u ON u.id = c.user_id
		WHERE c.post_id = $1
		ORDER BY c.created_at ASC
	`, req.PostID)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &CommentsMobileResponse{
		Comments: comments,
		Total:    len(comments),
	}, nil
}
