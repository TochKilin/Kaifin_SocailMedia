package course

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/utls"
)

type Course struct {
	ID            int64      `json:"id" db:"id"`
	Title         string     `json:"title" db:"title"`
	Subtitle      string     `json:"subtitle" db:"subtitle"`
	Description   string     `json:"description" db:"description"`
	Thumbnail     string     `json:"thumbnail" db:"thumbnail"`
	CategoryID    *int64     `json:"category_id" db:"category_id"`
	TypeID        *int64     `json:"type_id" db:"type_id"`
	LevelID       *int64     `json:"level_id" db:"level_id"`
	InstructorID  int64      `json:"instructor_id" db:"instructor_id"`
	CurrentPrice  float64    `json:"current_price" db:"current_price"`
	OriginalPrice float64    `json:"original_price" db:"original_price"`
	IsFree        bool       `json:"is_free" db:"is_free"`
	Rating        float64    `json:"rating" db:"rating"`
	RatingsCount  int        `json:"ratings_count" db:"ratings_count"`
	StudentsCount int        `json:"students_count" db:"students_count"`
	SectionsCount int        `json:"sections_count" db:"sections_count"`
	LecturesCount int        `json:"lectures_count" db:"lectures_count"`
	TotalLength   string     `json:"total_length" db:"total_length"`
	PromoText     string     `json:"promo_text" db:"promo_text"`
	ContentType   string     `json:"content_type" db:"content_type"`
	CreatedBy     int64      `json:"created_by" db:"created_by"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at"`

	PreviewVideoURL string `json:"preview_video_url" db:"preview_video_url"`
}

type CourseResponse struct {
	ID            int64      `json:"id"`
	Title         string     `json:"title"`
	Subtitle      string     `json:"subtitle"`
	Description   string     `json:"description"`
	Thumbnail     string     `json:"thumbnail"`
	CategoryID    *int64     `json:"category_id"`
	TypeID        *int64     `json:"type_id"`
	LevelID       *int64     `json:"level_id"`
	InstructorID  int64      `json:"instructor_id"`
	CurrentPrice  float64    `json:"current_price"`
	OriginalPrice float64    `json:"original_price"`
	IsFree        bool       `json:"is_free"`
	Rating        float64    `json:"rating"`
	RatingsCount  int        `json:"ratings_count"`
	StudentsCount int        `json:"students_count"`
	SectionsCount int        `json:"sections_count"`
	LecturesCount int        `json:"lectures_count"`
	TotalLength   string     `json:"total_length"`
	PromoText     string     `json:"promo_text"`
	ContentType   string     `json:"content_type"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`

	PreviewVideoURL string `json:"preview_video_url"`
}

type CourseListResponse struct {
	Courses []CourseResponse `json:"courses"`
	Total   int              `json:"total"`
	Page    int              `json:"page"`
	Limit   int              `json:"limit"`
}

type CreateCourseRequest struct {
	Title         string  `json:"title" form:"title" validate:"required"`
	Subtitle      string  `json:"subtitle" form:"subtitle"`
	Description   string  `json:"description" form:"description"`
	Thumbnail     string  `json:"thumbnail"`
	CategoryID    *int64  `json:"category_id" form:"category_id"`
	TypeID        *int64  `json:"type_id" form:"type_id"`
	LevelID       *int64  `json:"level_id" form:"level_id"`
	CurrentPrice  float64 `json:"current_price" form:"current_price"`
	OriginalPrice float64 `json:"original_price" form:"original_price"`
	IsFree        bool    `json:"is_free" form:"is_free"`
	TotalLength   string  `json:"total_length" form:"total_length"`
	PromoText     string  `json:"promo_text" form:"promo_text"`
	ContentType   string  `json:"content_type" form:"content_type"`

	// បន្ថែម Fields ទាំងនេះ៖
	StudentsCount   int     `json:"students_count" form:"students_count"`
	SectionsCount   int     `json:"sections_count" form:"sections_count"`
	LecturesCount   int     `json:"lectures_count" form:"lectures_count"`
	Rating          float64 `json:"rating" form:"rating"`
	PreviewVideoURL string  `json:"preview_video_url"`
}

