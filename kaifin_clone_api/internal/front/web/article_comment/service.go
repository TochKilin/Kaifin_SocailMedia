package articlecomment

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CommentsService interface {
	Create(comment *ArticleComment) *error_responses.ErrorResponse
	Update(id int64, userID int64, text string) *error_responses.ErrorResponse
	Delete(id int64, userID int64) *error_responses.ErrorResponse
	Show(req ShowCommentsRequest) (*CommentsResponse, *error_responses.ErrorResponse)
}

type CommentsServiceImpl struct {
	repo CommentsRepo
}

func NewCommentsServiceImpl(dbpool *sqlx.DB) *CommentsServiceImpl {
	return &CommentsServiceImpl{repo: NewCommentsRepoImpl(dbpool)}
}

func (s *CommentsServiceImpl) Create(comment *ArticleComment) *error_responses.ErrorResponse {
	return s.repo.Create(comment)
}

func (s *CommentsServiceImpl) Update(id int64, userID int64, text string) *error_responses.ErrorResponse {
	return s.repo.Update(id, userID, text)
}

func (s *CommentsServiceImpl) Delete(id int64, userID int64) *error_responses.ErrorResponse {
	return s.repo.Delete(id, userID)
}

func (s *CommentsServiceImpl) Show(req ShowCommentsRequest) (*CommentsResponse, *error_responses.ErrorResponse) {
	return s.repo.Show(req)
}
