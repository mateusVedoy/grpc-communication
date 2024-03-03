package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type ClientGRPC struct{}

func (C ClientGRPC) Server() *grpc.Server {
	listener, listenerErr := net.Listen("tcp", ":9000")

	if listenerErr != nil {
		log.Fatalf("Failed to listen on port 9090. Reason: %v", listenerErr)
		return nil
	}

	grpcServer := grpc.NewServer()

	serverErr := grpcServer.Serve(listener)

	if serverErr != nil {
		log.Fatalf("Fail to serve gRPC server on port 9000. Reason: %v", serverErr)
		return nil
	}

	return grpcServer
}
