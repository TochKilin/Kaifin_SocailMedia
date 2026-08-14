package quotehidden

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type QuoteHiddenService interface {
	Hide(req *HideRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	Unhide(quoteID int64, uctx *share.UserContext) *error_responses.ErrorResponse
	SetUserCtx(ctx *share.UserContext) bool
}

type QuoteHiddenServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Repo    QuoteHiddenRepo
}

func NewQuoteHiddenServiceImpl(dbpool *sqlx.DB) *QuoteHiddenServiceImpl {
	return &QuoteHiddenServiceImpl{
		dbpool: dbpool,
		Repo:   NewQuoteHiddenRepoImpl(dbpool),
	}
}

func (s *QuoteHiddenServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *QuoteHiddenServiceImpl) Hide(req *HideRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.Hide(req.QuoteID, uctx.UserID)
}

func (s *QuoteHiddenServiceImpl) Unhide(quoteID int64, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.Unhide(quoteID, uctx.UserID)
}
