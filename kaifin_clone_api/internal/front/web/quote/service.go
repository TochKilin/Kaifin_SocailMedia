package quote

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type QuoteService interface {
	Show(req QuoteShowRequest, viewerID int64) (*QuoteResponse, *error_responses.ErrorResponse)
	ShowOne(id int64, viewerID *int64) (*Quote, *error_responses.ErrorResponse)
	Create(req *CreateQuoteRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	Update(id int64, req *UpdateQuoteRequest) (*Quote, *error_responses.ErrorResponse)
	Delete(id int64) *error_responses.ErrorResponse
	SetUserCtx(ctx *share.UserContext) bool
}

type QuoteServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	Repo    QuoteRepo
}

func NewQuoteServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *QuoteServiceImpl {
	return &QuoteServiceImpl{
		dbpool: dbpool,
		Redis:  rdb,
		Repo:   NewQuoteRepoImpl(dbpool),
	}
}

func (s *QuoteServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *QuoteServiceImpl) Show(req QuoteShowRequest, viewerID int64) (*QuoteResponse, *error_responses.ErrorResponse) {
	return s.Repo.Show(req, viewerID)
}

func (s *QuoteServiceImpl) ShowOne(id int64, viewerID *int64) (*Quote, *error_responses.ErrorResponse) {
	q, err := s.Repo.ShowOne(id)
	if err != nil {
		return nil, err
	}
	// log ការមើល (មិន block response បើ insert fail)
	_ = s.Repo.IncrementView(id, viewerID)
	return q, nil
}

func (s *QuoteServiceImpl) Create(req *CreateQuoteRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	q := &Quote{}
	q.new(req, uctx)
	return s.Repo.Create(q)
}

func (s *QuoteServiceImpl) Update(id int64, req *UpdateQuoteRequest) (*Quote, *error_responses.ErrorResponse) {
	updates := map[string]any{}

	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Content != nil {
		updates["content"] = *req.Content
	}
	if req.Visibility != nil {
		updates["visibility"] = *req.Visibility
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if len(updates) == 0 {
		msg := error_responses.ErrorResponse{}
		return nil, msg.NewErrorResponse("no_updates_provided", fmt.Errorf("empty"))
	}

	return s.Repo.Update(id, updates)
}

func (s *QuoteServiceImpl) Delete(id int64) *error_responses.ErrorResponse {
	return s.Repo.Delete(id)
}
