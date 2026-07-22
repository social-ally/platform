// Package x provides an unofficial SDK for the X APIs.
package x

import (
	"errors"
)

const (
	DisplayName = "X"
	BaseURL     = "https://api.x.com/2"
)

var (
	ErrNotImplemented       = errors.New("X SDK endpoint is not implemented")
	ErrMissingClientID      = errors.New("X client ID is required")
	ErrMissingClientSecret  = errors.New("X client secret is required")
	ErrMissingRedirectURL   = errors.New("X redirect URL is required")
	ErrMissingScopes        = errors.New("X scopes are required")
	ErrNilOAuthClient       = errors.New("X OAuth client is nil")
	ErrNilOAuthRequest      = errors.New("X OAuth request is nil")
	ErrNilEndpointRequest   = errors.New("X endpoint request is nil")
	ErrMissingCode          = errors.New("X authorization code is required")
	ErrMissingCodeChallenge = errors.New("X PKCE code challenge is required")
	ErrMissingCodeVerifier  = errors.New("X PKCE code verifier is required")
	ErrMissingToken         = errors.New("X token is required")
	ErrMissingAccessToken   = errors.New("X access token is required")
	ErrMissingID            = errors.New("X resource ID is required")
	ErrMissingMedia         = errors.New("X media is required")
	ErrUnsupportedMedia     = errors.New("X media must be a string, byte slice, or reader")
	ErrMissingPostContent   = errors.New("X post content is required")
	ErrNilClient            = errors.New("X endpoint client is nil")
)

// IDPath is shared by requests whose path contains only an ID.
type IDPath struct {
	ID string `json:"id"`
}
