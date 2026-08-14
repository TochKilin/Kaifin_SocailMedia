package songs

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

// uploadsRoot / baseURL: adjust these to match your actual storage setup.
// This implementation saves to local disk under ./public/uploads and
// assumes that folder is served statically by your app (e.g.
// app.Static("/uploads", "./public/uploads") registered somewhere in
// your server setup). If you use S3 or another provider instead, this
// is the one place that needs to change — everything else (repo/service)
// only ever deals with the final URL string.
const uploadsRoot = "./uploads"

func baseURL() string {
	if v := os.Getenv("APP_BASE_URL"); v != "" {
		return v
	}
	return "http://localhost:7070"
}

// saveUploadedFile reads the multipart file under fieldName, writes it
// to disk under uploadsRoot/subDir, and returns the public URL to store
// in the DB. Returns an error if no file was sent under that field —
// callers decide whether that's fatal (required) or fine (optional).
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

type SongHandlerImpl struct {
	ss *SongServiceImpl
}

func NewSongHandlerImpl(dbpool *sqlx.DB, rdb *redis.Client) *SongHandlerImpl {
	return &SongHandlerImpl{
		ss: NewSongServiceImpl(dbpool, rdb),
	}
}

// currentUser mirrors the post module's helper — pulls the authenticated
// user from c.Locals("UserContext"), set by pkg/middleware/jwt.go.
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

// ---------- Create ----------

func (h *SongHandlerImpl) Create(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	var req CreateSongRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	fileURL, err := saveUploadedFile(c, "file_url", "songs")
	if err != nil {
		return badRequest(c, fmt.Errorf("file_url: an audio file upload is required (%w)", err))
	}
	req.FileURL = fileURL

	// cover_url is optional — only set it if a cover file was actually sent
	if coverURL, err := saveUploadedFile(c, "cover_url", "covers"); err == nil {
		req.CoverURL = coverURL
	}

	h.ss.UserCtx = userCtx
	newSong, e := h.ss.Create(req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "song_created")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, newSong),
	)
}

// ---------- Show / List ----------

func (h *SongHandlerImpl) Show(c fiber.Ctx) error {
	var req ShowSongRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	list, e := h.ss.Show(req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "songs_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, list),
	)
}

// ---------- Update ----------

func (h *SongHandlerImpl) Update(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}

	var req UpdateSongRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	// Only overwrite file_url / cover_url if a new file was actually
	// uploaded with this request — leaving the pointer nil means
	// repo.Update()'s COALESCE keeps whatever is already in the DB.
	if fileURL, err := saveUploadedFile(c, "file_url", "songs"); err == nil {
		req.FileURL = &fileURL
	}
	if coverURL, err := saveUploadedFile(c, "cover_url", "covers"); err == nil {
		req.CoverURL = &coverURL
	}

	h.ss.UserCtx = userCtx
	updated, e := h.ss.Update(id, req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "song_updated")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, updated),
	)
}

// ---------- Delete ----------

func (h *SongHandlerImpl) Delete(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}

	h.ss.UserCtx = userCtx
	if e := h.ss.Delete(id); e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "song_deleted")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}
