package post

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
	custom_sql "kaifin_clone_api/pkg/sql"
)

type PostRepo interface {
	Create(req *CreatePostRequest) *error_responses.ErrorResponse
	Delete(id int64) *error_responses.ErrorResponse
	Show(postRequest ShowPostRequest) (*PostResponse, *error_responses.ErrorResponse)
	IncrementView(id int64) (int, *error_responses.ErrorResponse)
	ShowWithReposts(userID int64, page, perPage int) (*PostResponse, *error_responses.ErrorResponse)
	CreateShare(postID, userID int64) *error_responses.ErrorResponse
	GetShareCount(postID int64) (int, *error_responses.ErrorResponse)
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

func (r *PostRepoImpl) Create(req *CreatePostRequest, uctx *share.UserContext) *error_responses.ErrorResponse {
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
		fmt.Println("INSERT TAG =>", hashtag.Name, "sticker:", hashtag.StickerID)

		_, tagErr := r.dbpool.Exec(`
        INSERT INTO tbl_post_hashtags (post_id, tag_name, sticker_id)
        VALUES ($1,$2,$3)
    `, post.ID, hashtag.Name, hashtag.StickerID)

		if tagErr != nil {
			fmt.Println("INSERT TAG ERROR:", tagErr)
			return msg.NewErrorResponse("database_error", tagErr)
		}
	}

	fmt.Println("STICKER IDS TO INSERT:", req.StickerIDs)
	for _, stickerID := range req.StickerIDs {
		fmt.Println("INSERT STICKER =>", stickerID)
		fmt.Println("STICKER IDS TO INSERT:", req.StickerIDs)

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

// func (r *PostRepoImpl) Show(postRequest ShowPostRequest) (*PostResponse, *error_responses.ErrorResponse) {
// 	var per_page = postRequest.PageOption.Perpage
// 	var page = postRequest.PageOption.Page
// 	var offset = (page - 1) * per_page
// 	var limit_clause = fmt.Sprintf(" LIMIT %d OFFSET %d", per_page, offset)
// 	var sql_orderby string

// 	if len(postRequest.Sorts) == 0 {
// 		sql_orderby = "ORDER BY sort_at DESC"
// 	} else {
// 		sql_orderby = custom_sql.BuildSQLSort(postRequest.Sorts)
// 	}

// 	sql_filters, args_filters := custom_sql.BuildSQLFilter(postRequest.Filters)

// 	branch1Filters := ""
// 	branch2Filters := ""

// 	if len(args_filters) > 0 {
// 		branch1Filters = " AND " + sql_filters
// 		branch2Filters = " AND " + strings.ReplaceAll(sql_filters, "p.user_id", "shr.user_id")
// 	}

// 	if searchClause, searchArgs := custom_sql.BuildSQLSearch(
// 		[]string{"p.caption"},
// 		postRequest.Search, len(args_filters)+1,
// 	); searchClause != "" {
// 		branch1Filters += " AND " + searchClause
// 		branch2Filters += " AND " + searchClause
// 		args_filters = append(args_filters, searchArgs...)
// 	}

// 	fmt.Println("branch1:", branch1Filters, "branch2:", branch2Filters)
// 	msg := error_responses.ErrorResponse{}
// 	var posts []Post

// 	query := fmt.Sprintf(`
//     WITH combined AS (
//         SELECT
//             p.id, p.user_id, u.user_name AS user_name, u.profile_images AS profile_images,
//             p.community_id, p.caption, p.post_type, p.code_content, p.link_url,
//             p.views_count, p.created_at, p.updated_at,
//             STRING_AGG(DISTINCT pi.image_url, ',') AS images,
//             STRING_AGG(
//                 DISTINCT ph.tag_name || '::' || COALESCE(ph.sticker_id::text, ''),
//                 ',' ORDER BY ph.tag_name || '::' || COALESCE(ph.sticker_id::text, '')
//             ) AS tag_data,
//             STRING_AGG(DISTINCT ps.sticker_id::text, ',') AS sticker_ids,
//             STRING_AGG(DISTINCT ph.sticker_id::text, ',') AS tag_sticker_ids,
//             pv.video_path AS video_path, pv.duration AS duration, pv.thumbnail_path AS thumbnail_path,
//             COALESCE(cc.comment_count, 0) AS comment_count,
//             NULL::bigint AS repost_id,
//             NULL::bigint AS reposted_by_user_id,
//             NULL::text AS reposted_by_username,
//             NULL::text AS reposted_by_profile_images,
//             NULL::timestamp AS reposted_at,
//             p.created_at AS sort_at
//         FROM tbl_posts p
//         LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
//         LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
//         LEFT JOIN tbl_post_stickers ps ON ps.post_id = p.id
//         LEFT JOIN tbl_users u ON u.id = p.user_id
//         LEFT JOIN (SELECT post_id, COUNT(*) AS comment_count FROM tbl_comments GROUP BY post_id) cc ON cc.post_id = p.id
//         LEFT JOIN tbl_post_videos pv ON pv.post_id = p.id
//         WHERE 1=1 %s
//         GROUP BY p.id, u.user_name, u.profile_images, cc.comment_count, pv.video_path, pv.duration, pv.thumbnail_path

//         UNION ALL

//         SELECT
//             p.id, p.user_id, u.user_name AS user_name, u.profile_images AS profile_images,
//             p.community_id, p.caption, p.post_type, p.code_content, p.link_url,
//             p.views_count, p.created_at, p.updated_at,
//             STRING_AGG(DISTINCT pi.image_url, ',') AS images,
//             STRING_AGG(
//                 DISTINCT ph.tag_name || '::' || COALESCE(ph.sticker_id::text, ''),
//                 ',' ORDER BY ph.tag_name || '::' || COALESCE(ph.sticker_id::text, '')
//             ) AS tag_data,
//             STRING_AGG(DISTINCT ps.sticker_id::text, ',') AS sticker_ids,
//             STRING_AGG(DISTINCT ph.sticker_id::text, ',') AS tag_sticker_ids,
//             pv.video_path AS video_path, pv.duration AS duration, pv.thumbnail_path AS thumbnail_path,
//             COALESCE(cc.comment_count, 0) AS comment_count,
//             shr.id AS repost_id,
//             shr.user_id AS reposted_by_user_id,
//             ru.user_name AS reposted_by_username,
//             ru.profile_images AS reposted_by_profile_images,
//             shr.created_at AS reposted_at,
//             shr.created_at AS sort_at
//         FROM tbl_post_shares shr
//         JOIN tbl_posts p ON p.id = shr.post_id
//         JOIN tbl_users ru ON ru.id = shr.user_id
//         LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
//         LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
//         LEFT JOIN tbl_post_stickers ps ON ps.post_id = p.id
//         LEFT JOIN tbl_users u ON u.id = p.user_id
//         LEFT JOIN (SELECT post_id, COUNT(*) AS comment_count FROM tbl_comments GROUP BY post_id) cc ON cc.post_id = p.id
//         LEFT JOIN tbl_post_videos pv ON pv.post_id = p.id
//         WHERE 1=1 %s
//         GROUP BY p.id, u.user_name, u.profile_images, cc.comment_count, pv.video_path, pv.duration, pv.thumbnail_path,
//                  shr.id, shr.user_id, ru.user_name, ru.profile_images, shr.created_at
//     )
//     SELECT
//         id, user_id, user_name, profile_images, community_id, caption,
//         post_type, code_content, link_url, views_count, created_at, updated_at,
//         images, tag_data, sticker_ids, tag_sticker_ids, video_path, duration, thumbnail_path,
//         comment_count, repost_id, reposted_by_user_id, reposted_by_username,
//         reposted_by_profile_images, reposted_at
//     FROM combined
//     %s
//     %s`,
// 		branch1Filters,
// 		branch2Filters,
// 		sql_orderby,
// 		limit_clause,
// 	)

// 	err := r.dbpool.Select(&posts, query, args_filters...)
// 	if err != nil {
// 		return nil, msg.NewErrorResponse("database_error", err)
// 	}

// 	var total int
// 	countQuery := fmt.Sprintf(`
//         SELECT
//             (SELECT COUNT(DISTINCT p.id) FROM tbl_posts p
//              LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
//              LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
//              WHERE 1=1 %s) +
//             (SELECT COUNT(*) FROM tbl_post_shares shr WHERE 1=1 %s)
//     `, branch1Filters, branch2Filters)

// 	err = r.dbpool.Get(&total, countQuery, args_filters...)
// 	if err != nil {
// 		return nil, msg.NewErrorResponse("database_error", err)
// 	}

//		return &PostResponse{Posts: posts, Total: total}, nil
//	}

// =====================================
// func (r *PostRepoImpl) Show(postRequest ShowPostRequest) (*PostResponse, *error_responses.ErrorResponse) {
// 	var per_page = postRequest.PageOption.Perpage
// 	var page = postRequest.PageOption.Page
// 	var offset = (page - 1) * per_page
// 	var limit_clause = fmt.Sprintf(" LIMIT %d OFFSET %d", per_page, offset)
// 	var sql_orderby string

// 	if len(postRequest.Sorts) == 0 {
// 		sql_orderby = "ORDER BY sort_at DESC"
// 	} else {
// 		sql_orderby = custom_sql.BuildSQLSort(postRequest.Sorts)
// 	}

// 	sql_filters, args_filters := custom_sql.BuildSQLFilter(postRequest.Filters)

// 	branch1Filters := ""
// 	branch2Filters := ""
// 	branch3Filters := "" // NEW
// 	if len(args_filters) > 0 {
// 		branch1Filters = " AND " + sql_filters
// 		branch2Filters = " AND " + strings.ReplaceAll(sql_filters, "p.user_id", "shr.user_id")
// 		branch3Filters = " AND " + strings.ReplaceAll(sql_filters, "p.user_id", "qs.user_id") // NEW
// 	}

// 	if searchClause, searchArgs := custom_sql.BuildSQLSearch(
// 		[]string{"p.caption"},
// 		postRequest.Search, len(args_filters)+1,
// 	); searchClause != "" {
// 		branch1Filters += " AND " + searchClause
// 		branch2Filters += " AND " + searchClause
// 		branch3Filters += " AND " + strings.ReplaceAll(searchClause, "p.caption", "q.title") // NEW
// 		args_filters = append(args_filters, searchArgs...)
// 	}

// 	msg := error_responses.ErrorResponse{}
// 	var posts []Post

// 	query := fmt.Sprintf(`
//     WITH combined AS (
//         SELECT
//             p.id, p.user_id, u.user_name AS user_name, u.profile_images AS profile_images,
//             p.community_id, p.caption, p.post_type, p.code_content, p.link_url,
//             p.views_count, p.created_at, p.updated_at,
//             STRING_AGG(DISTINCT pi.image_url, ',') AS images,
//             STRING_AGG(
//                 DISTINCT ph.tag_name || '::' || COALESCE(ph.sticker_id::text, ''),
//                 ',' ORDER BY ph.tag_name || '::' || COALESCE(ph.sticker_id::text, '')
//             ) AS tag_data,
//             STRING_AGG(DISTINCT ps.sticker_id::text, ',') AS sticker_ids,
//             STRING_AGG(DISTINCT ph.sticker_id::text, ',') AS tag_sticker_ids,
//             pv.video_path AS video_path, pv.duration AS duration, pv.thumbnail_path AS thumbnail_path,
//             COALESCE(cc.comment_count, 0) AS comment_count,
//             NULL::bigint AS repost_id,
//             NULL::bigint AS reposted_by_user_id,
//             NULL::text AS reposted_by_username,
//             NULL::text AS reposted_by_profile_images,
//             NULL::timestamp AS reposted_at,
//             p.created_at AS sort_at
//         FROM tbl_posts p
//         LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
//         LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
//         LEFT JOIN tbl_post_stickers ps ON ps.post_id = p.id
//         LEFT JOIN tbl_users u ON u.id = p.user_id
//         LEFT JOIN (SELECT post_id, COUNT(*) AS comment_count FROM tbl_comments GROUP BY post_id) cc ON cc.post_id = p.id
//         LEFT JOIN tbl_post_videos pv ON pv.post_id = p.id
//         WHERE 1=1 %s
//         GROUP BY p.id, u.user_name, u.profile_images, cc.comment_count, pv.video_path, pv.duration, pv.thumbnail_path

//         UNION ALL

//         SELECT
//             p.id, p.user_id, u.user_name AS user_name, u.profile_images AS profile_images,
//             p.community_id, p.caption, p.post_type, p.code_content, p.link_url,
//             p.views_count, p.created_at, p.updated_at,
//             STRING_AGG(DISTINCT pi.image_url, ',') AS images,
//             STRING_AGG(
//                 DISTINCT ph.tag_name || '::' || COALESCE(ph.sticker_id::text, ''),
//                 ',' ORDER BY ph.tag_name || '::' || COALESCE(ph.sticker_id::text, '')
//             ) AS tag_data,
//             STRING_AGG(DISTINCT ps.sticker_id::text, ',') AS sticker_ids,
//             STRING_AGG(DISTINCT ph.sticker_id::text, ',') AS tag_sticker_ids,
//             pv.video_path AS video_path, pv.duration AS duration, pv.thumbnail_path AS thumbnail_path,
//             COALESCE(cc.comment_count, 0) AS comment_count,
//             shr.id AS repost_id,
//             shr.user_id AS reposted_by_user_id,
//             ru.user_name AS reposted_by_username,
//             ru.profile_images AS reposted_by_profile_images,
//             shr.created_at AS reposted_at,
//             shr.created_at AS sort_at
//         FROM tbl_post_shares shr
//         JOIN tbl_posts p ON p.id = shr.post_id
//         JOIN tbl_users ru ON ru.id = shr.user_id
//         LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
//         LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
//         LEFT JOIN tbl_post_stickers ps ON ps.post_id = p.id
//         LEFT JOIN tbl_users u ON u.id = p.user_id
//         LEFT JOIN (SELECT post_id, COUNT(*) AS comment_count FROM tbl_comments GROUP BY post_id) cc ON cc.post_id = p.id
//         LEFT JOIN tbl_post_videos pv ON pv.post_id = p.id
//         WHERE 1=1 %s
//         GROUP BY p.id, u.user_name, u.profile_images, cc.comment_count, pv.video_path, pv.duration, pv.thumbnail_path,
//                  shr.id, shr.user_id, ru.user_name, ru.profile_images, shr.created_at

//         UNION ALL

//         SELECT
//             (2000000000 + qs.id) AS id,
//             qs.user_id AS user_id,
//             sharer.user_name AS user_name,
//             sharer.profile_images AS profile_images,
//             q.id AS community_id,
//             q.title AS caption,
//             'quote_share' AS post_type,
//             q.content AS code_content,
//             NULL::text AS link_url,
//             0 AS views_count,
//             qs.created_at AS created_at,
//             qs.created_at AS updated_at,
//             NULL::text AS images,
//             NULL::text AS tag_data,
//             NULL::text AS sticker_ids,
//             NULL::text AS tag_sticker_ids,
//             NULL::text AS video_path,
//             NULL::int AS duration,
//             NULL::text AS thumbnail_path,
//             0 AS comment_count,
//             qs.id AS repost_id,
//             qs.user_id AS reposted_by_user_id,
//             sharer.user_name AS reposted_by_username,
//             sharer.profile_images AS reposted_by_profile_images,
//             qs.created_at AS reposted_at,
//             qs.created_at AS sort_at
//         FROM quote_shares qs
//         JOIN quotes q ON q.id = qs.quote_id
//         JOIN tbl_users sharer ON sharer.id = qs.user_id
//         WHERE qs.channel = 'feed' %s
//     )
//     SELECT
//         id, user_id, user_name, profile_images, community_id, caption,
//         post_type, code_content, link_url, views_count, created_at, updated_at,
//         images, tag_data, sticker_ids, tag_sticker_ids, video_path, duration, thumbnail_path,
//         comment_count, repost_id, reposted_by_user_id, reposted_by_username,
//         reposted_by_profile_images, reposted_at
//     FROM combined
//     %s
//     %s`,
// 		branch1Filters,
// 		branch2Filters,
// 		branch3Filters,
// 		sql_orderby,
// 		limit_clause,
// 	)

// 	allArgs := append(append(append([]any{}, args_filters...), args_filters...), args_filters...)

// 	err := r.dbpool.Select(&posts, query, allArgs...)
// 	if err != nil {
// 		return nil, msg.NewErrorResponse("database_error", err)
// 	}

// 	var total int
// 	countQuery := fmt.Sprintf(`
//         SELECT
//             (SELECT COUNT(DISTINCT p.id) FROM tbl_posts p
//              LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
//              LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
//              WHERE 1=1 %s) +
//             (SELECT COUNT(*) FROM tbl_post_shares shr WHERE 1=1 %s) +
//             (SELECT COUNT(*) FROM quote_shares qs WHERE qs.channel = 'feed' %s)
//     `, branch1Filters, branch2Filters, branch3Filters)

// 	countArgs := append(append(append([]any{}, args_filters...), args_filters...), args_filters...)

// 	err = r.dbpool.Get(&total, countQuery, countArgs...)
// 	if err != nil {
// 		return nil, msg.NewErrorResponse("database_error", err)
// 	}

// 	return &PostResponse{Posts: posts, Total: total}, nil
// }

// =========================================
func (r *PostRepoImpl) Show(postRequest ShowPostRequest) (*PostResponse, *error_responses.ErrorResponse) {
	var per_page = postRequest.PageOption.Perpage
	var page = postRequest.PageOption.Page
	var offset = (page - 1) * per_page
	var limit_clause = fmt.Sprintf(" LIMIT %d OFFSET %d", per_page, offset)
	var sql_orderby string

	if len(postRequest.Sorts) == 0 {
		sql_orderby = "ORDER BY sort_at DESC"
	} else {
		sql_orderby = custom_sql.BuildSQLSort(postRequest.Sorts)
	}

	sql_filters, args_filters := custom_sql.BuildSQLFilter(postRequest.Filters)

	branch1Filters := ""
	branch2Filters := ""
	branch3Filters := "" // NEW — quote shares
	if len(args_filters) > 0 {
		branch1Filters = " AND " + sql_filters
		branch2Filters = " AND " + strings.ReplaceAll(sql_filters, "p.user_id", "shr.user_id")
		branch3Filters = " AND " + strings.ReplaceAll(sql_filters, "p.user_id", "qs.user_id") // NEW
	}

	if searchClause, searchArgs := custom_sql.BuildSQLSearch(
		[]string{"p.caption"},
		postRequest.Search, len(args_filters)+1,
	); searchClause != "" {
		branch1Filters += " AND " + searchClause
		branch2Filters += " AND " + searchClause
		branch3Filters += " AND " + strings.ReplaceAll(searchClause, "p.caption", "q.title") // NEW
		args_filters = append(args_filters, searchArgs...)
	}

	msg := error_responses.ErrorResponse{}
	var posts []Post

	query := fmt.Sprintf(`
    WITH combined AS (
        SELECT
            p.id, p.user_id, u.user_name AS user_name, u.profile_images AS profile_images,
            p.community_id, p.caption, p.post_type, p.code_content, p.link_url,
            p.views_count, p.created_at, p.updated_at,
            STRING_AGG(DISTINCT pi.image_url, ',') AS images,
            STRING_AGG(
                DISTINCT ph.tag_name || '::' || COALESCE(ph.sticker_id::text, ''),
                ',' ORDER BY ph.tag_name || '::' || COALESCE(ph.sticker_id::text, '')
            ) AS tag_data,
            STRING_AGG(DISTINCT ps.sticker_id::text, ',') AS sticker_ids,
            STRING_AGG(DISTINCT ph.sticker_id::text, ',') AS tag_sticker_ids,
            pv.video_path AS video_path, pv.duration AS duration, pv.thumbnail_path AS thumbnail_path,
            COALESCE(cc.comment_count, 0) AS comment_count,
            NULL::bigint AS repost_id,
            NULL::bigint AS reposted_by_user_id,
            NULL::text AS reposted_by_username,
            NULL::text AS reposted_by_profile_images,
            NULL::timestamp AS reposted_at,
            p.created_at AS sort_at
        FROM tbl_posts p
        LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
        LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
        LEFT JOIN tbl_post_stickers ps ON ps.post_id = p.id
        LEFT JOIN tbl_users u ON u.id = p.user_id
        LEFT JOIN (SELECT post_id, COUNT(*) AS comment_count FROM tbl_comments GROUP BY post_id) cc ON cc.post_id = p.id
        LEFT JOIN tbl_post_videos pv ON pv.post_id = p.id
        WHERE 1=1 %s
        GROUP BY p.id, u.user_name, u.profile_images, cc.comment_count, pv.video_path, pv.duration, pv.thumbnail_path

        UNION ALL

        SELECT
            p.id, p.user_id, u.user_name AS user_name, u.profile_images AS profile_images,
            p.community_id, p.caption, p.post_type, p.code_content, p.link_url,
            p.views_count, p.created_at, p.updated_at,
            STRING_AGG(DISTINCT pi.image_url, ',') AS images,
            STRING_AGG(
                DISTINCT ph.tag_name || '::' || COALESCE(ph.sticker_id::text, ''),
                ',' ORDER BY ph.tag_name || '::' || COALESCE(ph.sticker_id::text, '')
            ) AS tag_data,
            STRING_AGG(DISTINCT ps.sticker_id::text, ',') AS sticker_ids,
            STRING_AGG(DISTINCT ph.sticker_id::text, ',') AS tag_sticker_ids,
            pv.video_path AS video_path, pv.duration AS duration, pv.thumbnail_path AS thumbnail_path,
            COALESCE(cc.comment_count, 0) AS comment_count,
            shr.id AS repost_id,
            shr.user_id AS reposted_by_user_id,
            ru.user_name AS reposted_by_username,
            ru.profile_images AS reposted_by_profile_images,
            shr.created_at AS reposted_at,
            shr.created_at AS sort_at
        FROM tbl_post_shares shr
        JOIN tbl_posts p ON p.id = shr.post_id
        JOIN tbl_users ru ON ru.id = shr.user_id
        LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
        LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
        LEFT JOIN tbl_post_stickers ps ON ps.post_id = p.id
        LEFT JOIN tbl_users u ON u.id = p.user_id
        LEFT JOIN (SELECT post_id, COUNT(*) AS comment_count FROM tbl_comments GROUP BY post_id) cc ON cc.post_id = p.id
        LEFT JOIN tbl_post_videos pv ON pv.post_id = p.id
        WHERE 1=1 %s
        GROUP BY p.id, u.user_name, u.profile_images, cc.comment_count, pv.video_path, pv.duration, pv.thumbnail_path,
                 shr.id, shr.user_id, ru.user_name, ru.profile_images, shr.created_at

        UNION ALL

        SELECT
            (2000000000 + qs.id) AS id,
            qs.user_id AS user_id,
            sharer.user_name AS user_name,
            sharer.profile_images AS profile_images,
            q.id AS community_id,
            q.title AS caption,
            'quote_share' AS post_type,
            q.content AS code_content,
            NULL::text AS link_url,
            0 AS views_count,
            qs.created_at AS created_at,
            qs.created_at AS updated_at,
            NULL::text AS images,
            NULL::text AS tag_data,
            NULL::text AS sticker_ids,
            NULL::text AS tag_sticker_ids,
            NULL::text AS video_path,
            NULL::int AS duration,
            NULL::text AS thumbnail_path,
            0 AS comment_count,
            qs.id AS repost_id,
            qs.user_id AS reposted_by_user_id,
            sharer.user_name AS reposted_by_username,
            sharer.profile_images AS reposted_by_profile_images,
            qs.created_at AS reposted_at,
            qs.created_at AS sort_at
        FROM quote_shares qs
        JOIN quotes q ON q.id = qs.quote_id
        JOIN tbl_users sharer ON sharer.id = qs.user_id
        WHERE qs.channel = 'feed' %s
    )
    SELECT
        id, user_id, user_name, profile_images, community_id, caption,
        post_type, code_content, link_url, views_count, created_at, updated_at,
        images, tag_data, sticker_ids, tag_sticker_ids, video_path, duration, thumbnail_path,
        comment_count, repost_id, reposted_by_user_id, reposted_by_username,
        reposted_by_profile_images, reposted_at
    FROM combined
    %s
    %s`,
		branch1Filters,
		branch2Filters,
		branch3Filters,
		sql_orderby,
		limit_clause,
	)

	// ✅ args_filters ត្រូវផ្ញើតែម្តងគត់ — $1,$2... ដដែលត្រូវបានប្រើឡើងវិញរាល់ branch
	err := r.dbpool.Select(&posts, query, args_filters...)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var total int
	countQuery := fmt.Sprintf(`
        SELECT
            (SELECT COUNT(DISTINCT p.id) FROM tbl_posts p
             LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
             LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
             WHERE 1=1 %s) +
            (SELECT COUNT(*) FROM tbl_post_shares shr WHERE 1=1 %s) +
            (SELECT COUNT(*) FROM quote_shares qs WHERE qs.channel = 'feed' %s)
    `, branch1Filters, branch2Filters, branch3Filters)

	// ✅ ដូចគ្នា — args_filters តែម្តងគត់
	err = r.dbpool.Get(&total, countQuery, args_filters...)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &PostResponse{Posts: posts, Total: total}, nil
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

func (r *PostRepoImpl) ShowWithReposts(userID int64, page, perPage int) (*PostResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	offset := (page - 1) * perPage

	var posts []Post
	query := `
    WITH combined AS (
        SELECT
            p.id, p.user_id, u.user_name AS user_name, u.profile_images AS profile_images,
            p.community_id, p.caption, p.post_type, p.code_content, p.link_url,
            p.views_count, p.created_at, p.updated_at,
            STRING_AGG(DISTINCT pi.image_url, ',') AS images,
            STRING_AGG(DISTINCT ph.tag_name, ',') AS tag_name,
            STRING_AGG(DISTINCT ps.sticker_id::text, ',') AS sticker_ids,
            STRING_AGG(DISTINCT ph.sticker_id::text, ',') AS tag_sticker_ids,
            pv.video_path AS video_path, pv.duration AS duration, pv.thumbnail_path AS thumbnail_path,
            COALESCE(cc.comment_count, 0) AS comment_count,
            NULL::bigint AS repost_id,
            NULL::bigint AS reposted_by_user_id,
            NULL::text AS reposted_by_username,
            NULL::timestamp AS reposted_at,
            p.created_at AS sort_at
        FROM tbl_posts p
        LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
        LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
        LEFT JOIN tbl_post_stickers ps ON ps.post_id = p.id
        LEFT JOIN tbl_users u ON u.id = p.user_id
        LEFT JOIN (SELECT post_id, COUNT(*) AS comment_count FROM tbl_comments GROUP BY post_id) cc ON cc.post_id = p.id
        LEFT JOIN tbl_post_videos pv ON pv.post_id = p.id
        WHERE p.user_id = $1
        GROUP BY p.id, u.user_name, u.profile_images, cc.comment_count, pv.video_path, pv.duration, pv.thumbnail_path

        UNION ALL

        SELECT
            p.id, p.user_id, u.user_name AS user_name, u.profile_images AS profile_images,
            p.community_id, p.caption, p.post_type, p.code_content, p.link_url,
            p.views_count, p.created_at, p.updated_at,
            STRING_AGG(DISTINCT pi.image_url, ',') AS images,
            STRING_AGG(DISTINCT ph.tag_name, ',') AS tag_name,
            STRING_AGG(DISTINCT ps.sticker_id::text, ',') AS sticker_ids,
            STRING_AGG(DISTINCT ph.sticker_id::text, ',') AS tag_sticker_ids,
            pv.video_path AS video_path, pv.duration AS duration, pv.thumbnail_path AS thumbnail_path,
            COALESCE(cc.comment_count, 0) AS comment_count,
            shr.id AS repost_id,
            shr.user_id AS reposted_by_user_id,
            ru.user_name AS reposted_by_username,
            shr.created_at AS reposted_at,
            shr.created_at AS sort_at
        FROM tbl_post_shares shr
        JOIN tbl_posts p ON p.id = shr.post_id
        JOIN tbl_users ru ON ru.id = shr.user_id
        LEFT JOIN tbl_post_images pi ON pi.post_id = p.id
        LEFT JOIN tbl_post_hashtags ph ON ph.post_id = p.id
        LEFT JOIN tbl_post_stickers ps ON ps.post_id = p.id
        LEFT JOIN tbl_users u ON u.id = p.user_id
        LEFT JOIN (SELECT post_id, COUNT(*) AS comment_count FROM tbl_comments GROUP BY post_id) cc ON cc.post_id = p.id
        LEFT JOIN tbl_post_videos pv ON pv.post_id = p.id
        WHERE shr.user_id = $1
        GROUP BY p.id, u.user_name, u.profile_images, cc.comment_count, pv.video_path, pv.duration, pv.thumbnail_path,
                 shr.id, shr.user_id, ru.user_name, shr.created_at
    )
    SELECT * FROM combined
    ORDER BY sort_at DESC
    LIMIT $2 OFFSET $3
    `

	err := r.dbpool.Select(&posts, query, userID, perPage, offset)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var total int
	countQuery := `
        SELECT
            (SELECT COUNT(*) FROM tbl_posts WHERE user_id = $1) +
            (SELECT COUNT(*) FROM tbl_post_shares WHERE user_id = $1)
    `
	err = r.dbpool.Get(&total, countQuery, userID)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &PostResponse{Posts: posts, Total: total}, nil
}

func (r *PostRepoImpl) CreateShare(postID, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	_, err := r.dbpool.Exec(`
        INSERT INTO tbl_post_shares (post_id, user_id, created_at)
        VALUES ($1, $2, NOW())
        ON CONFLICT (post_id, user_id) DO NOTHING
    `, postID, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *PostRepoImpl) GetShareCount(postID int64) (int, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var count int
	err := r.dbpool.Get(&count, `
        SELECT COUNT(*) FROM tbl_post_shares WHERE post_id = $1
    `, postID)
	if err != nil {
		return 0, msg.NewErrorResponse("database_error", err)
	}
	return count, nil
}
