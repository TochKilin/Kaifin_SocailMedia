package post

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type PostService interface {
	Create(req *CreatePostRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	Show(postRequest ShowPostRequest) (*PostResponse, *error_responses.ErrorResponse)
	Delete(id int64) *error_responses.ErrorResponse
	SetUserCtx(ctx *share.UserContext) bool
	IncrementView(id int64) (int, *error_responses.ErrorResponse)
	CreateShare(postID int64, uctx *share.UserContext) *error_responses.ErrorResponse
	GetShareCount(postID int64) (int, *error_responses.ErrorResponse)
	ShowWithReposts(userID int64, page, perPage int) (*PostResponse, *error_responses.ErrorResponse)
}

type PostServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	Repo    *PostRepoImpl
}

func NewPostServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *PostServiceImpl {
	return &PostServiceImpl{
		dbpool: dbpool,
		Redis:  rdb,
		Repo:   NewPostRepoImpl(dbpool),
	}
}

func (s *PostServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *PostServiceImpl) Create(req *CreatePostRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	if err := s.Repo.Create(req, uctx); err != nil {
		return err
	}
	return nil
}

func (s *PostServiceImpl) Show(postRequest ShowPostRequest) (*PostResponse, *error_responses.ErrorResponse) {

	return s.Repo.Show(postRequest)
}

func (s *PostServiceImpl) Delete(id int64) *error_responses.ErrorResponse {
	return s.Repo.Delete(id)
}

func (s *PostServiceImpl) IncrementView(id int64) (int, *error_responses.ErrorResponse) {
	return s.Repo.IncrementView(id)
}

func (s *PostServiceImpl) CreateShare(postID int64, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.CreateShare(postID, uctx.UserID)
}

func (s *PostServiceImpl) GetShareCount(postID int64) (int, *error_responses.ErrorResponse) {
	return s.Repo.GetShareCount(postID)
}

func (s *PostServiceImpl) ShowWithReposts(userID int64, page, perPage int) (*PostResponse, *error_responses.ErrorResponse) {
	return s.Repo.ShowWithReposts(userID, page, perPage)
}
