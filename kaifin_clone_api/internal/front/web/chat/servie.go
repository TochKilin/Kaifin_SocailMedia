package chat

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
	"kaifin_clone_api/pkg/share"
)

type ChatsService interface {
	ShowConversations(req ShowConversationsRequest) (*ConversationsResponse, *error_responses.ErrorResponse)
	ShowMessages(req ShowMessagesRequest) (*MessagesResponse, *error_responses.ErrorResponse)
	SendMessage(req *SendMessageRequest, attachments []MessageAttachment) (*Message, *error_responses.ErrorResponse)
	ToggleReaction(req ToggleReactionRequest) (bool, *error_responses.ErrorResponse)
	MarkAsRead(conversationID int64) *error_responses.ErrorResponse
	StartConversation(targetUserID int64) (int64, *error_responses.ErrorResponse)
	SearchUsers(req SearchUsersRequest) ([]UserSearchItem, *error_responses.ErrorResponse)
}

type ChatsServiceImpl struct {
	UserCtx *share.UserContext
	dbpool  *sqlx.DB
	Repo    *ChatsRepoImpl
}

func NewChatsServiceImpl(dbpool *sqlx.DB) *ChatsServiceImpl {
	return &ChatsServiceImpl{
		dbpool: dbpool,
		Repo:   NewChatsRepoImpl(dbpool),
	}
}

func (s *ChatsServiceImpl) ShowConversations(req ShowConversationsRequest) (*ConversationsResponse, *error_responses.ErrorResponse) {
	return s.Repo.ShowConversations(req, s.UserCtx.UserID)
}

func (s *ChatsServiceImpl) ShowMessages(req ShowMessagesRequest) (*MessagesResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	ok, e := s.Repo.IsParticipant(req.ConversationID, s.UserCtx.UserID)
	if e != nil {
		return nil, e
	}
	if !ok {
		return nil, msg.NewErrorResponse("forbidden", errors.New("user is not a participant of this conversation"))
	}
	return s.Repo.ShowMessages(req, s.UserCtx.UserID)
}

func (s *ChatsServiceImpl) SendMessage(req *SendMessageRequest, attachments []MessageAttachment) (*Message, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	log.Printf(">>> DEBUG: ConversationID=%d UserID=%d", req.ConversationID, s.UserCtx.UserID)
	ok, e := s.Repo.IsParticipant(req.ConversationID, s.UserCtx.UserID)
	if e != nil {
		return nil, e
	}
	if !ok {
		return nil, msg.NewErrorResponse("forbidden", errors.New("user is not a participant of this conversation"))
	}

	m := req.new(s.UserCtx)
	if e := s.Repo.SendMessage(m, attachments); e != nil {
		return nil, e
	}
	m.IsMine = true
	return m, nil
}

func (s *ChatsServiceImpl) ToggleReaction(req ToggleReactionRequest) (bool, *error_responses.ErrorResponse) {
	return s.Repo.ToggleReaction(req.MessageID, s.UserCtx.UserID, req.ReactionTypeID)
}

func (s *ChatsServiceImpl) MarkAsRead(conversationID int64) *error_responses.ErrorResponse {
	return s.Repo.MarkAsRead(conversationID, s.UserCtx.UserID)
}

func (s *ChatsServiceImpl) StartConversation(targetUserID int64) (int64, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	if targetUserID == s.UserCtx.UserID {
		return 0, msg.NewErrorResponse("invalid_request", errors.New("cannot start a conversation with yourself"))
	}
	return s.Repo.FindOrCreateDirectConversation(s.UserCtx.UserID, targetUserID)
}

func (s *ChatsServiceImpl) SearchUsers(req SearchUsersRequest) ([]UserSearchItem, *error_responses.ErrorResponse) {
	return s.Repo.SearchUsers(s.UserCtx.UserID, req.Search, req.Limit)
}
