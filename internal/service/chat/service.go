package chat

import (
	"context"

	"github.com/marioscordia/chat/internal/model"
	repo "github.com/marioscordia/chat/internal/repository"
	"github.com/marioscordia/chat/internal/service"
)

// New is the function that returns ChatService object
func New(repo repo.ChatRepository) service.ChatService {
	return &useCase{
		repo: repo,
	}
}

type useCase struct {
	repo repo.ChatRepository
}

func (u *useCase) CreateChat(ctx context.Context, chat *model.Chat, members []int64) (int64, error) {
	return u.repo.CreateChat(ctx, chat, members)
}

func (u *useCase) DeleteMember(ctx context.Context, chatID, memberID int64) error {
	return u.repo.DeleteMember(ctx, chatID, memberID)
}

func (u *useCase) DeleteChat(ctx context.Context, chatID int64) error {
	return u.repo.DeleteChat(ctx, chatID)
}

func (u *useCase) CreateMessage(ctx context.Context, msg *model.Message) error {
	return u.repo.CreateMessage(ctx, msg)
}
