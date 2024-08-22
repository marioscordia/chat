package api

import (
	"errors"
	"strings"

	"github.com/marioscordia/chat/internal/utils"
	"github.com/marioscordia/chat/pkg/chat_v1"
)

func validateCreateChatReq(req *chat_v1.CreateRequest) error {
	if req.GetChatName() == "" {
		return errors.New("please enter chat name")
	}

	if !utils.ValidChannelType(req.GetChatType()) {
		return errors.New("please provide correct chat type")
	}

	return nil
}

func validateCreateMsgReq(req *chat_v1.Message) error {
	if strings.TrimSpace(req.GetText()) == "" {
		return errors.New("cannot create empty message")
	}

	return nil
}
