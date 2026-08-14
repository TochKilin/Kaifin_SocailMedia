package article

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type ArticlesServie interface {
	Create(req *CreateArticleRequest, uctx *share.UserContext) (*Article, *error_responses.ErrorResponse)
	Update(id int64, req *UpdateArticleRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	Show(req ShowArticlesRequest) (*ArticlesResponse, *error_responses.ErrorResponse)
	Detail(id int64) (*Article, *error_responses.ErrorResponse)
	Delete(id int64, uctx *share.UserContext) *error_responses.ErrorResponse
	SetUserCtx(ctx *share.UserContext) bool
	ToggleLike(articleID int64, uctx *share.UserContext) (bool, *error_responses.ErrorResponse)
	ToggleSave(articleID int64, uctx *share.UserContext) (bool, *error_responses.ErrorResponse)
	Report(req *ReportArticleRequest, uctx *share.UserContext) *error_responses.ErrorResponse
}

type ArticlesServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Repo    *ArticlesRepoImpl
}

func NewArticlesServiceImpl(dbpool *sqlx.DB) *ArticlesServiceImpl {
	return &ArticlesServiceImpl{
		dbpool: dbpool,
		Repo:   NewArticlesRepoImpl(dbpool),
	}
}

func (s *ArticlesServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *ArticlesServiceImpl) Create(req *CreateArticleRequest, uctx *share.UserContext) (*Article, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	article := Article{}
	if err := article.new(req, uctx); err != nil {
		return nil, msg.NewErrorResponse("invalid", err)
	}
	blocks := blocksFromInput(req.Blocks)
	if e := s.Repo.Create(&article, req.Tags, blocks); e != nil {
		return nil, e
	}
	article.Tags = req.Tags
	article.Blocks = blocks
	return &article, nil
}

func (s *ArticlesServiceImpl) Update(id int64, req *UpdateArticleRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	article := Article{ID: id, UserID: uctx.UserID}
	article.applyUpdate(req)
	blocks := blocksFromInput(req.Blocks)
	return s.Repo.Update(&article, req.Tags, blocks)
}

func (s *ArticlesServiceImpl) Show(req ShowArticlesRequest) (*ArticlesResponse, *error_responses.ErrorResponse) {
	var userID int64
	if s.UserCtx != nil {
		userID = s.UserCtx.UserID
	}
	return s.Repo.Show(req, userID)
}

func (s *ArticlesServiceImpl) Detail(id int64) (*Article, *error_responses.ErrorResponse) {
	var userID int64
	if s.UserCtx != nil {
		userID = s.UserCtx.UserID
	}
	return s.Repo.Detail(id, userID)
}

func (s *ArticlesServiceImpl) Delete(id int64, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.Delete(id, uctx.UserID)
}

func (s *ArticlesServiceImpl) ToggleLike(articleID int64, uctx *share.UserContext) (bool, *error_responses.ErrorResponse) {
	return s.Repo.ToggleLike(articleID, uctx.UserID)
}

func (s *ArticlesServiceImpl) ToggleSave(articleID int64, uctx *share.UserContext) (bool, *error_responses.ErrorResponse) {
	return s.Repo.ToggleSave(articleID, uctx.UserID)
}

func (s *ArticlesServiceImpl) Report(req *ReportArticleRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.Report(req.ArticleID, uctx.UserID, req.ReportType, req.Text)
}
