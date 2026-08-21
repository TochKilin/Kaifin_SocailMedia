package notificationbell

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type NotificationsServiceImpl struct {
	UserCtx *share.UserContext
	Repo    *NotificationsRepoImpl
}

func NewNotificationsServiceImpl(db *sqlx.DB) *NotificationsServiceImpl {
	return &NotificationsServiceImpl{Repo: NewNotificationsRepoImpl(db)}
}

func (s *NotificationsServiceImpl) SetUserCtx(ctx *share.UserContext) {
	s.UserCtx = ctx
}

func (s *NotificationsServiceImpl) Show() (*ShowNotificationsResponse, *error_responses.ErrorResponse) {
	return s.Repo.Show(s.UserCtx.UserID)
}

func (s *NotificationsServiceImpl) MarkAsRead(notificationID int64) *error_responses.ErrorResponse {
	return s.Repo.MarkAsRead(notificationID, s.UserCtx.UserID)
}

// ហៅដោយ posts handler ក្រោយ post ត្រូវបានបង្កើត
func (s *NotificationsServiceImpl) NotifyFollowersNewPost(actorID, postID int64) ([]Notification, *error_responses.ErrorResponse) {
	return s.Repo.CreateForNewPost(actorID, postID)
}
