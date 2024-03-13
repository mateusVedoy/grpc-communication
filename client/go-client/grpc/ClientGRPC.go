package grpc

import (
	"log"

	"google.golang.org/grpc"
)

type ClientGRPC struct {
	conn *grpc.ClientConn
}

func (C ClientGRPC) Server() *grpc.ClientConn {

	conn, connErr := grpc.Dial(":9000", grpc.WithInsecure())

	if connErr != nil {
		log.Fatalf("Client gRPC connection error. Reason: %s", connErr)
		return nil
	}

	C.conn = conn

	return C.conn
}

func (C ClientGRPC) Close() {
	C.conn.Close()
}
