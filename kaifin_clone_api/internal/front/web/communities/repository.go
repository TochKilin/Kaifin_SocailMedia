package communities

import (
	"strconv"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type CommunitiesRepo interface {
	Create(community *Communities) *error_responses.ErrorResponse
	Show(req ShowCommunitiesRequest, userID int64) (*CommunitiesResponse, *error_responses.ErrorResponse)
	ShowDetail(id int64, userID int64) (*Communities, *error_responses.ErrorResponse)
	ToggleJoin(communityID, userID int64) (bool, string, *error_responses.ErrorResponse)
	ShowMembers(req ShowMembersRequest) (*MembersResponse, *error_responses.ErrorResponse)
	UpdateAvatar(communityID int64, avatarURL string) *error_responses.ErrorResponse // ➜ ថ្មី
	UpdateCover(communityID int64, coverURL string) *error_responses.ErrorResponse
}

type CommunitiesRepoImpl struct {
	dbpool *sqlx.DB
}

func NewCommunitiesRepoImpl(db *sqlx.DB) *CommunitiesRepoImpl {
	return &CommunitiesRepoImpl{
		dbpool: db,
	}
}

func (r *CommunitiesRepoImpl) Create(community *Communities) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}

	tx, err := r.dbpool.Beginx()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	err = tx.QueryRow(`
		INSERT INTO tbl_communities
			(name, description, avatar_url, cover_url, category_id, privacy, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		RETURNING id, is_verified, member_count, hot_score, created_at, updated_at
	`, community.Name, community.Description, community.AvatarURL, community.CoverURL,
		community.CategoryID, community.Privacy, community.CreatedBy,
	).Scan(&community.ID, &community.IsVerified, &community.MemberCount, &community.HotScore,
		&community.CreatedAt, &community.UpdatedAt)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	if community.CreatedBy != nil {
		if _, err = tx.Exec(`
			INSERT INTO tbl_community_members (community_id, user_id, role, status, joined_at)
			VALUES ($1, $2, 'admin', 'approved', NOW())
		`, community.ID, *community.CreatedBy); err != nil {
			return msg.NewErrorResponse("database_error", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *CommunitiesRepoImpl) Show(req ShowCommunitiesRequest, userID int64) (*CommunitiesResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var communities []Communities

	query := `
		SELECT
			c.id, c.name, c.description, c.avatar_url, c.cover_url,
			c.category_id, c.is_verified, c.privacy, c.member_count,
			c.hot_score, c.created_at, c.updated_at, c.created_by,
			COALESCE(m.is_joined, false) AS is_joined
		FROM tbl_communities c
		LEFT JOIN (
			SELECT community_id, true AS is_joined
			FROM tbl_community_members
			WHERE user_id = $1 AND status = 'approved'
		) m ON m.community_id = c.id
		WHERE 1=1
	`
	args := []interface{}{userID}
	index := 2

	if req.Search != "" {
		query += " AND c.name ILIKE $" + strconv.Itoa(index)
		args = append(args, "%"+req.Search+"%")
		index++
	}

	if req.CategoryID != nil {
		query += " AND c.category_id = $" + strconv.Itoa(index)
		args = append(args, *req.CategoryID)
		index++
	}

	query += " ORDER BY c.hot_score DESC, c.created_at DESC"

	perpage := req.PageOption.Perpage
	page := req.PageOption.Page
	if perpage <= 0 {
		perpage = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * perpage

	query += " LIMIT $" + strconv.Itoa(index) + " OFFSET $" + strconv.Itoa(index+1)
	args = append(args, perpage, offset)

	if err := r.dbpool.Select(&communities, query, args...); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	countQuery := `SELECT COUNT(*) FROM tbl_communities c WHERE 1=1`
	countArgs := []interface{}{}
	cIndex := 1
	if req.Search != "" {
		countQuery += " AND c.name ILIKE $" + strconv.Itoa(cIndex)
		countArgs = append(countArgs, "%"+req.Search+"%")
		cIndex++
	}
	if req.CategoryID != nil {
		countQuery += " AND c.category_id = $" + strconv.Itoa(cIndex)
		countArgs = append(countArgs, *req.CategoryID)
		cIndex++
	}

	var total int
	if err := r.dbpool.Get(&total, countQuery, countArgs...); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &CommunitiesResponse{Communities: communities, Total: total}, nil
}

func (r *CommunitiesRepoImpl) ShowDetail(id int64, userID int64) (*Communities, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var community Communities
	err := r.dbpool.Get(&community, `
		SELECT
			c.id, c.name, c.description, c.avatar_url, c.cover_url,
			c.category_id, c.is_verified, c.privacy, c.member_count,
			c.hot_score, c.created_at, c.updated_at, c.created_by,
			COALESCE(m.is_joined, false) AS is_joined
		FROM tbl_communities c
		LEFT JOIN (
			SELECT community_id, true AS is_joined
			FROM tbl_community_members
			WHERE user_id = $2 AND status = 'approved'
		) m ON m.community_id = c.id
		WHERE c.id = $1
	`, id, userID)
	if err != nil {
		return nil, msg.NewErrorResponse("community_not_found", err)
	}

	return &community, nil
}

func (r *CommunitiesRepoImpl) ToggleJoin(communityID, userID int64) (bool, string, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var exists bool
	err := r.dbpool.Get(&exists, `
		SELECT EXISTS(
			SELECT 1 FROM tbl_community_members
			WHERE community_id = $1 AND user_id = $2
		)
	`, communityID, userID)
	if err != nil {
		return false, "", msg.NewErrorResponse("database_error", err)
	}

	if exists {
		_, err = r.dbpool.Exec(`
			DELETE FROM tbl_community_members
			WHERE community_id = $1 AND user_id = $2
		`, communityID, userID)
		if err != nil {
			return false, "", msg.NewErrorResponse("database_error", err)
		}
		return false, "left", nil
	}

	var privacy string
	if err := r.dbpool.Get(&privacy, `SELECT privacy FROM tbl_communities WHERE id = $1`, communityID); err != nil {
		return false, "", msg.NewErrorResponse("community_not_found", err)
	}

	status := "approved"
	if privacy == "private" {
		status = "pending"
	}

	_, err = r.dbpool.Exec(`
		INSERT INTO tbl_community_members (community_id, user_id, role, status, joined_at)
		VALUES ($1, $2, 'member', $3, NOW())
	`, communityID, userID, status)
	if err != nil {
		return false, "", msg.NewErrorResponse("database_error", err)
	}

	return status == "approved", status, nil
}

func (r *CommunitiesRepoImpl) ShowMembers(req ShowMembersRequest) (*MembersResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	perpage := req.PageOption.Perpage
	page := req.PageOption.Page
	if perpage <= 0 {
		perpage = 30
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * perpage

	var members []CommunityMember
	err := r.dbpool.Select(&members, `
		SELECT
			m.id, m.community_id, m.user_id, u.user_name AS user_name,
			u.profile_images AS profile_images, m.role, m.status, m.joined_at
		FROM tbl_community_members m
		LEFT JOIN tbl_users u ON u.id = m.user_id
		WHERE m.community_id = $1 AND m.status = 'approved'
		ORDER BY
			CASE m.role WHEN 'admin' THEN 0 WHEN 'moderator' THEN 1 ELSE 2 END,
			m.joined_at ASC
		LIMIT $2 OFFSET $3
	`, req.CommunityID, perpage, offset)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var total int
	if err := r.dbpool.Get(&total, `
		SELECT COUNT(*) FROM tbl_community_members
		WHERE community_id = $1 AND status = 'approved'
	`, req.CommunityID); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &MembersResponse{Members: members, Total: total}, nil
}

func (r *CommunitiesRepoImpl) UpdateAvatar(communityID int64, avatarURL string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	res, err := r.dbpool.Exec(`
		UPDATE tbl_communities SET avatar_url = $1, updated_at = NOW() WHERE id = $2
	`, avatarURL, communityID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("community_not_found", nil)
	}
	return nil
}

func (r *CommunitiesRepoImpl) UpdateCover(communityID int64, coverURL string) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	res, err := r.dbpool.Exec(`
		UPDATE tbl_communities SET cover_url = $1, updated_at = NOW() WHERE id = $2
	`, coverURL, communityID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return msg.NewErrorResponse("community_not_found", nil)
	}
	return nil
}
