package main

import (
	"context"
	"github.com/ZER020/go-grpc-beginner/server/chatpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func closeConnection(conn *grpc.ClientConn) {
	if err := conn.Close(); err != nil {
		log.Fatalf("could not close connection: %v", err)
	}
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer closeConnection(conn)

	c := chatpb.NewChatServiceClient(conn)
	response, err := c.SayHello(context.Background(), &chatpb.Message{Body: "Hello from client!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response from server: %v", response.Body)

}
