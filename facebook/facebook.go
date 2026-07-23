// Package facebook provides an unofficial SDK for the Facebook APIs.
package facebook

import "errors"

const (
	DisplayName          = "Facebook"
	BaseURL              = "https://graph.facebook.com"
	APIBaseURL           = BaseURL + "/v24.0"
	VideoBaseURL         = "https://graph-video.facebook.com/v24.0"
	AuthorizationBaseURL = "https://www.facebook.com/v24.0"
)

var (
	ErrMissingClientID     = errors.New("Facebook client ID is required")
	ErrMissingClientSecret = errors.New("Facebook client secret is required")
	ErrMissingRedirectURL  = errors.New("Facebook redirect URL is required")
	ErrMissingScopes       = errors.New("Facebook scopes are required")
	ErrMissingAccessToken  = errors.New("Facebook access token is required")
	ErrNilEndpointRequest  = errors.New("Facebook endpoint request is nil")
	ErrMissingID           = errors.New("Facebook resource ID is required")
	ErrNilClient           = errors.New("Facebook endpoint client is nil")
	ErrMissingCode         = errors.New("Facebook authorization code is required")
)

// IDPath is shared by requests whose path contains only an ID.
type IDPath struct {
	ID string `json:"id"`
}
