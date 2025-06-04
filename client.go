package main

import (
	"context"
	"log"

	"github.com/petermazzocco/grpc-broadcast-cli/broadcast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunClient(msg *string) {
	var conn *grpc.ClientConn
	creds := insecure.NewCredentials()
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := broadcast.NewBroadcastClient(conn)

	response, err := c.SendMessage(context.Background(), &broadcast.MessageRequest{Message: *msg})
	if err != nil {
		log.Fatalf("Error when calling SendMessage: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)
}
