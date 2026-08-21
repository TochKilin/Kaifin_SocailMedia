package streaklevel

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type LevelService interface {
	ListLevels() ([]LevelInfo, *error_responses.ErrorResponse)
	GetStatus(userID int64) (*UserLevelStatus, *error_responses.ErrorResponse)
	CheckIn(userID int64) (*UserLevelStatus, *error_responses.ErrorResponse)
	GetLeaderboard(limit int) ([]LeaderboardUser, *error_responses.ErrorResponse)
}

type LevelServiceImpl struct {
	repo LevelRepo
}

func NewLevelServiceImpl(dbpool *sqlx.DB) *LevelServiceImpl {
	return &LevelServiceImpl{repo: NewLevelRepoImpl(dbpool)}
}

func (s *LevelServiceImpl) ListLevels() ([]LevelInfo, *error_responses.ErrorResponse) {
	return s.repo.ListLevels()
}

func (s *LevelServiceImpl) GetStatus(userID int64) (*UserLevelStatus, *error_responses.ErrorResponse) {
	return s.repo.GetStatus(userID)
}

func (s *LevelServiceImpl) CheckIn(userID int64) (*UserLevelStatus, *error_responses.ErrorResponse) {
	return s.repo.CheckIn(userID)
}

func (s *LevelServiceImpl) GetLeaderboard(limit int) ([]LeaderboardUser, *error_responses.ErrorResponse) {
	return s.repo.GetLeaderboard(limit)
}
