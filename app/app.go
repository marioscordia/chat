package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/marioscordia/chat"
	chatGrpc "github.com/marioscordia/chat/delivery/grpc"
	"github.com/marioscordia/chat/facility"
	"github.com/marioscordia/chat/pkg/chat_v1"
	"github.com/marioscordia/chat/repository/postgres"
)

// Run is ...
func Run(ctx context.Context, postgresDb *sqlx.DB, server *grpc.Server, config *facility.Config) error {
	repo, err := postgres.New(ctx, postgresDb)
	if err != nil {
		return err
	}

	useCase := chat.New(repo)

	handler := chatGrpc.New(useCase)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}
	defer func() {
		if err = lis.Close(); err != nil {
			log.Panicf("error closing the listener: %v", err)
		}
	}()

	reflection.Register(server)
	chat_v1.RegisterChatV1Server(server, handler)

	log.Printf("server listening at %v", lis.Addr())

	if err = server.Serve(lis); err != nil {
		log.Panicf("failed to serve: %v", err)
	}

	return nil
}
