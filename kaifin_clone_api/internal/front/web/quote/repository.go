package quote

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	custom_sql "kaifin_clone_api/pkg/sql"
)

type QuoteRepo interface {
	Show(req QuoteShowRequest, viewerID int64) (*QuoteResponse, *error_responses.ErrorResponse)
	ShowOne(id int64) (*Quote, *error_responses.ErrorResponse)
	Create(q *Quote) *error_responses.ErrorResponse
	Update(id int64, updates map[string]any) (*Quote, *error_responses.ErrorResponse)
	Delete(id int64) *error_responses.ErrorResponse
	IncrementView(quoteID int64, userID *int64) *error_responses.ErrorResponse
}

type QuoteRepoImpl struct {
	dbpool *sqlx.DB
	redis  *redis.Client
}

func NewQuoteRepoImpl(db *sqlx.DB) QuoteRepo {
	return &QuoteRepoImpl{
		dbpool: db,
	}
}
func (r *QuoteRepoImpl) Show(req QuoteShowRequest, viewerID int64) (*QuoteResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var per_page = req.PageOption.Perpage
	var page = req.PageOption.Page
	var offset = (page - 1) * per_page
	limit_clause := fmt.Sprintf(" LIMIT %d OFFSET %d", per_page, offset)

	orderby := " ORDER BY q.created_at DESC"
	if req.Tab == "popular" {
		orderby = " ORDER BY q.views_count DESC"
	}

	sql_filters, args_filters := custom_sql.BuildSQLFilter(req.Filters)
	if len(args_filters) > 0 {
		sql_filters = " AND " + sql_filters
	}

	if searchClause, searchArgs := custom_sql.BuildSQLSearch(
		[]string{"q.title", "q.content"},
		req.Search, len(args_filters)+1,
	); searchClause != "" {
		sql_filters += " AND " + searchClause
		args_filters = append(args_filters, searchArgs...)
	}

	viewerParamIdx := len(args_filters) + 1
	args_filters = append(args_filters, viewerID)

	var quotes []Quote
	query := fmt.Sprintf(
		`SELECT q.id, q.user_id, q.title, q.content, q.visibility, q.status,
     q.views_count, q.likes_count, q.created_at, q.updated_at, u.user_name AS username,
     u.profile_images AS profile_images,
     qr.reaction_type_id AS my_reaction_type_id
     FROM quotes q
     JOIN tbl_users u ON u.id = q.user_id
     LEFT JOIN quote_reactions qr ON qr.quote_id = q.id AND qr.user_id = $%d
     WHERE q.status = 'published'
     %s %s %s`, viewerParamIdx, sql_filters, orderby, limit_clause)

	err := r.dbpool.Select(&quotes, query, args_filters...)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	countArgs := args_filters[:len(args_filters)-1]
	var total int
	countQuery := fmt.Sprintf(
		`SELECT COUNT(*) FROM quotes q WHERE q.status = 'published' %s`, sql_filters)

	err = r.dbpool.Get(&total, countQuery, countArgs...)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &QuoteResponse{Quotes: quotes, Total: total}, nil
}

func (r *QuoteRepoImpl) ShowOne(id int64) (*Quote, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var q Quote

	err := r.dbpool.Get(&q,
		`SELECT q.id, q.user_id, q.title, q.content, q.visibility, q.status,
		 q.views_count, q.likes_count, q.created_at, q.updated_at, u.user_name AS username
		 FROM quotes q
		 JOIN tbl_users u ON u.id = q.user_id
		 WHERE q.id = $1 LIMIT 1`, id,
	)
	if err != nil {
		return nil, msg.NewErrorResponse("quote_not_found", err)
	}
	return &q, nil
}

func (r *QuoteRepoImpl) Create(q *Quote) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	err := r.dbpool.QueryRow(
		`INSERT INTO quotes (user_id, title, content, visibility, status)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id, created_at, updated_at`,
		q.UserID, q.Title, q.Content, q.Visibility, q.Status,
	).Scan(&q.ID, &q.CreatedAt, &q.UpdatedAt)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *QuoteRepoImpl) Update(id int64, updates map[string]any) (*Quote, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	setClauses := []string{}
	args := []any{}
	i := 1
	for col, val := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, i))
		args = append(args, val)
		i++
	}
	if len(setClauses) == 0 {
		return nil, msg.NewErrorResponse("no_updates_provided", fmt.Errorf("empty update"))
	}
	setClauses = append(setClauses, "updated_at = now()")

	query := fmt.Sprintf(
		`UPDATE quotes SET %s WHERE id = $%d RETURNING *`,
		strings.Join(setClauses, ", "), i,
	)
	args = append(args, id)

	var q Quote
	err := r.dbpool.Get(&q, query, args...)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	return &q, nil
}

func (r *QuoteRepoImpl) Delete(id int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	result, err := r.dbpool.Exec(`DELETE FROM quotes WHERE id = $1`, id)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("quote_not_found", fmt.Errorf("quote %d not found", id))
	}
	return nil
}

func (r *QuoteRepoImpl) IncrementView(quoteID int64, userID *int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	_, err := r.dbpool.Exec(
		`INSERT INTO quote_views (quote_id, user_id) VALUES ($1, $2)`,
		quoteID, userID,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}
