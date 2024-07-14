package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/marioscordia/chat/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50052

type server struct {
	chat_v1.UnimplementedChatV1Server
}

func (s *server) Create(_ context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	log.Printf("Creating new chat with these usernames: %v", req.Usernames)

	return &chat_v1.CreateResponse{
		Id: int64(gofakeit.Uint64()),
	}, nil
}

func (s *server) Delete(_ context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Deleting chat with id: %d", req.Id)
	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, msg *chat_v1.Message) (*emptypb.Empty, error) {
	log.Printf("From: %s Text: %s Time: %v", msg.Form, msg.Text, msg.Timestamp)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_v1.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
