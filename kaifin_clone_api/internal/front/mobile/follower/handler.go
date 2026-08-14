package follower_mobile

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type FollowersMobileHandlerImpl struct {
	Service *FollowersMobileServiceImpl
}

func NewFollowersMobileHandlerImpl(dbpool *sqlx.DB) *FollowersMobileHandlerImpl {
	return &FollowersMobileHandlerImpl{
		Service: NewFollowersMobileServiceImpl(dbpool),
	}
}

func (h *FollowersMobileHandlerImpl) setUserCtx(c fiber.Ctx) (share.UserContext, bool) {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if ok {
		h.Service.SetUserCtx(&uCtx)
	}
	return uCtx, ok
}

func (h *FollowersMobileHandlerImpl) Toggle(c fiber.Ctx) error {
	if _, ok := h.setUserCtx(c); !ok {
		msg, _ := translate.TranslateWithError(c, "unauthorized")
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, errors.New("user context not found")),
		)
	}

	req := &ToggleFollowRequest{}
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}

	result, e := h.Service.Toggle(*req)
	// 🟢 ការពារយ៉ាងតឹងរឹង៖ ឆែកមើលក្រែងលោ e មិនមែនជា nil ផ្ទាល់ ប៉ុន្តែជា Object ទទេដែលគ្មាន Error សោះ
	if e != nil && (e.MessageID != "" || e.Err != nil) {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			errStr := "translation error"
			var rawErr error = nil
			if e_msg.Err != nil {
				errStr = e_msg.Err.Error()
				rawErr = e_msg.Err
			}
			if rawErr == nil {
				rawErr = errors.New(errStr)
			}
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(errStr, constants.Translate_Failed, rawErr),
			)
		}

		var serviceErr error = e.Err
		if serviceErr == nil {
			serviceErr = errors.New("unknown service error")
		}

		if msg == "" {
			msg = "bad request error from service"
		}

		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_error, serviceErr),
		)
	}

	msg, e_msg := translate.TranslateWithError(c, "follow_status_updated")
	if e_msg != nil {
		errStr := "translation error"
		var rawErr error = nil
		if e_msg.Err != nil {
			errStr = e_msg.Err.Error()
			rawErr = e_msg.Err
		}
		if rawErr == nil {
			rawErr = errors.New(errStr)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(errStr, constants.Translate_Failed, rawErr),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, result),
	)
}

func (h *FollowersMobileHandlerImpl) Show(c fiber.Ctx) error {
	if _, ok := h.setUserCtx(c); !ok {
		msg, _ := translate.TranslateWithError(c, "unauthorized")
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, errors.New("user context not found")),
		)
	}

	req := &FollowShowRequest{}
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}

	result, e := h.Service.Show(*req)
	// 🟢 ការពារយ៉ាងតឹងរឹង៖ ឆែកមើលក្រែងលោ e មិនមែនជា nil ផ្ទាល់ ប៉ុន្តែជា Object ទទេដែលគ្មាន Error សោះ
	if e != nil && (e.MessageID != "" || e.Err != nil) {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			errStr := "translation error"
			var rawErr error = nil
			if e_msg.Err != nil {
				errStr = e_msg.Err.Error()
				rawErr = e_msg.Err
			}
			if rawErr == nil {
				rawErr = errors.New(errStr)
			}
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(errStr, constants.Translate_Failed, rawErr),
			)
		}

		var serviceErr error = e.Err
		if serviceErr == nil {
			serviceErr = errors.New("unknown service error")
		}

		if msg == "" {
			msg = "internal server error from service"
		}

		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, serviceErr),
		)
	}

	msg, e_msg := translate.TranslateWithError(c, "follow_status_retrieved")
	if e_msg != nil {
		errStr := "translation error"
		var rawErr error = nil
		if e_msg.Err != nil {
			errStr = e_msg.Err.Error()
			rawErr = e_msg.Err
		}
		if rawErr == nil {
			rawErr = errors.New(errStr)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(errStr, constants.Translate_Failed, rawErr),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, result),
	)
}
