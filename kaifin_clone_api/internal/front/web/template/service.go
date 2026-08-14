package template

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type TemplateService interface {
	List() (*TemplatesResponse, *error_responses.ErrorResponse)
	GetImage(id int64) ([]byte, string, error)
	Create(name string, fileData []byte, fileType string) *error_responses.ErrorResponse
}

type TemplateServiceImpl struct {
	repo TemplateRepo
}

func NewTemplateServiceImpl(dbpool *sqlx.DB) *TemplateServiceImpl {
	return &TemplateServiceImpl{
		repo: NewTemplateRepoImpl(dbpool),
	}
}

func (s *TemplateServiceImpl) List() (*TemplatesResponse, *error_responses.ErrorResponse) {
	return s.repo.List()
}

func (s *TemplateServiceImpl) GetImage(id int64) ([]byte, string, error) {
	return s.repo.GetImage(id)
}

func (s *TemplateServiceImpl) Create(name string, fileData []byte, fileType string) *error_responses.ErrorResponse {
	return s.repo.Create(name, fileData, fileType)
}
