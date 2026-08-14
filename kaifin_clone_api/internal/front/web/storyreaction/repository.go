package storyreaction

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type StoryReactionRepo interface {
	Upsert(req *CreateStoryReactionRequest, uctx *share.UserContext) (*StoryReaction, *error_responses.ErrorResponse)
	Show(req ShowStoryReactionRequest) (*StoryReactionResponse, *error_responses.ErrorResponse)
	Delete(storyID int64, userID int64) *error_responses.ErrorResponse
}

type StoryReactionRepoImpl struct {
	dbpool *sqlx.DB
}

func NewStoryReactionRepoImpl(db *sqlx.DB) *StoryReactionRepoImpl {
	return &StoryReactionRepoImpl{
		dbpool: db,
	}
}

func (r *StoryReactionRepoImpl) Upsert(req *CreateStoryReactionRequest, uctx *share.UserContext) (*StoryReaction, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	sr := StoryReaction{}

	if err := sr.new(req, uctx); err != nil {
		return nil, msg.NewErrorResponse("invalid_reaction_type", err)
	}

	fmt.Printf("DEBUG: about to insert story_id=%d user_id=%d reaction_type=%s\n", sr.StoryID, sr.UserID, sr.ReactionType)
	err := r.dbpool.QueryRow(`
		INSERT INTO tbl_story_reactions (story_id, user_id, reaction_type)
		VALUES ($1, $2, $3)
		ON CONFLICT (story_id, user_id)
		DO UPDATE SET reaction_type = EXCLUDED.reaction_type, updated_at = NOW()
		RETURNING id, story_id, user_id, reaction_type, created_at, updated_at
	`,
		sr.StoryID,
		sr.UserID,
		sr.ReactionType,
	).Scan(&sr.ID, &sr.StoryID, &sr.UserID, &sr.ReactionType, &sr.CreatedAt, &sr.UpdatedAt)

	if err != nil {
		fmt.Println("UPSERT STORY REACTION ERROR:", err)
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &sr, nil
}

func (r *StoryReactionRepoImpl) Show(req ShowStoryReactionRequest) (*StoryReactionResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var reactions []StoryReaction
	listQuery := `
		SELECT
			sr.id,
			sr.story_id,
			sr.user_id,
			u.user_name AS user_name,
			sr.reaction_type,
			sr.created_at,
			sr.updated_at
		FROM tbl_story_reactions sr
		LEFT JOIN tbl_users u ON u.id = sr.user_id
		WHERE sr.story_id = $1
		ORDER BY sr.created_at DESC
	`
	if err := r.dbpool.Select(&reactions, listQuery, req.StoryID); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var summary []ReactionCount
	summaryQuery := `
		SELECT reaction_type, COUNT(*) AS count
		FROM tbl_story_reactions
		WHERE story_id = $1
		GROUP BY reaction_type
		ORDER BY count DESC
	`
	if err := r.dbpool.Select(&summary, summaryQuery, req.StoryID); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	total := 0
	for _, s := range summary {
		total += s.Count
	}

	return &StoryReactionResponse{
		Reactions:  reactions,
		Summary:    summary,
		TotalCount: total,
	}, nil
}

func (r *StoryReactionRepoImpl) Delete(storyID int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	result, err := r.dbpool.Exec(`
		DELETE FROM tbl_story_reactions WHERE story_id = $1 AND user_id = $2
	`, storyID, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("reaction_not_found", fmt.Errorf("reaction on story %d not found for user %d", storyID, userID))
	}
	return nil
}
