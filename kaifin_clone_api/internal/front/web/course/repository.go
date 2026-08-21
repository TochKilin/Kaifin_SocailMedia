package course

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CourseRepo interface {
	Create(instructorID int64, req CreateCourseRequest) (*CourseResponse, *error_responses.ErrorResponse)
	Show(req ShowCourseRequest) (*CourseListResponse, *error_responses.ErrorResponse)
	ShowByID(id int64) (*CourseResponse, *error_responses.ErrorResponse)
	GetByID(id int64) (*Course, *error_responses.ErrorResponse)
	Update(id int64, instructorID int64, req UpdateCourseRequest) (*CourseResponse, *error_responses.ErrorResponse)
	Delete(id int64, instructorID int64) *error_responses.ErrorResponse
	GetFullDetail(id int64) (*CourseFullDetailResponse, *error_responses.ErrorResponse)
}

type CourseRepoImpl struct {
	dbpool *sqlx.DB
}

func NewCourseRepoImpl(db *sqlx.DB) *CourseRepoImpl {
	return &CourseRepoImpl{
		dbpool: db,
	}
}

func (r *CourseRepoImpl) Create(instructorID int64, req CreateCourseRequest) (*CourseResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var newCourse Course
	query := `
		INSERT INTO tbl_courses (
			title, subtitle, description, thumbnail,
			category_id, type_id, level_id, instructor_id,
			current_price, original_price, is_free,
			total_length, promo_text, content_type,
			students_count, sections_count, lectures_count, rating,
			preview_video_url,
			created_by
		)
		VALUES (
			$1, $2, $3, $4,
			$5, $6, $7, $8,
			$9, $10, $11,
			$12, $13, $14,
			$15, $16, $17, $18,
			$19,
			$8
		)
		RETURNING
			id, title, subtitle, description, thumbnail,
			category_id, type_id, level_id, instructor_id,
			current_price, original_price, is_free,
			rating, ratings_count, students_count, sections_count, lectures_count,
			total_length, promo_text, content_type,
			preview_video_url,
			created_by, created_at, updated_at`

	err := r.dbpool.QueryRowx(
		query,
		req.Title, req.Subtitle, req.Description, req.Thumbnail,
		req.CategoryID, req.TypeID, req.LevelID, instructorID,
		req.CurrentPrice, req.OriginalPrice, req.IsFree,
		req.TotalLength, req.PromoText, req.ContentType,
		req.StudentsCount, req.SectionsCount, req.LecturesCount, req.Rating,
		req.PreviewVideoURL,
	).StructScan(&newCourse)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	resp := toCourseResponse(newCourse)
	return &resp, nil
}

func (r *CourseRepoImpl) Show(req ShowCourseRequest) (*CourseListResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	offset := (req.Page - 1) * req.Limit

	baseQuery := `
		FROM tbl_courses c
		WHERE c.deleted_at IS NULL
		  AND ($1 = '' OR c.title ILIKE '%' || $1 || '%')
		  AND ($2 = 0 OR c.category_id = $2)
		  AND ($3 = 0 OR c.level_id = $3)
		  AND ($4 = '' OR c.is_free = ($4 = 'true'))`

	var total int
	countQuery := `SELECT COUNT(*) ` + baseQuery
	if err := r.dbpool.Get(&total, countQuery, req.Search, req.CategoryID, req.LevelID, req.IsFree); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var list []Course
	listQuery := `
		SELECT
			c.id, c.title, c.subtitle, c.description, c.thumbnail,
			c.category_id, c.type_id, c.level_id, c.instructor_id,
			c.current_price, c.original_price, c.is_free,
			c.rating, c.ratings_count, c.students_count, c.sections_count, c.lectures_count,
			c.total_length, c.promo_text, c.content_type,
			c.preview_video_url,
			c.created_by, c.created_at, c.updated_at
		` + baseQuery + `
		ORDER BY c.created_at DESC
		LIMIT $5 OFFSET $6`
	if err := r.dbpool.Select(&list, listQuery, req.Search, req.CategoryID, req.LevelID, req.IsFree, req.Limit, offset); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	responses := make([]CourseResponse, 0, len(list))
	for _, c := range list {
		responses = append(responses, toCourseResponse(c))
	}

	return &CourseListResponse{
		Courses: responses,
		Total:   total,
		Page:    req.Page,
		Limit:   req.Limit,
	}, nil
}

