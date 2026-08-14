package auth

import (
	// internal package
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/configs/radis"
	error_response "kaifin_clone_api/pkg/responses"
	error_responses "kaifin_clone_api/pkg/responses"
)

type AuthRepo interface {
	Login(req AuthLoginRequest) (*Auth, *error_response.ErrorResponse)
	UpdateLoginSession(userID int64, loginSession string) *error_responses.ErrorResponse
	CheckSession(loginSession string, userID float64) (bool, error)
}

type AuthRepoImpl struct {
	db    *sqlx.DB
	Redis *redis.Client
}

func NewAuthRepoImpl(db *sqlx.DB, rdb *redis.Client) *AuthRepoImpl {
	return &AuthRepoImpl{
		db:    db,
		Redis: rdb,
	}
}

func (r *AuthRepoImpl) Login(req AuthLoginRequest) (*Auth, *error_response.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var user Auth
	err := r.db.Get(&user, `SELECT id, user_name, password, role_id FROM tbl_users WHERE user_name = $1 AND password = $2 AND deleted_at IS NULL LIMIT 1`, req.UserName, req.Password)
	if err != nil {
		return nil, msg.NewErrorResponse("user_not_found", fmt.Errorf("invaild username or password"))
	}

	return &user, nil
}

func (r *AuthRepoImpl) UpdateLoginSession(userID int64, loginSession string) *error_responses.ErrorResponse {
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

func (r *AuthRepoImpl) CheckSession(loginSession string, userID float64) (bool, error) {
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
