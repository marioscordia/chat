package service

import (
	"context"

	"github.com/marioscordia/chat/internal/model"
)

// ChatService is an interface through which Handler layer communicates with business layer
type ChatService interface {
	CreateChat(ctx context.Context, chat *model.Chat, members []int64) (int64, error)
	DeleteMember(ctx context.Context, chatID, memberID int64) error
	DeleteChat(ctx context.Context, chatID int64) error
	CreateMessage(ctx context.Context, msg *model.Message) error
}
