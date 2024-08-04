package chat

import (
	"context"

	"github.com/marioscordia/chat/pkg/chat_v1"
	repo "github.com/marioscordia/chat/repository"
	"github.com/marioscordia/chat/service"
)

// New is ...
func New(repo repo.ChatRepository) service.ChatService {
	return &useCase{
		repo: repo,
	}
}

type useCase struct {
	repo repo.ChatRepository
}

func (u *useCase) CreateChat(ctx context.Context, chat *chat_v1.CreateRequest) (int64, error) {
	return u.repo.CreateChat(ctx, chat)
}

func (u *useCase) DeleteMember(ctx context.Context, chatID, memberID int64) error {
	return u.repo.DeleteMember(ctx, chatID, memberID)
}

func (u *useCase) DeleteChat(ctx context.Context, chatID int64) error {
	return u.repo.DeleteChat(ctx, chatID)
}

func (u *useCase) CreateMessage(ctx context.Context, msg *chat_v1.Message) error {
	return u.repo.CreateMessage(ctx, msg)
}
