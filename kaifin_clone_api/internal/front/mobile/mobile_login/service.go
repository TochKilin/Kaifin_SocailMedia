package mobile_login

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

type AuthUserMobileService interface {
	UserLogin(userreq AuthLoginMobileRequest) (*AuthUserLoginMobileResponse, *error_responses.ErrorResponse)
	Profile(userID int64) (*AuthProfileResponse, *error_responses.ErrorResponse)
}

type AuthUserMobileServiceImpl struct {
	re    AuthUserRepo
	db    *sqlx.DB
	Redis *redis.Client
}

func NewAuthUserServiceImpl(db *sqlx.DB, rdb *redis.Client) *AuthUserMobileServiceImpl {
	ar := NewAuthUserRepoImpl(db, rdb)
	return &AuthUserMobileServiceImpl{
		re:    ar,
		db:    db,
		Redis: rdb,
	}
}

// User login
func (s *AuthUserMobileServiceImpl) UserLogin(userreq AuthLoginMobileRequest) (*AuthUserLoginMobileResponse, *error_responses.ErrorResponse) {
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

	var au AuthUserLoginMobileResponse
	au.Auth.Token = tokenString
	au.Auth.TokenType = "jwt"
	au.Auth.UserName = rst.UserName
	return &au, nil

}

// Get user profile
func (s *AuthUserMobileServiceImpl) Profile(userID int64) (*AuthProfileResponse, *error_responses.ErrorResponse) {
	return s.re.Profile(userID)
}
