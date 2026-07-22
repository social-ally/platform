// Package tiktok provides an unofficial SDK for the TikTok APIs.
package tiktok

import "errors"

const (
	// DisplayName is the platform's public name.
	DisplayName = "TikTok"
	// BaseURL is the base URL used by the platform API.
	BaseURL = "https://open.tiktokapis.com"
)

var (
	ErrNotImplemented       = errors.New("TikTok SDK endpoint is not implemented")
	ErrMissingClientID      = errors.New("TikTok client ID is required")
	ErrMissingClientSecret  = errors.New("TikTok client secret is required")
	ErrMissingRedirectURL   = errors.New("TikTok redirect URL is required")
	ErrMissingScopes        = errors.New("TikTok scopes are required")
	ErrNilOAuthClient       = errors.New("TikTok OAuth client is nil")
	ErrNilOAuthRequest      = errors.New("TikTok OAuth request is nil")
	ErrMissingCode          = errors.New("TikTok authorization code is required")
	ErrMissingToken         = errors.New("TikTok token is required")
	ErrMissingCodeVerifier  = errors.New("TikTok PKCE code verifier is required")
	ErrMissingCodeChallenge = errors.New("TikTok PKCE code challenge is required")
	ErrMissingAccessToken   = errors.New("TikTok access token is required")
	ErrNilClient            = errors.New("TikTok endpoint client is nil")
)

// IDPath is shared by requests whose path contains only an ID.
type IDPath struct {
	ID string `json:"id"`
}
