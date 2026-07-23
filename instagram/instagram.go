// Package instagram provides an unofficial SDK for the Instagram APIs.
package instagram

import "errors"

const (
	DisplayName          = "Instagram"
	BaseURL              = "https://graph.instagram.com"
	APIBaseURL           = BaseURL + "/v24.0"
	AuthorizationBaseURL = "https://www.instagram.com"
	OAuthBaseURL         = "https://api.instagram.com"
)

var (
	ErrMissingClientID     = errors.New("Instagram client ID is required")
	ErrMissingClientSecret = errors.New("Instagram client secret is required")
	ErrMissingRedirectURL  = errors.New("Instagram redirect URL is required")
	ErrMissingScopes       = errors.New("Instagram scopes are required")
	ErrMissingAccessToken  = errors.New("Instagram access token is required")
	ErrNilEndpointRequest  = errors.New("Instagram endpoint request is nil")
	ErrMissingID           = errors.New("Instagram resource ID is required")
	ErrNilClient           = errors.New("Instagram endpoint client is nil")
	ErrMissingCode         = errors.New("Instagram authorization code is required")
)

// IDPath is shared by requests whose path contains only an ID.
type IDPath struct {
	ID string `json:"id"`
}
