package quotehidden

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type QuoteHiddenRepo interface {
	Hide(quoteID int64, userID int64) *error_responses.ErrorResponse
	Unhide(quoteID int64, userID int64) *error_responses.ErrorResponse
}

type QuoteHiddenRepoImpl struct {
	dbpool *sqlx.DB
}

func NewQuoteHiddenRepoImpl(db *sqlx.DB) QuoteHiddenRepo {
	return &QuoteHiddenRepoImpl{dbpool: db}
}

func (r *QuoteHiddenRepoImpl) Hide(quoteID int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	_, err := r.dbpool.Exec(
		`INSERT INTO quote_hidden (quote_id, user_id) VALUES ($1, $2)
		 ON CONFLICT (quote_id, user_id) DO NOTHING`,
		quoteID, userID,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *QuoteHiddenRepoImpl) Unhide(quoteID int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	result, err := r.dbpool.Exec(
		`DELETE FROM quote_hidden WHERE quote_id = $1 AND user_id = $2`,
		quoteID, userID,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("hidden_entry_not_found", fmt.Errorf("not hidden"))
	}
	return nil
}
