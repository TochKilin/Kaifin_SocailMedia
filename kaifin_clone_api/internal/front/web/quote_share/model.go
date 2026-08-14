package quoteshare

type TrackShareRequest struct {
	QuoteID int64  `json:"quote_id" validate:"required"`
	Channel string `json:"channel" validate:"required,oneof=feed message telegram whatsapp copy_link"`
}