func (r *CourseRepoImpl) ShowByID(id int64) (*CourseResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var c Course
	query := `
		SELECT
			id, title, subtitle, description, thumbnail,
			category_id, type_id, level_id, instructor_id,
			current_price, original_price, is_free,
			rating, ratings_count, students_count, sections_count, lectures_count,
			total_length, promo_text, content_type,
			preview_video_url,
			created_by, created_at, updated_at
		FROM tbl_courses
		WHERE id = $1 AND deleted_at IS NULL`
	if err := r.dbpool.Get(&c, query, id); err != nil {
		return nil, msg.NewErrorResponse("course_not_found", err)
	}

	resp := toCourseResponse(c)
	return &resp, nil
}

func (r *CourseRepoImpl) GetByID(id int64) (*Course, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var c Course
	query := `
		SELECT
			id, title, subtitle, description, thumbnail,
			category_id, type_id, level_id, instructor_id,
			current_price, original_price, is_free,
			rating, ratings_count, students_count, sections_count, lectures_count,
			total_length, promo_text, content_type,
			preview_video_url,
			created_by, created_at, updated_at
		FROM tbl_courses
		WHERE id = $1 AND deleted_at IS NULL`
	if err := r.dbpool.Get(&c, query, id); err != nil {
		return nil, msg.NewErrorResponse("course_not_found", err)
	}
	return &c, nil
}

func (r *CourseRepoImpl) Update(id int64, instructorID int64, req UpdateCourseRequest) (*CourseResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var updated Course
	query := `
		UPDATE tbl_courses SET
			title              = COALESCE($1, title),
			subtitle           = COALESCE($2, subtitle),
			description        = COALESCE($3, description),
			thumbnail          = COALESCE($4, thumbnail),
			category_id        = COALESCE($5, category_id),
			type_id            = COALESCE($6, type_id),
			level_id           = COALESCE($7, level_id),
			current_price      = COALESCE($8, current_price),
			original_price     = COALESCE($9, original_price),
			is_free            = COALESCE($10, is_free),
			total_length       = COALESCE($11, total_length),
			promo_text         = COALESCE($12, promo_text),
			content_type       = COALESCE($13, content_type),
			preview_video_url  = COALESCE($14, preview_video_url),
			updated_at         = NOW()
		WHERE id = $15 AND created_by = $16 AND deleted_at IS NULL
		RETURNING
			id, title, subtitle, description, thumbnail,
			category_id, type_id, level_id, instructor_id,
			current_price, original_price, is_free,
			rating, ratings_count, students_count, sections_count, lectures_count,
			total_length, promo_text, content_type,
			preview_video_url,
			created_by, created_at, updated_at`

	err := r.dbpool.QueryRowx(
		query,
		req.Title, req.Subtitle, req.Description, req.Thumbnail,
		req.CategoryID, req.TypeID, req.LevelID,
		req.CurrentPrice, req.OriginalPrice, req.IsFree,
		req.TotalLength, req.PromoText, req.ContentType,
		req.PreviewVideoURL,
		id, instructorID,
	).StructScan(&updated)
	if err != nil {

		return nil, msg.NewErrorResponse("course_not_found_or_forbidden", err)
	}

	resp := toCourseResponse(updated)
	return &resp, nil
}

