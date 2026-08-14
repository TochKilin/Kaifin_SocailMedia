package mystickerset

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
)

type MysetStickerRouteImpl struct {
	hh     *MysetStickerHandlerImpl
	dbpool *sqlx.DB
}

func NewMysetStickerRouteImpl(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *MysetStickerRouteImpl {
	h := NewMysetStickerHanslerImpl(dbpool, rdb)
	r := &MysetStickerRouteImpl{
		hh:     h,
		dbpool: dbpool,
	}
	stickergroup := app.Group("/api/v1/front/stickers")
	stickergroup.Get("/packs", h.ListPacks)
	stickergroup.Get("/show", h.ShowPackStickers)
	stickergroup.Post("/create", h.CreateSticker)
	stickergroup.Get("/my-sets", h.Show)
	stickergroup.Delete("/my-sets/:pack_id", h.Delete)
	stickergroup.Post("/my-sets/:pack_id", h.Create)
	stickergroup.Get("/image/:id", h.ServeImage)
	stickergroup.Post("/packs", h.CreatePack)

	return r
}
