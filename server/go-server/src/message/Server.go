package message

import (
	"context"
	"log"

	"github.com/golesson/go-grpc-server/src/protobuf"
)

type Server struct {
	protobuf.UnimplementedMessageServiceServer
}

func (S *Server) mustEmbedUnimplementedMessageServiceServer() {}

func (S *Server) Create(ctx context.Context, message *protobuf.CreateMessageRequest) (*protobuf.Response, error) {

	log.Printf("Received message from client: %s", message.Content)
	return &protobuf.Response{
		StatusCode: 201,
		Message:    "Ok",
	}, nil
}

func (S *Server) Read(ctx context.Context, messageIds *protobuf.MessageIds) (*protobuf.Messages, error) {

	log.Print("Received request to send unread messages")

	message_one := &protobuf.Message{
		Identifier: "one",
		Content:    "Hy from message one",
	}

	message_two := &protobuf.Message{
		Identifier: "two",
		Content:    "Hy from message two",
	}

	var msgs []*protobuf.Message

	msgs = append(msgs, message_one, message_two)

	return &protobuf.Messages{
		Messages: msgs,
	}, nil
}

func (S *Server) MarkAsRead(ctx context.Context, messageId *protobuf.MessageId) (*protobuf.Response, error) {

	log.Printf("Message %s marked as read.", messageId)

	return &protobuf.Response{
		StatusCode: 200,
		Message:    "Ok",
	}, nil
}
