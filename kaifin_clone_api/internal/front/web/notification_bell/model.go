package notificationbell

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/utls"
)

type Notification struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	ActorID     int64     `json:"actor_id" db:"actor_id"`
	ActorName   string    `json:"actor_name" db:"actor_name"`
	ActorAvatar *string   `json:"actor_avatar" db:"actor_avatar"`
	Type        string    `json:"type" db:"type"`
	PostID      *int64    `json:"post_id" db:"post_id"`
	IsRead      bool      `json:"is_read" db:"is_read"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type ShowNotificationsResponse struct {
	Notifications []Notification `json:"notifications"`
	UnreadCount   int64          `json:"unread_count"`
}

type MarkAsReadRequest struct {
	NotificationID int64 `json:"notification_id" validate:"required"`
}

func (r *MarkAsReadRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}
