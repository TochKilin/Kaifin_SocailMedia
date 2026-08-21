package addcard

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CartServiceImpl struct {
	Repo *CartRepoImpl
}

func NewCartServiceImpl(dbpool *sqlx.DB) *CartServiceImpl {
	return &CartServiceImpl{Repo: NewCartRepoImpl(dbpool)}
}

func (s *CartServiceImpl) AddItem(userID, courseID int64) *error_responses.ErrorResponse {
	return s.Repo.AddItem(userID, courseID)
}

func (s *CartServiceImpl) RemoveItem(userID, courseID int64) *error_responses.ErrorResponse {
	return s.Repo.RemoveItem(userID, courseID)
}

func (s *CartServiceImpl) Show(userID int64) (*CartResponse, *error_responses.ErrorResponse) {
	return s.Repo.Show(userID)
}
