package sponsor

import "time"

type Sponsor struct {
	ID         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	LogoImage  string    `json:"logo_image" db:"logo_image"`
	WebsiteURL *string   `json:"website_url" db:"website_url"`
	IsVerified bool      `json:"is_verified" db:"is_verified"`
	IsActive   bool      `json:"is_active" db:"is_active"`
	SortOrder  int       `json:"sort_order" db:"sort_order"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// ---------- Create ----------

type CreateSponsorRequest struct {
	Name       string `json:"name" validate:"required"`
	WebsiteURL string `json:"website_url"`
	IsVerified bool   `json:"is_verified"`
	SortOrder  int    `json:"sort_order"`
}

type UpdateSponsorRequest struct {
	Name       string `json:"name" validate:"required"`
	WebsiteURL string `json:"website_url"`
	IsVerified bool   `json:"is_verified"`
	IsActive   bool   `json:"is_active"`
	SortOrder  int    `json:"sort_order"`
}

type ShowSponsorsRequest struct {
	Page    int `query:"page"`
	PerPage int `query:"per_page"`
}

type SponsorsResponse struct {
	Sponsors []Sponsor `json:"sponsors"`
	Total    int       `json:"total"`
}
