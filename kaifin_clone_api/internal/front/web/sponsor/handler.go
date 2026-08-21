package sponsor

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/utls"
)

func saveSponsorLogo(c fiber.Ctx, fh *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(fh.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), randomHex(8), ext)
	dst := filepath.Join("uploads", "sponsors", filename)
	if err := c.SaveFile(fh, dst); err != nil {
		log.Printf("SaveFile error: %v", err)
		return "", err
	}
	return filepath.Join("sponsors", filename), nil
}

func randomHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return hex.EncodeToString(b)
}

type SponsorsHandlerImpl struct {
	ss SponsorsServiceImpl
}

func NewSponsorsHandlerImpl(dbpool *sqlx.DB) *SponsorsHandlerImpl {
	return &SponsorsHandlerImpl{
		ss: *NewSponsorsServiceImpl(dbpool),
	}
}

func (h *SponsorsHandlerImpl) Create(c fiber.Ctx) error {
	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Name is required", constants.Generic_invalid, nil),
		)
	}

	sortOrder, _ := strconv.Atoi(c.FormValue("sort_order"))
	isVerified := c.FormValue("is_verified") == "true"

	req := CreateSponsorRequest{
		Name:       name,
		WebsiteURL: c.FormValue("website_url"),
		IsVerified: isVerified,
		SortOrder:  sortOrder,
	}

	form, ferr := c.MultipartForm()
	if ferr != nil || form == nil || len(form.File["logo_image"]) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("logo_image file is required", constants.Generic_invalid, nil),
		)
	}
	logoPath, serr := saveSponsorLogo(c, form.File["logo_image"][0])
	if serr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Failed to save logo image", constants.Generic_error, serr),
		)
	}

	v := utls.NewValidator()
	if verr := v.Validate(&req); verr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, verr),
		)
	}

	sp, e := h.ss.Create(&req, logoPath)
	if e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Failed to create sponsor", constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse("sponsor_created", constants.Generic_success, sp),
	)
}

func (h *SponsorsHandlerImpl) Update(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid sponsor id", constants.Generic_invalid, err),
		)
	}

	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Name is required", constants.Generic_invalid, nil),
		)
	}

	sortOrder, _ := strconv.Atoi(c.FormValue("sort_order"))
	isVerified := c.FormValue("is_verified") == "true"
	isActive := c.FormValue("is_active") == "true"

	req := UpdateSponsorRequest{
		Name:       name,
		WebsiteURL: c.FormValue("website_url"),
		IsVerified: isVerified,
		IsActive:   isActive,
		SortOrder:  sortOrder,
	}

	logoPath := ""
	if form, ferr := c.MultipartForm(); ferr == nil && form != nil {
		if files := form.File["logo_image"]; len(files) > 0 {
			path, serr := saveSponsorLogo(c, files[0])
			if serr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					response.NewResponseError("Failed to save logo image", constants.Generic_error, serr),
				)
			}
			logoPath = path
		}
	}

	if e := h.ss.Update(id, &req, logoPath); e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Failed to update sponsor", constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("sponsor_updated", constants.Generic_success, true),
	)
}

func (h *SponsorsHandlerImpl) Delete(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid sponsor id", constants.Generic_invalid, err),
		)
	}

	if e := h.ss.Delete(id); e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Failed to delete sponsor", constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("sponsor_deleted", constants.Generic_success, nil),
	)
}

func (h *SponsorsHandlerImpl) Show(c fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 || perPage > 50 {
		perPage = 20
	}

	data, e := h.ss.Show(ShowSponsorsRequest{Page: page, PerPage: perPage})
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Failed to fetch sponsors", constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("sponsors_retrieved", constants.Generic_success, data),
	)
}
