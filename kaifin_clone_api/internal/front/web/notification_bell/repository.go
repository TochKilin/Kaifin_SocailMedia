package notificationbell

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CommunitiesRepoImpl = struct{} // (placeholder removed below)

type NotificationsRepoImpl struct {
	db *sqlx.DB
}

func NewNotificationsRepoImpl(db *sqlx.DB) *NotificationsRepoImpl {
	return &NotificationsRepoImpl{db: db}
}

func (r *NotificationsRepoImpl) Show(userID int64) (*ShowNotificationsResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var list []Notification
	query := `
		SELECT
			n.id, n.user_id, n.actor_id,
			COALESCE(NULLIF(TRIM(CONCAT(u.first_name, ' ', u.last_name)), ''), u.user_name) AS actor_name,
			u.profile_images AS actor_avatar,
			n.type, n.post_id, n.is_read, n.created_at
		FROM notifications_bell n
		JOIN tbl_users u ON u.id = n.actor_id
		WHERE n.user_id = $1
		ORDER BY n.created_at DESC
		LIMIT 50
	`
	if err := r.db.Select(&list, query, userID); err != nil {
		return nil, msg.NewErrorResponse("generic_error", err)
	}

	var unread int64
	if err := r.db.Get(&unread, `SELECT COUNT(*) FROM notifications_bell WHERE user_id = $1 AND is_read = FALSE`, userID); err != nil {
		return nil, msg.NewErrorResponse("generic_error", err)
	}

	return &ShowNotificationsResponse{Notifications: list, UnreadCount: unread}, nil
}

func (r *NotificationsRepoImpl) MarkAsRead(notificationID, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	_, err := r.db.Exec(
		`UPDATE notifications_bell SET is_read = TRUE WHERE id = $1 AND user_id = $2`,
		notificationID, userID,
	)
	if err != nil {
		return msg.NewErrorResponse("generic_error", err)
	}
	return nil
}

// Create one notification row, returns the inserted row (need actor info joined back)
func (r *NotificationsRepoImpl) CreateForNewPost(actorID, postID int64) ([]Notification, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	// បញ្ចូល notification សម្រាប់ follower ទាំងអស់នៃ actor (ម្ចាស់ post)
	// សន្មតតារាង followers(follower_id, following_id) ពី package `follower`
	insertQuery := `
		INSERT INTO notifications_bell (user_id, actor_id, type, post_id)
		SELECT follower_id, $1, 'new_post', $2
		FROM followers
		WHERE following_id = $1
		RETURNING id, user_id, actor_id, type, post_id, is_read, created_at
	`
	var inserted []Notification
	if err := r.db.Select(&inserted, insertQuery, actorID, postID); err != nil {
		return nil, msg.NewErrorResponse("generic_error", err)
	}

	if len(inserted) == 0 {
		return inserted, nil
	}

	// ទាញ actor name/avatar ដើម្បីភ្ជាប់ចូល row ដែល insert រួច (សម្រាប់ផ្ញើតាម WS)
	var actorName string
	var actorAvatar *string
	err := r.db.QueryRow(
		`SELECT COALESCE(NULLIF(TRIM(CONCAT(first_name, ' ', last_name)), ''), user_name), profile_images FROM tbl_users WHERE id = $1`,
		actorID,
	).Scan(&actorName, &actorAvatar)
	if err != nil {
		return nil, msg.NewErrorResponse("generic_error", err)
	}

	for i := range inserted {
		inserted[i].ActorName = actorName
		inserted[i].ActorAvatar = actorAvatar
	}

	return inserted, nil
}
