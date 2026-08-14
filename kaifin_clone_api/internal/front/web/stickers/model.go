package stickers

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/utls"
)

type Sticker struct {
	ID          int64     `json:"id" db:"id"`
	PackID      int64     `json:"pack_id" db:"pack_id"`
	FileName    string    `json:"file_name" db:"file_name"`
	FileType    string    `json:"file_type" db:"file_type"`
	FileData    []byte    `json:"-" db:"file_data"`
	TriggerCode *string   `json:"trigger_code" db:"trigger_code"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	URL         string    `json:"url"`
}

type CreateStickerRequest struct {
	PackID      int64  `form:"pack_id"`
	TriggerCode string `form:"trigger_code"`
}

type ShowStickersRequest struct {
	PackID int64   `query:"pack_id"`
	IDs    []int64 `query:"-"`
}

type UpdateStickerRequest struct {
	ID          int64  `form:"id" validate:"required"`
	PackID      int64  `form:"pack_id"`
	TriggerCode string `form:"trigger_code"`
}

func (r *CreateStickerRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *ShowStickersRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(r); err != nil {
		return err
	}

	if idsParam := c.Query("ids"); idsParam != "" {
		for _, s := range strings.Split(idsParam, ",") {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			if id, err := strconv.ParseInt(s, 10, 64); err == nil {
				r.IDs = append(r.IDs, id)
			}
		}
	}

	if r.PackID == 0 && len(r.IDs) == 0 {
		msg := error_responses.ErrorResponse{}
		return msg.NewErrorResponse(
			"invalid_request",
			fmt.Errorf("either pack_id or ids is required"),
		).Err
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type StickersResponse struct {
	Stickers []Sticker `json:"stickers"`
	Total    int       `json:"total"`
}

func (s *Sticker) New(req *CreateStickerRequest) error {
	s.PackID = req.PackID
	if req.TriggerCode != "" {
		s.TriggerCode = &req.TriggerCode
	}
	return nil
}

func (r *UpdateStickerRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

func (s *Sticker) FormatBase64URL() {
	base64Str := base64.StdEncoding.EncodeToString(s.FileData)
	s.URL = fmt.Sprintf("data:%s;base64,%s", s.FileType, base64Str)
}