func (r *CourseRepoImpl) Delete(id int64, instructorID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	result, err := r.dbpool.Exec(
		`UPDATE tbl_courses SET deleted_at = NOW(), deleted_by = $1
		 WHERE id = $2 AND created_by = $1 AND deleted_at IS NULL`,
		instructorID, id,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	if rows == 0 {
		return msg.NewErrorResponse("course_not_found_or_forbidden", fmt.Errorf("course %d not found for instructor %d", id, instructorID))
	}
	return nil
}

func (r *CourseRepoImpl) GetFullDetail(id int64) (*CourseFullDetailResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var row CourseDetailRow
	courseQuery := `
		SELECT
			c.id, c.title, c.subtitle, c.description, c.thumbnail,
			c.category_id, c.type_id, c.level_id, c.instructor_id,
			c.current_price, c.original_price, c.is_free,
			c.rating, c.ratings_count, c.students_count, c.sections_count, c.lectures_count,
			c.total_length, c.promo_text, c.content_type,
			c.preview_video_url,
			c.created_by, c.created_at, c.updated_at,
			u.first_name  AS instructor_first_name,
			u.last_name   AS instructor_last_name,
			u.user_name   AS instructor_username,
			u.profile_images AS instructor_avatar,
			ip.headline     AS instructor_headline,
			ip.description  AS instructor_description,
			ip.rating       AS instructor_rating,
			ip.reviews_count  AS instructor_reviews_count,
			ip.students_count AS instructor_students_count,
			ip.courses_count  AS instructor_courses_count
		FROM tbl_courses c
		LEFT JOIN tbl_users u ON u.id = c.instructor_id
		LEFT JOIN tbl_instructor_profiles ip ON ip.user_id = c.instructor_id
		WHERE c.id = $1 AND c.deleted_at IS NULL`

	if err := r.dbpool.Get(&row, courseQuery, id); err != nil {
		return nil, msg.NewErrorResponse("course_not_found", err)
	}

	var sections []CourseSection
	sectionsQuery := `
		SELECT id, course_id, title, lectures_count, length, sort_order
		FROM tbl_course_sections
		WHERE course_id = $1
		ORDER BY sort_order, id`
	if err := r.dbpool.Select(&sections, sectionsQuery, id); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var lectures []CourseLecture
	lecturesQuery := `
		SELECT l.id, l.section_id, l.title, l.duration, l.video_url, l.sort_order
		FROM tbl_course_lectures l
		JOIN tbl_course_sections s ON s.id = l.section_id
		WHERE s.course_id = $1
		ORDER BY s.sort_order, l.sort_order, l.id`
	if err := r.dbpool.Select(&lectures, lecturesQuery, id); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	lecturesBySection := make(map[int64][]CourseLecture)
	for _, l := range lectures {
		lecturesBySection[l.SectionID] = append(lecturesBySection[l.SectionID], l)
	}

	sectionsOut := make([]SectionWithLectures, 0, len(sections))
	for _, s := range sections {
		sectionsOut = append(sectionsOut, SectionWithLectures{
			ID:            s.ID,
			Title:         s.Title,
			LecturesCount: s.LecturesCount,
			Length:        s.Length,
			Lectures:      lecturesBySection[s.ID],
		})
	}

	var reviews []CourseReview
	reviewsQuery := `
		SELECT
			r.id, r.course_id, r.user_id, r.rating, r.comment, r.created_at,
			COALESCE(NULLIF(TRIM(u.first_name || ' ' || u.last_name), ''), u.user_name) AS reviewer_name,
			COALESCE(u.profile_images, '') AS reviewer_avatar
		FROM tbl_course_reviews r
		LEFT JOIN tbl_users u ON u.id = r.user_id
		WHERE r.course_id = $1
		ORDER BY r.created_at DESC
		LIMIT 20`
	if err := r.dbpool.Select(&reviews, reviewsQuery, id); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var includes []CourseInclude
	includesQuery := `
		SELECT id, course_id, icon, text, sort_order
		FROM tbl_course_includes
		WHERE course_id = $1
		ORDER BY sort_order, id`
	if err := r.dbpool.Select(&includes, includesQuery, id); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	// assemble
	resp := CourseFullDetailResponse{
		CourseResponse: toCourseResponse(row.Course),
		Instructor: InstructorInfo{
			Name:          strOr(row.InstructorFirstName, "") + " " + strOr(row.InstructorLastName, ""),
			Avatar:        strOr(row.InstructorAvatar, ""),
			Headline:      strOr(row.InstructorHeadline, ""),
			Description:   strOr(row.InstructorDescription, ""),
			Rating:        f64Or(row.InstructorRating),
			ReviewsCount:  intOr(row.InstructorReviewsCount),
			StudentsCount: intOr(row.InstructorStudentsCount),
			CoursesCount:  intOr(row.InstructorCoursesCount),
		},
		Sections: sectionsOut,
		Reviews:  reviews,
		Includes: includes,
	}

	if row.InstructorFirstName == nil && row.InstructorLastName == nil {
		resp.Instructor.Name = strOr(row.InstructorUsername, "Instructor")
	}

	return &resp, nil
}
