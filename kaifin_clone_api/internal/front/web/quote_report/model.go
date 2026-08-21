package quotereport

import "time"

type QuoteReport struct {
	ID         int64     `json:"id" db:"id"`
	QuoteID    int64     `json:"quote_id" db:"quote_id"`
	ReporterID int64     `json:"reporter_id" db:"reporter_id"`
	Reason     *string   `json:"reason" db:"reason"`
	Status     string    `json:"status" db:"status"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type CreateReportRequest struct {
	QuoteID int64   `json:"quote_id" validate:"required"`
	Reason  *string `json:"reason"`
}

type UpdateReportStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=pending reviewed dismissed"`
}
