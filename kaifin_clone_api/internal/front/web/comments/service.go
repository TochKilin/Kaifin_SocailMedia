package comments

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type CommentsServie interface {
	Create(req *CreateCommentRequest, uctx *share.UserContext, imagePaths []string) *error_responses.ErrorResponse
	Show(req ShowCommentsRequest) (*CommentsResponse, *error_responses.ErrorResponse)
	Delete(id int64, uctx *share.UserContext) *error_responses.ErrorResponse
	SetUserCtx(ctx *share.UserContext) bool
	ToggleLike(commentID int64, uctx *share.UserContext) (bool, *error_responses.ErrorResponse)
}

type CommentsServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Repo    *CommentsRepoImpl
}

func NewCommentsServiceImpl(dbpool *sqlx.DB) *CommentsServiceImpl {
	return &CommentsServiceImpl{
		dbpool: dbpool,
		Repo:   NewCommentsRepoImpl(dbpool),
	}
}

func (s *CommentsServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *CommentsServiceImpl) Create(req *CreateCommentRequest, uctx *share.UserContext, imagePaths []string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	comment := Comment{}
	if err := comment.new(req, uctx); err != nil {
		return msg.NewErrorResponse("invalid", err)
	}
	return s.Repo.Create(&comment, imagePaths)
}

func (s *CommentsServiceImpl) Show(req ShowCommentsRequest) (*CommentsResponse, *error_responses.ErrorResponse) {
	var userID int64
	if s.UserCtx != nil {
		userID = s.UserCtx.UserID
	}
	return s.Repo.Show(req, userID)
}

func (s *CommentsServiceImpl) Delete(id int64, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.Delete(id, uctx.UserID)
}

func (s *CommentsServiceImpl) ToggleLike(commentID int64, uctx *share.UserContext) (bool, *error_responses.ErrorResponse) {
	return s.Repo.ToggleLike(commentID, uctx.UserID)
}
