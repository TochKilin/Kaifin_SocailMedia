package playlist

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

// Same local-disk upload approach as the song module — see the NOTE
// there for swapping to S3/cloud storage if that's what you actually use.
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

type PlaylistHandlerImpl struct {
	ps *PlaylistServiceImpl
}

func NewPlaylistHandlerImpl(dbpool *sqlx.DB, rdb *redis.Client) *PlaylistHandlerImpl {
	return &PlaylistHandlerImpl{
		ps: NewPlaylistServiceImpl(dbpool, rdb),
	}
}

// currentUser mirrors the post/song modules' helper — pulls the
// authenticated user from c.Locals("UserContext"), set by
// pkg/middleware/jwt.go.
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

func (h *PlaylistHandlerImpl) Create(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	var req CreatePlaylistRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	// cover_url is optional — only set it if a cover file was actually sent
	if coverURL, err := saveUploadedFile(c, "cover_url", "playlists"); err == nil {
		req.CoverURL = coverURL
	}

	h.ps.UserCtx = userCtx
	newPlaylist, e := h.ps.Create(req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "playlist_created")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, newPlaylist),
	)
}

// ---------- Show / List ----------

func (h *PlaylistHandlerImpl) Show(c fiber.Ctx) error {
	var req ShowPlaylistRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	h.ps.UserCtx = currentUser(c) // public endpoint — nil is fine

	list, e := h.ps.Show(req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "playlists_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, list),
	)
}

// ---------- Top ----------
// Public endpoint — this is what the "Top Playlist" section on the
// homepage calls.

func (h *PlaylistHandlerImpl) Top(c fiber.Ctx) error {
	var req TopPlaylistRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	top, e := h.ps.Top(req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "playlists_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, fiber.Map{"playlists": top}),
	)
}

// ---------- Update ----------

func (h *PlaylistHandlerImpl) Update(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}

	var req UpdatePlaylistRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	// Only overwrite cover_url if a new file was actually uploaded with
	// this request — leaving the pointer nil means repo.Update()'s
	// COALESCE keeps whatever is already in the DB.
	if coverURL, err := saveUploadedFile(c, "cover_url", "playlists"); err == nil {
		req.CoverURL = &coverURL
	}

	h.ps.UserCtx = userCtx
	updated, e := h.ps.Update(id, req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "playlist_updated")
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

func (h *PlaylistHandlerImpl) Delete(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}

	h.ps.UserCtx = userCtx
	if e := h.ps.Delete(id); e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "playlist_deleted")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}

// ---------- Add song ----------
// POST /api/v1/front/playlists/:id/songs — used by the "Add to
// playlist" button, adds one song without replacing the whole list.

func (h *PlaylistHandlerImpl) AddSong(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	playlistID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}

	var req AddPlaylistSongRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	h.ps.UserCtx = userCtx
	if e := h.ps.AddSong(playlistID, req); e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "playlist_song_added")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}

// ---------- Remove song ----------
// DELETE /api/v1/front/playlists/:id/songs/:songId

func (h *PlaylistHandlerImpl) RemoveSong(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	playlistID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}
	songID, err := strconv.ParseInt(c.Params("songId"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}

	h.ps.UserCtx = userCtx
	if e := h.ps.RemoveSong(playlistID, songID); e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "playlist_song_removed")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}

func (h *PlaylistHandlerImpl) ShowDetail(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}

	h.ps.UserCtx = currentUser(c)

	detail, e := h.ps.ShowDetail(id)
	if e != nil {
		fmt.Println("🔴 ShowDetail error:", e.Err)
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "playlists_retrieved") // ← ប្តូរត្រង់នេះ
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, detail),
	)
}
