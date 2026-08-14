package storyreaction_mobile

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type StoryReactionMobileService interface {
	React(req *CreateStoryReactionRequest, uctx *share.UserContext) (*StoryReaction, *error_responses.ErrorResponse)
	Show(req ShowStoryReactionRequest) (*StoryReactionResponse, *error_responses.ErrorResponse)
	Delete(storyID int64, userID int64) *error_responses.ErrorResponse
}

type StoryReactionMobileServiceImpl struct {
	repo StoryReactionMobileRepo
}

func NewStoryReactionServiceImpl(dbpool *sqlx.DB) *StoryReactionMobileServiceImpl {
	return &StoryReactionMobileServiceImpl{
		repo: NewStoryReactionMobileRepoImpl(dbpool),
	}
}

func (s *StoryReactionMobileServiceImpl) React(req *CreateStoryReactionRequest, uctx *share.UserContext) (*StoryReaction, *error_responses.ErrorResponse) {
	return s.repo.Upsert(req, uctx)
}

func (s *StoryReactionMobileServiceImpl) Show(req ShowStoryReactionRequest) (*StoryReactionResponse, *error_responses.ErrorResponse) {
	return s.repo.Show(req)
}

func (s *StoryReactionMobileServiceImpl) Delete(storyID int64, userID int64) *error_responses.ErrorResponse {
	return s.repo.Delete(storyID, userID)
}
