package comments_mobile

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type CommentsMobileServie interface {
	Create(req *CreateCommentMobileRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	Show(req ShowCommentsMobileRequest) (*CommentsMobileResponse, *error_responses.ErrorResponse)
	Delete(id int64, uctx *share.UserContext) *error_responses.ErrorResponse
	SetUserCtx(ctx *share.UserContext) bool
}

type CommentsMobileServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Repo    *CommentsMobileRepoImpl
}

func NewCommentsServiceImpl(dbpool *sqlx.DB) *CommentsMobileServiceImpl {
	return &CommentsMobileServiceImpl{
		dbpool: dbpool,
		Repo:   NewCommentsMobileRepoImpl(dbpool),
	}
}

func (s *CommentsMobileServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *CommentsMobileServiceImpl) Create(req *CreateCommentMobileRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	comment := CommentMobile{}

	if err := comment.new(req, uctx); err != nil {
		return msg.NewErrorResponse("invalid", err)
	}

	return s.Repo.Create(&comment)
}

func (s *CommentsMobileServiceImpl) Show(req ShowCommentsMobileRequest) (*CommentsMobileResponse, *error_responses.ErrorResponse) {
	return s.Repo.Show(req)
}

func (s *CommentsMobileServiceImpl) Delete(id int64, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.Delete(id, uctx.UserID)
}
