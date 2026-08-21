package notificationbell

import (
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v3"

	wsManager "kaifin_clone_api/internal/admin/websocket"
	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type NotificationsHandlerImpl struct {
	ns *NotificationsServiceImpl
	ws *wsManager.WebSocketManager
}

func NewNotificationsHandlerImpl(svc *NotificationsServiceImpl, ws *wsManager.WebSocketManager) *NotificationsHandlerImpl {
	return &NotificationsHandlerImpl{ns: svc, ws: ws}
}

func (h *NotificationsHandlerImpl) Show(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}
	h.ns.SetUserCtx(&uCtx)

	data, e := h.ns.Show()
	if e != nil {
		msg, _ := translate.TranslateWithError(c, e.MessageID)
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	msg, _ := translate.TranslateWithError(c, "notifications_retrieved")
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, data),
	)
}

func (h *NotificationsHandlerImpl) MarkAsRead(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, nil),
		)
	}
	h.ns.SetUserCtx(&uCtx)

	var req MarkAsReadRequest
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_request")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Invalid_request, err),
		)
	}

	if e := h.ns.MarkAsRead(req.NotificationID); e != nil {
		msg, _ := translate.TranslateWithError(c, e.MessageID)
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	msg, _ := translate.TranslateWithError(c, "notification_marked_read")
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, true),
	)
}

// WebSocket upgrade handler — client តភ្ជាប់ជាមួយ ?token=xxx
func (h *NotificationsHandlerImpl) HandleWS(c *websocket.Conn) {
	userIDStr, ok := c.Locals("wsUserID").(string)
	if !ok || userIDStr == "" {
		c.Close()
		return
	}

	client := &wsManager.Client{
		Conn: c,
		ID:   "user-" + userIDStr,
	}
	h.ws.AddClient(client)
	defer h.ws.RemoveClient(client.ID)

	for {
		if _, _, err := c.ReadMessage(); err != nil {
			break
		}
	}
}

var _ = strconv.Itoa // keep import if unused elsewhere
