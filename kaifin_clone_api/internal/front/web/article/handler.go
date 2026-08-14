package article

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"regexp"
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

// saveArticleImage stores an uploaded cover image under uploads/articles
// and returns the relative path to store in tbl_articles.cover_image.
func saveArticleImage(c fiber.Ctx, fh *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(fh.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), randomHex(8), ext)
	dst := filepath.Join("uploads", "articles", filename)
	if err := c.SaveFile(fh, dst); err != nil {
		log.Printf("SaveFile error: %v", err)
		return "", err
	}
	return filepath.Join("articles", filename), nil
}

func randomHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		// fallback ដែលមិនប្រើ crypto/rand ក្នុងករណីកម្រកើតមាន error
		return strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return hex.EncodeToString(b)
}

// parseTagsForm splits a comma-separated "tags" form field into a slice.
func parseTagsForm(raw string) []string {
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	tags := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			tags = append(tags, p)
		}
	}
	return tags
}

// parseBlocksForm decodes a JSON-encoded "blocks" form field, e.g.
// [{"block_type":"text","title":"Intro","content":"..."},{"block_type":"image","content":"https://..."}]
func parseBlocksForm(raw string) ([]ArticleBlockInput, error) {
	if raw == "" {
		return nil, nil
	}
	var blocks []ArticleBlockInput
	if err := json.Unmarshal([]byte(raw), &blocks); err != nil {
		return nil, err
	}
	return blocks, nil
}

// ============================================================
// ➕ NEW (Option B): auto-derive cover_image from the first image
// found in the article's blocks, when no cover file was uploaded.
//
// This covers two cases:
//  1. block_type == "image"  -> Content is already an image URL/path.
//  2. block_type == "text"   -> Content is WYSIWYG HTML that may contain
//     an inline <img src="..."> inserted via the editor's image toolbar.
// ============================================================

var imgSrcRegex = regexp.MustCompile(`<img[^>]+src=["']([^"']+)["']`)

// extractCoverImageFromBlocks scans blocks in order and returns a value
// suitable for tbl_articles.cover_image, or "" if no image is found.
func extractCoverImageFromBlocks(blocks []ArticleBlockInput) string {
	for _, b := range blocks {
		var src string

		switch strings.ToLower(strings.TrimSpace(b.BlockType)) {
		case "image":
			src = strings.TrimSpace(b.Content)
		default:
			if m := imgSrcRegex.FindStringSubmatch(b.Content); len(m) == 2 {
				src = strings.TrimSpace(m[1])
			}
		}

		if src == "" {
			continue
		}
		return normalizeCoverImagePath(src)
	}
	return ""
}

// normalizeCoverImagePath makes an inline image URL match the same
// relative-path format saveArticleImage() already produces (e.g.
// "articles/169..._ab12cd34.png"), regardless of whether the source was:
//   - a full URL:      http://localhost:7070/uploads/articles/x.png
//   - a rooted path:    /uploads/articles/x.png
//   - already relative: articles/x.png
func normalizeCoverImagePath(src string) string {
	if idx := strings.Index(src, "://"); idx != -1 {
		rest := src[idx+3:]
		if slash := strings.Index(rest, "/"); slash != -1 {
			src = rest[slash:]
		} else {
			src = ""
		}
	}
	src = strings.TrimPrefix(src, "/uploads/")
	src = strings.TrimPrefix(src, "uploads/")
	src = strings.TrimPrefix(src, "/")
	return src
}

type ArticlesHandlerImpl struct {
	as ArticlesServiceImpl
}

func NewArticlesHandlerImpl(dbpool *sqlx.DB) *ArticlesHandlerImpl {
	return &ArticlesHandlerImpl{
		as: *NewArticlesServiceImpl(dbpool),
	}
}

func (h *ArticlesHandlerImpl) Create(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	title := strings.TrimSpace(c.FormValue("title"))
	if title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Title is required", constants.Generic_invalid, nil),
		)
	}

	req := CreateArticleRequest{
		Title:           title,
		Summary:         c.FormValue("summary"),
		Category:        c.FormValue("category"),
		CodeSubcategory: c.FormValue("code_subcategory"),
		Visibility:      c.FormValue("visibility"),
		Tags:            parseTagsForm(c.FormValue("tags")),
	}

	blocks, err := parseBlocksForm(c.FormValue("blocks"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid blocks JSON", constants.Generic_invalid, err),
		)
	}
	req.Blocks = blocks

	// Explicit cover file upload (if the client ever sends one) still wins.
	if form, ferr := c.MultipartForm(); ferr == nil && form != nil {
		if files := form.File["cover_image"]; len(files) > 0 {
			path, serr := saveArticleImage(c, files[0])
			if serr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					response.NewResponseError("Failed to save cover image", constants.Generic_error, serr),
				)
			}
			req.CoverImage = path
		}
	}

	// ➕ NEW (Option B): no explicit cover was uploaded — fall back to the
	// first image found in the article's blocks.
	if req.CoverImage == "" {
		if auto := extractCoverImageFromBlocks(blocks); auto != "" {
			req.CoverImage = auto
		}
	}

	v := utls.NewValidator()
	if verr := v.Validate(&req); verr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, verr),
		)
	}

	article, e := h.as.Create(&req, &uCtx)
	if e != nil {
		log.Printf("Create article error: %v", e.Err)
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

	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse("article_created", constants.Generic_success, article),
	)
}

