package likes

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type LikesRepo interface {
	Create(like *Like) *error_responses.ErrorResponse
	Delete(postID, userID int64) *error_responses.ErrorResponse
	Show(req ShowLikesRequest, currentUserID int64) (*LikesResponse, *error_responses.ErrorResponse)
	FindByPostAndUser(postID, userID int64) (*Like, error)
}

type LikesRepoImpl struct {
	dbpool *sqlx.DB
}

func NewLikesRepoImpl(db *sqlx.DB) *LikesRepoImpl {
	return &LikesRepoImpl{
		dbpool: db,
	}
}

func (r *LikesRepoImpl) FindByPostAndUser(postID, userID int64) (*Like, error) {
	var like Like
	err := r.dbpool.Get(&like, `
		SELECT id, post_id, user_id, reaction_type, created_at
		FROM tbl_reactions
		WHERE post_id = $1 AND user_id = $2
	`, postID, userID)
	if err != nil {
		return nil, err
	}
	return &like, nil
}

func (r *LikesRepoImpl) Create(like *Like) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	err := r.dbpool.QueryRow(`
		INSERT INTO tbl_reactions (post_id, user_id, reaction_type, created_at)
		VALUES ($1, $2, $3, NOW())
		ON CONFLICT (post_id, user_id)
		DO UPDATE SET reaction_type = EXCLUDED.reaction_type
		RETURNING id, created_at
	`, like.PostID, like.UserID, like.Type).Scan(&like.ID, &like.CreatedAt)

	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *LikesRepoImpl) Delete(postID, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	result, err := r.dbpool.Exec(`
		DELETE FROM tbl_reactions WHERE post_id = $1 AND user_id = $2
	`, postID, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("like_not_found", nil)
	}
	return nil
}

func (r *LikesRepoImpl) Show(req ShowLikesRequest, currentUserID int64) (*LikesResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var likes []Like
	err := r.dbpool.Select(&likes, `
		SELECT tbl_reactions.id, tbl_reactions.post_id, tbl_reactions.user_id, 
		       tbl_reactions.reaction_type, tbl_reactions.created_at, 
		       u.profile_images, u.user_name
		FROM tbl_reactions
		LEFT JOIN tbl_users u ON u.id = tbl_reactions.user_id
		WHERE post_id = $1
		ORDER BY created_at DESC
	`, req.PostID)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var summary []LikeSummary
	err = r.dbpool.Select(&summary, `
		SELECT reaction_type, COUNT(*) AS count
		FROM tbl_reactions
		WHERE post_id = $1
		GROUP BY reaction_type
	`, req.PostID)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	likedByMe := false
	myReaction := ""
	for _, l := range likes {
		if l.UserID == currentUserID {
			likedByMe = true
			myReaction = l.Type
			break
		}
	}

	return &LikesResponse{
		Likes:      likes,
		Summary:    summary,
		Total:      len(likes),
		LikedByMe:  likedByMe,
		MyReaction: myReaction,
	}, nil
}

func (r *LikesRepoImpl) ToggleLike(req *CreateLikeRequest, userID int64) (bool, *error_responses.ErrorResponse) {
	existing, err := r.FindByPostAndUser(req.PostID, userID)
	if err == nil && existing != nil {
		if existing.Type == req.Type {
			if e := r.Delete(req.PostID, userID); e != nil {
				return false, e
			}
			return false, nil
		}
	}

	like := &Like{
		PostID: req.PostID,
		UserID: userID,
		Type:   req.Type,
	}
	if e := r.Create(like); e != nil {
		return false, e
	}
	return true, nil
}
