package auth

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_response "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type AuthService interface {
	Login(req AuthLoginRequest) (*AuthLoginResponse, *error_response.ErrorResponse)
}

type AuthServiceImpl struct {
	re    AuthRepo
	db    *sqlx.DB
	Redis *redis.Client
}

func NewAuthServiceImpl(db *sqlx.DB, redis *redis.Client) *AuthServiceImpl {
	ar := NewAuthRepoImpl(db, redis)
	return &AuthServiceImpl{
		re:    ar,
		db:    db,
		Redis: redis,
	}
}

func (s *AuthServiceImpl) Login(req AuthLoginRequest) (*AuthLoginResponse, *error_response.ErrorResponse) {
	msg := error_response.ErrorResponse{}
	rst, err := s.re.Login(req)

	if err != nil {
		return nil, err
	}

	loginSession := uuid.New().String()

	if errRes := s.re.UpdateLoginSession(int64(rst.ID), loginSession); errRes != nil {
		return nil, msg.NewErrorResponse("update_user_failed", fmt.Errorf("failed to update login session"))
	}

	// jwt generate store info payload as id username login sessoin....
	tokenString, _, tokenErr := share.GenerateToken(float64(rst.ID), rst.UserName, loginSession, int(rst.RoleID))
	if tokenErr != nil {
		return nil, msg.NewErrorResponse("token_error", fmt.Errorf("failed to generate token"))
	}

	var au AuthLoginResponse
	au.Auth.Token = tokenString
	au.Auth.TokenType = "jwt"
	au.Auth.UserName = rst.UserName
	return &au, nil

}
