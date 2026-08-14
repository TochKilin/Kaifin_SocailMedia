package template

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/utls"
)

type ShowTemplateRequest struct {
	// មិនត្រូវការ field ណាមួយសម្រាប់ list templates ទាំងអស់
}

type Template struct {
	ID           int64     `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	ThumbnailURL string    `json:"thumbnail_url" db:"-"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type TemplatesResponse struct {
	Templates []Template `json:"templates"`
}

func (u *ShowTemplateRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(u); err != nil {
		return err
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}
