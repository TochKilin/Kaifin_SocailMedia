package courseinroll

import "time"

type ShowEnrolledRequest struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

func (r *ShowEnrolledRequest) normalize() {
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.Limit <= 0 {
		r.Limit = 20
	}
}

type CourseEnrollment struct {
	ID             int64      `db:"id"`
	CourseID       int64      `db:"course_id"`
	UserID         int64      `db:"user_id"`
	Progress       int16      `db:"progress"`
	EnrolledAt     time.Time  `db:"enrolled_at"`
	CompletedAt    *time.Time `db:"completed_at"`
	CertificateURL *string    `db:"certificate_url"`
}

type EnrolledCourseRow struct {
	// enrollment fields
	EnrollmentID   int64      `db:"enrollment_id"`
	Progress       int16      `db:"progress"`
	EnrolledAt     time.Time  `db:"enrolled_at"`
	CompletedAt    *time.Time `db:"completed_at"`
	CertificateURL *string    `db:"certificate_url"`

	// course fields
	ID            int64    `db:"id"`
	Title         string   `db:"title"`
	Subtitle      *string  `db:"subtitle"`
	Description   *string  `db:"description"`
	Thumbnail     *string  `db:"thumbnail"`
	CurrentPrice  float64  `db:"current_price"`
	OriginalPrice *float64 `db:"original_price"`
	IsFree        bool     `db:"is_free"`
	Rating        *float64 `db:"rating"`
	LevelID       *int16   `db:"level_id"`
	PromoText     *string  `db:"promo_text"`

	// joined fields
	CourseType          *string `db:"course_type"`
	InstructorFirstName *string `db:"first_name"`
	InstructorLastName  *string `db:"last_name"`
	InstructorUsername  *string `db:"instructor_username"`

	PreviewVideoURL *string `db:"preview_video_url"`
}

type EnrolledCourseResponse struct {
	ID              int64    `json:"id"`
	EnrollmentID    int64    `json:"enrollment_id"`
	Type            string   `json:"type"` // "video" | "book" | "image"
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Instructor      string   `json:"instructor"`
	Level           int16    `json:"level"`
	Price           float64  `json:"price"`
	OldPrice        *float64 `json:"old_price,omitempty"`
	Promo           string   `json:"promo"`
	MediaUrl        string   `json:"media_url"`
	Progress        int16    `json:"progress"`
	EnrolledAt      string   `json:"enrolled_at"`
	CompletedAt     *string  `json:"completed_at,omitempty"`
	CertificateURL  *string  `json:"certificate_url,omitempty"`
	PreviewVideoURL string   `json:"preview_video_url"`
}

type EnrolledListResponse struct {
	Courses []EnrolledCourseResponse `json:"courses"`
	Total   int                      `json:"total"`
	Page    int                      `json:"page"`
	Limit   int                      `json:"limit"`
}

func toEnrolledCourseResponse(row EnrolledCourseRow) EnrolledCourseResponse {
	courseType := "image"
	if row.CourseType != nil && *row.CourseType != "" {
		courseType = *row.CourseType
	}

	instructor := "Instructor"
	if row.InstructorFirstName != nil || row.InstructorLastName != nil {
		name := strOr(row.InstructorFirstName) + " " + strOr(row.InstructorLastName)
		if trimmed := trimSpace(name); trimmed != "" {
			instructor = trimmed
		}
	} else if row.InstructorUsername != nil {
		instructor = *row.InstructorUsername
	}

	var level int16
	if row.LevelID != nil {
		level = *row.LevelID
	}

	var completedAt *string
	if row.CompletedAt != nil {
		s := row.CompletedAt.Format("2006-01-02T15:04:05Z07:00")
		completedAt = &s
	}

	return EnrolledCourseResponse{
		ID:              row.ID,
		EnrollmentID:    row.EnrollmentID,
		Type:            courseType,
		Title:           row.Title,
		Description:     strOr(row.Description),
		Instructor:      instructor,
		Level:           level,
		Price:           row.CurrentPrice,
		OldPrice:        row.OriginalPrice,
		Promo:           strOr(row.PromoText),
		MediaUrl:        strOr(row.Thumbnail),
		Progress:        row.Progress,
		EnrolledAt:      row.EnrolledAt.Format("2006-01-02T15:04:05Z07:00"),
		CompletedAt:     completedAt,
		CertificateURL:  row.CertificateURL,
		PreviewVideoURL: strOr(row.PreviewVideoURL),
	}
}

func strOr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func trimSpace(s string) string {
	start, end := 0, len(s)
	for start < end && s[start] == ' ' {
		start++
	}
	for end > start && s[end-1] == ' ' {
		end--
	}
	return s[start:end]
}

type CreateEnrollmentRequest struct {
	CourseIDs []int64 `json:"course_ids"`
}

type CreateEnrollmentResponse struct {
	Enrolled []int64 `json:"enrolled_course_ids"`
}
