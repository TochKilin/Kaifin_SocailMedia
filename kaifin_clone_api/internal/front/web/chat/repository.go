package chat

import (
	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

type ChatsRepo interface {
	ShowConversations(req ShowConversationsRequest, userID int64) (*ConversationsResponse, *error_responses.ErrorResponse)
	ShowMessages(req ShowMessagesRequest, userID int64) (*MessagesResponse, *error_responses.ErrorResponse)
	SendMessage(msg *Message, attachments []MessageAttachment) *error_responses.ErrorResponse
	ToggleReaction(messageID, userID, reactionTypeID int64) (bool, *error_responses.ErrorResponse)
	MarkAsRead(conversationID, userID int64) *error_responses.ErrorResponse
	IsParticipant(conversationID, userID int64) (bool, *error_responses.ErrorResponse)
	FindOrCreateDirectConversation(userID, targetUserID int64) (int64, *error_responses.ErrorResponse)
	SearchUsers(currentUserID int64, search string, limit int) ([]UserSearchItem, *error_responses.ErrorResponse)
}

type ChatsRepoImpl struct {
	dbpool *sqlx.DB
}

func NewChatsRepoImpl(db *sqlx.DB) *ChatsRepoImpl {
	return &ChatsRepoImpl{dbpool: db}
}

func (r *ChatsRepoImpl) ShowConversations(req ShowConversationsRequest, userID int64) (*ConversationsResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	query := `
		SELECT
			c.id AS conversation_id,
			(c.type = 'group') AS is_group,
			CASE WHEN c.type = 'group' THEN c.name ELSE ou.user_name END AS name,
			CASE WHEN c.type = 'group' THEN c.avatar ELSE ou.profile_images END AS avatar,
			COALESCE(ou.is_online, false) AS online,
			lm.content AS last_message,
			lm.type AS last_message_type,
			lm.created_at AS last_message_at,
			COALESCE(uc.unread_count, 0) AS unread_count,
			ou.id AS other_user_id
		FROM conversation_participants cp
		JOIN conversations c ON c.id = cp.conversation_id
		LEFT JOIN LATERAL (
			SELECT cp2.user_id AS id, u.user_name, u.profile_images,
			       (u.last_login > NOW() - INTERVAL '5 minutes') AS is_online
			FROM conversation_participants cp2
			JOIN tbl_users u ON u.id = cp2.user_id
			WHERE cp2.conversation_id = c.id AND cp2.user_id != $1 AND c.type != 'group'
			LIMIT 1
		) ou ON true
		LEFT JOIN LATERAL (
			SELECT m.content, m.type, m.created_at
			FROM messages m
			WHERE m.conversation_id = c.id
			ORDER BY m.created_at DESC
			LIMIT 1
		) lm ON true
		LEFT JOIN LATERAL (
			SELECT COUNT(*) AS unread_count
			FROM messages m
			WHERE m.conversation_id = c.id
			AND m.id > COALESCE(cp.last_read_message_id, 0)
			AND m.sender_id != $1
		) uc ON true
		WHERE cp.user_id = $1
		AND cp.status = 'accepted'
		AND cp.is_archived = false
	`

	args := []interface{}{userID}
	switch req.Tab {
	case "unread":
		query += ` AND COALESCE(uc.unread_count, 0) > 0`
	case "groups":
		query += ` AND c.type = 'group'`
	}
	if req.Search != "" {
		query += ` AND (CASE WHEN c.type = 'group' THEN c.name ELSE ou.user_name END) ILIKE $2`
		args = append(args, "%"+req.Search+"%")
	}
	query += ` ORDER BY COALESCE(lm.created_at, c.created_at) DESC`

	var items []ChatListItem
	if err := r.dbpool.Select(&items, query, args...); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &ConversationsResponse{Conversations: items, Total: len(items)}, nil
}

func (r *ChatsRepoImpl) IsParticipant(conversationID, userID int64) (bool, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var exists bool
	err := r.dbpool.Get(&exists, `
		SELECT EXISTS(
			SELECT 1 FROM conversation_participants
			WHERE conversation_id = $1 AND user_id = $2
		)
	`, conversationID, userID)
	if err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}
	return exists, nil
}

