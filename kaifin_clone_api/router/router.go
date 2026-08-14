package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func New() *fiber.App {
	f := fiber.New(fiber.Config{
		BodyLimit:      200 * 1024 * 1024,
		ReadBufferSize: 16 * 1024,
	})
	f.Use(logger.New(logger.Config{}))
	f.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowHeaders: []string{
			"Origins",
			"Content-Type",
			"Accept",
			"Authorization",
			"Accept-language",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"HEAD",
			"PATCH",
			"DELETE",
		},
	}))

	return f
}
