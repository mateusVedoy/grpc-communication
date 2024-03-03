package main

import (
	"log"
	"net"

	"github.com/golesson/go-grpc-server/src/message"
	"github.com/golesson/go-grpc-server/src/protobuf"
	google_gRPC "google.golang.org/grpc"
)

func main() {

	log.Print("Running gRPC server")

	listener, listenerErr := net.Listen("tcp", ":9000")

	if listenerErr != nil {
		log.Fatalf("Failed to listen on port 9090. Reason: %v", listenerErr)
	}

	grpcServer := google_gRPC.NewServer()

	server := &message.Server{}

	protobuf.RegisterMessageServiceServer(grpcServer, server)

	if servErr := grpcServer.Serve(listener); servErr != nil {
		log.Fatalf("Fail to serve gRPC server on port 9000. Reason: %v", servErr)
	}
}
