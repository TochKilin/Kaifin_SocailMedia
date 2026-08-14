package article

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type ArticlesRepo interface {
	Create(article *Article, tags []string, blocks []ArticleBlock) *error_responses.ErrorResponse
	Update(article *Article, tags []string, blocks []ArticleBlock) *error_responses.ErrorResponse
	Delete(id int64, userID int64) *error_responses.ErrorResponse
	Show(req ShowArticlesRequest, userID int64) (*ArticlesResponse, *error_responses.ErrorResponse)
	Detail(id int64, userID int64) (*Article, *error_responses.ErrorResponse)
	ToggleLike(articleID, userID int64) (bool, *error_responses.ErrorResponse)
	ToggleSave(articleID, userID int64) (bool, *error_responses.ErrorResponse)
	Report(articleID, userID int64, reportType, text string) *error_responses.ErrorResponse
}

type ArticlesRepoImpl struct {
	dbpool *sqlx.DB
}

func NewArticlesRepoImpl(db *sqlx.DB) *ArticlesRepoImpl {
	return &ArticlesRepoImpl{dbpool: db}
}

// ---------------- Create ----------------

func (r *ArticlesRepoImpl) Create(article *Article, tags []string, blocks []ArticleBlock) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	tx, err := r.dbpool.Beginx()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	err = tx.QueryRow(`
		INSERT INTO tbl_articles
			(user_id, title, summary, cover_image, category, code_subcategory, visibility, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`, article.UserID, article.Title, article.Summary, article.CoverImage,
		article.Category, article.CodeSubcategory, article.Visibility, article.Status,
	).Scan(&article.ID, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	if err := insertTags(tx, article.ID, tags); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	if err := insertBlocks(tx, article.ID, blocks); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	if err := tx.Commit(); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

// ---------------- Update ----------------

func (r *ArticlesRepoImpl) Update(article *Article, tags []string, blocks []ArticleBlock) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	tx, err := r.dbpool.Beginx()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	result, err := tx.Exec(`
		UPDATE tbl_articles
		SET title = $1, summary = $2, cover_image = COALESCE($3, cover_image), category = $4,
			code_subcategory = $5, visibility = $6, updated_at = NOW()
		WHERE id = $7 AND user_id = $8
	`, article.Title, article.Summary, article.CoverImage, article.Category,
		article.CodeSubcategory, article.Visibility, article.ID, article.UserID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("article_not_found", nil)
	}

	if _, err := tx.Exec(`DELETE FROM tbl_article_tags WHERE article_id = $1`, article.ID); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	if err := insertTags(tx, article.ID, tags); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	if _, err := tx.Exec(`DELETE FROM tbl_article_blocks WHERE article_id = $1`, article.ID); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	if err := insertBlocks(tx, article.ID, blocks); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	if err := tx.Commit(); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func insertTags(tx *sqlx.Tx, articleID int64, tags []string) error {
	for _, name := range tags {
		if name == "" {
			continue
		}
		var tagID int64
		err := tx.QueryRow(`
			INSERT INTO tbl_tags (name)
			VALUES ($1)
			ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name
			RETURNING id
		`, name).Scan(&tagID)
		if err != nil {
			return err
		}
		if _, err := tx.Exec(`
			INSERT INTO tbl_article_tags (article_id, tag_id)
			VALUES ($1, $2)
			ON CONFLICT DO NOTHING
		`, articleID, tagID); err != nil {
			return err
		}
	}
	return nil
}

func insertBlocks(tx *sqlx.Tx, articleID int64, blocks []ArticleBlock) error {
	for i, b := range blocks {
		if _, err := tx.Exec(`
			INSERT INTO tbl_article_blocks (article_id, block_type, title, content, position, created_at)
			VALUES ($1, $2, $3, $4, $5, NOW())
		`, articleID, b.BlockType, b.Title, b.Content, i); err != nil {
			return err
		}
	}
	return nil
}

// ---------------- Delete ----------------

func (r *ArticlesRepoImpl) Delete(id int64, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	result, err := r.dbpool.Exec(`
		DELETE FROM tbl_articles WHERE id = $1 AND user_id = $2
	`, id, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("article_not_found", nil)
	}
	return nil
}

// ---------------- Show (feed) ----------------

// articleSortColumns whitelists which "property" values a client may sort
// by, and maps each one to the actual SQL expression to use. This is
// required because req.Sorts[i].Property comes straight from the client —
// interpolating it into SQL directly would be a SQL-injection vector.
var articleSortColumns = map[string]string{
	"created_at":  "a.created_at",
	"views_count": "a.views_count",
	"like_count":  "COALESCE(lc.cnt, 0)",
	"title":       "a.title",
}

// buildArticleOrderBy turns validated share.Sort entries into a safe
// ORDER BY clause. Unknown/unsafe property names are silently skipped.
// Falls back to "a.created_at DESC" when nothing usable was supplied.
func buildArticleOrderBy(sorts []share.Sort) string {
	if len(sorts) == 0 {
		return "a.created_at DESC"
	}
	parts := make([]string, 0, len(sorts))
	for _, s := range sorts {
		col, ok := articleSortColumns[s.Property]
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
		return "a.created_at DESC"
	}
	return strings.Join(parts, ", ")
}

func (r *ArticlesRepoImpl) Show(req ShowArticlesRequest, userID int64) (*ArticlesResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	where := `WHERE a.status = 'published' AND a.visibility = 'public'`
	args := []interface{}{}
	argN := 1

	if req.Category != "" {
		where += fmt.Sprintf(" AND a.category = $%d", argN)
		args = append(args, req.Category)
		argN++
	}
	if req.CodeSubcategory != "" {
		where += fmt.Sprintf(" AND a.code_subcategory = $%d", argN)
		args = append(args, req.CodeSubcategory)
		argN++
	}
	if req.Search != "" {
		where += fmt.Sprintf(" AND a.title ILIKE $%d", argN)
		args = append(args, "%"+req.Search+"%")
		argN++
	}

	var total int
	countQuery := `SELECT COUNT(*) FROM tbl_articles a ` + where
	if err := r.dbpool.Get(&total, countQuery, args...); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	userIDParam := argN
	limitParam := argN + 1
	offsetParam := argN + 2
	offset := (req.Page - 1) * req.PerPage
	args = append(args, userID, req.PerPage, offset)

	orderBy := buildArticleOrderBy(req.Sorts)

	query := fmt.Sprintf(`
		SELECT
			a.id, a.user_id, u.user_name AS user_name, u.profile_images AS profile_images,
			a.title, a.summary, a.cover_image, a.category, a.code_subcategory,
			a.visibility, a.status, a.views_count, a.created_at, a.updated_at,
			COALESCE(lc.cnt, 0) AS like_count,
			COALESCE(ul.liked, false) AS liked,
			COALESCE(cc.cnt, 0) AS comment_count,
			COALESCE(sc.cnt, 0) AS save_count,
			COALESCE(us.saved, false) AS saved
		FROM tbl_articles a
		LEFT JOIN tbl_users u ON u.id = a.user_id
		LEFT JOIN (
			SELECT article_id, COUNT(*) AS cnt FROM tbl_article_likes GROUP BY article_id
		) lc ON lc.article_id = a.id
		LEFT JOIN (
			SELECT article_id, true AS liked FROM tbl_article_likes WHERE user_id = $%d
		) ul ON ul.article_id = a.id
		LEFT JOIN (
			SELECT article_id, COUNT(*) AS cnt FROM tbl_article_comments GROUP BY article_id
		) cc ON cc.article_id = a.id
		LEFT JOIN (
			SELECT article_id, COUNT(*) AS cnt FROM tbl_article_saves GROUP BY article_id
		) sc ON sc.article_id = a.id
		LEFT JOIN (
			SELECT article_id, true AS saved FROM tbl_article_saves WHERE user_id = $%d
		) us ON us.article_id = a.id
		%s
		ORDER BY %s
		LIMIT $%d OFFSET $%d
	`, userIDParam, userIDParam, where, orderBy, limitParam, offsetParam)

	var articleList []Article
	if err := r.dbpool.Select(&articleList, query, args...); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	if err := attachTags(r.dbpool, articleList); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &ArticlesResponse{
		Articles: articleList,
		Total:    total,
		Page:     req.Page,
		PerPage:  req.PerPage,
	}, nil
}

// ---------------- Detail ----------------

func (r *ArticlesRepoImpl) Detail(id int64, userID int64) (*Article, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var article Article
	err := r.dbpool.Get(&article, `
		SELECT
			a.id, a.user_id, u.user_name AS user_name, u.profile_images AS profile_images,
			a.title, a.summary, a.cover_image, a.category, a.code_subcategory,
			a.visibility, a.status, a.views_count, a.created_at, a.updated_at,
			COALESCE(lc.cnt, 0) AS like_count,
			COALESCE(ul.liked, false) AS liked,
			COALESCE(cc.cnt, 0) AS comment_count,
			COALESCE(sc.cnt, 0) AS save_count,
			COALESCE(us.saved, false) AS saved
		FROM tbl_articles a
		LEFT JOIN tbl_users u ON u.id = a.user_id
		LEFT JOIN (
			SELECT article_id, COUNT(*) AS cnt FROM tbl_article_likes GROUP BY article_id
		) lc ON lc.article_id = a.id
		LEFT JOIN (
			SELECT article_id, true AS liked FROM tbl_article_likes WHERE user_id = $2
		) ul ON ul.article_id = a.id
		LEFT JOIN (
			SELECT article_id, COUNT(*) AS cnt FROM tbl_article_comments GROUP BY article_id
		) cc ON cc.article_id = a.id
		LEFT JOIN (
			SELECT article_id, COUNT(*) AS cnt FROM tbl_article_saves GROUP BY article_id
		) sc ON sc.article_id = a.id
		LEFT JOIN (
			SELECT article_id, true AS saved FROM tbl_article_saves WHERE user_id = $2
		) us ON us.article_id = a.id
		WHERE a.id = $1
	`, id, userID)
	if err != nil {
		return nil, msg.NewErrorResponse("article_not_found", err)
	}

	var blocks []ArticleBlock
	if err := r.dbpool.Select(&blocks, `
		SELECT id, article_id, block_type, title, content, position, created_at
		FROM tbl_article_blocks
		WHERE article_id = $1
		ORDER BY position ASC
	`, id); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	article.Blocks = blocks

	articleList := []Article{article}
	if err := attachTags(r.dbpool, articleList); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	article.Tags = articleList[0].Tags

	// Fire-and-forget view count bump (best effort, ignore error)
	_, _ = r.dbpool.Exec(`UPDATE tbl_articles SET views_count = views_count + 1 WHERE id = $1`, id)

	return &article, nil
}

func attachTags(db *sqlx.DB, articleList []Article) error {
	if len(articleList) == 0 {
		return nil
	}
	ids := make([]int64, len(articleList))
	idx := map[int64]int{}
	for i, a := range articleList {
		ids[i] = a.ID
		idx[a.ID] = i
	}

	type tagRow struct {
		ArticleID int64  `db:"article_id"`
		Name      string `db:"name"`
	}
	var rows []tagRow
	q, args, err := sqlx.In(`
		SELECT at.article_id, t.name
		FROM tbl_article_tags at
		JOIN tbl_tags t ON t.id = at.tag_id
		WHERE at.article_id IN (?)
		ORDER BY t.name ASC
	`, ids)
	if err != nil {
		return err
	}
	q = db.Rebind(q)
	if err := db.Select(&rows, q, args...); err != nil {
		return err
	}
	for _, row := range rows {
		i := idx[row.ArticleID]
		articleList[i].Tags = append(articleList[i].Tags, row.Name)
	}
	return nil
}

// ---------------- Like ----------------

func (r *ArticlesRepoImpl) ToggleLike(articleID, userID int64) (bool, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var exists bool
	err := r.dbpool.Get(&exists, `
		SELECT EXISTS(SELECT 1 FROM tbl_article_likes WHERE article_id = $1 AND user_id = $2)
	`, articleID, userID)
	if err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}

	if exists {
		if _, err = r.dbpool.Exec(`
			DELETE FROM tbl_article_likes WHERE article_id = $1 AND user_id = $2
		`, articleID, userID); err != nil {
			return false, msg.NewErrorResponse("database_error", err)
		}
		return false, nil
	}

	if _, err = r.dbpool.Exec(`
		INSERT INTO tbl_article_likes (article_id, user_id, created_at) VALUES ($1, $2, NOW())
	`, articleID, userID); err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}
	return true, nil
}

// ---------------- Save ----------------

func (r *ArticlesRepoImpl) ToggleSave(articleID, userID int64) (bool, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var exists bool
	err := r.dbpool.Get(&exists, `
		SELECT EXISTS(SELECT 1 FROM tbl_article_saves WHERE article_id = $1 AND user_id = $2)
	`, articleID, userID)
	if err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}

	if exists {
		if _, err = r.dbpool.Exec(`
			DELETE FROM tbl_article_saves WHERE article_id = $1 AND user_id = $2
		`, articleID, userID); err != nil {
			return false, msg.NewErrorResponse("database_error", err)
		}
		return false, nil
	}

	if _, err = r.dbpool.Exec(`
		INSERT INTO tbl_article_saves (article_id, user_id, created_at) VALUES ($1, $2, NOW())
	`, articleID, userID); err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}
	return true, nil
}

// ---------------- Report ----------------

func (r *ArticlesRepoImpl) Report(articleID, userID int64, reportType, text string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	_, err := r.dbpool.Exec(`
		INSERT INTO tbl_article_reports (article_id, user_id, report_type, text, status, created_at)
		VALUES ($1, $2, $3, $4, 'pending', NOW())
	`, articleID, userID, reportType, text)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}
