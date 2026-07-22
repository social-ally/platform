// Package youtube provides an unofficial SDK for the YouTube APIs.
package youtube

import "errors"

const (
	DisplayName          = "YouTube"
	BaseURL              = "https://www.googleapis.com/youtube/v3"
	UploadBaseURL        = "https://www.googleapis.com/upload/youtube/v3"
	AnalyticsBaseURL     = "https://youtubeanalytics.googleapis.com/v2"
	AuthorizationBaseURL = "https://accounts.google.com/o/oauth2/v2"
	OAuthBaseURL         = "https://oauth2.googleapis.com"
)

var (
	ErrNotImplemented      = errors.New("YouTube SDK endpoint is not implemented")
	ErrMissingClientID     = errors.New("YouTube client ID is required")
	ErrMissingClientSecret = errors.New("YouTube client secret is required")
	ErrMissingRedirectURL  = errors.New("YouTube redirect URL is required")
	ErrMissingScopes       = errors.New("YouTube scopes are required")
	ErrMissingAccessToken  = errors.New("YouTube access token is required")
	ErrNilEndpointRequest  = errors.New("YouTube endpoint request is nil")
	ErrMissingID           = errors.New("YouTube resource ID is required")
	ErrNilClient           = errors.New("YouTube endpoint client is nil")
	ErrMissingCode         = errors.New("YouTube authorization code is required")
)

// IDPath is shared by requests whose path contains only an ID.
type IDPath struct {
	ID string `json:"id"`
}
