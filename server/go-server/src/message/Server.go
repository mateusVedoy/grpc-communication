package message

import (
	"context"
	"errors"
	"log"
	"reflect"

	"github.com/golesson/go-grpc-server/src/protobuf"
	"github.com/google/uuid"
)

var messages []*protobuf.Message

type Server struct {
	protobuf.UnimplementedMessageServiceServer
}

func (S *Server) mustEmbedUnimplementedMessageServiceServer() {}

func (S *Server) Create(ctx context.Context, message *protobuf.CreateMessageRequest) (*protobuf.Response, error) {

	log.Printf("Received message from client: %s", message.Content)

	msg := &protobuf.Message{
		Identifier: uuid.NewString(),
		Content:    message.Content,
		Read:       message.Read,
	}

	messages = append(messages, msg)

	return &protobuf.Response{
		StatusCode: 201,
		Message:    "Ok",
	}, nil
}

func (S *Server) Read(ctx context.Context, isRead *protobuf.UnreadMessageRequest) (*protobuf.Messages, error) {

	log.Print("Received request to send unread messages")

	var unread []*protobuf.Message

	for _, msg := range messages {
		if msg.Read == isRead.Read {
			unread = append(unread, msg)
		}
	}

	if unread == nil {
		return nil, errors.New("there's no unread messages")
	}

	return &protobuf.Messages{
		Messages: unread,
	}, nil
}

func (S *Server) MarkAsRead(ctx context.Context, messageId *protobuf.MessageId) (*protobuf.Response, error) {

	log.Printf("Message %s marked as read.", messageId)

	var updatedMessages []*protobuf.Message

	for _, msg := range messages {
		if msg.Identifier != messageId.Identifier {
			updatedMessages = append(updatedMessages, msg)
		}
	}

	if reflect.DeepEqual(messages, updatedMessages) {
		return nil, errors.New("there's no message for given identifier")
	}

	messages = updatedMessages

	return &protobuf.Response{
		StatusCode: 200,
		Message:    "Ok",
	}, nil
}
