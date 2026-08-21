package chat

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/utls"
)

type ChatListItem struct {
	ConversationID  int64      `json:"conversation_id" db:"conversation_id"`
	IsGroup         bool       `json:"is_group" db:"is_group"`
	Name            string     `json:"name" db:"name"`
	Avatar          *string    `json:"avatar" db:"avatar"`
	Online          bool       `json:"online" db:"online"`
	LastMessage     *string    `json:"last_message" db:"last_message"`
	LastMessageType *string    `json:"last_message_type" db:"last_message_type"`
	LastMessageAt   *time.Time `json:"last_message_at" db:"last_message_at"`
	UnreadCount     int        `json:"unread_count" db:"unread_count"`
	OtherUserID     *int64     `json:"other_user_id" db:"other_user_id"`
}

type ShowConversationsRequest struct {
	Tab    string `query:"tab"`
	Search string `query:"search"`
}

func (r *ShowConversationsRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Query(r); err != nil {
		return err
	}
	if r.Tab == "" {
		r.Tab = "all"
	}
	return nil
}

type ConversationsResponse struct {
	Conversations []ChatListItem `json:"conversations"`
	Total         int            `json:"total"`
}

type MessageAttachment struct {
	ID        int64   `json:"id" db:"id"`
	MessageID int64   `json:"message_id" db:"message_id"`
	Type      string  `json:"type" db:"type"`
	URL       string  `json:"url" db:"url"`
	FileName  *string `json:"file_name" db:"file_name"`
	FileSize  *string `json:"file_size" db:"file_size"`
	Duration  *string `json:"duration" db:"duration"`
	SortOrder int     `json:"sort_order" db:"sort_order"`
}

type ReactionSummary struct {
	MessageID      int64 `json:"message_id" db:"message_id"`
	ReactionTypeID int64 `json:"reaction_type_id" db:"reaction_type_id"`
	Count          int   `json:"count" db:"count"`
}

type Message struct {
	ID              int64               `json:"id" db:"id"`
	ConversationID  int64               `json:"conversation_id" db:"conversation_id"`
	SenderID        int64               `json:"sender_id" db:"sender_id"`
	SenderName      string              `json:"sender_name" db:"sender_name"`
	SenderAvatar    *string             `json:"sender_avatar" db:"sender_avatar"`
	Content         *string             `json:"content" db:"content"`
	Type            string              `json:"type" db:"type"`
	ForwardedFromID *int64              `json:"forwarded_from_id" db:"forwarded_from_id"`
	CreatedAt       time.Time           `json:"created_at" db:"created_at"`
	Attachments     []MessageAttachment `json:"attachments" db:"-"`
	Reactions       []ReactionSummary   `json:"reactions" db:"-"`
	IsMine          bool                `json:"is_mine" db:"-"`
}

type ShowMessagesRequest struct {
	ConversationID int64 `uri:"conversation_id" validate:"required"`
	BeforeID       int64 `query:"before_id"`
	Limit          int   `query:"limit"`
}

func (r *ShowMessagesRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().URI(r); err != nil {
		return err
	}
	if err := c.Bind().Query(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	if r.Limit <= 0 || r.Limit > 200 {
		r.Limit = 200
	}
	return nil
}

type MessagesResponse struct {
	Messages []Message `json:"messages"`
	Total    int       `json:"total"`
}

type SendMessageRequest struct {
	ConversationID  int64  `json:"conversation_id" form:"conversation_id" validate:"required"`
	Content         string `json:"content" form:"content"`
	Type            string `json:"type" form:"type"`
	ForwardedFromID *int64 `json:"forwarded_from_id" form:"forwarded_from_id"`
}

func (r *SendMessageRequest) new(uctx *share.UserContext) *Message {
	t := r.Type
	if t == "" {
		t = "text"
	}
	var content *string
	if r.Content != "" {
		content = &r.Content
	}
	return &Message{
		ConversationID:  r.ConversationID,
		SenderID:        uctx.UserID,
		Content:         content,
		Type:            t,
		ForwardedFromID: r.ForwardedFromID,
	}
}

type ToggleReactionRequest struct {
	MessageID      int64 `uri:"message_id" validate:"required"`
	ReactionTypeID int64 `json:"reaction_type_id" form:"reaction_type_id" validate:"required"`
}

func (r *ToggleReactionRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().URI(r); err != nil {
		return err
	}
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type UserSearchItem struct {
	ID            int64   `json:"id" db:"id"`
	UserName      string  `json:"user_name" db:"user_name"`
	FirstName     *string `json:"first_name" db:"first_name"`
	LastName      *string `json:"last_name" db:"last_name"`
	ProfileImages *string `json:"profile_images" db:"profile_images"`
}

type SearchUsersRequest struct {
	Search string `query:"search"`
	Limit  int    `query:"limit"`
}

func (r *SearchUsersRequest) bind(c fiber.Ctx) error {
	if err := c.Bind().Query(r); err != nil {
		return err
	}
	if r.Limit <= 0 || r.Limit > 200 {
		r.Limit = 200
	}
	return nil
}
