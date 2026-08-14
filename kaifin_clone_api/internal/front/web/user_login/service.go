package user_login

import (
	"fmt"
	// "uuid"
	// "uuid"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type AuthUserService interface {
	UserLogin(userreq AuthLoginRequest) (*AuthUserLoginResponse, *error_responses.ErrorResponse)
	Profile(userID int64) (*AuthProfileResponse, *error_responses.ErrorResponse)
}

type AuthUserServiceImpl struct {
	re    AuthUserRepo
	db    *sqlx.DB
	Redis *redis.Client
}

func NewAuthUserServiceImpl(db *sqlx.DB, rdb *redis.Client) *AuthUserServiceImpl {
	ar := NewAuthUserRepoImpl(db, rdb)
	return &AuthUserServiceImpl{
		re:    ar,
		db:    db,
		Redis: rdb,
	}
}

func (s *AuthUserServiceImpl) UserLogin(userreq AuthLoginRequest) (*AuthUserLoginResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	rst, err := s.re.UserLogin(userreq)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(rst.Password),
		[]byte(userreq.Password),
	); err != nil {

		return nil, msg.NewErrorResponse(
			"invalid_password",
			fmt.Errorf("invalid username or password"),
		)
	}

	loginSession := uuid.New().String()

	if errRes := s.re.UpdateLoginSession(int64(rst.ID), loginSession); errRes != nil {
		return nil, msg.NewErrorResponse("update_user_failed", fmt.Errorf("failed to update login session"))
	}

	tokenString, _, tokenErr := share.GenerateToken(float64(rst.ID), rst.UserName, loginSession, int(rst.RoleID))
	if tokenErr != nil {
		return nil, msg.NewErrorResponse("token_error", fmt.Errorf("failed to generate token"))
	}

	var au AuthUserLoginResponse
	au.Auth.Token = tokenString
	au.Auth.TokenType = "jwt"
	au.Auth.UserName = rst.UserName
	return &au, nil

}

func (s *AuthUserServiceImpl) Profile(userID int64) (*AuthProfileResponse, *error_responses.ErrorResponse) {
	return s.re.Profile(userID)
}
