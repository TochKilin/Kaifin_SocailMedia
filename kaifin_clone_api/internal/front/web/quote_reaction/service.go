package quotereaction

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type QuoteReactionService interface {
	React(req *ReactRequest, uctx *share.UserContext) (*QuoteReaction, *error_responses.ErrorResponse)
	Unreact(req *UnreactRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	ListTypes() ([]ReactionType, *error_responses.ErrorResponse)
	SetUserCtx(ctx *share.UserContext) bool
	GetByUser(quoteID int64, uctx *share.UserContext) (*QuoteReaction, *error_responses.ErrorResponse)
}

type QuoteReactionServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	Repo    QuoteReactionRepo
}

func NewQuoteReactionServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *QuoteReactionServiceImpl {
	return &QuoteReactionServiceImpl{
		dbpool: dbpool,
		Redis:  rdb,
		Repo:   NewQuoteReactionRepoImpl(dbpool),
	}
}

func (s *QuoteReactionServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *QuoteReactionServiceImpl) React(req *ReactRequest, uctx *share.UserContext) (*QuoteReaction, *error_responses.ErrorResponse) {
	return s.Repo.Upsert(req.QuoteID, uctx.UserID, req.ReactionTypeID)
}

func (s *QuoteReactionServiceImpl) Unreact(req *UnreactRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.Delete(req.QuoteID, uctx.UserID)
}

func (s *QuoteReactionServiceImpl) ListTypes() ([]ReactionType, *error_responses.ErrorResponse) {
	return s.Repo.ListTypes()
}

func (s *QuoteReactionServiceImpl) GetByUser(quoteID int64, uctx *share.UserContext) (*QuoteReaction, *error_responses.ErrorResponse) {
	return s.Repo.GetByUser(quoteID, uctx.UserID)
}
