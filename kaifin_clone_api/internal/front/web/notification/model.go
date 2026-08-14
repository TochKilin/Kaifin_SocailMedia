package notification

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/utls"
)

type Notification struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	ActorID     int64     `json:"actor_id" db:"actor_id"`
	Type        string    `json:"type" db:"type"`
	ReferenceID *int64    `json:"reference_id" db:"reference_id"`
	IsRead      bool      `json:"is_read" db:"is_read"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	ActorName   string    `json:"actor_name" db:"actor_name"`
	ActorAvatar string    `json:"actor_avatar" db:"actor_avatar"`
}

type NotificationListResponse struct {
	Notifications []Notification `json:"notifications"`
	UnreadCount   int            `json:"unread_count"`
}

type ReadNotificationRequest struct {
	NotificationID int64 `json:"notification_id" validate:"required"`
}

func (r *ReadNotificationRequest) Bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	return v.Validate(r)
}
