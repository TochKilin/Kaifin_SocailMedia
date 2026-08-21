package communities

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

func saveCommunityImage(c fiber.Ctx, fh *multipart.FileHeader, subDir string) (string, error) {
	dir := filepath.Join("uploads", "communities", subDir)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		log.Printf("MkdirAll error: %v", err)
		return "", err
	}

	ext := filepath.Ext(fh.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), randomCommunityHex(8), ext)
	dst := filepath.Join(dir, filename)
	if err := c.SaveFile(fh, dst); err != nil {
		log.Printf("SaveFile error: %v", err)
		return "", err
	}
	return filepath.Join("communities", subDir, filename), nil
}

func randomCommunityHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return hex.EncodeToString(b)
}

type CommunitiesHandlerImpl struct {
	cs CommunitiesServiceImpl
}

func NewCommunitiesHandlerImpl(app *fiber.App, dbpool *sqlx.DB) *CommunitiesHandlerImpl {
	return &CommunitiesHandlerImpl{
		cs: *NewCommunitiesServiceImpl(dbpool),
	}
}

func (h *CommunitiesHandlerImpl) Create(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	req := CreatCommunitiesRequest{
		Name:    strings.TrimSpace(c.FormValue("name")),
		Privacy: strings.TrimSpace(c.FormValue("privacy")),
	}

	if req.Name == "" {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, nil),
		)
	}

	if desc := strings.TrimSpace(c.FormValue("description")); desc != "" {
		req.Description = &desc
	}

	if catStr := c.FormValue("category_id"); catStr != "" {
		catStr = strings.TrimSpace(catStr)
		if catID, err := strconv.ParseInt(catStr, 10, 64); err == nil {
			req.CategoryID = &catID
			log.Printf("SUCCESS PARSED CATEGORY ID: %d", catID)
		} else {
			// បន្ថែម log នេះដើម្បីเช็คថាវា Error អ្វីពី strconv
			log.Printf("ERROR PARSING CATEGORY ID ('%s'): %v", catStr, err)
		}
	}

	if form, ferr := c.MultipartForm(); ferr == nil && form != nil {
		if files := form.File["avatar"]; len(files) > 0 {
			path, serr := saveCommunityImage(c, files[0], "avatars")
			if serr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					response.NewResponseError("Failed to save avatar", constants.Generic_error, serr),
				)
			}
			req.AvatarURL = &path
		}
		if files := form.File["cover"]; len(files) > 0 {
			path, serr := saveCommunityImage(c, files[0], "covers")
			if serr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					response.NewResponseError("Failed to save cover", constants.Generic_error, serr),
				)
			}
			req.CoverURL = &path
		}
	}

	if req.AvatarURL == nil {
		if v := strings.TrimSpace(c.FormValue("avatar_url")); v != "" {
			req.AvatarURL = &v
		}
	}
	if req.CoverURL == nil {
		if v := strings.TrimSpace(c.FormValue("cover_url")); v != "" {
			req.CoverURL = &v
		}
	}

	v := utls.NewValidator()
	if err := req.validateOnly(v); err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, err),
		)
	}

	if e := h.cs.Create(&req, &uCtx); e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	msg, e_msg := translate.TranslateWithError(c, "community_created")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, true),
	)
}

func (h *CommunitiesHandlerImpl) Show(c fiber.Ctx) error {
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		b, _ := json.MarshalIndent(uCtx, "", "  ")
		h.cs.SetUserCtx(&uCtx)
		fmt.Println("jwt_data:", string(b))
	}

	var req ShowCommunitiesRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, err),
		)
	}

	data, e := h.cs.Show(req)
	if e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "communities_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, data),
	)
}

func (h *CommunitiesHandlerImpl) ShowDetail(c fiber.Ctx) error {
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		h.cs.SetUserCtx(&uCtx)
	}

	var req ShowCommunityDetailRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, err),
		)
	}

	data, e := h.cs.ShowDetail(req.ID)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusNotFound).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	msg, e_msg := translate.TranslateWithError(c, "community_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, data),
	)
}

func (h *CommunitiesHandlerImpl) ToggleJoin(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}

	joined, status, e := h.cs.ToggleJoin(id, &uCtx)
	if e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "join_toggled")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, fiber.Map{
			"joined": joined,
			"status": status,
		}),
	)
}

func (h *CommunitiesHandlerImpl) ShowMembers(c fiber.Ctx) error {
	var req ShowMembersRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, err),
		)
	}

	data, e := h.cs.ShowMembers(req)
	if e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "members_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, data),
	)
}
func (h *CommunitiesHandlerImpl) UpdateAvatar(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	communityID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}

	fileHeader, ferr := c.FormFile("avatar")
	if ferr != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, ferr),
		)
	}

	path, serr := saveCommunityImage(c, fileHeader, "avatars")
	if serr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Failed to save avatar", constants.Generic_error, serr),
		)
	}

	if e := h.cs.UpdateAvatar(communityID, path, &uCtx); e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "avatar_updated")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, fiber.Map{
			"avatar_url": path,
		}),
	)
}

func (h *CommunitiesHandlerImpl) UpdateCover(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	communityID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}

	fileHeader, ferr := c.FormFile("cover")
	if ferr != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, ferr),
		)
	}

	path, serr := saveCommunityImage(c, fileHeader, "covers")
	if serr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Failed to save cover", constants.Generic_error, serr),
		)
	}

	if e := h.cs.UpdateCover(communityID, path, &uCtx); e != nil {
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

	msg, e_msg := translate.TranslateWithError(c, "cover_updated")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, fiber.Map{
			"cover_url": path,
		}),
	)
}
