package x

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
	ErrNilHTTPClient    = errors.New("X HTTP client is nil")
	ErrNilRequest       = errors.New("X request is nil")
	ErrUnexpectedStatus = errors.New("X API returned an unexpected status")
)

type (
	// Option configures an X client.
	Option func(*XClient) error

	// XClient is the shared authenticated HTTP client for X endpoint groups.
	XClient struct {
		httpClient   *http.Client
		clientID     string
		clientSecret string
		redirectURL  string
		scopes       []Scope
		accessToken  string
		confidential bool
	}

	// APIError describes a non-success response from the X API.
	APIError struct {
		StatusCode int
		Status     string
		Body       []byte
	}
)

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
func WithScopes(scopes ...Scope) Option {
	return func(client *XClient) error {
		client.scopes = append([]Scope(nil), scopes...)
		return nil
	}
}

// WithHTTPClient configures the HTTP client used to execute requests.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(client *XClient) error {
		if httpClient == nil {
			return ErrNilHTTPClient
		}
		client.httpClient = httpClient
		return nil
	}
}

// WithAccessToken configures the bearer token used by authenticated API calls.
func WithAccessToken(accessToken string) Option {
	return func(client *XClient) error {
		if accessToken == "" {
			return ErrMissingAccessToken
		}
		client.accessToken = accessToken
		return nil
	}
}

// WithConfidentialClient configures OAuth token requests to use HTTP Basic authentication.
func WithConfidentialClient() Option {
	return func(client *XClient) error {
		client.confidential = true
		return nil
	}
}

// Do implements [platform.Client].
func (x *XClient) Do(request *http.Request, response any) error {
	if request == nil {
		return ErrNilRequest
	}
	if x.httpClient == nil {
		return ErrNilHTTPClient
	}

	httpResponse, err := x.httpClient.Do(request)
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
func (x *XClient) NewRequest(ctx context.Context, method string, rawURL string, body any) (*http.Request, error) {
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
	if x.accessToken != "" {
		request.Header.Set("Authorization", "Bearer "+x.accessToken)
	}
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	return request, nil
}

func (x *XClient) authenticatedRequest(ctx context.Context, method string, rawURL string, body any) (*http.Request, error) {
	if x == nil {
		return nil, ErrNilClient
	}
	if x.accessToken == "" {
		return nil, ErrMissingAccessToken
	}
	return x.NewRequest(ctx, method, rawURL, body)
}

func addOptionalQuery(values url.Values, key string, value any) {
	switch typed := value.(type) {
	case nil:
	case string:
		if typed != "" {
			values.Set(key, typed)
		}
	case []string:
		if len(typed) != 0 {
			values.Set(key, strings.Join(typed, ","))
		}
	default:
		values.Set(key, fmt.Sprint(typed))
	}
}

// NewXClient creates an X OAuth client.
func NewXClient(clientID, clientSecret, redirectURL string, options ...Option) (*XClient, error) {
	if clientID == "" {
		return nil, ErrMissingClientID
	}
	if redirectURL == "" {
		return nil, ErrMissingRedirectURL
	}

	client := &XClient{
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
	if client.confidential && client.clientSecret == "" {
		return nil, ErrMissingClientSecret
	}
	return client, nil
}

var _ platform.Client = (*XClient)(nil)
