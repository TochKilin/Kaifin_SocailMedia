package userregister

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type UserRegisterServie interface {
	Show(UserShowRequest) (*UserResponse, *error_responses.ErrorResponse)
	Create(req *AuthRegisterRequest) *error_responses.ErrorResponse
	Update(id int64, req *UpdateUserRequest, updatedBy int64) (*User, *error_responses.ErrorResponse)
	SetUserCtx(ctx *share.UserContext) bool
	Delete(id int64, deletedBy int64) *error_responses.ErrorResponse
	Profile(userID int64) (*UserProfileResponse, *error_responses.ErrorResponse)
}

type UserRegisterServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	Repo    UserRegisterRepoImpl
}

func NewUserRegisterServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *UserRegisterServiceImpl {
	return &UserRegisterServiceImpl{
		// UserCtx: userCnt,
		dbpool: dbpool,
		Redis:  rdb,
		Repo:   *NewUserRegisterRepoImpl(dbpool, rdb),
	}
}

func (s *UserRegisterServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *UserRegisterServiceImpl) Create(c fiber.Ctx, req *AuthRegisterRequest) *error_responses.ErrorResponse {
	if err := s.Repo.Create(c, req, s.UserCtx); err != nil {
		return err
	}
	return nil
}

func (s *UserRegisterServiceImpl) Show(userRequest UserShowRequest) (*UserResponse, *error_responses.ErrorResponse) {

	return s.Repo.Show(userRequest)
}

func (s *UserRegisterServiceImpl) Update(c fiber.Ctx, id int64, req *UpdateUserRequest, updatedBy int64) (*User, *error_responses.ErrorResponse) {
	// create update map = null
	updates := map[string]any{}
	// if client send f-name
	if req.FirstName != nil {
		updates["first_name"] = *req.FirstName
	}
	if req.LastName != nil {
		updates["last_name"] = *req.LastName
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.RoleID != nil {
		updates["role_id"] = *req.RoleID
	}
	if req.RoleName != nil {
		updates["role_name"] = *req.RoleName
	}
	if req.StatusID != nil {
		updates["status_id"] = *req.StatusID
	}

	if req.ProfileImages != nil {
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), req.ProfileImages.Filename)
		savePath := "./uploads/" + filename

		if err := c.SaveFile(req.ProfileImages, savePath); err != nil {
			msg := error_responses.ErrorResponse{}
			return nil, msg.NewErrorResponse("upload_failed", err)
		}

		updates["profile_images"] = filename
	}

	// if client send null = 0
	if len(updates) == 0 {
		msg := error_responses.ErrorResponse{}
		return nil, msg.NewErrorResponse("no_updates_provided", fmt.Errorf("empty"))
	}
	// take note update by someone
	updates["updated_by"] = updatedBy
	// call repository to update in db
	return s.Repo.Update(id, updates)
}

func (s *UserRegisterServiceImpl) Delete(id int64, deletedBy int64) *error_responses.ErrorResponse {
	return s.Repo.Delete(id, deletedBy)
}

func (s *UserRegisterServiceImpl) Profile(userID int64) (*UserProfileResponse, *error_responses.ErrorResponse) {
	return s.Repo.Profile(userID)
}
