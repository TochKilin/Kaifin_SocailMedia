package storyreaction

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type StoryReactionService interface {
	React(req *CreateStoryReactionRequest, uctx *share.UserContext) (*StoryReaction, *error_responses.ErrorResponse)
	Show(req ShowStoryReactionRequest) (*StoryReactionResponse, *error_responses.ErrorResponse)
	Delete(storyID int64, userID int64) *error_responses.ErrorResponse
}

type StoryReactionServiceImpl struct {
	repo StoryReactionRepo
}

func NewStoryReactionServiceImpl(dbpool *sqlx.DB) *StoryReactionServiceImpl {
	return &StoryReactionServiceImpl{
		repo: NewStoryReactionRepoImpl(dbpool),
	}
}

func (s *StoryReactionServiceImpl) React(req *CreateStoryReactionRequest, uctx *share.UserContext) (*StoryReaction, *error_responses.ErrorResponse) {
	return s.repo.Upsert(req, uctx)
}

func (s *StoryReactionServiceImpl) Show(req ShowStoryReactionRequest) (*StoryReactionResponse, *error_responses.ErrorResponse) {
	return s.repo.Show(req)
}

func (s *StoryReactionServiceImpl) Delete(storyID int64, userID int64) *error_responses.ErrorResponse {
	return s.repo.Delete(storyID, userID)
}
