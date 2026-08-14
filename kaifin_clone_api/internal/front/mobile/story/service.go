package story_mobile

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type StoryMobileService interface {
	Create(req *CreateStoryRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	Show(req ShowStoryRequest) (*StoryResponse, *error_responses.ErrorResponse)
	Delete(id int64, userID int64) *error_responses.ErrorResponse
}

type StoryMobileServiceImpl struct {
	repo  StoryMobileRepo
	redis *redis.Client
}

func NewStoryMobileServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *StoryMobileServiceImpl {
	return &StoryMobileServiceImpl{
		repo:  NewStoryMobileRepoImpl(dbpool),
		redis: rdb,
	}
}

func (s *StoryMobileServiceImpl) Create(req *CreateStoryRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.repo.Create(req, uctx)
}

func (s *StoryMobileServiceImpl) Show(req ShowStoryRequest) (*StoryResponse, *error_responses.ErrorResponse) {
	return s.repo.Show(req)
}

func (s *StoryMobileServiceImpl) Delete(id int64, userID int64) *error_responses.ErrorResponse {
	return s.repo.Delete(id, userID)
}
