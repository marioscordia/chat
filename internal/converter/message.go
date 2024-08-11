package converter

import (
	"github.com/marioscordia/chat/internal/model"
	"github.com/marioscordia/chat/pkg/chat_v1"
)

// ToMessageFromCreateRequest is the method that converts GRPC Create request to Message model
func ToMessageFromCreateRequest(req *chat_v1.Message) *model.Message {
	return &model.Message{
		ChatID: req.ChatId,
		UserID: req.AuthorId,
		Text:   req.Text,
	}
}
