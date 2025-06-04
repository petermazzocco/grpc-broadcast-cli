package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/petermazzocco/grpc-broadcast-cli/broadcast"
	"google.golang.org/grpc"
)

type server struct {
	broadcast.UnimplementedBroadcastServer
}

func (s *server) SendMessage(ctx context.Context, req *broadcast.MessageRequest) (*broadcast.MessageResponse, error) {
	log.Printf("Received message: %s", req.Message)
	return &broadcast.MessageResponse{Message: "Your message has been broadcast to the server!"}, nil
}

func RunServer() {
	fmt.Println("Broadcast gRPC started")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &server{}

	grpcServer := grpc.NewServer()

	broadcast.RegisterBroadcastServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %s", err)
	}
}
