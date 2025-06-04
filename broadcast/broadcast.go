package broadcast

import (
	"context"
	"log"
)

type Server struct{}

func (s *Server) SendMessage(ctx context.Context, in *MessageRequest) (*MessageResponse, error) {
	log.Printf("Receive message body from client: %s", in.Message)
	return &MessageResponse{Message: "Hello From the Server!"}, nil
}
