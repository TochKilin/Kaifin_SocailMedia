package quoteview

import (
	error_responses "kaifin_clone_api/pkg/responses"
)

type QuoteViewService interface {
	Track(
		quoteID int64,
		userID *int64,
		ip string,
	) *error_responses.ErrorResponse
}

type QuoteViewServiceImpl struct {
	Repo QuoteViewRepo
}

func NewQuoteViewServiceImpl(repo QuoteViewRepo) *QuoteViewServiceImpl {
	return &QuoteViewServiceImpl{
		Repo: repo,
	}
}

func (s *QuoteViewServiceImpl) Track(
	quoteID int64,
	userID *int64,
	ip string,
) *error_responses.ErrorResponse {

	return s.Repo.Track(quoteID, userID, ip)
}
