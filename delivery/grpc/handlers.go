package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marioscordia/chat"
	"github.com/marioscordia/chat/pkg/chat_v1"
)

// New is ...
func New(usecase chat.UseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

// Handler is ...
type Handler struct {
	chat_v1.UnimplementedChatV1Server
	usecase chat.UseCase
}

// Create is ...
func (h *Handler) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	if err := validateCreateChatReq(req); err != nil {
		return nil, err
	}

	id, err := h.usecase.CreateChat(ctx, req)
	if err != nil {
		return nil, err
	}

	return &chat_v1.CreateResponse{
		Id: id,
	}, nil
}

// DeleteChat is ...
func (h *Handler) DeleteChat(ctx context.Context, req *chat_v1.DeleteChatRequest) (*emptypb.Empty, error) {
	if err := h.usecase.DeleteChat(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteMember is ...
func (h *Handler) DeleteMember(ctx context.Context, req *chat_v1.DeleteMemberRequest) (*emptypb.Empty, error) {
	if err := h.usecase.DeleteMember(ctx, req.ChatId, req.MemberId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// SendMessage is ...
func (h *Handler) SendMessage(ctx context.Context, msg *chat_v1.Message) (*emptypb.Empty, error) {
	if err := validateCreateMsgReq(msg); err != nil {
		return nil, err
	}

	if err := h.usecase.CreateMessage(ctx, msg); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
