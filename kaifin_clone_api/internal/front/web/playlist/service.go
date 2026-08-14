package playlist

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type PlaylistService interface {
	Create(req CreatePlaylistRequest) (*PlaylistResponse, *error_responses.ErrorResponse)
	Show(req ShowPlaylistRequest) (*PlaylistListResponse, *error_responses.ErrorResponse)
	Top(req TopPlaylistRequest) ([]PlaylistResponse, *error_responses.ErrorResponse)
	Update(id int64, req UpdatePlaylistRequest) (*PlaylistResponse, *error_responses.ErrorResponse)
	Delete(id int64) *error_responses.ErrorResponse
	AddSong(playlistID int64, req AddPlaylistSongRequest) *error_responses.ErrorResponse
	RemoveSong(playlistID int64, songID int64) *error_responses.ErrorResponse
	ShowDetail(id int64) (*PlaylistDetailResponse, *error_responses.ErrorResponse)
}

type PlaylistServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	Repo    *PlaylistRepoImpl
}

func NewPlaylistServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *PlaylistServiceImpl {
	return &PlaylistServiceImpl{
		dbpool: dbpool,
		Redis:  rdb,
		Repo:   NewPlaylistRepoImpl(dbpool),
	}
}

func (s *PlaylistServiceImpl) Create(req CreatePlaylistRequest) (*PlaylistResponse, *error_responses.ErrorResponse) {
	return s.Repo.Create(s.UserCtx.UserID, req)
}

func (s *PlaylistServiceImpl) Show(req ShowPlaylistRequest) (*PlaylistListResponse, *error_responses.ErrorResponse) {
	var requesterID int64
	if s.UserCtx != nil {
		requesterID = s.UserCtx.UserID
	}
	return s.Repo.Show(requesterID, req)
}

func (s *PlaylistServiceImpl) Top(req TopPlaylistRequest) ([]PlaylistResponse, *error_responses.ErrorResponse) {
	return s.Repo.Top(req.Limit)
}

func (s *PlaylistServiceImpl) Update(id int64, req UpdatePlaylistRequest) (*PlaylistResponse, *error_responses.ErrorResponse) {
	return s.Repo.Update(id, s.UserCtx.UserID, req)
}

func (s *PlaylistServiceImpl) Delete(id int64) *error_responses.ErrorResponse {
	return s.Repo.Delete(id, s.UserCtx.UserID)
}

func (s *PlaylistServiceImpl) AddSong(playlistID int64, req AddPlaylistSongRequest) *error_responses.ErrorResponse {
	return s.Repo.AddSong(playlistID, s.UserCtx.UserID, req.SongID)
}

func (s *PlaylistServiceImpl) RemoveSong(playlistID int64, songID int64) *error_responses.ErrorResponse {
	return s.Repo.RemoveSong(playlistID, s.UserCtx.UserID, songID)
}

func (s *PlaylistServiceImpl) ShowDetail(id int64) (*PlaylistDetailResponse, *error_responses.ErrorResponse) {
	return s.Repo.GetDetailByID(id)
}
