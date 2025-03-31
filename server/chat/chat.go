package chat

import (
	"context"
	"fmt"
	"github.com/ZER020/go-grpc-beginner/server/chatpb"
)

type Server struct {
	chatpb.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *chatpb.Message) (*chatpb.Message, error) {
	fmt.Println("hello world")
	return &chatpb.Message{Body: "Hello from the Server!"}, nil
}
