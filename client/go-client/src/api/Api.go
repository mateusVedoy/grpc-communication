package api

import "net/http"

type CreateMessageRequest struct {
	Message string `protobuf:"bytes,3,opt,name=content,proto3" json:"message"`
	Aside   any    `protobuf:"bytes,3,opt,name=aside,proto3" json:"aside,omitempty"`
}

type Message struct {
	Message string `json:"message"`
	Aside   any    `json:"aside"`
}

type Messages struct {
	Messages []Message `json:"messages"`
}

// Bind implements render.Binder.
func (*CreateMessageRequest) Bind(r *http.Request) error {
	return nil
}

type ApiResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type MarkAsReadRequest struct {
	Identifiers []string `json:"identifiers"`
}
