package music

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type PostMusicRepo interface {
	Create(userID int64, req CreatePostRequest) (*PostResponse, *error_responses.ErrorResponse)
	Show(requesterID int64, req ShowPostRequest) (*PostListResponse, *error_responses.ErrorResponse)
	GetByID(id int64) (*Post, *error_responses.ErrorResponse)
	Update(id int64, userID int64, req UpdatePostRequest) (*PostResponse, *error_responses.ErrorResponse)
	Delete(id int64, userID int64) *error_responses.ErrorResponse
}

type PostMusicRepoImpl struct {
	dbpool *sqlx.DB
}

func NewPostRepoImpl(db *sqlx.DB) *PostMusicRepoImpl {
	return &PostMusicRepoImpl{
		dbpool: db,
	}
}

// ---------- Create ----------

func (r *PostMusicRepoImpl) Create(userID int64, req CreatePostRequest) (*PostResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	tx, err := r.dbpool.Beginx()
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	var newPost Post
	insertPostQuery := `
		INSERT INTO music_posts (user_id, song_id, content, type, audience, disable_comments)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, user_id, song_id, content, type, audience, disable_comments, created_at`

	err = tx.QueryRowx(
		insertPostQuery,
		userID, req.SongID, req.Content, req.Type, req.Audience, req.DisableComments,
	).StructScan(&newPost)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	taggedUserIDs, err := insertTags(tx, newPost.ID, userID, req.TaggedUserIDs)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return toPostResponse(newPost, taggedUserIDs), nil
}

// ---------- Show / List ----------

