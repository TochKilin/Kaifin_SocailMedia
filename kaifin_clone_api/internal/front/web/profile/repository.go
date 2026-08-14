package profile

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type ProfileRepo interface {
	GetProfile(userID int64) (*ProfileResponse, *error_responses.ErrorResponse)
	Update(userID int64, filename string) *error_responses.ErrorResponse
	UpdateCover(userID int64, filename string) *error_responses.ErrorResponse
	UpdateInfo(userID int64, req *UpdateProfileInfoRequest) *error_responses.ErrorResponse
}

type ProfileRepoImpl struct {
	db *sqlx.DB
}

func NewProfileRepoImpl(db *sqlx.DB) *ProfileRepoImpl {
	return &ProfileRepoImpl{
		db: db,
	}
}

func (r *ProfileRepoImpl) GetProfile(userID int64) (*ProfileResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var profile ProfileResponse
	err := r.db.Get(&profile, `
		SELECT
			u.id,
			u.user_name,
			u.first_name,
			u.last_name,
			u.email,
			u.profile_images,
			u.cover_images, 
			u.location, 
			u.relationship_status, 
			u.bio, 
			u.role_id,
			u.role_name,
			u.last_login,
			u.created_at,

			(
				SELECT COUNT(*)
				FROM tbl_posts p
				WHERE p.user_id = u.id
			) AS post_count,

			(
				SELECT COUNT(*)
				FROM tbl_follows f
				WHERE f.following_id = u.id
			) AS follower_count,

			(
				SELECT COUNT(*)
				FROM tbl_follows f
				WHERE f.follower_id = u.id
			) AS following_count

		FROM tbl_users u
		WHERE u.id = $1
		AND u.deleted_at IS NULL

		LIMIT 1

	`, userID)

	if err != nil {
		return nil, msg.NewErrorResponse(
			"user_not_found",
			fmt.Errorf("profile not found"),
		)
	}

	return &profile, nil
}

func (r *ProfileRepoImpl) Update(userID int64, filename string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	if userID == 0 {
		return msg.NewErrorResponse(
			"invalid_user_id",
			fmt.Errorf("userID is zero — refusing update"),
		)
	}

	res, err := r.db.Exec(
		`UPDATE tbl_users SET profile_images = $1 WHERE id = $2 AND deleted_at IS NULL`,
		filename,
		userID,
	)

	if err != nil {
		return msg.NewErrorResponse(
			"update_failed",
			fmt.Errorf("failed to update profile image: %w", err),
		)
	}

	rows, rowsErr := res.RowsAffected()
	if rowsErr != nil {
		return msg.NewErrorResponse(
			"update_failed",
			fmt.Errorf("failed to get rows affected: %w", rowsErr),
		)
	}

	fmt.Println("UpdateProfileImage — userID:", userID, "rows affected:", rows)

	if rows == 0 {
		return msg.NewErrorResponse(
			"user_not_found",
			fmt.Errorf("no user updated for id %d", userID),
		)
	}

	if rows > 1 {
		return msg.NewErrorResponse(
			"update_error",
			fmt.Errorf("unexpected: %d rows updated for id %d", rows, userID),
		)
	}

	return nil
}

func (r *ProfileRepoImpl) UpdateCover(userID int64, filename string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	if userID == 0 {
		return msg.NewErrorResponse(
			"invalid_user_id",
			fmt.Errorf("userID is zero — refusing update"),
		)
	}

	res, err := r.db.Exec(
		`UPDATE tbl_users SET cover_images = $1 WHERE id = $2 AND deleted_at IS NULL`,
		filename,
		userID,
	)

	if err != nil {
		return msg.NewErrorResponse(
			"update_failed",
			fmt.Errorf("failed to update cover image: %w", err),
		)
	}

	rows, rowsErr := res.RowsAffected()
	if rowsErr != nil {
		return msg.NewErrorResponse(
			"update_failed",
			fmt.Errorf("failed to get rows affected: %w", rowsErr),
		)
	}

	if rows == 0 {
		return msg.NewErrorResponse(
			"user_not_found",
			fmt.Errorf("no user updated for id %d", userID),
		)
	}

	return nil
}

func (r *ProfileRepoImpl) UpdateInfo(userID int64, req *UpdateProfileInfoRequest) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	if userID == 0 {
		return msg.NewErrorResponse(
			"invalid_user_id",
			fmt.Errorf("userID is zero — refusing update"),
		)
	}

	setClauses := []string{}
	args := []interface{}{}
	i := 1

	add := func(column string, value interface{}) {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, i))
		args = append(args, value)
		i++
	}

	if req.FirstName != nil {
		add("first_name", *req.FirstName)
	}
	if req.LastName != nil {
		add("last_name", *req.LastName)
	}
	if req.UserName != nil {
		add("user_name", *req.UserName)
	}
	if req.Bio != nil {
		add("bio", *req.Bio)
	}
	if req.RelationshipStatus != nil {
		add("relationship_status", *req.RelationshipStatus)
	}
	if req.Location != nil {
		add("location", *req.Location)
	}

	if len(setClauses) == 0 {
		return msg.NewErrorResponse("no_fields", fmt.Errorf("no fields to update"))
	}

	query := fmt.Sprintf(
		"UPDATE tbl_users SET %s WHERE id = $%d AND deleted_at IS NULL",
		strings.Join(setClauses, ", "),
		i,
	)
	args = append(args, userID)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		return msg.NewErrorResponse("update_failed", fmt.Errorf("failed to update profile info: %w", err))
	}

	rows, rowsErr := res.RowsAffected()
	if rowsErr != nil {
		return msg.NewErrorResponse("update_failed", fmt.Errorf("failed to get rows affected: %w", rowsErr))
	}

	if rows == 0 {
		return msg.NewErrorResponse("user_not_found", fmt.Errorf("no user updated for id %d", userID))
	}

	return nil
}
