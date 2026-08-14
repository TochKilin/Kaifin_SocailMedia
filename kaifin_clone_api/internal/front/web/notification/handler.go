package notification

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	wsManager "kaifin_clone_api/internal/admin/websocket"
	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type NotificationsHandlerImpl struct {
	Service *NotificationsServiceImpl
	Ws      *wsManager.WebSocketManager
}

func NewNotificationsHandlerImpl(dbpool *sqlx.DB, ws *wsManager.WebSocketManager) *NotificationsHandlerImpl {
	return &NotificationsHandlerImpl{
		Service: NewNotificationsServiceImpl(dbpool, ws),
		Ws:      ws,
	}
}

func (h *NotificationsHandlerImpl) setUserCtx(c fiber.Ctx) (share.UserContext, bool) {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if ok {
		h.Service.SetUserCtx(&uCtx)
	}
	return uCtx, ok
}

func (h *NotificationsHandlerImpl) GetList(c fiber.Ctx) error {
	if _, ok := h.setUserCtx(c); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, errors.New("unauthorized")),
		)
	}

	result, err := h.Service.GetList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Failed to fetch notifications", constants.Generic_error, err.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("notifications_retrieved", constants.Generic_success, result),
	)
}

func (h *NotificationsHandlerImpl) Read(c fiber.Ctx) error {
	if _, ok := h.setUserCtx(c); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, errors.New("unauthorized")),
		)
	}

	req := &ReadNotificationRequest{}
	v := utls.NewValidator()
	if err := req.Bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request", constants.Invalid_request, err),
		)
	}

	if err := h.Service.Read(req.NotificationID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("Failed to update notification", constants.Generic_error, err.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("notification_marked_read", constants.Generic_success, true),
	)
}