func (r *PostMusicRepoImpl) Show(requesterID int64, req ShowPostRequest) (*PostListResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	offset := (req.Page - 1) * req.Limit

	// A post is visible if it's public, it's the requester's own post,
	// or the requester is an accepted friend of the author (for
	// audience = 'friend' posts). requesterID = 0 (no logged-in user)
	// never matches p.user_id or a friends row, so anonymous requests
	// naturally fall back to public posts only.
	baseQuery := `
		FROM music_posts p
		LEFT JOIN friends f ON (
			p.audience = 'friend' AND f.status = 'accepted' AND (
				(f.user_id = $1 AND f.friend_id = p.user_id) OR
				(f.friend_id = $1 AND f.user_id = p.user_id)
			)
		)
		WHERE (p.audience = 'everyone' OR p.user_id = $1 OR f.id IS NOT NULL)
		AND ($2 = '' OR p.content ILIKE '%' || $2 || '%')`

	var total int
	countQuery := `SELECT COUNT(DISTINCT p.id) ` + baseQuery
	if err := r.dbpool.Get(&total, countQuery, requesterID, req.Search); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var posts []Post
	listQuery := `
		SELECT DISTINCT p.id, p.user_id, p.song_id, p.content, p.type,
			p.audience, p.disable_comments, p.created_at
		` + baseQuery + `
		ORDER BY p.created_at DESC
		LIMIT $3 OFFSET $4`
	if err := r.dbpool.Select(&posts, listQuery, requesterID, req.Search, req.Limit, offset); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	tagsByPost, err := r.fetchTags(postIDs(posts))
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	responses := make([]PostResponse, 0, len(posts))
	for _, p := range posts {
		responses = append(responses, *toPostResponse(p, tagsByPost[p.ID]))
	}

	return &PostListResponse{
		Posts: responses,
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}

// ---------- GetByID ----------
// Used internally (and available to other modules, e.g. comments/likes)
// to check ownership or existence before acting on a post.

func (r *PostMusicRepoImpl) GetByID(id int64) (*Post, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var p Post
	query := `
		SELECT id, user_id, song_id, content, type, audience, disable_comments, created_at
		FROM music_posts WHERE id = $1`
	if err := r.dbpool.Get(&p, query, id); err != nil {
		return nil, msg.NewErrorResponse("post_not_found", err)
	}
	return &p, nil
}

// ---------- Update ----------

func (r *PostMusicRepoImpl) Update(id int64, userID int64, req UpdatePostRequest) (*PostResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	tx, err := r.dbpool.Beginx()
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	var updated Post
	updateQuery := `
		UPDATE music_posts SET
			content          = COALESCE($1, content),
			type             = COALESCE($2, type),
			audience         = COALESCE($3, audience),
			disable_comments = COALESCE($4, disable_comments)
		WHERE id = $5 AND user_id = $6
		RETURNING id, user_id, song_id, content, type, audience, disable_comments, created_at`

	err = tx.QueryRowx(
		updateQuery,
		req.Content, req.Type, req.Audience, req.DisableComments, id, userID,
	).StructScan(&updated)
	if err != nil {
		// Either the row doesn't exist, or it exists but belongs to
		// someone else — both surface as the same "not yours" error so
		// we don't leak which one it was.
		return nil, msg.NewErrorResponse("post_not_found_or_forbidden", err)
	}

	var taggedUserIDs []int64
	if req.TaggedUserIDs != nil {
		if _, err := tx.Exec(`DELETE FROM music_post_tags WHERE music_post_id = $1`, id); err != nil {
			return nil, msg.NewErrorResponse("database_error", err)
		}
		taggedUserIDs, err = insertTags(tx, id, userID, *req.TaggedUserIDs)
		if err != nil {
			return nil, msg.NewErrorResponse("database_error", err)
		}
	} else {
		taggedUserIDs, err = fetchTagsTx(tx, id)
		if err != nil {
			return nil, msg.NewErrorResponse("database_error", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return toPostResponse(updated, taggedUserIDs), nil
}

// ---------- Delete ----------

func (r *PostMusicRepoImpl) Delete(id int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	tx, err := r.dbpool.Beginx()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	// No ON DELETE CASCADE was declared on these FKs in the schema, so
	// child rows are removed explicitly before the post itself.
	if _, err := tx.Exec(`DELETE FROM music_post_tags WHERE music_post_id = $1`, id); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	if _, err := tx.Exec(`DELETE FROM music_post_comments WHERE music_post_id = $1`, id); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	if _, err := tx.Exec(`DELETE FROM music_post_likes WHERE music_post_id = $1`, id); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	result, err := tx.Exec(`DELETE FROM music_posts WHERE id = $1 AND user_id = $2`, id, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	if rows == 0 {
		return msg.NewErrorResponse("post_not_found_or_forbidden", fmt.Errorf("post %d not found for user %d", id, userID))
	}

	if err := tx.Commit(); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

// ---------- shared helpers ----------

func insertTags(tx *sqlx.Tx, postID int64, ownerID int64, taggedUserIDs []int64) ([]int64, error) {
	result := make([]int64, 0, len(taggedUserIDs))
	if len(taggedUserIDs) == 0 {
		return result, nil
	}
	query := `INSERT INTO music_post_tags (music_post_id, tagged_user_id) VALUES ($1, $2)`
	for _, taggedUserID := range taggedUserIDs {
		if taggedUserID == ownerID {
			continue // a user can't tag themselves
		}
		if _, err := tx.Exec(query, postID, taggedUserID); err != nil {
			return nil, err
		}
		result = append(result, taggedUserID)
	}
	return result, nil
}

func fetchTagsTx(tx *sqlx.Tx, postID int64) ([]int64, error) {
	var ids []int64
	err := tx.Select(&ids, `SELECT tagged_user_id FROM music_post_tags WHERE music_post_id = $1`, postID)
	return ids, err
}

func (r *PostMusicRepoImpl) fetchTags(postIDsList []int64) (map[int64][]int64, error) {
	tagsByPost := make(map[int64][]int64)
	if len(postIDsList) == 0 {
		return tagsByPost, nil
	}

	var tags []PostTag
	query, args, err := sqlx.In(
		`SELECT id, music_post_id, tagged_user_id, created_at
		 FROM music_post_tags WHERE music_post_id IN (?)`, postIDsList,
	)
	if err != nil {
		return nil, err
	}
	query = r.dbpool.Rebind(query)
	if err := r.dbpool.Select(&tags, query, args...); err != nil {
		return nil, err
	}

	for _, t := range tags {
		tagsByPost[t.PostID] = append(tagsByPost[t.PostID], t.TaggedUserID)
	}
	return tagsByPost, nil
}

func postIDs(posts []Post) []int64 {
	ids := make([]int64, len(posts))
	for i, p := range posts {
		ids[i] = p.ID
	}
	return ids
}

func toPostResponse(p Post, taggedUserIDs []int64) *PostResponse {
	return &PostResponse{
		ID:              p.ID,
		UserID:          p.UserID,
		SongID:          p.SongID,
		Content:         p.Content,
		Type:            p.Type,
		Audience:        p.Audience,
		DisableComments: p.DisableComments,
		TaggedUserIDs:   taggedUserIDs,
		CreatedAt:       p.CreatedAt,
	}
}
