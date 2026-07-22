package youtube

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/social-ally/platform"
)

var (
	ErrNilHTTPClient    = errors.New("YouTube HTTP client is nil")
	ErrNilRequest       = errors.New("YouTube request is nil")
	ErrUnexpectedStatus = errors.New("YouTube API returned an unexpected status")
)

// ClientOption configures a YouTube client.
type ClientOption func(*YouTubeClient) error

// YouTubeClient is the shared authenticated HTTP client for YouTube endpoint groups.
type YouTubeClient struct {
	httpClient   *http.Client
	clientID     string
	clientSecret string
	redirectURL  string
	scopes       []Scope
	accessToken  string
}

// APIError describes a non-success response from the YouTube API.
type APIError struct {
	StatusCode int
	Status     string
	Body       []byte
}

// Error implements error.
func (e *APIError) Error() string {
	if len(e.Body) == 0 {
		return fmt.Sprintf("%s: %s", ErrUnexpectedStatus, e.Status)
	}
	return fmt.Sprintf("%s: %s: %s", ErrUnexpectedStatus, e.Status, strings.TrimSpace(string(e.Body)))
}

// Unwrap makes APIError match ErrUnexpectedStatus with errors.Is.
func (e *APIError) Unwrap() error {
	return ErrUnexpectedStatus
}

// WithScopes configures the OAuth scopes requested by the client.
func WithScopes(scopes ...Scope) ClientOption {
	return func(client *YouTubeClient) error {
		client.scopes = append([]Scope(nil), scopes...)
		return nil
	}
}

// WithHTTPClient configures the HTTP client used to execute requests.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(client *YouTubeClient) error {
		if httpClient == nil {
			return ErrNilHTTPClient
		}
		client.httpClient = httpClient
		return nil
	}
}

// WithAccessToken configures the bearer token used by API endpoint requests.
func WithAccessToken(accessToken string) ClientOption {
	return func(client *YouTubeClient) error { client.accessToken = accessToken; return nil }
}

// Do implements [platform.Client].
func (c *YouTubeClient) Do(request *http.Request, response any) error {
	if request == nil {
		return ErrNilRequest
	}
	if c.httpClient == nil {
		return ErrNilHTTPClient
	}

	httpResponse, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode < http.StatusOK || httpResponse.StatusCode >= http.StatusMultipleChoices {
		body, readErr := io.ReadAll(httpResponse.Body)
		if readErr != nil {
			return readErr
		}
		return &APIError{StatusCode: httpResponse.StatusCode, Status: httpResponse.Status, Body: body}
	}
	if response == nil || httpResponse.StatusCode == http.StatusNoContent {
		return nil
	}
	if err := json.NewDecoder(httpResponse.Body).Decode(response); err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}

// NewRequest implements [platform.Client].
func (c *YouTubeClient) NewRequest(ctx context.Context, method string, rawURL string, body any) (*http.Request, error) {
	var reader io.Reader
	var contentType string

	switch value := body.(type) {
	case nil:
	case io.Reader:
		reader = value
	case []byte:
		reader = bytes.NewReader(value)
	case url.Values:
		reader = strings.NewReader(value.Encode())
		contentType = "application/x-www-form-urlencoded"
	default:
		encoded, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(encoded)
		contentType = "application/json"
	}

	request, err := http.NewRequestWithContext(ctx, method, rawURL, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")
	if c.accessToken != "" {
		request.Header.Set("Authorization", "Bearer "+c.accessToken)
	}
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	return request, nil
}

func (c *YouTubeClient) authenticatedRequest(ctx context.Context, method, rawURL string, body any) (*http.Request, error) {
	if c == nil {
		return nil, ErrNilClient
	}
	if c.accessToken == "" {
		return nil, ErrMissingAccessToken
	}
	return c.NewRequest(ctx, method, rawURL, body)
}

// NewYouTubeClient creates a YouTube OAuth client.
func NewYouTubeClient(clientID, clientSecret, redirectURL string, options ...ClientOption) (*YouTubeClient, error) {
	if clientID == "" {
		return nil, ErrMissingClientID
	}
	if clientSecret == "" {
		return nil, ErrMissingClientSecret
	}
	if redirectURL == "" {
		return nil, ErrMissingRedirectURL
	}

	client := &YouTubeClient{
		httpClient:   http.DefaultClient,
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
	}
	for _, option := range options {
		if option == nil {
			continue
		}
		if err := option(client); err != nil {
			return nil, err
		}
	}
	if len(client.scopes) == 0 {
		return nil, ErrMissingScopes
	}
	return client, nil
}

var _ platform.Client = (*YouTubeClient)(nil)
