package course

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

const uploadsRoot = "./uploads"

func baseURL() string {
	if v := os.Getenv("APP_BASE_URL"); v != "" {
		return v
	}
	return "http://localhost:7070"
}

func saveUploadedFile(c fiber.Ctx, fieldName string, subDir string) (string, error) {
	fileHeader, err := c.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	dir := filepath.Join(uploadsRoot, subDir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	destPath := filepath.Join(dir, filename)
	if err := c.SaveFile(fileHeader, destPath); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/uploads/%s/%s", baseURL(), subDir, filename), nil
}

type CourseHandlerImpl struct {
	cs *CourseServiceImpl
}

func NewCourseHandlerImpl(dbpool *sqlx.DB, rdb *redis.Client) *CourseHandlerImpl {
	return &CourseHandlerImpl{
		cs: NewCourseServiceImpl(dbpool, rdb),
	}
}

func currentUser(c fiber.Ctx) *share.UserContext {
	raw := c.Locals("UserContext")
	if raw == nil {
		return nil
	}
	userCtx, ok := raw.(share.UserContext)
	if !ok {
		return nil
	}
	return &userCtx
}

func unauthorized(c fiber.Ctx) error {
	msg, err_msg := translate.TranslateWithError(c, "unauthorized")
	if err_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(err_msg.ErrorString(), constants.Translate_Failed, err_msg.Err),
		)
	}
	return c.Status(fiber.StatusUnauthorized).JSON(
		response.NewResponseError(msg, constants.Invalid_request, fmt.Errorf("unauthorized: no user context on request")),
	)
}

func badRequest(c fiber.Ctx, err error) error {
	msg, err_msg := translate.TranslateWithError(c, "invalid_request")
	if err_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(err_msg.ErrorString(), constants.Translate_Failed, err_msg.Err),
		)
	}
	return c.Status(fiber.StatusBadRequest).JSON(
		response.NewResponseError(msg, constants.Invalid_request, err),
	)
}

func respondError(c fiber.Ctx, e *error_responses.ErrorResponse) error {
	msg, e_msg := translate.TranslateWithError(c, e.MessageID)
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusInternalServerError).JSON(
		response.NewResponseError(msg, constants.Generic_error, e.Err),
	)
}

func (h *CourseHandlerImpl) Create(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	var req CreateCourseRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	if thumb, err := saveUploadedFile(c, "thumbnail", "courses"); err == nil {
		req.Thumbnail = thumb
	}

	if preview, err := saveUploadedFile(c, "preview_video", "courses/previews"); err == nil {
		req.PreviewVideoURL = preview
	}

	h.cs.UserCtx = userCtx
	newCourse, e := h.cs.Create(req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "course_created")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, newCourse),
	)
}

func (h *CourseHandlerImpl) Show(c fiber.Ctx) error {
	var req ShowCourseRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	list, e := h.cs.Show(req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "courses_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, list),
	)
}

func (h *CourseHandlerImpl) Update(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}

	var req UpdateCourseRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	if thumb, err := saveUploadedFile(c, "thumbnail", "courses"); err == nil {
		req.Thumbnail = &thumb
	}

	if preview, err := saveUploadedFile(c, "preview_video", "courses/previews"); err == nil {
		req.PreviewVideoURL = &preview
	}

	h.cs.UserCtx = userCtx
	updated, e := h.cs.Update(id, req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "course_updated")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, updated),
	)
}

func (h *CourseHandlerImpl) Delete(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}

	h.cs.UserCtx = userCtx
	if e := h.cs.Delete(id); e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "course_deleted")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}

func (h *CourseHandlerImpl) ShowByID(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}

	detail, e := h.cs.ShowByID(id)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "courses_retrieved") // ← ប្តូរពី "course_retrieved"
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, detail),
	)
}
