package user

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type UserService interface {
	Show(UserShowRequest) (*UserResponse, *error_responses.ErrorResponse)
	ShowOne(id int64) (*UserResponse, *error_responses.ErrorResponse)
	Create(req *CreateUserRequest) *error_responses.ErrorResponse
	Update(id int64, req *UpdateUserRequest, updatedBy int64) (*User, *error_responses.ErrorResponse)
	Delete(id int64, deletedBy int64) *error_responses.ErrorResponse
	GetCreateForm() any
	GetUpdateForm(id int64) (any, *error_responses.ErrorResponse)
	SetUserCtx(ctx *share.UserContext) bool
}

type UserServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Redis   *redis.Client
	Repo    UserRepo
}

func NewUserServiceImpl(dbpool *sqlx.DB, rdb *redis.Client) *UserServiceImpl {
	return &UserServiceImpl{
		// UserCtx: userCnt,
		dbpool: dbpool,
		Redis:  rdb,
		Repo:   NewUserRepoImpl(dbpool),
	}
}

func (s *UserServiceImpl) SetUserCtx(ctx *share.UserContext) bool {
	s.UserCtx = ctx
	return true
}

func (s *UserServiceImpl) Show(userRequest UserShowRequest) (*UserResponse, *error_responses.ErrorResponse) {

	return s.Repo.Show(userRequest)
}

func (s *UserServiceImpl) ShowOne(id int64) (*UserResponse, *error_responses.ErrorResponse) {
	return s.Repo.ShowOne(id)
}

func (s *UserServiceImpl) Create(req *CreateUserRequest) *error_responses.ErrorResponse {
	if err := s.Repo.Create(req); err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) Delete(id int64, deletedBy int64) *error_responses.ErrorResponse {
	return s.Repo.Delete(id, deletedBy)
}

func (s *UserServiceImpl) Update(id int64, req *UpdateUserRequest, updatedBy int64) (*User, *error_responses.ErrorResponse) {
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

//	func (s *UserServiceImpl) GetCreateForm() any {
//		return map[string]any{
//			"fields": map[string]any{
//				"user_name":  map[string]string{"type": "text", "required": "true", "min": "4"},
//				"password":   map[string]string{"type": "password", "required": "true", "min": "6"},
//				"first_name": map[string]string{"type": "text"},
//				"last_name":  map[string]string{"type": "text"},
//				"email":      map[string]string{"type": "email"},
//				"role_id":    map[string]string{"type": "number", "required": "true"},
//				"role_name":  map[string]string{"type": "text", "required": "true"},
//			},
//		}
//	}
func (s *UserServiceImpl) GetCreateForm() any {
	return map[string]any{
		"title":  "Create New User",
		"method": http.MethodPost,
		"fields": map[string]any{
			"user_name":  map[string]any{"type": "text", "required": true, "min": 4, "label": "Username"},
			"password":   map[string]any{"type": "password", "required": true, "min": 6, "label": "Password"},
			"first_name": map[string]any{"type": "text", "required": false, "label": "First Name"},
			"last_name":  map[string]any{"type": "text", "required": false, "label": "Last Name"},
			"email":      map[string]any{"type": "email", "required": false, "label": "Email Address"},
			"role_id":    map[string]any{"type": "number", "required": true, "label": "Role ID"},
			"role_name":  map[string]any{"type": "text", "required": true, "label": "Role Name"},
		},
	}
}

// func (s *UserServiceImpl) GetUpdateForm(id int64) (*UserResponse, *error_responses.ErrorResponse) {

// 	return s.Repo.ShowOne(id)
// }

func (s *UserServiceImpl) GetUpdateForm(id int64) (any, *error_responses.ErrorResponse) {
	res, err := s.Repo.ShowOne(id)
	if err != nil {
		return nil, err
	}

	user := res.Users[0]

	return map[string]any{
		"title": "Update User Information",
		"fields": map[string]any{
			"first_name": map[string]any{"type": "text", "required": false, "label": "First Name", "value": user.FirstName},
			"last_name":  map[string]any{"type": "text", "required": false, "label": "Last Name", "value": user.LastName},
			"email":      map[string]any{"type": "email", "required": false, "label": "Email Address", "value": user.Email},
			"role_id":    map[string]any{"type": "number", "required": true, "label": "Role ID", "value": user.RoleID},
			"role_name":  map[string]any{"type": "text", "required": true, "label": "Role Name", "value": user.RoleName},
			"status_id":  map[string]any{"type": "number", "required": false, "label": "Status ID", "value": user.StatusID},
		},
	}, nil
}
