package quotereport

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type QuoteReportRepo interface {
	Create(quoteID int64, reporterID int64, reason *string) *error_responses.ErrorResponse
	Show(status string, page, perPage int) ([]QuoteReport, int, *error_responses.ErrorResponse)
	UpdateStatus(id int64, status string) *error_responses.ErrorResponse
}

type QuoteReportRepoImpl struct {
	dbpool *sqlx.DB
}

func NewQuoteReportRepoImpl(db *sqlx.DB) QuoteReportRepo {
	return &QuoteReportRepoImpl{dbpool: db}
}

func (r *QuoteReportRepoImpl) Create(quoteID int64, reporterID int64, reason *string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	_, err := r.dbpool.Exec(
		`INSERT INTO quote_reports (quote_id, reporter_id, reason) VALUES ($1, $2, $3)`,
		quoteID, reporterID, reason,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *QuoteReportRepoImpl) Show(status string, page, perPage int) ([]QuoteReport, int, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var reports []QuoteReport
	offset := (page - 1) * perPage

	filterClause := ""
	args := []any{}
	if status != "" {
		filterClause = "WHERE status = $1"
		args = append(args, status)
	}

	query := fmt.Sprintf(
		`SELECT id, quote_id, reporter_id, reason, status, created_at
		 FROM quote_reports %s ORDER BY created_at DESC LIMIT %d OFFSET %d`,
		filterClause, perPage, offset,
	)
	if err := r.dbpool.Select(&reports, query, args...); err != nil {
		return nil, 0, msg.NewErrorResponse("database_error", err)
	}

	var total int
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM quote_reports %s`, filterClause)
	if err := r.dbpool.Get(&total, countQuery, args...); err != nil {
		return nil, 0, msg.NewErrorResponse("database_error", err)
	}

	return reports, total, nil
}

func (r *QuoteReportRepoImpl) UpdateStatus(id int64, status string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	result, err := r.dbpool.Exec(
		`UPDATE quote_reports SET status = $1 WHERE id = $2`,
		status, id,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("report_not_found", fmt.Errorf("report %d not found", id))
	}
	return nil
}
