package mystickerset

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type MysetStickerService interface {
}

type MysetStickerServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	repo    *MysetStickerRepoImpl
}

func NewMysetStickerServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *MysetStickerServiceImpl {
	return &MysetStickerServiceImpl{
		dbpool: dbpool,
		Redis:  rdb,
		repo:   NewMysetStickerRepoImpl(dbpool),
	}
}

func (s *MysetStickerServiceImpl) ShowMySets(userID int64) (*MyStickerSetsResponse, *error_responses.ErrorResponse) {
	return s.repo.Show(userID)
}
func (s *MysetStickerServiceImpl) RemoveMySet(packID, userID int64) *error_responses.ErrorResponse {
	return s.repo.Delete(packID, userID)
}
func (s *MysetStickerServiceImpl) AddPackToMySets(packID, userID int64) *error_responses.ErrorResponse {
	return s.repo.Create(packID, userID)
}
func (s *MysetStickerServiceImpl) GetStickerImage(id int64) ([]byte, string, error) {
	return s.repo.GetStickerImage(id)
}

func (s *MysetStickerServiceImpl) ListPacks() (*StickerPacksResponse, *error_responses.ErrorResponse) {
	return s.repo.ListPacks()
}
func (s *MysetStickerServiceImpl) ListStickersByPack(packID int64) (*StickerListResponse, *error_responses.ErrorResponse) {
	return s.repo.ListStickersByPack(packID)
}
func (s *MysetStickerServiceImpl) CreateSticker(packID int64, fileData []byte, fileType string) *error_responses.ErrorResponse {
	return s.repo.CreateSticker(packID, fileData, fileType)
}

func (s *MysetStickerServiceImpl) CreatePack(name string, userID int64) (int64, *error_responses.ErrorResponse) {
	return s.repo.CreatePack(name, userID)
}
