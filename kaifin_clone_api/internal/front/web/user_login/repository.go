package user_login

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/configs/radis"
	error_responses "kaifin_clone_api/pkg/responses"
)

type AuthUserRepo interface {
	UserLogin(userreq AuthLoginRequest) (*AuthUser, *error_responses.ErrorResponse)
	UpdateLoginSession(userID int64, loginSession string) *error_responses.ErrorResponse
	CheckSession(loginSession string, userID float64) (bool, error)
	Profile(userID int64) (*AuthProfileResponse, *error_responses.ErrorResponse)
}

type AuthUserRepoImpl struct {
	db    *sqlx.DB
	Redis *redis.Client
}

func NewAuthUserRepoImpl(db *sqlx.DB, rdb *redis.Client) *AuthUserRepoImpl {
	return &AuthUserRepoImpl{
		db:    db,
		Redis: rdb,
	}
}

func (ur *AuthUserRepoImpl) UserLogin(userreq AuthLoginRequest) (*AuthUser, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var user AuthUser

	// កែប្រែ SQL Query ត្រង់ WHERE clause ឱ្យឆែកទាំង user_name និង email
	err := ur.db.Get(&user, `
	SELECT id, user_name, email, password, role_id
	FROM tbl_users
WHERE (LOWER(user_name) = LOWER($1) OR LOWER(email) = LOWER($1))
  AND role_id = $2
  AND deleted_at IS NULL
	LIMIT 1
	`, userreq.UserName, userreq.RoleID)

	if err != nil {
		return nil, msg.NewErrorResponse("user_not_found", fmt.Errorf("invalid username/email or password"))
	}
	return &user, nil
}

func (r *AuthUserRepoImpl) UpdateLoginSession(userID int64, loginSession string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	_, err := r.db.Exec(
		`UPDATE tbl_users
         SET login_session = $1,
             last_login = NOW()
         WHERE id = $2`,
		loginSession,
		userID,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	key := fmt.Sprintf("user_info_id:%d", userID)
	claims := map[string]interface{}{
		"user_id":       userID,
		"login_session": loginSession,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	redisClient := radis.NewRedisClient()
	rdb := radis.NewRedisUtil(redisClient)
	if err = rdb.SetCacheKey(key, claims, ctx); err != nil {
		fmt.Printf("redis warning: %v\n", err)
	}

	return nil
}

func (r *AuthUserRepoImpl) CheckSession(loginSession string, userID float64) (bool, error) {
	ctx := context.Background()
	redisKey := fmt.Sprintf("user_session:%d", int64(userID))

	val, err := r.Redis.Get(ctx, redisKey).Result()
	if err == nil {
		return val == loginSession, nil
	}

	var exists bool
	err = r.db.QueryRow(
		`SELECT EXISTS(SELECT 1 FROM tbl_users WHERE id = $1 AND login_session = $2)`,
		int(userID), loginSession,
	).Scan(&exists)
	if err != nil {
		return false, err
	}

	if exists {
		r.Redis.Set(ctx, redisKey, loginSession, 24*time.Hour)
	}

	return exists, nil
}

func (ur *AuthUserRepoImpl) Profile(userID int64) (*AuthProfileResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var profile AuthProfileResponse
	err := ur.db.Get(&profile, `
        SELECT id, user_name, first_name, last_name, email,
               profile_images, role_id, role_name, last_login, created_at
        FROM tbl_users
        WHERE id = $1
          AND deleted_at IS NULL
        LIMIT 1
    `, userID)

	if err != nil {
		return nil, msg.NewErrorResponse("user_not_found", fmt.Errorf("profile not found"))
	}

	return &profile, nil
}
