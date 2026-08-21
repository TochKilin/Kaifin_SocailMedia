package course

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type CourseService interface {
	Create(req CreateCourseRequest) (*CourseResponse, *error_responses.ErrorResponse)
	Show(req ShowCourseRequest) (*CourseListResponse, *error_responses.ErrorResponse)
	ShowByID(id int64) (*CourseFullDetailResponse, *error_responses.ErrorResponse)
	Update(id int64, req UpdateCourseRequest) (*CourseResponse, *error_responses.ErrorResponse)
	Delete(id int64) *error_responses.ErrorResponse
}

type CourseServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	Repo    *CourseRepoImpl
}

func NewCourseServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *CourseServiceImpl {
	return &CourseServiceImpl{
		dbpool: dbpool,
		Redis:  rdb,
		Repo:   NewCourseRepoImpl(dbpool),
	}
}

func (s *CourseServiceImpl) Create(req CreateCourseRequest) (*CourseResponse, *error_responses.ErrorResponse) {
	return s.Repo.Create(s.UserCtx.UserID, req)
}

func (s *CourseServiceImpl) Show(req ShowCourseRequest) (*CourseListResponse, *error_responses.ErrorResponse) {
	return s.Repo.Show(req)
}

func (s *CourseServiceImpl) Update(id int64, req UpdateCourseRequest) (*CourseResponse, *error_responses.ErrorResponse) {
	return s.Repo.Update(id, s.UserCtx.UserID, req)
}

func (s *CourseServiceImpl) Delete(id int64) *error_responses.ErrorResponse {
	return s.Repo.Delete(id, s.UserCtx.UserID)
}

func (s *CourseServiceImpl) ShowByID(id int64) (*CourseFullDetailResponse, *error_responses.ErrorResponse) {
	return s.Repo.GetFullDetail(id)
}
