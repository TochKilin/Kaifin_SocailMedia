package story_mobile

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type StoryMobileRepo interface {
	Create(req *CreateStoryRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	Show(req ShowStoryRequest) (*StoryResponse, *error_responses.ErrorResponse)
	Delete(id int64, userID int64) *error_responses.ErrorResponse
}

type StoryMobileRepoImpl struct {
	dbpool *sqlx.DB
}

func NewStoryMobileRepoImpl(db *sqlx.DB) *StoryMobileRepoImpl {
	return &StoryMobileRepoImpl{
		dbpool: db,
	}
}

func (r *StoryMobileRepoImpl) Create(req *CreateStoryRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	st := Story{}

	if err := st.new(req, uctx); err != nil {
		return msg.NewErrorResponse("invalid", err)
	}

	err := r.dbpool.QueryRow(`
		INSERT INTO tbl_stories
		(user_id, media_url, media_type, expires_at)
		VALUES ($1,$2,$3, NOW() + INTERVAL '24 hours')
		RETURNING id, created_at, expires_at
	`,
		st.UserID,
		st.MediaURL,
		st.MediaType,
	).Scan(&st.ID, &st.CreatedAt, &st.ExpiresAt)

	if err != nil {
		fmt.Println("INSERT STORY ERROR:", err)
		return msg.NewErrorResponse("database_error", err)
	}

	return nil
}

func (r *StoryMobileRepoImpl) Show(req ShowStoryRequest) (*StoryResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	perPage := req.PageOption.Perpage
	page := req.PageOption.Page
	offset := (page - 1) * perPage
	filter := " AND s.expires_at > NOW()"
	args := []interface{}{}
	argIdx := 1
	if req.UserID != nil {
		filter += fmt.Sprintf(" AND s.user_id = $%d", argIdx)
		args = append(args, *req.UserID)
		argIdx++
	}

	var stories []Story
	query := fmt.Sprintf(`
		SELECT
			s.id,
			s.user_id,
			u.user_name AS user_name,
			u.profile_images AS profile_images,
			s.media_url,
			s.media_type,
			s.created_at,
			s.expires_at
		FROM tbl_stories s
		LEFT JOIN tbl_users u ON u.id = s.user_id
		WHERE 1=1
		%s
		ORDER BY s.created_at DESC
		LIMIT %d OFFSET %d
	`, filter, perPage, offset)

	if err := r.dbpool.Select(&stories, query, args...); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var total int
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM tbl_stories s WHERE 1=1 %s`, filter)
	if err := r.dbpool.Get(&total, countQuery, args...); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &StoryResponse{Stories: stories, Total: total}, nil
}

func (r *StoryMobileRepoImpl) Delete(id int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	result, err := r.dbpool.Exec(`DELETE FROM tbl_stories WHERE id = $1 AND user_id = $2`, id, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("story_not_found", fmt.Errorf("story %d not found", id))
	}
	return nil
}
