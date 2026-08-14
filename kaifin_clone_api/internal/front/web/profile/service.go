package profile

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type ProfileService interface {
	GetProfile(userID int64) (*ProfileResponse, *error_responses.ErrorResponse)
	Update(userID int64, filename string) *error_responses.ErrorResponse
	UpdateCover(userID int64, filename string) *error_responses.ErrorResponse
	UpdateInfo(userID int64, req *UpdateProfileInfoRequest) *error_responses.ErrorResponse
}

type ProfileServiceImpl struct {
	re ProfileRepo
}

func NewProfileServiceImpl(db *sqlx.DB) *ProfileServiceImpl {
	repo := NewProfileRepoImpl(db)
	return &ProfileServiceImpl{
		re: repo,
	}
}

func (s *ProfileServiceImpl) Update(userID int64, filename string) *error_responses.ErrorResponse {
	return s.re.Update(userID, filename)
}

func (s *ProfileServiceImpl) GetProfile(userID int64) (*ProfileResponse, *error_responses.ErrorResponse) {

	return s.re.GetProfile(userID)
}

func (s *ProfileServiceImpl) UpdateCover(userID int64, filename string) *error_responses.ErrorResponse {
	return s.re.UpdateCover(userID, filename)
}

func (s *ProfileServiceImpl) UpdateInfo(userID int64, req *UpdateProfileInfoRequest) *error_responses.ErrorResponse {
	return s.re.UpdateInfo(userID, req)
}
