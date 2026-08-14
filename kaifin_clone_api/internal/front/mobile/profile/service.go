package profile_mobile

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type ProfileMobileService interface {
	GetProfile(userID int64) (*ProfileResponse, *error_responses.ErrorResponse)
	Update(userID int64, filename string) *error_responses.ErrorResponse
}

type ProfileMobileServiceImpl struct {
	re ProfileMobileRepo
}

func NewProfileServiceImpl(db *sqlx.DB) *ProfileMobileServiceImpl {
	repo := NewProfileRepoImpl(db)
	return &ProfileMobileServiceImpl{
		re: repo,
	}
}

func (s *ProfileMobileServiceImpl) Update(userID int64, filename string) *error_responses.ErrorResponse {
	return s.re.Update(userID, filename)
}

func (s *ProfileMobileServiceImpl) GetProfile(userID int64) (*ProfileResponse, *error_responses.ErrorResponse) {

	return s.re.GetProfile(userID)
}