func (r *CreateCourseRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type UpdateCourseRequest struct {
	Title         *string  `json:"title" form:"title"`
	Subtitle      *string  `json:"subtitle" form:"subtitle"`
	Description   *string  `json:"description" form:"description"`
	Thumbnail     *string  `json:"thumbnail"`
	CategoryID    *int64   `json:"category_id" form:"category_id"`
	TypeID        *int64   `json:"type_id" form:"type_id"`
	LevelID       *int64   `json:"level_id" form:"level_id"`
	CurrentPrice  *float64 `json:"current_price" form:"current_price"`
	OriginalPrice *float64 `json:"original_price" form:"original_price"`
	IsFree        *bool    `json:"is_free" form:"is_free"`
	TotalLength   *string  `json:"total_length" form:"total_length"`
	PromoText     *string  `json:"promo_text" form:"promo_text"`
	ContentType   *string  `json:"content_type" form:"content_type"`

	PreviewVideoURL *string `json:"preview_video_url"`
}

func (r *UpdateCourseRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type ShowCourseRequest struct {
	Search     string `query:"search"`
	CategoryID int64  `query:"category_id"`
	LevelID    int64  `query:"level_id"`
	IsFree     string `query:"is_free"`
	Page       int    `query:"page"`
	Limit      int    `query:"limit"`
}

func (r *ShowCourseRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(r); err != nil {
		return err
	}
	if r.Search == "" {
		r.Search = c.Query("q")
	}
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.Limit <= 0 || r.Limit > 200 {
		r.Limit = 200
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

func toCourseResponse(c Course) CourseResponse {
	return CourseResponse{
		ID:              c.ID,
		Title:           c.Title,
		Subtitle:        c.Subtitle,
		Description:     c.Description,
		Thumbnail:       c.Thumbnail,
		CategoryID:      c.CategoryID,
		TypeID:          c.TypeID,
		LevelID:         c.LevelID,
		InstructorID:    c.InstructorID,
		CurrentPrice:    c.CurrentPrice,
		OriginalPrice:   c.OriginalPrice,
		IsFree:          c.IsFree,
		Rating:          c.Rating,
		RatingsCount:    c.RatingsCount,
		StudentsCount:   c.StudentsCount,
		SectionsCount:   c.SectionsCount,
		LecturesCount:   c.LecturesCount,
		TotalLength:     c.TotalLength,
		PromoText:       c.PromoText,
		ContentType:     c.ContentType,
		CreatedAt:       c.CreatedAt,
		UpdatedAt:       c.UpdatedAt,
		PreviewVideoURL: c.PreviewVideoURL,
	}
}

type CourseDetailResponse struct {
	CourseResponse
	InstructorName     string `json:"instructor_name"`
	InstructorAvatar   string `json:"instructor_avatar"`
	InstructorHeadline string `json:"instructor_headline"`
	InstructorBio      string `json:"instructor_bio"`
}

type CourseDetailRow struct {
	Course
	InstructorFirstName     *string  `db:"instructor_first_name"`
	InstructorLastName      *string  `db:"instructor_last_name"`
	InstructorUsername      *string  `db:"instructor_username"`
	InstructorAvatar        *string  `db:"instructor_avatar"`
	InstructorHeadline      *string  `db:"instructor_headline"`
	InstructorDescription   *string  `db:"instructor_description"`
	InstructorRating        *float64 `db:"instructor_rating"`
	InstructorReviewsCount  *int     `db:"instructor_reviews_count"`
	InstructorStudentsCount *int     `db:"instructor_students_count"`
	InstructorCoursesCount  *int     `db:"instructor_courses_count"`
}

type CourseSection struct {
	ID            int64  `db:"id" json:"id"`
	CourseID      int64  `db:"course_id" json:"course_id"`
	Title         string `db:"title" json:"title"`
	LecturesCount int    `db:"lectures_count" json:"lectures_count"`
	Length        string `db:"length" json:"length"`
	SortOrder     int16  `db:"sort_order" json:"sort_order"`
}

type CourseLecture struct {
	ID        int64  `db:"id" json:"id"`
	SectionID int64  `db:"section_id" json:"section_id"`
	Title     string `db:"title" json:"title"`
	Duration  string `db:"duration" json:"duration"`
	VideoURL  string `db:"video_url" json:"video_url"`
	SortOrder int16  `db:"sort_order" json:"sort_order"`
}

type CourseReview struct {
	ID             int64     `db:"id" json:"id"`
	CourseID       int64     `db:"course_id" json:"course_id"`
	UserID         int64     `db:"user_id" json:"user_id"`
	Rating         int16     `db:"rating" json:"rating"`
	Comment        string    `db:"comment" json:"comment"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	ReviewerName   string    `db:"reviewer_name" json:"reviewer_name"`
	ReviewerAvatar string    `db:"reviewer_avatar" json:"reviewer_avatar"`
}

type CourseInclude struct {
	ID        int64  `db:"id" json:"id"`
	CourseID  int64  `db:"course_id" json:"course_id"`
	Icon      string `db:"icon" json:"icon"`
	Text      string `db:"text" json:"text"`
	SortOrder int16  `db:"sort_order" json:"sort_order"`
}

type SectionWithLectures struct {
	ID            int64           `json:"id"`
	Title         string          `json:"title"`
	LecturesCount int             `json:"lectures_count"`
	Length        string          `json:"length"`
	Lectures      []CourseLecture `json:"lectures"`
}

type InstructorInfo struct {
	Name          string  `json:"name"`
	Avatar        string  `json:"avatar"`
	Headline      string  `json:"headline"`
	Description   string  `json:"description"`
	Rating        float64 `json:"rating"`
	ReviewsCount  int     `json:"reviews_count"`
	StudentsCount int     `json:"students_count"`
	CoursesCount  int     `json:"courses_count"`
}

type CourseFullDetailResponse struct {
	CourseResponse
	Instructor InstructorInfo        `json:"instructor"`
	Sections   []SectionWithLectures `json:"sections"`
	Reviews    []CourseReview        `json:"reviews"`
	Includes   []CourseInclude       `json:"includes"`
}

func strOr(p *string, fallback string) string {
	if p == nil || *p == "" {
		return fallback
	}
	return *p
}
func f64Or(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}
func intOr(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}
