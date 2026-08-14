package profile

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
)

type ProfileHandlerImpl struct {
	se *ProfileServiceImpl
}

func NewProfileHandlerImpl(
	db *sqlx.DB,
) *ProfileHandlerImpl {

	return &ProfileHandlerImpl{
		se: NewProfileServiceImpl(db),
	}
}

// func (h *ProfileHandlerImpl) Profile(c fiber.Ctx) error {
// 	userCtx, ok := c.Locals(
// 		"UserContext",
// 	).(share.UserContext)
// 	if !ok {
// 		return c.Status(401).JSON(
// 			response.NewResponseError(
// 				"invalid user context",
// 				constants.Login_failed,
// 				fmt.Errorf("user context missing"),
// 			),
// 		)
// 	}

// 	profile, errRes := h.se.GetProfile(
// 		userCtx.UserID,
// 	)

// 	if errRes != nil {
// 		return c.Status(400).JSON(
// 			response.NewResponseError(
// 				errRes.MessageID,
// 				constants.Login_failed,
// 				errRes,
// 			),
// 		)
// 	}

// 	return c.Status(200).JSON(
// 		response.NewResponse(
// 			"profile fetched successfully",
// 			constants.Login_success,
// 			profile,
// 		),
// 	)
// }

func (h *ProfileHandlerImpl) Profile(c fiber.Ctx) error {
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

	idParam := c.Query("id")
	fmt.Println("🔍 [BACKEND] c.Query(\"id\"):", idParam, "| userCtx.UserID:", userCtx.UserID)

	// ✅ អាន id ពី query string; បើគ្មាន ប្រើ user ដែល login ជា default
	targetUserID := userCtx.UserID
	if idParam := c.Query("id"); idParam != "" {
		if parsedID, err := strconv.ParseInt(idParam, 10, 64); err == nil {
			targetUserID = parsedID
		}
	}

	profile, errRes := h.se.GetProfile(
		targetUserID, // ✅ ប្រើ target ជំនួស userCtx.UserID ជានិច្ច
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

func (h *ProfileHandlerImpl) Update(c fiber.Ctx) error {
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

var allowedCoverExtensions = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
	// .svg មិនអនុញ្ញាត ដោយសារ XSS risk (SVG អាចផ្ទុក script tag)
}

const maxCoverFileSize = 5 * 1024 * 1024 // 5MB

func (h *ProfileHandlerImpl) UpdateCover(c fiber.Ctx) error {
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

	file, err := c.FormFile("cover")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("no_file_uploaded", constants.Generic_invalid, err),
		)
	}

	// Validate extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedCoverExtensions[ext] {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(
				fmt.Sprintf("file type %s not allowed, use png/jpg/jpeg/gif", ext),
				constants.Generic_invalid,
				fmt.Errorf("invalid extension: %s", ext),
			),
		)
	}

	// Validate size
	if file.Size > maxCoverFileSize {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(
				"file too large, max 5MB",
				constants.Generic_invalid,
				fmt.Errorf("file size %d exceeds limit", file.Size),
			),
		)
	}

	uniqueName := fmt.Sprintf("cover_%d_%d%s", userCtx.UserID, time.Now().UnixNano(), ext)
	savePath := filepath.Join("uploads", uniqueName)

	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("upload_failed", constants.Generic_error, err),
		)
	}

	if errRes := h.se.UpdateCover(userCtx.UserID, uniqueName); errRes != nil {
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
			"cover image updated successfully",
			constants.Generic_success,
			profile,
		),
	)
}

func timeNowUnix() int64 {
	return time.Now().UnixNano()
}

func (h *ProfileHandlerImpl) UpdateInfo(c fiber.Ctx) error {
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

	req := &UpdateProfileInfoRequest{}
	if err := c.Bind().Body(req); err != nil {
		return c.Status(400).JSON(
			response.NewResponseError("invalid request body", constants.Invalid_request, err),
		)
	}

	if errRes := h.se.UpdateInfo(userCtx.UserID, req); errRes != nil {
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
			"profile info updated successfully",
			constants.Generic_success,
			profile,
		),
	)
}
