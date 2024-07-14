package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/marioscordia/chat/pkg/chat_v1"
)

const grpcPort = 50052

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_v1.RegisterChatV1Server(s, &server{})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("server listening at %v", lis.Addr())

		if err = s.Serve(lis); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()

	<-signalChan
	log.Println("received shutdown signal")

	s.GracefulStop()

	log.Println("server shutdown complete")
}
