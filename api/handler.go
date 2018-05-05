package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server represents gRPC server.
type Server struct{}

// SayHello foo bar
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}

func (s *Server) Authorize(ctx context.Context, in *AuthorizeRequest) (*AuthorizeResponse, error) {
	log.Printf("Receive Authorize Request – resource: %s, action: %s", in.Resource, in.Action)
	return &AuthorizeResponse{Permitted: false}, nil
}
