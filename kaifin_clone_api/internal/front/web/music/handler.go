package music

import (
	"fmt"
	"strconv"

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

type PostMusicHandlerImpl struct {
	ps *PostMusicServiceImpl
}

func NewPostHandlerImpl(dbpool *sqlx.DB, rdb *redis.Client) *PostMusicHandlerImpl {
	return &PostMusicHandlerImpl{
		ps: NewPostServiceImpl(dbpool, rdb),
	}
}

// currentUser pulls the authenticated user out of c.Locals.
//
// Matches pkg/middleware/jwt.go's handleUserContext, which stores the
// claims as a *value* (not pointer) under the key "UserContext":
//
//	c.Locals("UserContext", uCtx) // uCtx is types.UserContext
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

// unauthorized is mostly a defensive fallback: the routes in this module
// aren't in middleware's publicPaths list, so the JWT middleware already
// rejects unauthenticated requests before they ever reach this handler.
// This only fires if currentUser() somehow can't read what the
// middleware set (e.g. a type/key mismatch), so it's still worth
// keeping — just never expected to trigger in normal use.
func unauthorized(c fiber.Ctx) error {
	msg, err_msg := translate.TranslateWithError(c, "unauthorized")
	if err_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(err_msg.ErrorString(), constants.Translate_Failed, err_msg.Err),
		)
	}
	// NEVER pass nil as the error argument here — pkg/http/errors.go:30
	// dereferences it directly and panics on nil, which previously took
	// the whole process down instead of just failing this request.
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

func (h *PostMusicHandlerImpl) Create(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	var req CreatePostRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	h.ps.UserCtx = userCtx
	newPost, e := h.ps.Create(req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "post_created")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, newPost),
	)
}

// ---------- Show / List ----------

func (h *PostMusicHandlerImpl) Show(c fiber.Ctx) error {
	var req ShowPostRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	// Public endpoint: an anonymous request still works, it just only
	// sees audience = 'everyone' posts (see PostServiceImpl.Show).
	h.ps.UserCtx = currentUser(c)

	list, e := h.ps.Show(req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "posts_retrieved")
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

func (h *PostMusicHandlerImpl) Update(c fiber.Ctx) error {
	userCtx := currentUser(c)
	if userCtx == nil {
		return unauthorized(c)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return badRequest(c, err)
	}

	var req UpdatePostRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return badRequest(c, err)
	}

	h.ps.UserCtx = userCtx
	updated, e := h.ps.Update(id, req)
	if e != nil {
		return respondError(c, e)
	}

	msg, e_msg := translate.TranslateWithError(c, "post_updated")
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

func (h *PostMusicHandlerImpl) Delete(c fiber.Ctx) error {
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

	msg, e_msg := translate.TranslateWithError(c, "post_deleted")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}
