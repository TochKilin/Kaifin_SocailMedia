package courseinroll

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CourseEnrollmentService interface {
	ShowEnrolled(userID int64, req ShowEnrolledRequest) (*EnrolledListResponse, *error_responses.ErrorResponse)
	CreateEnrollments(userID int64, req CreateEnrollmentRequest) (*CreateEnrollmentResponse, *error_responses.ErrorResponse)
}

type CourseEnrollmentServiceImpl struct {
	Repo CourseEnrollmentRepo
}

func NewCourseEnrollmentServiceImpl(dbpool *sqlx.DB) *CourseEnrollmentServiceImpl {
	return &CourseEnrollmentServiceImpl{
		Repo: NewCourseEnrollmentRepoImpl(dbpool),
	}
}

func (s *CourseEnrollmentServiceImpl) ShowEnrolled(userID int64, req ShowEnrolledRequest) (*EnrolledListResponse, *error_responses.ErrorResponse) {
	return s.Repo.ShowEnrolled(userID, req)
}

func (s *CourseEnrollmentServiceImpl) CreateEnrollments(userID int64, req CreateEnrollmentRequest) (*CreateEnrollmentResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	if len(req.CourseIDs) == 0 {
		return nil, msg.NewErrorResponse("course_ids_required", nil)
	}

	enrolled, err := s.Repo.CreateEnrollments(userID, req.CourseIDs)
	if err != nil {
		return nil, err
	}

	return &CreateEnrollmentResponse{Enrolled: enrolled}, nil
}
