package chat

import (
	"context"

	"github.com/marioscordia/chat/pkg/chat_v1"
)

// Repository is ...
type Repository interface {
	CreateChat(ctx context.Context, chat *chat_v1.CreateRequest) (int64, error)
	DeleteMember(ctx context.Context, chatID, memberID int64) error
	DeleteChat(ctx context.Context, chatID int64) error
	CreateMessage(ctx context.Context, msg *chat_v1.Message) error
}

// UseCase is ...
type UseCase interface {
	CreateChat(ctx context.Context, chat *chat_v1.CreateRequest) (int64, error)
	DeleteMember(ctx context.Context, chatID, memberID int64) error
	DeleteChat(ctx context.Context, chatID int64) error
	CreateMessage(ctx context.Context, msg *chat_v1.Message) error
}
