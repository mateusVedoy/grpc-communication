package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	RenderChi "github.com/go-chi/render"
	"github.com/golesson/go-grpc-client/grpc"
	"github.com/golesson/go-grpc-client/src/protobuf"
	RenderPkg "github.com/unrolled/render"
)

type Controller struct {
	render      *RenderPkg.Render
	client_grpc *grpc.ClientGRPC
}

func (C *Controller) CreateMessage(w http.ResponseWriter, r *http.Request) {

	request := &CreateMessageRequest{}

	if bindErr := RenderChi.Bind(r, request); bindErr != nil {
		response := &ApiResponse{
			StatusCode: 400,
			Message:    bindErr.Error(),
		}
		C.render.JSON(w, 400, response)
		return
	}

	byt, errJson := json.Marshal(request.Aside)

	if errJson != nil {
		C.render.JSON(w, 400, errJson)
		return
	}

	createMessage := &protobuf.CreateMessageRequest{
		Content: request.Message,
		Read:    false,
		Aside:   byt,
	}

	conn := C.client_grpc.Server()
	defer conn.Close()

	serverConn := protobuf.NewMessageServiceClient(conn)

	_, responseErr := serverConn.Create(context.Background(), createMessage)

	if responseErr != nil {
		response := &ApiResponse{
			StatusCode: 400,
			Message:    responseErr.Error(),
		}
		C.render.JSON(w, 400, response)
		return
	}

	C.render.JSON(w, 200, &ApiResponse{
		StatusCode: 200,
		Message:    "Message delivered successfully",
	})
}

func (C *Controller) ReadMessage(w http.ResponseWriter, r *http.Request) {
	conn := C.client_grpc.Server()
	defer conn.Close()

	serverConn := protobuf.NewMessageServiceClient(conn)

	unread := &protobuf.UnreadMessageRequest{
		Read: false,
	}

	response, responseErr := serverConn.Read(context.Background(), unread)

	if responseErr != nil {
		response := &ApiResponse{
			StatusCode: 400,
			Message:    responseErr.Error(),
		}
		C.render.JSON(w, 400, response)
		return
	}

	var messages Messages

	for _, msg := range response.Messages {

		var asd any

		err := json.Unmarshal(msg.Aside, &asd)

		if err != nil {
			fmt.Println(err.Error())
		}

		newMsg := Message{
			Message: msg.Content,
			Aside:   asd,
		}

		messages.Messages = append(messages.Messages, newMsg)
	}

	C.render.JSON(w, 200, messages)
}

func (C *Controller) MarkAsRead(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "identifier")

	conn := C.client_grpc.Server()
	defer conn.Close()

	serverConn := protobuf.NewMessageServiceClient(conn)

	messageId := &protobuf.MessageId{
		Identifier: id,
	}

	_, responseErr := serverConn.MarkAsRead(context.Background(), messageId)

	if responseErr != nil {
		response := &ApiResponse{
			StatusCode: 400,
			Message:    responseErr.Error(),
		}
		C.render.JSON(w, 400, response)
		return
	}

	C.render.JSON(w, 200, &ApiResponse{
		StatusCode: 200,
		Message:    "Message marked as read successfully",
	})
}

func NewController() *Controller {
	return &Controller{
		render:      RenderPkg.New(),
		client_grpc: &grpc.ClientGRPC{},
	}
}
