package stickers

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type StickersService interface {
	Show(req ShowStickersRequest) (*StickersResponse, *error_responses.ErrorResponse)
	SetUserCtx(ctx *share.UserContext) bool
	Create(req *CreateStickerRequest, fileData []byte, fileName string, fileType string) *error_responses.ErrorResponse
	Update(req *UpdateStickerRequest, fileData []byte, fileName string, fileType string, updateFile bool) *error_responses.ErrorResponse
}

type StickersServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Repo    *StickersRepoImpl
}

func NewStickerServiceImpl(dbpool *sqlx.DB) *StickersServiceImpl {
	return &StickersServiceImpl{
		dbpool: dbpool,
		Repo:   NewStickersRepoImp(dbpool),
	}
}

func (s *StickersServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *StickersServiceImpl) Show(req ShowStickersRequest) (*StickersResponse, *error_responses.ErrorResponse) {
	return s.Repo.Show(req)
}

func (s *StickersServiceImpl) Create(req *CreateStickerRequest, sticker *Sticker) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	if err := sticker.New(req); err != nil {
		return msg.NewErrorResponse("invalid", err)
	}

	return s.Repo.Create(sticker)
}

func (s *StickersServiceImpl) Update(req *UpdateStickerRequest, sticker *Sticker, updateFile bool) *error_responses.ErrorResponse {
	sticker.ID = req.ID
	sticker.PackID = req.PackID
	if req.TriggerCode != "" {
		sticker.TriggerCode = &req.TriggerCode
	}

	return s.Repo.Update(sticker, updateFile)
}
