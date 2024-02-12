package core

import (
	"context"
	"net/http"
)

type Handler struct {
	Endpoint string
	Method   string
	Handle   func(ctx context.Context, req *http.Request) (*http.Response, error)
}
