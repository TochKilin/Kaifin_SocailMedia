package articlecomment

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type CommentsRepo interface {
	Create(comment *ArticleComment) *error_responses.ErrorResponse
	Update(id int64, userID int64, text string) *error_responses.ErrorResponse
	Delete(id int64, userID int64) *error_responses.ErrorResponse
	Show(req ShowCommentsRequest) (*CommentsResponse, *error_responses.ErrorResponse)
}

type CommentsRepoImpl struct {
	dbpool *sqlx.DB
}

func NewCommentsRepoImpl(db *sqlx.DB) *CommentsRepoImpl {
	return &CommentsRepoImpl{dbpool: db}
}

// ---------------- Create ----------------

func (r *CommentsRepoImpl) Create(comment *ArticleComment) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	var exists bool
	if err := r.dbpool.Get(&exists, `
        SELECT EXISTS(SELECT 1 FROM tbl_articles WHERE id = $1)
    `, comment.ArticleID); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	if !exists {
		return msg.NewErrorResponse("article_not_found", nil)
	}

	if comment.ParentCommentID != nil {
		var parentOk bool
		if err := r.dbpool.Get(&parentOk, `
            SELECT EXISTS(
                SELECT 1 FROM tbl_article_comments
                WHERE id = $1 AND article_id = $2
            )
        `, *comment.ParentCommentID, comment.ArticleID); err != nil {
			return msg.NewErrorResponse("database_error", err)
		}
		if !parentOk {
			return msg.NewErrorResponse("parent_comment_not_found", nil)
		}
	}

	tx, err := r.dbpool.Beginx()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	// 1. Insert ចូល tbl_article_comments
	err = tx.QueryRow(`
        INSERT INTO tbl_article_comments (article_id, user_id, parent_comment_id, text, created_at)
        VALUES ($1, $2, $3, $4, NOW())
        RETURNING id, created_at
    `, comment.ArticleID, comment.UserID, comment.ParentCommentID, comment.Text,
	).Scan(&comment.ID, &comment.CreatedAt)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	// 2. Insert ចូល tbl_article_comment_images សម្រាប់ ImageFiles ទាំងអស់ដែលបាន Upload
	for _, imageURL := range comment.SavedImageURLs {
		_, err = tx.Exec(`
        INSERT INTO tbl_article_comment_images (comment_id, image_url)
        VALUES ($1, $2)
    `, comment.ID, imageURL)
		if err != nil {
			return msg.NewErrorResponse("database_error", err)
		}
	}

	// 3. Insert ចូល tbl_article_comment_stickers សម្រាប់ StickerIDs ទាំងអស់
	for _, stickerID := range comment.StickerIDs {
		_, err = tx.Exec(`
            INSERT INTO tbl_article_comment_stickers (comment_id, sticker_id)
            VALUES ($1, $2)
        `, comment.ID, stickerID)
		if err != nil {
			return msg.NewErrorResponse("database_error", err)
		}
	}

	// 4. Load display info (username, avatar) for the response — the same
	// source Show() joins in — inside the same transaction, so the Create
	// response looks exactly like a row from the list endpoint. This is
	// what makes the username/avatar show up immediately, no page refresh
	// needed. Do NOT overwrite these from uCtx/JWT claims in the handler.
	if err := tx.QueryRow(`
        SELECT user_name, profile_images FROM tbl_users WHERE id = $1
    `, comment.UserID).Scan(&comment.UserName, &comment.ProfileImage); err != nil {
		// Not fatal — the comment itself was created successfully.
		log.Printf("Create comment: failed to load user info for user %d: %v", comment.UserID, err)
	}

	if err := tx.Commit(); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	// Also fill ImageURLs on the response so it matches what Show() would
	// return for this comment (Show gets it via string_agg + splitImageURLs).
	comment.ImageURLs = append([]string{}, comment.SavedImageURLs...)

	return nil
}

// ---------------- Update ----------------

func (r *CommentsRepoImpl) Update(id int64, userID int64, text string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	result, err := r.dbpool.Exec(`
        UPDATE tbl_article_comments
        SET text = $1
        WHERE id = $2 AND user_id = $3
    `, text, id, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("comment_not_found", nil)
	}
	return nil
}

// ---------------- Delete ----------------

func (r *CommentsRepoImpl) Delete(id int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	result, err := r.dbpool.Exec(`
        DELETE FROM tbl_article_comments WHERE id = $1 AND user_id = $2
    `, id, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("comment_not_found", nil)
	}
	return nil
}

// ---------------- Show ----------------

var commentSortColumns = map[string]string{
	"created_at": "c.created_at",
}

func buildCommentOrderBy(sorts []share.Sort) string {
	if len(sorts) == 0 {
		return "c.created_at DESC"
	}
	parts := make([]string, 0, len(sorts))
	for _, s := range sorts {
		col, ok := commentSortColumns[s.Property]
		if !ok {
			continue
		}
		dir := "ASC"
		if strings.EqualFold(s.Direction, "desc") {
			dir = "DESC"
		}
		parts = append(parts, col+" "+dir)
	}
	if len(parts) == 0 {
		return "c.created_at DESC"
	}
	return strings.Join(parts, ", ")
}

func (r *CommentsRepoImpl) Show(req ShowCommentsRequest) (*CommentsResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var total int
	if err := r.dbpool.Get(&total, `
        SELECT COUNT(*) FROM tbl_article_comments WHERE article_id = $1
    `, req.ArticleID); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	offset := (req.Page - 1) * req.PerPage
	orderBy := buildCommentOrderBy(req.Sorts)

	query := fmt.Sprintf(`
        SELECT
            c.id,
            c.article_id,
            c.user_id,
            c.parent_comment_id,
            c.text,
            c.created_at,
            u.user_name AS user_name,
            u.profile_images AS profile_images,
            string_agg(ci.image_url, ',') AS image_url
        FROM tbl_article_comments c
        LEFT JOIN tbl_users u ON u.id = c.user_id
        LEFT JOIN tbl_article_comment_images ci ON ci.comment_id = c.id
        WHERE c.article_id = $1
        GROUP BY c.id, u.user_name, u.profile_images
        ORDER BY %s
        LIMIT $2 OFFSET $3
    `, orderBy)

	var comments []ArticleComment
	if err := r.dbpool.Select(&comments, query, req.ArticleID, req.PerPage, offset); err != nil {
		log.Printf("SHOW COMMENTS SQL ERROR: %v", err)
		return nil, msg.NewErrorResponse("database_error", err)
	}

	splitImageURLs(comments)

	return &CommentsResponse{
		Comments: comments,
		Total:    total,
		Page:     req.Page,
		PerPage:  req.PerPage,
	}, nil
}

func splitImageURLs(comments []ArticleComment) {
	for i := range comments {
		if comments[i].ImageURLRaw != nil && *comments[i].ImageURLRaw != "" {
			comments[i].ImageURLs = strings.Split(*comments[i].ImageURLRaw, ",")
		} else {
			comments[i].ImageURLs = []string{}
		}
	}
}

func saveCommentImage(c fiber.Ctx, fh *multipart.FileHeader) (string, error) {
	dir := filepath.Join("uploads", "comments")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}

	ext := filepath.Ext(fh.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), randomHex(8), ext)
	dst := filepath.Join(dir, filename)
	if err := c.SaveFile(fh, dst); err != nil {
		return "", err
	}
	return fmt.Sprintf("/uploads/comments/%s", filename), nil
}

func randomHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return hex.EncodeToString(b)
}
