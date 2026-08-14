package articlecomment

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type CommentsHandlerImpl struct {
	cs CommentsService
}

func NewCommentsHandlerImpl(dbpool *sqlx.DB) *CommentsHandlerImpl {
	return &CommentsHandlerImpl{
		cs: NewCommentsServiceImpl(dbpool),
	}
}

// Create adds a comment to the article identified by the :id path param.
// POST /articles/:id/comments (Supports multipart/form-data for text, images, and stickers)
func (h *CommentsHandlerImpl) Create(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	articleID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid article id", constants.Generic_invalid, err),
		)
	}

	req := &CreateCommentRequest{}
	req.ArticleID = articleID
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	comment := &ArticleComment{}
	comment.new(articleID, req, &uCtx)

	// Save uploaded files to disk first, collect their URLs.
	var savedImageURLs []string
	for _, fh := range req.ImageFiles {
		url, serr := saveCommentImage(c, fh)
		if serr != nil {
			log.Printf("saveCommentImage error: %v", serr)
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError("Failed to save image", constants.Generic_error, serr),
			)
		}
		savedImageURLs = append(savedImageURLs, url)
	}
	comment.SavedImageURLs = savedImageURLs

	if e := h.cs.Create(comment); e != nil {
		log.Printf("Create comment error: %v", e.Err)
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

	// ⚠️ DO NOT do this here — it's what caused the avatar/username to
	// only show up after a page refresh:
	//
	//   comment.UserName = &uCtx.UserName
	//   comment.ProfileImage = &uCtx.UserName   // bug: username copied into the avatar field
	//
	// comment.UserName / comment.ProfileImage are already populated
	// correctly straight from tbl_users inside CommentsRepoImpl.Create()
	// (same source Show() uses), so this response already matches what a
	// follow-up list fetch would return.
	comment.ImageURLs = comment.SavedImageURLs

	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse("comment_created", constants.Generic_success, comment),
	)
}

// Update edits a comment the caller owns.
// PUT /comments/:id
func (h *CommentsHandlerImpl) Update(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid comment id", constants.Generic_invalid, err),
		)
	}

	var req UpdateCommentRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	if e := h.cs.Update(id, uCtx.UserID, req.Text); e != nil {
		log.Printf("Update comment error: %v", e.Err)
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
		response.NewResponse("comment_updated", constants.Generic_success, true),
	)
}

// Delete removes a comment the caller owns.
// DELETE /comments/:id
func (h *CommentsHandlerImpl) Delete(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid comment id", constants.Generic_invalid, err),
		)
	}

	if e := h.cs.Delete(id, uCtx.UserID); e != nil {
		log.Printf("Delete comment error: %v", e.Err)
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
		response.NewResponse("comment_deleted", constants.Generic_success, nil),
	)
}

// Show lists comments for the article identified by the :id path param.
// GET /articles/:id/comments
func (h *CommentsHandlerImpl) Show(c fiber.Ctx) error {
	articleID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid article id", constants.Generic_invalid, err),
		)
	}

	var req ShowCommentsRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}
	req.ArticleID = articleID

	data, e := h.cs.Show(req)
	if e != nil {
		log.Printf("Show comments error: %v", e.Err)
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
		response.NewResponse("comments_retrieved", constants.Generic_success, data),
	)
}
