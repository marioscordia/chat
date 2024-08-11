package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marioscordia/chat/internal/converter"
	"github.com/marioscordia/chat/internal/service"
	"github.com/marioscordia/chat/pkg/chat_v1"
)

// New is a function that returns Handler object
func New(usecase service.ChatService) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

// Handler is an object, which have methods that receive GRPC requests
type Handler struct {
	chat_v1.UnimplementedChatV1Server
	usecase service.ChatService
}

// Create is the method that receives GRPC Create request
func (h *Handler) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	if err := validateCreateChatReq(req); err != nil {
		return nil, err
	}

	chat := converter.ToChatFromCreateRequest(req)

	id, err := h.usecase.CreateChat(ctx, chat, req.UserIds)
	if err != nil {
		return nil, err
	}

	return &chat_v1.CreateResponse{
		Id: id,
	}, nil
}

// DeleteChat is the method that receives GRPC Delete request
func (h *Handler) DeleteChat(ctx context.Context, req *chat_v1.DeleteChatRequest) (*emptypb.Empty, error) {
	if err := h.usecase.DeleteChat(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteMember is the method that receives GRPC Delete request
func (h *Handler) DeleteMember(ctx context.Context, req *chat_v1.DeleteMemberRequest) (*emptypb.Empty, error) {
	if err := h.usecase.DeleteMember(ctx, req.ChatId, req.MemberId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// CreateMessage is the method that receives GRPC create request
func (h *Handler) CreateMessage(ctx context.Context, req *chat_v1.Message) (*emptypb.Empty, error) {
	if err := validateCreateMsgReq(req); err != nil {
		return nil, err
	}

	msg := converter.ToMessageFromCreateRequest(req)

	if err := h.usecase.CreateMessage(ctx, msg); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
