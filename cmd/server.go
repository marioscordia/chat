package main

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marioscordia/chat/pkg/chat_v1"
)

type server struct {
	chat_v1.UnimplementedChatV1Server
}

func (s *server) Create(_ context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	log.Printf("Creating new chat '%s' with these user ids: %v", req.ChatName, req.UserIds)

	return &chat_v1.CreateResponse{
		Id: int64(gofakeit.Uint64()),
	}, nil
}

func (s *server) Delete(_ context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Deleting chat with id: %d", req.Id)

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, msg *chat_v1.Message) (*emptypb.Empty, error) {
	log.Printf("From: %s Text: %s Time: %v", msg.From, msg.Text, msg.Timestamp)

	return &emptypb.Empty{}, nil
}
