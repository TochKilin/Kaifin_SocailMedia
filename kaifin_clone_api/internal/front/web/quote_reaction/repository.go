package quotereaction

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
)

type QuoteReactionRepo interface {
	Upsert(quoteID int64, userID int64, reactionTypeID int16) (*QuoteReaction, *error_responses.ErrorResponse)
	Delete(quoteID int64, userID int64) *error_responses.ErrorResponse
	GetByUser(quoteID int64, userID int64) (*QuoteReaction, *error_responses.ErrorResponse)
	ListTypes() ([]ReactionType, *error_responses.ErrorResponse)
}

type QuoteReactionRepoImpl struct {
	dbpool *sqlx.DB
	redis  *redis.Client
}

func NewQuoteReactionRepoImpl(db *sqlx.DB) QuoteReactionRepo {
	return &QuoteReactionRepoImpl{
		dbpool: db,
	}
}

func (r *QuoteReactionRepoImpl) Upsert(quoteID int64, userID int64, reactionTypeID int16) (*QuoteReaction, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var reaction QuoteReaction
	err := r.dbpool.Get(&reaction,
		`INSERT INTO quote_reactions (quote_id, user_id, reaction_type_id)
		 VALUES ($1, $2, $3)
		 ON CONFLICT (quote_id, user_id)
		 DO UPDATE SET reaction_type_id = EXCLUDED.reaction_type_id
		 RETURNING id, quote_id, user_id, reaction_type_id, created_at`,
		quoteID, userID, reactionTypeID,
	)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	return &reaction, nil
}

func (r *QuoteReactionRepoImpl) Delete(quoteID int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	result, err := r.dbpool.Exec(
		`DELETE FROM quote_reactions WHERE quote_id = $1 AND user_id = $2`,
		quoteID, userID,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("reaction_not_found", err)
	}
	return nil
}

func (r *QuoteReactionRepoImpl) GetByUser(quoteID int64, userID int64) (*QuoteReaction, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var reaction QuoteReaction

	err := r.dbpool.Get(&reaction,
		`SELECT id, quote_id, user_id, reaction_type_id, created_at
		 FROM quote_reactions WHERE quote_id = $1 AND user_id = $2 LIMIT 1`,
		quoteID, userID,
	)
	if err != nil {
		return nil, msg.NewErrorResponse("reaction_not_found", err)
	}
	return &reaction, nil
}

func (r *QuoteReactionRepoImpl) ListTypes() ([]ReactionType, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var types []ReactionType

	err := r.dbpool.Select(&types,
		`SELECT id, name, icon_type, icon_value, sort_order
		 FROM reaction_types ORDER BY sort_order ASC`,
	)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	return types, nil
}
