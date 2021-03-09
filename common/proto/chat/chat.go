package chat

import (
	"context"
	"log"
)

type Server struct {}

func (s *Server) SeyHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from: %s", message.Body)

	return &Message{Body: "Hello from the Server!"}, nil
}
