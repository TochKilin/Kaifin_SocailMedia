package user

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type UserRoute struct {
	handler *UserHandler
	ws      *websocket.WebSocketManager
}

func NewUserRoute(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *UserRoute {
	h := NewUserHandler(dbpool, rdb, ws)
	v1 := app.Group("/api/v1/admin")
	users := v1.Group("/users")
	users.Get("/", h.Show)
	users.Get("/:id", h.ShowOne)
	users.Post("/create", h.Create)
	users.Put("/update/:id", h.Update)
	users.Delete("/delete/:id", h.Delete)

	users.Get("/form/create", h.GetUserFormCreate)
	users.Get("/form/update/:id", h.GetUserFormUpdate)

	return &UserRoute{
		handler: h,
		ws:      ws,
	}

}
