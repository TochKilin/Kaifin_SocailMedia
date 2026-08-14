package storyreaction

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type CreateStoryReactionRequest struct {
	StoryID      int64  `json:"story_id" form:"story_id" validate:"required"`
	ReactionType string `json:"reaction_type" form:"reaction_type" validate:"required"`
}

type ShowStoryReactionRequest struct {
	StoryID int64 `query:"story_id" validate:"required"`
}

type StoryReaction struct {
	ID           int64     `json:"id" db:"id"`
	StoryID      int64     `json:"story_id" db:"story_id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	Username     string    `json:"user_name" db:"user_name"`
	ReactionType string    `json:"reaction_type" db:"reaction_type"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type ReactionCount struct {
	ReactionType string `json:"reaction_type" db:"reaction_type"`
	Count        int    `json:"count" db:"count"`
}

type StoryReactionResponse struct {
	Reactions  []StoryReaction `json:"reactions"`
	Summary    []ReactionCount `json:"summary"`
	TotalCount int             `json:"total_count"`
}

var allowedReactionTypes = map[string]bool{
	"heart":    true,
	"congrate": true,
	"cool":     true,
	"thinking": true,
}

func (u *CreateStoryReactionRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(u); err != nil {
		return err
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

func (u *ShowStoryReactionRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(u); err != nil {
		return err
	}
	if err := v.Validate(u); err != nil {
		return err
	}
	return nil
}

func (r *StoryReaction) new(req *CreateStoryReactionRequest, uctx *share.UserContext) error {
	if !allowedReactionTypes[req.ReactionType] {
		return fmt.Errorf("invalid reaction type: %s", req.ReactionType)
	}

	r.StoryID = req.StoryID
	r.UserID = uctx.UserID
	r.ReactionType = req.ReactionType

	return nil
}
