package quoteshare

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type QuoteShareHandler struct {
	Repo QuoteShareRepo
}

func NewQuoteShareHandler(dbpool *sqlx.DB) *QuoteShareHandler {
	return &QuoteShareHandler{Repo: NewQuoteShareRepoImpl(dbpool)}
}

func (h *QuoteShareHandler) Track(c fiber.Ctx) error {
	req := &TrackShareRequest{}
	v := utls.NewValidator()
	if err := c.Bind().Body(req); err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}
	if err := v.Validate(req); err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}

	var userID *int64
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		userID = &uCtx.UserID
	}

	if e := h.Repo.Track(req.QuoteID, userID, req.Channel); e != nil {
		msg, _ := translate.TranslateWithError(c, e.MessageID)
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	msg, _ := translate.TranslateWithError(c, "share_tracked")
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}
