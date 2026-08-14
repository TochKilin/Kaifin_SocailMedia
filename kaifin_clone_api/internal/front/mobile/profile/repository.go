package profile_mobile

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type ProfileMobileRepo interface {
	GetProfile(userID int64) (*ProfileResponse, *error_responses.ErrorResponse)
	Update(userID int64, filename string) *error_responses.ErrorResponse
}

type ProfileMobileRepoImpl struct {
	db *sqlx.DB
}

func NewProfileRepoImpl(db *sqlx.DB) *ProfileMobileRepoImpl {
	return &ProfileMobileRepoImpl{
		db: db,
	}
}

func (r *ProfileMobileRepoImpl) GetProfile(userID int64) (*ProfileResponse, *error_responses.ErrorResponse) {
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

func (r *ProfileMobileRepoImpl) Update(userID int64, filename string) *error_responses.ErrorResponse {
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
