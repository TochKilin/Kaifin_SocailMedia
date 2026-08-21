package notification

import (
	"github.com/jmoiron/sqlx"

	"kaifin_clone_api/internal/admin/websocket" // ហៅយក WebSocket Manager
	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type NotificationsServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Repo    *NotificationsRepoImpl
	Ws      *websocket.WebSocketManager
}

func NewNotificationsServiceImpl(dbpool *sqlx.DB, ws *websocket.WebSocketManager) *NotificationsServiceImpl {
	return &NotificationsServiceImpl{
		dbpool: dbpool,
		Repo:   NewNotificationsRepoImpl(dbpool, ws),
		Ws:     ws,
	}
}

func (s *NotificationsServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *NotificationsServiceImpl) GetList() (*NotificationListResponse, *error_responses.ErrorResponse) {
	if s.UserCtx == nil {
		return nil, &error_responses.ErrorResponse{}
	}
	notifs, unread, err := s.Repo.GetNotifications(s.UserCtx.UserID)
	if err != nil {
		return nil, err
	}
	return &NotificationListResponse{
		Notifications: notifs,
		UnreadCount:   unread,
	}, nil
}

func (s *NotificationsServiceImpl) Read(notificationID int64) *error_responses.ErrorResponse {
	if s.UserCtx == nil {
		return &error_responses.ErrorResponse{}
	}
	return s.Repo.MarkAsRead(notificationID, s.UserCtx.UserID)
}
