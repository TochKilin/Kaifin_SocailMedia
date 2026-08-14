package bookmark

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type BookMarkService interface {
	Toggle(userID int64, req ToggleBookmarkRequest) (*BookmarStatus, *error_responses.ErrorResponse)
}

type BookMarkServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	Repo    *BookMarkRepoImpl
}

func NewBookMarkServiceImpl(db *sqlx.DB, rdb *redis.Client) *BookMarkServiceImpl {
	return &BookMarkServiceImpl{
		dbpool: db,
		Redis:  rdb,
		Repo:   NewBookMarkRepoImpl(db),
	}
}

func (s *BookMarkServiceImpl) Toggle(userID int64, req ToggleBookmarkRequest) (*BookmarStatus, *error_responses.ErrorResponse) {
	isBookmarked, e := s.Repo.ToggleBookmark(userID, req.PostID)
	if e != nil {
		return nil, e
	}
	total, e := s.Repo.CountByPost(req.PostID)
	if e != nil {
		return nil, e
	}
	return &BookmarStatus{
		BookMarked: isBookmarked,
		Total:      total,
	}, nil
}

func (s *BookMarkServiceImpl) Show(userID int64, req ShowBookmarkRequest) (*BookmarkResponse, *error_responses.ErrorResponse) {
	return s.Repo.Show(userID, req)
}
