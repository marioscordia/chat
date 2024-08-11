package converter

import (
	"github.com/marioscordia/chat/internal/model"
	"github.com/marioscordia/chat/pkg/chat_v1"
)

// ToChatFromCreateRequest is the method that converts GRPC Create request to Chat model
func ToChatFromCreateRequest(req *chat_v1.CreateRequest) *model.Chat {
	return &model.Chat{
		Name:      req.ChatName,
		CreatorID: req.CreatorId,
		Type:      req.ChatType,
	}
}
