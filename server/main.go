package main

import (
	"fmt"
	"github.com/ZER020/go-grpc-beginner/server/chat"
	"github.com/ZER020/go-grpc-beginner/server/chatpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Go gRPC Beginners Tutorial")
	list, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chatpb.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
