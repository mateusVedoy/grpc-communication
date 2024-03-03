package main

import (
	"context"
	"log"

	"github.com/golesson/go-grpc-client/grpc"
	"github.com/golesson/go-grpc-client/src/protobuf"
)

func main() {

	grpcCli := &grpc.ClientGRPC{}

	conn := grpcCli.Server()
	defer conn.Close()

	connectionToServer := protobuf.NewMessageServiceClient(conn)

	request := &protobuf.CreateMessageRequest{
		Content: "Nova mensagem de buffer",
		Read:    false,
	}

	response, responseErr := connectionToServer.Create(context.Background(), request)

	if responseErr != nil {
		log.Fatalf("Error reading message from client. Reason: %s", responseErr)
	}

	log.Printf("Response from server is %v", response)
}
