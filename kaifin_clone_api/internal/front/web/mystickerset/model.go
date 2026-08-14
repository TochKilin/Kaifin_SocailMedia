package mystickerset

import "time"

type MyStickerSet struct {
	PackID   int64           `json:"pack_id" db:"pack_id"`
	PackName string          `json:"pack_name" db:"pack_name"`
	AddedAt  time.Time       `json:"added_at" db:"added_at"`
	Stickers []StickerInPack `json:"stickers"`
}

type StickerPack struct {
	ID           int64  `json:"id" db:"id"`
	Name         string `json:"name" db:"pack_name"`
	ThumbnailURL string `json:"thumbnail_url" db:"-"`
	StickerCount int    `json:"sticker_count" db:"sticker_count"`
}

type StickerPacksResponse struct {
	Packs []StickerPack `json:"packs"`
}

type StickerListResponse struct {
	PackID   int64           `json:"pack_id"`
	Stickers []StickerInPack `json:"stickers"`
}

type StickerInPack struct {
	ID  int64  `json:"id" db:"id"`
	URL string `json:"url" db:"-"`
}

type MyStickerSetsResponse struct {
	Sets []MyStickerSet `json:"sets"`
}

type createPackRequest struct {
	Name string `json:"name" form:"name"`
}
