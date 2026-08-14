package articlecomment

import (
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type ArticleComment struct {
	ID              int64     `json:"id" db:"id"`
	ArticleID       int64     `json:"article_id" db:"article_id"`
	UserID          int64     `json:"user_id" db:"user_id"`
	ParentCommentID *int64    `json:"parent_comment_id" db:"parent_comment_id"`
	Text            string    `json:"text" db:"text"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`

	UserName     *string `json:"user_name" db:"user_name"`
	ProfileImage *string `json:"profile_images" db:"profile_images"`

	// Raw comma-separated string ពី string_agg() — scan ត្រង់នេះ
	ImageURLRaw *string `json:"-" db:"image_url"`
	// Parsed array សម្រាប់ JSON response ទៅ frontend
	ImageURLs []string `json:"image_urls" db:"-"`

	StickerURL *string `json:"sticker_url" db:"sticker_url"`
	ImageID    *int64  `json:"image_id" db:"-"`
	StickerID  *int64  `json:"sticker_id" db:"-"`

	ImageFiles     []*multipart.FileHeader `db:"-"`
	StickerIDs     []int64                 `db:"-"`
	SavedImageURLs []string                `db:"-"`
}

// ---------- Create ----------

type CreateCommentRequest struct {
	ArticleID       int64                 `json:"article_id" form:"article_id" validate:"required"`
	ParentCommentID *int64                `json:"parent_comment_id" form:"parent_comment_id"`
	Text            string                `json:"text" form:"text"`
	ImageID         *int64                `json:"image_id" form:"image_id"`
	StickerID       *int64                `json:"sticker_id" form:"sticker_id"`
	ImageFile       *multipart.FileHeader `form:"image_file"` // single-file legacy field, kept for backward compat

	ImageFiles []*multipart.FileHeader `form:"-"`
	StickerIDs []int64                 `form:"-"`
}

func (r *CreateCommentRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	contentType := c.Get("Content-Type")

	if strings.Contains(contentType, "multipart/form-data") {
		if textVal := c.FormValue("text"); textVal != "" {
			r.Text = textVal
		}

		if parentVal := c.FormValue("parent_comment_id"); parentVal != "" {
			if id, err := strconv.ParseInt(parentVal, 10, 64); err == nil {
				r.ParentCommentID = &id
			}
		}

		form, err := c.MultipartForm()
		if err == nil && form != nil {
			if files, ok := form.File["images"]; ok && len(files) > 0 {
				r.ImageFiles = files
			}

			if stickerVals, ok := form.Value["sticker_ids"]; ok && len(stickerVals) > 0 {
				for _, sv := range stickerVals {
					if id, convErr := strconv.ParseInt(sv, 10, 64); convErr == nil {
						r.StickerIDs = append(r.StickerIDs, id)
					}
				}
			}
		}

		// Backward-compat fallback: single-file "image_file" / "image" key
		if len(r.ImageFiles) == 0 {
			file, ferr := c.FormFile("image_file")
			if ferr != nil {
				file, _ = c.FormFile("image")
			}
			if file != nil {
				r.ImageFile = file
				r.ImageFiles = append(r.ImageFiles, file)
			}
		}
	} else {
		if err := c.Bind().Body(r); err != nil {
			return err
		}
	}

	if r.Text == "" && r.ImageID == nil && r.StickerID == nil &&
		r.ImageFile == nil && len(r.ImageFiles) == 0 && len(r.StickerIDs) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Comment cannot be empty")
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

// ---------- Update ----------

type UpdateCommentRequest struct {
	Text string `json:"text" validate:"required,min=1,max=2000"`
}

func (r *UpdateCommentRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

// ---------- Show (list comments for an article) ----------

type ShowCommentsRequest struct {
	ArticleID int64        `query:"-"`
	Page      int          `query:"page"`
	PerPage   int          `query:"per_page"`
	Sorts     []share.Sort `query:"sorts"`
}

func (u *ShowCommentsRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(u); err != nil {
		return err
	}
	if u.Page <= 0 {
		u.Page = 1
	}
	if u.PerPage <= 0 || u.PerPage > 50 {
		u.PerPage = 20
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

type CommentsResponse struct {
	Comments []ArticleComment `json:"comments"`
	Total    int              `json:"total"`
	Page     int              `json:"page"`
	PerPage  int              `json:"per_page"`
}

// ---------- helpers ----------

func (cm *ArticleComment) new(articleID int64, req *CreateCommentRequest, uctx *share.UserContext) {
	cm.ArticleID = articleID
	cm.UserID = uctx.UserID
	cm.ParentCommentID = req.ParentCommentID
	cm.Text = req.Text
	cm.ImageID = req.ImageID
	cm.StickerID = req.StickerID
	cm.ImageFiles = req.ImageFiles
	cm.StickerIDs = req.StickerIDs
}
