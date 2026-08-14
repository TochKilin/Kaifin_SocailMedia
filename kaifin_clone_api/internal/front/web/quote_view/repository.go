package quoteview

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type QuoteViewRepo interface {
	Track(quoteID int64, userID *int64, ip string) *error_responses.ErrorResponse
}

type QuoteViewRepoImpl struct {
	dbpool *sqlx.DB
}

func NewQuoteViewRepoImpl(db *sqlx.DB) QuoteViewRepo {
	return &QuoteViewRepoImpl{dbpool: db}
}

func (r *QuoteViewRepoImpl) Track(quoteID int64, userID *int64, ip string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	_, err := r.dbpool.Exec(
		`INSERT INTO quote_views (quote_id, user_id, ip_address) VALUES ($1, $2, $3)`,
		quoteID, userID, ip,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}
