package course

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type CourseRouteImpl struct {
	Handler *CourseHandlerImpl
	ws      *websocket.WebSocketManager
}

func NewCourseRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *CourseRouteImpl {
	handler := NewCourseHandlerImpl(dbpool, rdb)
	route := &CourseRouteImpl{
		Handler: handler,
		ws:      ws,
	}

	courses := app.Group("/api/v1/front/courses")
	courses.Get("/show", handler.Show)
	courses.Post("/create", handler.Create)
	courses.Put("/update/:id", handler.Update)
	courses.Delete("/delete/:id", handler.Delete)
	courses.Get("/show/:id", handler.ShowByID)

	return route
}
