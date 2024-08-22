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
		useCase: usecase,
	}
}

// Handler is an object, which have methods that receive GRPC requests
type Handler struct {
	chat_v1.UnimplementedChatV1Server
	useCase service.ChatService
}

// Create is the method that receives GRPC Create request
func (h *Handler) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	if err := validateCreateChatReq(req); err != nil {
		return nil, err
	}

	chat := converter.ToChatCreateFromCreateRequest(req)

	id, err := h.useCase.CreateChat(ctx, chat)
	if err != nil {
		return nil, err
	}

	return &chat_v1.CreateResponse{
		Id: id,
	}, nil
}

// DeleteChat is the method that receives GRPC Delete request
func (h *Handler) DeleteChat(ctx context.Context, req *chat_v1.DeleteChatRequest) (*emptypb.Empty, error) {
	if err := h.useCase.DeleteChat(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteMember is the method that receives GRPC Delete request
func (h *Handler) DeleteMember(ctx context.Context, req *chat_v1.DeleteMemberRequest) (*emptypb.Empty, error) {
	if err := h.useCase.DeleteMember(ctx, req.GetChatId(), req.GetMemberId()); err != nil {
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

	if err := h.useCase.CreateMessage(ctx, msg); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
