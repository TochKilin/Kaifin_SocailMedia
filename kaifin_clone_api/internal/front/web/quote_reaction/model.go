package quotereaction

import (
	"time"
)

type QuoteReaction struct {
	ID             int64     `json:"id" db:"id"`
	QuoteID        int64     `json:"quote_id" db:"quote_id"`
	UserID         int64     `json:"user_id" db:"user_id"`
	ReactionTypeID int16     `json:"reaction_type_id" db:"reaction_type_id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`

	ReactionName string `json:"reaction_name,omitempty" db:"reaction_name"`
}

type ReactRequest struct {
	QuoteID        int64 `json:"quote_id" validate:"required"`
	ReactionTypeID int16 `json:"reaction_type_id" validate:"required"`
}

type UnreactRequest struct {
	QuoteID int64 `json:"quote_id" validate:"required"`
}

type ReactionType struct {
	ID        int16  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	IconType  string `json:"icon_type" db:"icon_type"`
	IconValue string `json:"icon_value" db:"icon_value"`
	SortOrder int16  `json:"sort_order" db:"sort_order"`
}
