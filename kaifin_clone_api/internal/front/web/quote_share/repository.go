package quoteshare

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type QuoteShareRepo interface {
	Track(quoteID int64, userID *int64, channel string) *error_responses.ErrorResponse
}

type QuoteShareRepoImpl struct {
	dbpool *sqlx.DB
}

func NewQuoteShareRepoImpl(db *sqlx.DB) QuoteShareRepo {
	return &QuoteShareRepoImpl{dbpool: db}
}

func (r *QuoteShareRepoImpl) Track(quoteID int64, userID *int64, channel string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	_, err := r.dbpool.Exec(
		`INSERT INTO quote_shares (quote_id, user_id, channel) VALUES ($1, $2, $3)`,
		quoteID, userID, channel,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}
