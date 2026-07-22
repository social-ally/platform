// Package threads provides an unofficial SDK for the Threads APIs.
package threads

import "errors"

const (
	DisplayName = "Threads"
	BaseURL     = "https://graph.threads.net"
)

var (
	ErrNotImplemented      = errors.New("Threads SDK endpoint is not implemented")
	ErrMissingClientID     = errors.New("Threads client ID is required")
	ErrMissingClientSecret = errors.New("Threads client secret is required")
	ErrMissingRedirectURL  = errors.New("Threads redirect URL is required")
	ErrMissingScopes       = errors.New("Threads scopes are required")
	ErrNilOAuthClient      = errors.New("Threads OAuth client is nil")
	ErrNilOAuthRequest     = errors.New("Threads OAuth request is nil")
	ErrMissingCode         = errors.New("Threads authorization code is required")
	ErrMissingToken        = errors.New("Threads token is required")
	ErrMissingAccessToken  = errors.New("Threads access token is required")
	ErrMissingID           = errors.New("Threads resource ID is required")
	ErrNilClient           = errors.New("Threads endpoint client is nil")
)

// IDPath is shared by requests whose path contains only an ID.
type IDPath struct {
	ID string `json:"id"`
}
