package main

import (
	//communiy package
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/joho/godotenv"

	config "kaifin_clone_api/configs/database" // internal library
	"kaifin_clone_api/configs/radis"
	"kaifin_clone_api/handler"
	"kaifin_clone_api/internal/admin/websocket"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/router"
)

func main() {
	err := godotenv.Load()
	app := router.New()
	if err := translate.Init(); err != nil {
		log.Fatalf("Failed to initialize i18n Translation: %v", err.Err)
	}

	appConfigs := config.NewConfig()
	ctx := context.Background()
	dbPool, err := config.ConnectPostgres(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer dbPool.Close()
	radis.NewRedisClient()
	rdb := radis.NewRedisClient()

	ws := websocket.NewWebSocketManager()
	app.Get("/uploads*", static.New("./uploads"))
	handler.NewServiceHandler(app, dbPool, rdb, ws)
	log.Println("Server starting...")

	// event loop
	err = app.Listen(
		fmt.Sprintf("%s:%d",
			appConfigs.AppHost,
			appConfigs.AppPort,
		))
	if err != nil {
		log.Fatal(err)
	}

	// judy ediit or modify the pointer of fiber
	// handler.NewServiceHandler(app)
	// app.Listen(":6000")
}