func (r *ChatsRepoImpl) ShowMessages(req ShowMessagesRequest, userID int64) (*MessagesResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	query := `
		SELECT
			m.id, m.conversation_id, m.sender_id,
			u.user_name AS sender_name, u.profile_images AS sender_avatar,
			m.content, m.type, m.forwarded_from_id, m.created_at
		FROM messages m
		JOIN tbl_users u ON u.id = m.sender_id
		WHERE m.conversation_id = $1
	`
	args := []interface{}{req.ConversationID}
	if req.BeforeID > 0 {
		query += ` AND m.id < $2 ORDER BY m.created_at DESC LIMIT $3`
		args = append(args, req.BeforeID, req.Limit)
	} else {
		query += ` ORDER BY m.created_at DESC LIMIT $2`
		args = append(args, req.Limit)
	}

	var messages []Message
	if err := r.dbpool.Select(&messages, query, args...); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	for i := range messages {
		messages[i].IsMine = messages[i].SenderID == userID
	}

	if len(messages) > 0 {
		ids := make([]int64, len(messages))
		idx := map[int64]int{}
		for i, m := range messages {
			ids[i] = m.ID
			idx[m.ID] = i
		}

		var attachments []MessageAttachment
		q, args2, _ := sqlx.In(`
			SELECT id, message_id, type, url, file_name, file_size, duration, sort_order
			FROM message_attachments
			WHERE message_id IN (?)
			ORDER BY sort_order ASC, id ASC
		`, ids)
		q = r.dbpool.Rebind(q)
		if err := r.dbpool.Select(&attachments, q, args2...); err != nil {
			return nil, msg.NewErrorResponse("database_error", err)
		}
		for _, a := range attachments {
			messages[idx[a.MessageID]].Attachments = append(messages[idx[a.MessageID]].Attachments, a)
		}

		var reactions []ReactionSummary
		q2, args3, _ := sqlx.In(`
			SELECT message_id, reaction_type_id, COUNT(*) AS count
			FROM message_reactions
			WHERE message_id IN (?)
			GROUP BY message_id, reaction_type_id
		`, ids)
		q2 = r.dbpool.Rebind(q2)
		if err := r.dbpool.Select(&reactions, q2, args3...); err != nil {
			return nil, msg.NewErrorResponse("database_error", err)
		}
		for _, rc := range reactions {
			messages[idx[rc.MessageID]].Reactions = append(messages[idx[rc.MessageID]].Reactions, rc)
		}
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return &MessagesResponse{Messages: messages, Total: len(messages)}, nil
}

func (r *ChatsRepoImpl) SendMessage(m *Message, attachments []MessageAttachment) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	tx, err := r.dbpool.Beginx()
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	defer tx.Rollback()

	err = tx.QueryRow(`
		INSERT INTO messages (conversation_id, sender_id, content, type, forwarded_from_id, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
		RETURNING id, created_at
	`, m.ConversationID, m.SenderID, m.Content, m.Type, m.ForwardedFromID,
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	for i := range attachments {
		attachments[i].MessageID = m.ID
		if _, err = tx.Exec(`
			INSERT INTO message_attachments (message_id, type, url, file_name, file_size, duration, sort_order)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, m.ID, attachments[i].Type, attachments[i].URL, attachments[i].FileName,
			attachments[i].FileSize, attachments[i].Duration, i,
		); err != nil {
			return msg.NewErrorResponse("database_error", err)
		}
	}

	if _, err = tx.Exec(`UPDATE conversations SET updated_at = NOW() WHERE id = $1`, m.ConversationID); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}

	if err := tx.Commit(); err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	m.Attachments = attachments
	return nil
}

func (r *ChatsRepoImpl) ToggleReaction(messageID, userID, reactionTypeID int64) (bool, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var exists bool
	err := r.dbpool.Get(&exists, `
		SELECT EXISTS(SELECT 1 FROM message_reactions WHERE message_id = $1 AND user_id = $2)
	`, messageID, userID)
	if err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}

	if exists {
		if _, err = r.dbpool.Exec(`
			DELETE FROM message_reactions WHERE message_id = $1 AND user_id = $2
		`, messageID, userID); err != nil {
			return false, msg.NewErrorResponse("database_error", err)
		}
		return false, nil
	}

	if _, err = r.dbpool.Exec(`
		INSERT INTO message_reactions (message_id, user_id, reaction_type_id)
		VALUES ($1, $2, $3)
	`, messageID, userID, reactionTypeID); err != nil {
		return false, msg.NewErrorResponse("database_error", err)
	}
	return true, nil
}

func (r *ChatsRepoImpl) MarkAsRead(conversationID, userID int64) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	_, err := r.dbpool.Exec(`
		UPDATE conversation_participants
		SET last_read_message_id = (SELECT MAX(id) FROM messages WHERE conversation_id = $1),
			last_read_at = NOW()
		WHERE conversation_id = $1 AND user_id = $2
	`, conversationID, userID)
	if err != nil {
		return msg.NewErrorResponse("database_error", err)
	}
	return nil
}

func (r *ChatsRepoImpl) FindOrCreateDirectConversation(userID, targetUserID int64) (int64, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var existingID int64
	err := r.dbpool.Get(&existingID, `
		SELECT cp1.conversation_id
		FROM conversation_participants cp1
		JOIN conversation_participants cp2 ON cp2.conversation_id = cp1.conversation_id
		JOIN conversations c ON c.id = cp1.conversation_id
		WHERE cp1.user_id = $1 AND cp2.user_id = $2 AND c.type = 'direct'
		LIMIT 1
	`, userID, targetUserID)
	if err == nil {
		return existingID, nil
	}

	tx, txErr := r.dbpool.Beginx()
	if txErr != nil {
		return 0, msg.NewErrorResponse("database_error", txErr)
	}
	defer tx.Rollback()

	var newID int64
	if err := tx.QueryRow(`
		INSERT INTO conversations (type, created_at, updated_at)
		VALUES ('direct', NOW(), NOW())
		RETURNING id
	`).Scan(&newID); err != nil {
		return 0, msg.NewErrorResponse("database_error", err)
	}

	if _, err := tx.Exec(`
		INSERT INTO conversation_participants (conversation_id, user_id, status, is_archived)
		VALUES ($1, $2, 'accepted', false), ($1, $3, 'accepted', false)
	`, newID, userID, targetUserID); err != nil {
		return 0, msg.NewErrorResponse("database_error", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, msg.NewErrorResponse("database_error", err)
	}
	return newID, nil
}

// បន្ថែមក្នុង chats_repo.go
func (r *ChatsRepoImpl) SearchUsers(currentUserID int64, search string, limit int) ([]UserSearchItem, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var users []UserSearchItem

	query := `
		SELECT id, user_name, first_name, last_name, profile_images
		FROM tbl_users
		WHERE id != $1
		  AND deleted_at IS NULL
		  AND (
		      $2 = '' OR
		      user_name ILIKE '%' || $2 || '%' OR
		      first_name ILIKE '%' || $2 || '%' OR
		      last_name ILIKE '%' || $2 || '%'
		  )
		ORDER BY user_name
		LIMIT $3
	`
	if err := r.dbpool.Select(&users, query, currentUserID, search, limit); err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	return users, nil
}
