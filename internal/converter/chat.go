package converter

import (
	"github.com/marioscordia/chat/internal/model"
	"github.com/marioscordia/chat/pkg/chat_v1"
)

// ToChatCreateFromCreateRequest is the method that converts GRPC Create request to Chat model
func ToChatCreateFromCreateRequest(req *chat_v1.CreateRequest) *model.ChatCreate {
	return &model.ChatCreate{
		Name:      req.GetChatName(),
		CreatorID: req.GetCreatorId(),
		Type:      req.GetChatType(),
		UserIDs:   req.GetUserIds(),
	}
}
