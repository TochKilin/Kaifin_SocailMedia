package profile_mobile

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
)

type ProfileMobileHandlerImpl struct {
	se *ProfileMobileServiceImpl
}

func NewProfileMobileHandlerImpl(
	db *sqlx.DB,
) *ProfileMobileHandlerImpl {

	return &ProfileMobileHandlerImpl{
		se: NewProfileServiceImpl(db),
	}
}

func (h *ProfileMobileHandlerImpl) Profile(c fiber.Ctx) error {
	userCtx, ok := c.Locals(
		"UserContext",
	).(share.UserContext)
	if !ok {
		return c.Status(401).JSON(
			response.NewResponseError(
				"invalid user context",
				constants.Login_failed,
				fmt.Errorf("user context missing"),
			),
		)
	}

	profile, errRes := h.se.GetProfile(
		userCtx.UserID,
	)

	if errRes != nil {
		return c.Status(400).JSON(
			response.NewResponseError(
				errRes.MessageID,
				constants.Login_failed,
				errRes,
			),
		)
	}

	return c.Status(200).JSON(
		response.NewResponse(
			"profile fetched successfully",
			constants.Login_success,
			profile,
		),
	)
}

func (h *ProfileMobileHandlerImpl) Update(c fiber.Ctx) error {
	userCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(401).JSON(
			response.NewResponseError(
				"invalid user context",
				constants.Login_failed,
				fmt.Errorf("user context missing"),
			),
		)
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(
			response.NewResponseError(
				"file is required",
				constants.Invalid_request,
				err,
			),
		)
	}

	ext := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("profile_%d_%d%s", userCtx.UserID, time.Now().UnixNano(), ext)

	savePath := filepath.Join("uploads", filename)
	if err := c.SaveFile(fileHeader, savePath); err != nil {
		return c.Status(500).JSON(
			response.NewResponseError(
				"failed to save file",
				constants.Generic_error,
				err,
			),
		)
	}

	if errRes := h.se.Update(userCtx.UserID, filename); errRes != nil {
		return c.Status(400).JSON(
			response.NewResponseError(
				errRes.MessageID,
				constants.Generic_error,
				errRes,
			),
		)
	}

	profile, errRes := h.se.GetProfile(userCtx.UserID)
	if errRes != nil {
		return c.Status(400).JSON(
			response.NewResponseError(
				errRes.MessageID,
				constants.Generic_error,
				errRes,
			),
		)
	}

	return c.Status(200).JSON(
		response.NewResponse(
			"profile image updated successfully",
			constants.Generic_success,
			profile,
		),
	)
}
