package repo

import (
	"context"

	"github.com/marioscordia/chat/internal/model"
)

// ChatRepository is an interface through which Service layer communicates with database
type ChatRepository interface {
	CreateChat(ctx context.Context, chat *model.Chat, members []int64) (int64, error)
	DeleteMember(ctx context.Context, chatID, memberID int64) error
	DeleteChat(ctx context.Context, chatID int64) error
	CreateMessage(ctx context.Context, msg *model.Message) error
}
