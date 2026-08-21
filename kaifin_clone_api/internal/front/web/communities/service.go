package communities

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type CommunitiesService interface {
	Create(req *CreatCommunitiesRequest, uctx *share.UserContext) *error_responses.ErrorResponse
	Show(req ShowCommunitiesRequest) (*CommunitiesResponse, *error_responses.ErrorResponse)
	ShowDetail(id int64) (*Communities, *error_responses.ErrorResponse)
	SetUserCtx(ctx *share.UserContext) bool
	ToggleJoin(communityID int64, uctx *share.UserContext) (bool, string, *error_responses.ErrorResponse)
	ShowMembers(req ShowMembersRequest) (*MembersResponse, *error_responses.ErrorResponse)
	UpdateAvatar(communityID int64, avatarURL string, uctx *share.UserContext) *error_responses.ErrorResponse // ➜ ថ្មី
	UpdateCover(communityID int64, coverURL string, uctx *share.UserContext) *error_responses.ErrorResponse
}

type CommunitiesServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Repo    *CommunitiesRepoImpl
}

func NewCommunitiesServiceImpl(dbpool *sqlx.DB) *CommunitiesServiceImpl {
	return &CommunitiesServiceImpl{
		dbpool: dbpool,
		Repo:   NewCommunitiesRepoImpl(dbpool),
	}
}

func (s *CommunitiesServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *CommunitiesServiceImpl) Create(req *CreatCommunitiesRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	community := Communities{
		Name:        req.Name,
		Description: req.Description,
		AvatarURL:   req.AvatarURL,
		CoverURL:    req.CoverURL,
		CategoryID:  req.CategoryID,
		Privacy:     req.Privacy,
		CreatedBy:   &uctx.UserID,
	}
	if e := s.Repo.Create(&community); e != nil {
		return msg.NewErrorResponse(e.MessageID, e.Err)
	}
	return nil
}

func (s *CommunitiesServiceImpl) Show(req ShowCommunitiesRequest) (*CommunitiesResponse, *error_responses.ErrorResponse) {
	var userID int64
	if s.UserCtx != nil {
		userID = s.UserCtx.UserID
	}
	return s.Repo.Show(req, userID)
}

func (s *CommunitiesServiceImpl) ShowDetail(id int64) (*Communities, *error_responses.ErrorResponse) {
	var userID int64
	if s.UserCtx != nil {
		userID = s.UserCtx.UserID
	}
	return s.Repo.ShowDetail(id, userID)
}

func (s *CommunitiesServiceImpl) ToggleJoin(communityID int64, uctx *share.UserContext) (bool, string, *error_responses.ErrorResponse) {
	return s.Repo.ToggleJoin(communityID, uctx.UserID)
}

func (s *CommunitiesServiceImpl) ShowMembers(req ShowMembersRequest) (*MembersResponse, *error_responses.ErrorResponse) {
	return s.Repo.ShowMembers(req)
}

func (s *CommunitiesServiceImpl) UpdateAvatar(communityID int64, avatarURL string, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.UpdateAvatar(communityID, avatarURL)
}

func (s *CommunitiesServiceImpl) UpdateCover(communityID int64, coverURL string, uctx *share.UserContext) *error_responses.ErrorResponse {
	return s.Repo.UpdateCover(communityID, coverURL)
}
