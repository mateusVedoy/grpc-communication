package api

import "net/http"

type CreateMessageRequest struct {
	Message string `json:"message"`
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
