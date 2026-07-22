package platform

import (
	"context"
	"net/http"
)

// Client creates and executes authenticated API requests.
type Client interface {
	NewRequest(ctx context.Context, method, url string, body any) (*http.Request, error)
	Do(request *http.Request, response any) error
}
