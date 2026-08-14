package notification

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"kaifin_clone_api/internal/admin/websocket"
	error_responses "kaifin_clone_api/pkg/responses"
)

type NotificationsRepo interface {
	GetNotifications(userID int64) ([]Notification, int, *error_responses.ErrorResponse)
	MarkAsRead(notificationID, userID int64) *error_responses.ErrorResponse
	CreateNotification(userID, actorID int64, nType string, refID *int64) *error_responses.ErrorResponse
}

type NotificationsRepoImpl struct {
	dbpool *sqlx.DB
	Ws     *websocket.WebSocketManager
}

func NewNotificationsRepoImpl(db *sqlx.DB, ws *websocket.WebSocketManager) *NotificationsRepoImpl {
	return &NotificationsRepoImpl{
		dbpool: db,
		Ws:     ws,
	}
}

func (r *NotificationsRepoImpl) GetNotifications(userID int64) ([]Notification, int, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	query := `
        SELECT 
            n.id, n.user_id, n.actor_id, n.type, n.reference_id, n.is_read, n.created_at,
            COALESCE(u.user_name, '') AS actor_name,
            COALESCE(u.profile_images, '') AS actor_avatar
        FROM notifications n
        LEFT JOIN tbl_users u ON u.id = n.actor_id
        WHERE n.user_id = $1
        ORDER BY n.created_at DESC
        LIMIT 30
    `
	var notifications []Notification
	if err := r.dbpool.Select(&notifications, query, userID); err != nil {
		return nil, 0, msg.NewErrorResponse("database_error", err)
	}

	var unreadCount int
	countQuery := `SELECT COUNT(*) FROM notifications WHERE user_id = $1 AND is_read = false`
	_ = r.dbpool.Get(&unreadCount, countQuery, userID)

	return notifications, unreadCount, nil
}

func (r *NotificationsRepoImpl) MarkAsRead(notificationID, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	_, err := r.dbpool.Exec(`
        UPDATE notifications 
        SET is_read = true 
        WHERE id = $1 AND user_id = $2
    `, notificationID, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *NotificationsRepoImpl) CreateNotification(userID, actorID int64, nType string, refID *int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	query := `
        INSERT INTO notifications (user_id, actor_id, type, reference_id, is_read, created_at)
        VALUES ($1, $2, $3, $4, false, NOW())
        RETURNING id, created_at
    `

	var notifID int64
	var createdAt string
	err := r.dbpool.QueryRow(query, userID, actorID, nType, refID).Scan(&notifID, &createdAt)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	if r.Ws != nil {
		var actorName, actorAvatar string
		_ = r.dbpool.Get(&actorName, `SELECT COALESCE(user_name,'') FROM tbl_users WHERE id = $1`, actorID)
		_ = r.dbpool.Get(&actorAvatar, `SELECT COALESCE(profile_images,'') FROM tbl_users WHERE id = $1`, actorID)

		payload := map[string]interface{}{
			"type":  "notification",
			"event": nType,
			"data": map[string]interface{}{
				"id":           notifID,
				"user_id":      userID,
				"actor_id":     actorID,
				"actor_name":   actorName,
				"actor_avatar": actorAvatar,
				"type":         nType,
				"reference_id": refID,
				"is_read":      false,
				"created_at":   createdAt,
			},
		}

		r.Ws.NotifyUser(fmt.Sprintf("%d", userID), payload)
	}

	return nil
}
