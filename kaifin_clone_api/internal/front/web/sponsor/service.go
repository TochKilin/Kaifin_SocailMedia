package sponsor

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type SponsorsServiceImpl struct {
	dbpool *sqlx.DB
	Repo   *SponsorsRepoImpl
}

func NewSponsorsServiceImpl(dbpool *sqlx.DB) *SponsorsServiceImpl {
	return &SponsorsServiceImpl{
		dbpool: dbpool,
		Repo:   NewSponsorsRepoImpl(dbpool),
	}
}

func (s *SponsorsServiceImpl) Create(req *CreateSponsorRequest, logoPath string) (*Sponsor, *error_responses.ErrorResponse) {
	sp := &Sponsor{
		Name:       req.Name,
		LogoImage:  logoPath,
		IsVerified: req.IsVerified,
		SortOrder:  req.SortOrder,
	}
	if req.WebsiteURL != "" {
		sp.WebsiteURL = &req.WebsiteURL
	}
	if e := s.Repo.Create(sp); e != nil {
		return nil, e
	}
	return sp, nil
}

func (s *SponsorsServiceImpl) Update(id int64, req *UpdateSponsorRequest, logoPath string) *error_responses.ErrorResponse {
	sp := &Sponsor{
		ID:         id,
		Name:       req.Name,
		LogoImage:  logoPath,
		IsVerified: req.IsVerified,
		IsActive:   req.IsActive,
		SortOrder:  req.SortOrder,
	}
	if req.WebsiteURL != "" {
		sp.WebsiteURL = &req.WebsiteURL
	}
	return s.Repo.Update(sp)
}

func (s *SponsorsServiceImpl) Delete(id int64) *error_responses.ErrorResponse {
	return s.Repo.Delete(id)
}

func (s *SponsorsServiceImpl) Show(req ShowSponsorsRequest) (*SponsorsResponse, *error_responses.ErrorResponse) {
	return s.Repo.Show(req)
}
