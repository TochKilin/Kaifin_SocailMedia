package profile

import (
	"time"
)

type ProfileResponse struct {
	ID                 int64      `json:"id" db:"id"`
	UserName           string     `json:"user_name" db:"user_name"`
	FirstName          string     `json:"first_name" db:"first_name"`
	LastName           string     `json:"last_name" db:"last_name"`
	Email              string     `json:"email" db:"email"`
	ProfileImages      *string    `json:"profile_images" db:"profile_images"`
	CoverImages        *string    `json:"cover_images" db:"cover_images"`
	Location           *string    `json:"location" db:"location"`
	RelationshipStatus *string    `json:"relationship_status" db:"relationship_status"`
	Bios               *string    `json:"bios" db:"bio"`
	RoleID             int        `json:"role_id" db:"role_id"`
	RoleName           string     `json:"role_name" db:"role_name"`
	LastLogin          *time.Time `json:"last_login" db:"last_login"`
	CreatedAt          time.Time  `json:"created_at" db:"created_at"`
	PostCount          int        `json:"post_count" db:"post_count"`

	FollowerCount  int `db:"follower_count" json:"follower_count"`
	FollowingCount int `db:"following_count" json:"following_count"`
}

type UpdateProfileInfoRequest struct {
	FirstName          *string `json:"first_name"`
	LastName           *string `json:"last_name"`
	UserName           *string `json:"user_name"`
	Bio                *string `json:"bio"`
	RelationshipStatus *string `json:"relationship_status"`
	Location           *string `json:"location"`
}
