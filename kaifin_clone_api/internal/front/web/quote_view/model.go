package quoteview

type TrackViewRequest struct {
	QuoteID int64 `json:"quote_id" validate:"required"`
}
