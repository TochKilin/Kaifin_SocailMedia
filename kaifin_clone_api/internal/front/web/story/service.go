package story

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type StoryService interface {
	Create(req *CreateStoryRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	Show(req ShowStoryRequest) (*StoryResponse, *error_responses.ErrorResponse)
	Delete(id int64, userID int64) *error_responses.ErrorResponse
}

type StoryServiceImpl struct {
	repo  StoryRepo
	redis *redis.Client
}

func NewStoryServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *StoryServiceImpl {
	return &StoryServiceImpl{
		repo:  NewStoryRepoImpl(dbpool),
		redis: rdb,
	}
}

func (s *StoryServiceImpl) Create(req *CreateStoryRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.repo.Create(req, uctx)
}

func (s *StoryServiceImpl) Show(req ShowStoryRequest) (*StoryResponse, *error_responses.ErrorResponse) {
	return s.repo.Show(req)
}

func (s *StoryServiceImpl) Delete(id int64, userID int64) *error_responses.ErrorResponse {
	return s.repo.Delete(id, userID)
}
