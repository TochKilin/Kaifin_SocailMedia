package quotehidden

type HideRequest struct {
	QuoteID int64 `json:"quote_id" validate:"required"`
}