func (h *ArticlesHandlerImpl) Update(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid article id", constants.Generic_invalid, err),
		)
	}

	title := strings.TrimSpace(c.FormValue("title"))
	if title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Title is required", constants.Generic_invalid, nil),
		)
	}

	req := UpdateArticleRequest{
		Title:           title,
		Summary:         c.FormValue("summary"),
		Category:        c.FormValue("category"),
		CodeSubcategory: c.FormValue("code_subcategory"),
		Visibility:      c.FormValue("visibility"),
		Tags:            parseTagsForm(c.FormValue("tags")),
	}

	blocks, berr := parseBlocksForm(c.FormValue("blocks"))
	if berr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid blocks JSON", constants.Generic_invalid, berr),
		)
	}
	req.Blocks = blocks

	// CoverImage stays empty when no new file is uploaded — the repository
	// keeps the existing cover_image in that case (see COALESCE in Update query),
	// UNLESS we can auto-derive a new one from the updated blocks below.
	if form, ferr := c.MultipartForm(); ferr == nil && form != nil {
		if files := form.File["cover_image"]; len(files) > 0 {
			path, serr := saveArticleImage(c, files[0])
			if serr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					response.NewResponseError("Failed to save cover image", constants.Generic_error, serr),
				)
			}
			req.CoverImage = path
		}
	}

	// ➕ NEW (Option B): no explicit cover was uploaded on update either —
	// fall back to the first image found in the (possibly edited) blocks.
	// If you'd rather KEEP the old cover untouched when the user didn't
	// change it, remove this block and rely on the repository's COALESCE.
	if req.CoverImage == "" {
		if auto := extractCoverImageFromBlocks(blocks); auto != "" {
			req.CoverImage = auto
		}
	}

	v := utls.NewValidator()
	if verr := v.Validate(&req); verr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, verr),
		)
	}

	if e := h.as.Update(id, &req, &uCtx); e != nil {
		log.Printf("Update article error: %v", e.Err)
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

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("article_updated", constants.Generic_success, true),
	)
}

func (h *ArticlesHandlerImpl) Show(c fiber.Ctx) error {
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		h.as.SetUserCtx(&uCtx)
	}

	var req ShowArticlesRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	data, e := h.as.Show(req)
	if e != nil {
		log.Printf("Show articles error: %v", e.Err)
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

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("articles_retrieved", constants.Generic_success, data),
	)
}

func (h *ArticlesHandlerImpl) Detail(c fiber.Ctx) error {
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		h.as.SetUserCtx(&uCtx)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid article id", constants.Generic_invalid, err),
		)
	}

	article, e := h.as.Detail(id)
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

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("article_retrieved", constants.Generic_success, article),
	)
}

func (h *ArticlesHandlerImpl) Delete(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid article id", constants.Generic_invalid, err),
		)
	}

	if e := h.as.Delete(id, &uCtx); e != nil {
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

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("article_deleted", constants.Generic_success, nil),
	)
}

func (h *ArticlesHandlerImpl) ToggleLike(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid article id", constants.Generic_invalid, err),
		)
	}

	liked, e := h.as.ToggleLike(id, &uCtx)
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

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("like_toggled", constants.Generic_success, fiber.Map{"liked": liked}),
	)
}

func (h *ArticlesHandlerImpl) ToggleSave(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid article id", constants.Generic_invalid, err),
		)
	}

	saved, e := h.as.ToggleSave(id, &uCtx)
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

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("save_toggled", constants.Generic_success, fiber.Map{"saved": saved}),
	)
}

func (h *ArticlesHandlerImpl) Report(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	var req ReportArticleRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	if e := h.as.Report(&req, &uCtx); e != nil {
		log.Printf("Report article error: %v", e.Err)
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

	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse("article_reported", constants.Generic_success, true),
	)
}

// UploadImage handles inline image uploads from the article editor toolbar.
// Returns the accessible URL so the frontend can insert it into the markdown content.
func (h *ArticlesHandlerImpl) UploadImage(c fiber.Ctx) error {
	_, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	fh, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Image file is required", constants.Generic_invalid, err),
		)
	}

	path, serr := saveArticleImage(c, fh)
	if serr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Failed to save image", constants.Generic_error, serr),
		)
	}

	// កែ base URL នេះតាម config ពិតរបស់អ្នក (ឧ. static file server / CDN)
	url := "/uploads/" + path

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("image_uploaded", constants.Generic_success, fiber.Map{
			"url": url,
		}),
	)
}
