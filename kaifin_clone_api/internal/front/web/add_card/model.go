package addcard

import "time"

type CartItem struct {
	ID               int64     `json:"id" db:"id"`
	CartID           int64     `json:"cart_id" db:"cart_id"`
	CourseID         int64     `json:"course_id" db:"course_id"`
	PriceAtAdd       float64   `json:"price_at_add" db:"price_at_add"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	Title            string    `json:"title" db:"title"`
	Thumbnail        string    `json:"thumbnail" db:"thumbnail"`
	CurrentPrice     float64   `json:"current_price" db:"current_price"`
	OriginalPrice    float64   `json:"original_price" db:"original_price"`
	Rating           float64   `json:"rating" db:"rating"`
	InstructorName   string    `json:"instructor_name" db:"instructor_name"`
	InstructorAvatar string    `json:"instructor_avatar" db:"instructor_avatar"`
	LevelID          *int64    `json:"level_id" db:"level_id"`
}

type AddToCartRequest struct {
	CourseID int64 `json:"course_id" validate:"required"`
}

type CartResponse struct {
	Items []CartItem `json:"items"`
	Total float64    `json:"total"`
}
