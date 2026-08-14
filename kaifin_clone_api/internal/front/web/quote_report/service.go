package quotereport

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type QuoteReportService interface {
	Create(req *CreateReportRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	Show(status string, page, perPage int) ([]QuoteReport, int, *error_responses.ErrorResponse)
	UpdateStatus(id int64, req *UpdateReportStatusRequest) *error_responses.ErrorResponse
	SetUserCtx(ctx *share.UserContext) bool
}

type QuoteReportServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Repo    QuoteReportRepo
}

func NewQuoteReportServiceImpl(dbpool *sqlx.DB) *QuoteReportServiceImpl {
	return &QuoteReportServiceImpl{
		dbpool: dbpool,
		Repo:   NewQuoteReportRepoImpl(dbpool),
	}
}

func (s *QuoteReportServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *QuoteReportServiceImpl) Create(req *CreateReportRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.Create(req.QuoteID, uctx.UserID, req.Reason)
}

func (s *QuoteReportServiceImpl) Show(status string, page, perPage int) ([]QuoteReport, int, *error_responses.ErrorResponse) {
	return s.Repo.Show(status, page, perPage)
}

func (s *QuoteReportServiceImpl) UpdateStatus(id int64, req *UpdateReportStatusRequest) *error_responses.ErrorResponse {
	return s.Repo.UpdateStatus(id, req.Status)
}
