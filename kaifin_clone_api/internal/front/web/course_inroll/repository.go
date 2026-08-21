package courseinroll

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CourseEnrollmentRepo interface {
	ShowEnrolled(userID int64, req ShowEnrolledRequest) (*EnrolledListResponse, *error_responses.ErrorResponse)
	IsEnrolled(userID int64, courseID int64) (bool, *error_responses.ErrorResponse)
	CreateEnrollments(userID int64, courseIDs []int64) ([]int64, *error_responses.ErrorResponse)
}

type CourseEnrollmentRepoImpl struct {
	dbpool *sqlx.DB
}

func NewCourseEnrollmentRepoImpl(db *sqlx.DB) CourseEnrollmentRepo {
	return &CourseEnrollmentRepoImpl{dbpool: db}
}

func (r *CourseEnrollmentRepoImpl) ShowEnrolled(userID int64, req ShowEnrolledRequest) (*EnrolledListResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	req.normalize()
	offset := (req.Page - 1) * req.Limit

	var total int
	countQuery := `
		SELECT COUNT(*)
		FROM tbl_course_enrollments e
		JOIN tbl_courses c ON c.id = e.course_id AND c.deleted_at IS NULL
		WHERE e.user_id = $1`
	if err := r.dbpool.Get(&total, countQuery, userID); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var rows []EnrolledCourseRow
	listQuery := `
		SELECT
			e.id AS enrollment_id, e.progress, e.enrolled_at, e.completed_at, e.certificate_url,
			c.id, c.title, c.subtitle, c.description, c.thumbnail,
			c.current_price, c.original_price, c.is_free, c.rating, c.level_id, c.promo_text,
			c.preview_video_url,
			t.name AS course_type,
			u.first_name, u.last_name, u.user_name AS instructor_username
		FROM tbl_course_enrollments e
		JOIN tbl_courses c ON c.id = e.course_id AND c.deleted_at IS NULL
		LEFT JOIN tbl_course_types t ON t.id = c.type_id
		LEFT JOIN tbl_users u ON u.id = c.instructor_id
		WHERE e.user_id = $1
		ORDER BY e.enrolled_at DESC
		LIMIT $2 OFFSET $3`
	if err := r.dbpool.Select(&rows, listQuery, userID, req.Limit, offset); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	courses := make([]EnrolledCourseResponse, 0, len(rows))
	for _, row := range rows {
		courses = append(courses, toEnrolledCourseResponse(row))
	}

	return &EnrolledListResponse{
		Courses: courses,
		Total:   total,
		Page:    req.Page,
		Limit:   req.Limit,
	}, nil
}

func (r *CourseEnrollmentRepoImpl) IsEnrolled(userID int64, courseID int64) (bool, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var count int
	err := r.dbpool.Get(&count,
		`SELECT COUNT(*) FROM tbl_course_enrollments WHERE user_id = $1 AND course_id = $2`,
		userID, courseID,
	)
	if err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}
	return count > 0, nil
}

func (r *CourseEnrollmentRepoImpl) CreateEnrollments(userID int64, courseIDs []int64) ([]int64, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	tx, err := r.dbpool.Beginx()
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	insertQuery := `
		INSERT INTO tbl_course_enrollments (user_id, course_id, progress, enrolled_at)
		VALUES ($1, $2, 0, NOW())
		ON CONFLICT (user_id, course_id) DO NOTHING
		RETURNING course_id`

	enrolled := make([]int64, 0, len(courseIDs))
	for _, courseID := range courseIDs {
		var insertedID int64
		err := tx.QueryRowx(insertQuery, userID, courseID).Scan(&insertedID)
		if err != nil && err.Error() != "sql: no rows in result set" {
			return nil, msg.NewErrorResponse("database_error", err)
		}
		if insertedID != 0 {
			enrolled = append(enrolled, insertedID)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return enrolled, nil
}
