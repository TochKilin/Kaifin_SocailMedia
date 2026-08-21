package addcard

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CartRepoImpl struct {
	dbpool *sqlx.DB
}

func NewCartRepoImpl(db *sqlx.DB) *CartRepoImpl {
	return &CartRepoImpl{dbpool: db}
}

func (r *CartRepoImpl) getOrCreateCart(userID int64) (int64, error) {
	var cartID int64
	err := r.dbpool.Get(&cartID, `SELECT id FROM tbl_cart WHERE user_id = $1`, userID)
	if err == nil {
		return cartID, nil
	}

	err = r.dbpool.QueryRow(`
		INSERT INTO tbl_cart (user_id, created_at)
		VALUES ($1, NOW())
		RETURNING id
	`, userID).Scan(&cartID)
	return cartID, err
}

func (r *CartRepoImpl) AddItem(userID, courseID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	cartID, err := r.getOrCreateCart(userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	var price float64
	if err := r.dbpool.Get(&price, `SELECT current_price FROM tbl_courses WHERE id = $1`, courseID); err != nil {
		return msg.NewErrorResponse("course_not_found", err)
	}

	_, err = r.dbpool.Exec(`
		INSERT INTO tbl_cart_items (cart_id, course_id, price_at_add, added_at)
		VALUES ($1, $2, $3, NOW())
		ON CONFLICT (cart_id, course_id) DO NOTHING
	`, cartID, courseID, price)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *CartRepoImpl) RemoveItem(userID, courseID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	cartID, err := r.getOrCreateCart(userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	_, err = r.dbpool.Exec(`
		DELETE FROM tbl_cart_items WHERE cart_id = $1 AND course_id = $2
	`, cartID, courseID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *CartRepoImpl) Show(userID int64) (*CartResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	cartID, err := r.getOrCreateCart(userID)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var items []CartItem
	query := `
		SELECT
			ci.id, ci.cart_id, ci.course_id, ci.price_at_add, ci.added_at AS created_at,
			c.title, c.thumbnail, c.current_price, c.original_price, c.rating, c.level_id,
			COALESCE(NULLIF(TRIM(u.first_name || ' ' || u.last_name), ''), u.user_name) AS instructor_name,
			COALESCE(u.profile_images, '') AS instructor_avatar
		FROM tbl_cart_items ci
		JOIN tbl_courses c ON c.id = ci.course_id
		LEFT JOIN tbl_users u ON u.id = c.instructor_id
		WHERE ci.cart_id = $1
		ORDER BY ci.added_at DESC
	`
	if err := r.dbpool.Select(&items, query, cartID); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var total float64
	for _, it := range items {
		total += it.CurrentPrice
	}

	return &CartResponse{Items: items, Total: total}, nil
}
