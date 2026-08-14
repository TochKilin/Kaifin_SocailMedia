package post_mobile

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
	custom_sql "kaifin_clone_api/pkg/sql"
)

type PostRepo interface {
	Create(req *CreatePostMobileRequest) *error_responses.ErrorResponse
	Delete(id int64) *error_responses.ErrorResponse
	Show(postRequest ShowPostMobileRequest) (*PostResponse, *error_responses.ErrorResponse)
	IncrementView(id int64) (int, *error_responses.ErrorResponse)
}

type PostRepoImpl struct {
	dbpool *sqlx.DB
	redis  *redis.Client
}

func NewPostRepoImpl(db *sqlx.DB) *PostRepoImpl {
	return &PostRepoImpl{
		dbpool: db,
	}
}

func (r *PostRepoImpl) Create(req *CreatePostMobileRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	post := Post{}
	if err := post.new(req, uctx); err != nil {
		return msg.NewErrorResponse("invalid", err)
	}
	err := r.dbpool.QueryRow(
		`
		INSERT INTO tbl_posts
		(
			user_id,
			community_id,
			caption,
			post_type,
			code_content,
			link_url
		)
		VALUES
		($1,$2,$3,$4,$5,$6)
		RETURNING id, created_at, updated_at
		`,
		post.UserID,
		post.CommunityID,
		post.Caption,
		post.PostType,
		post.CodeContent,
		post.LinkURL,
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	fmt.Println("POST CREATED WITH ID:", post.ID)

	for _, image := range req.ImageURLs {
		_, imgErr := r.dbpool.Exec(`
			INSERT INTO tbl_post_images (post_id, image_url, created_at)
			VALUES ($1,$2,NOW())
		`, post.ID, image)

		if imgErr != nil {
			fmt.Println("INSERT IMAGE ERROR:", imgErr)
			return msg.NewErrorResponse("database_error", imgErr)
		}
	}

	for _, hashtag := range req.Hashtags {
		fmt.Println("INSERT TAG =>", hashtag)

		_, tagErr := r.dbpool.Exec(`
			INSERT INTO tbl_post_hashtags (post_id, tag_name)
			VALUES ($1,$2)
		`, post.ID, hashtag)

		if tagErr != nil {
			fmt.Println("INSERT TAG ERROR:", tagErr)
			return msg.NewErrorResponse("database_error", tagErr)
		}
	}

	fmt.Println("STICKER IDS TO INSERT:", req.StickerIDs)
	for _, stickerID := range req.StickerIDs {
		fmt.Println("INSERT STICKER =>", stickerID)

		_, stkErr := r.dbpool.Exec(`
			INSERT INTO tbl_post_stickers (post_id, sticker_id, created_at)
			VALUES ($1,$2,NOW())
		`, post.ID, stickerID)

		if stkErr != nil {
			fmt.Println("INSERT STICKER ERROR:", stkErr)
			return msg.NewErrorResponse("database_error", stkErr)
		}
	}

	if req.VideoURL != nil {
		_, vidErr := r.dbpool.Exec(`
		INSERT INTO tbl_post_videos (post_id, video_path, created_at)
		VALUES ($1,$2,NOW())
	`, post.ID, *req.VideoURL)
		if vidErr != nil {
			return msg.NewErrorResponse("database_error", vidErr)
		}
	}
	return nil
}

func (r *PostRepoImpl) Show(postRequest ShowPostMobileRequest) (*PostResponse, *error_responses.ErrorResponse) {
	var per_page = postRequest.PageOption.Perpage
	var page = postRequest.PageOption.Page
	var offset = (page - 1) * per_page
	var limit_clause = fmt.Sprintf(" LIMIT %d OFFSET %d", per_page, offset)
	var sql_orderby string

	if len(postRequest.Sorts) == 0 {
		sql_orderby = "ORDER BY p.created_at DESC"
	} else {
		sql_orderby = custom_sql.BuildSQLSort(postRequest.Sorts)
	}

	sql_filters, args_filters := custom_sql.BuildSQLFilter(postRequest.Filters)
	if len(args_filters) > 0 {
		sql_filters = " AND " + sql_filters
	}
	fmt.Printf("DEBUG sql_orderby: %q\n", sql_orderby)

	if searchClause, searchArgs := custom_sql.BuildSQLSearch(
		[]string{"p.caption"},
		postRequest.Search, len(args_filters)+1,
	); searchClause != "" {
		sql_filters += " AND " + searchClause
		args_filters = append(args_filters, searchArgs...)
	}
	msg := error_responses.ErrorResponse{}
	var posts []Post
	query := fmt.Sprintf(
		`SELECT
		p.id,
		p.user_id,
		u.user_name AS user_name, 
		u.profile_images AS profile_images,
		p.community_id,
		p.caption,
		p.post_type,
		p.code_content,
		p.link_url,
		p.views_count,
		p.created_at,
		p.updated_at,
		STRING_AGG(DISTINCT pi.image_url, ',') AS images,
		STRING_AGG(DISTINCT ph.tag_name, ',') AS tag_name,
		STRING_AGG(DISTINCT ps.sticker_id::text, ',') AS sticker_ids,
		pv.video_path AS video_path,
        pv.duration AS duration,
        pv.thumbnail_path AS thumbnail_path,
		COALESCE(cc.comment_count, 0) AS comment_count
	FROM tbl_posts p
	LEFT JOIN tbl_post_images pi
	ON pi.post_id = p.id
	LEFT JOIN tbl_post_hashtags ph
	ON ph.post_id = p.id
	LEFT JOIN tbl_post_stickers ps
	ON ps.post_id = p.id
	LEFT JOIN tbl_users u
    ON u.id = p.user_id
	LEFT JOIN (
		SELECT post_id, COUNT(*) AS comment_count
		FROM tbl_comments
		GROUP BY post_id
	) cc ON cc.post_id = p.id
	LEFT JOIN tbl_post_videos pv ON pv.post_id = p.id
	WHERE 1=1
	%s
	GROUP BY p.id, u.user_name, u.profile_images, cc.comment_count, pv.video_path, pv.duration, pv.thumbnail_path
	%s %s`,
		sql_filters,
		sql_orderby,
		limit_clause,
	)
	err := r.dbpool.Select(&posts, query, args_filters...)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var total int
	countQuery := fmt.Sprintf(
		`SELECT COUNT(DISTINCT p.id)
	 FROM tbl_posts p
	 LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
	 LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
	 WHERE 1=1
	 %s`,
		sql_filters,
	)

	err = r.dbpool.Get(&total, countQuery, args_filters...)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &PostResponse{
		Posts: posts,
		Total: total}, nil

}

func (r *PostRepoImpl) Delete(id int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	result, err := r.dbpool.Exec(
		`DELETE FROM tbl_posts WHERE id = $1`, id,
	)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("post_not_found", fmt.Errorf("posts %d not found", id))
	}
	return nil
}

func (r *PostRepoImpl) IncrementView(id int64) (int, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var newCount int
	err := r.dbpool.QueryRow(`
		UPDATE tbl_posts SET views_count = views_count + 1
		WHERE id = $1
		RETURNING views_count
	`, id).Scan(&newCount)

	if err != nil {
		return 0, msg.NewErrorResponse("database_error", err)
	}

	return newCount, nil
}
