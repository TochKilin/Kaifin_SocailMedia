// repo.go
package userregister

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
	custom_sql "kaifin_clone_api/pkg/sql"
)

type UserRegisterRepo interface {
	Show(UserShowRequest) (*UserResponse, *error_responses.ErrorResponse)
	Create(req *AuthRegisterRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	Update(id int64, updates map[string]any) (*User, *error_responses.ErrorResponse)
	GetByUserName(userName string) (*User, *error_responses.ErrorResponse)
	Delete(id int64, deletedBy int64) *error_responses.ErrorResponse
}

type UserRegisterRepoImpl struct {
	dbpool *sqlx.DB
	// db     *sqlx.DB
	Redis *redis.Client
}

func NewUserRegisterRepoImpl(db *sqlx.DB, rdb *redis.Client) *UserRegisterRepoImpl {
	return &UserRegisterRepoImpl{
		dbpool: db,
		Redis:  rdb,
	}
}

func (r *UserRegisterRepoImpl) Create(c fiber.Ctx, req *AuthRegisterRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	fmt.Printf("REQ IN REPO: %+v\n", req)
	msg := error_responses.ErrorResponse{}
	user := User{}

	safeCtx := &share.UserContext{UserID: 0}
	if uctx != nil {
		safeCtx = uctx
	}

	if err := user.new(req, c, r.dbpool, r.Redis, uctx); err != nil {
		return msg.NewErrorResponse("invalid", err)
	}
	fmt.Println("USER BEFORE INSERT:")
	fmt.Printf("%+v\n", user)

	var createdBy int64 = safeCtx.UserID
	user.CreatedBy = &createdBy
	user.RoleName = "user"
	user.RoleID = 4
	user.IsAdmin = false

	err := r.dbpool.QueryRow(
		`
		INSERT INTO tbl_users (
			first_name,
			last_name,
			user_name,
			email,
			password,
			role_name,
			role_id,
			is_admin,
			created_by,
			profile_images,
			created_at
		)
		VALUES (
			$1,$2,$3,$4,$5,
			$6,$7,$8,$9,$10,
			NOW()
		)
		RETURNING id, created_at
		`,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Email,
		user.Password,
		user.RoleName,
		user.RoleID,
		user.IsAdmin,
		user.CreatedBy,
		user.ProfileImages,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	return nil
}

func (r *UserRegisterRepoImpl) Show(userRequest UserShowRequest) (*UserResponse, *error_responses.ErrorResponse) {
	var per_page = userRequest.PageOption.Perpage
	var page = userRequest.PageOption.Page // value get ត្រង់នេះ
	var offset = (page - 1) * per_page
	fmt.Printf("offset:%d", offset)
	var limit_clause = fmt.Sprintf(" LIMIT %d OFFSET %d", per_page, offset)
	var sql_orderby = custom_sql.BuildSQLSort(userRequest.Sorts)
	sql_filters, args_filters := custom_sql.BuildSQLFilter(userRequest.Filters)
	if len(args_filters) > 0 {
		sql_filters = " AND " + sql_filters
	}
	if searchClause, searchArgs := custom_sql.BuildSQLSearch(
		[]string{"u.user_name", "u.first_name", "u.email"},
		userRequest.Search, len(args_filters)+1,
	); searchClause != "" {
		sql_filters += " AND " + searchClause
		args_filters = append(args_filters, searchArgs...)
	}
	msg := error_responses.ErrorResponse{}
	var users []User
	query := fmt.Sprintf(
		`SELECT id, user_name, first_name, last_name, email,profile_images role_name, role_id, is_admin,
		 login_session, last_login, currency_id, language_id, status_id, created_at, updated_at
		 FROM tbl_users u
		WHERE deleted_at IS NULL
		%s %s %s`, sql_filters, sql_orderby, limit_clause)

	err := r.dbpool.Select(&users, query, args_filters...)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var total int

	countQuery := fmt.Sprintf(
		`SELECT COUNT(*) 
	 FROM tbl_users u
	 WHERE deleted_at IS NULL
	 %s`,
		sql_filters,
	)

	err = r.dbpool.Get(&total, countQuery, args_filters...)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	// find total
	return &UserResponse{
		Users: users,
		Total: total}, nil
}

func (r *UserRegisterRepoImpl) Update(id int64, updates map[string]any) (*User, *error_responses.ErrorResponse) {
	//
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

	query := fmt.Sprintf(
		`UPDATE tbl_users SET %s WHERE id = $%d AND deleted_at IS NULL RETURNING *`,
		strings.Join(setClauses, ", "), i,
	)
	args = append(args, id)

	var user User
	err := r.dbpool.Get(&user, query, args...)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	return &user, nil
}

func (r *UserRegisterRepoImpl) Delete(id int64, deletedBy int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	result, err := r.dbpool.Exec(
		`UPDATE tbl_users SET deleted_at = NOW(), deleted_by = $1 WHERE id = $2 AND deleted_at IS NULL`,
		deletedBy, id,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("user_not_found", fmt.Errorf("user %d not found", id))
	}
	return nil
}

func (r *UserRegisterRepoImpl) GetByUserName(userName string) (*User, *error_responses.ErrorResponse) {
	var user User
	err := r.dbpool.Get(&user,
		`SELECT * FROM tbl_users WHERE user_name = $1 LIMIT 1`, userName,
	)
	if err != nil {
		return nil, nil
	}
	return &user, nil
}

func (r *UserRegisterRepoImpl) Profile(userID int64) (*UserProfileResponse, *error_responses.ErrorResponse) {

	var user UserProfileResponse

	err := r.dbpool.Get(&user, `
		SELECT
			id,
			user_name,
			first_name,
			last_name,
			email,
			profile_images,
			role_id,
			role_name,
			last_login,
			created_at
		FROM tbl_users
		WHERE id = $1
		  AND deleted_at IS NULL
	`, userID)

	if err != nil {
		msg := error_responses.ErrorResponse{}
		return nil, msg.NewErrorResponse("user_not_found", err)
	}

	return &user, nil
}
