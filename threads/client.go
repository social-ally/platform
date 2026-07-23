package threads

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
	ErrNilHTTPClient    = errors.New("Threads HTTP client is nil")
	ErrNilRequest       = errors.New("Threads request is nil")
	ErrUnexpectedStatus = errors.New("Threads API returned an unexpected status")
)

// option configures a Threads client.
type option func(*threadsClient) error

// threadsClient is the shared authenticated HTTP client for Threads endpoint groups.
type threadsClient struct {
	httpClient   *http.Client
	clientID     string
	clientSecret string
	redirectURL  string
	scopes       []Scope
	accessToken  string
}

// APIError describes a non-success response from the Threads API.
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
func WithScopes(scopes ...Scope) option {
	return func(client *threadsClient) error {
		client.scopes = append([]Scope(nil), scopes...)
		return nil
	}
}

// WithHTTPClient configures the HTTP client used to execute requests.
func WithHTTPClient(httpClient *http.Client) option {
	return func(client *threadsClient) error {
		if httpClient == nil {
			return ErrNilHTTPClient
		}
		client.httpClient = httpClient
		return nil
	}
}

// WithAccessToken configures the bearer token for authenticated API requests.
func WithAccessToken(accessToken string) option {
	return func(client *threadsClient) error { client.accessToken = accessToken; return nil }
}

// WithAccessToken returns a copy of c configured for authenticated API requests.
func (c *threadsClient) WithAccessToken(accessToken string) (*threadsClient, error) {
	if c == nil {
		return nil, ErrNilClient
	}
	if accessToken == "" {
		return nil, ErrMissingAccessToken
	}
	copy := *c
	copy.scopes = append([]Scope(nil), c.scopes...)
	copy.accessToken = accessToken
	return &copy, nil
}

// Do implements [platform.Client].
func (c *threadsClient) Do(request *http.Request, response any) error {
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
func (c *threadsClient) NewRequest(ctx context.Context, method string, rawURL string, body any) (*http.Request, error) {
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

func (c *threadsClient) authenticatedRequest(ctx context.Context, method, rawURL string, body any) (*http.Request, error) {
	if c == nil {
		return nil, ErrNilClient
	}
	if c.accessToken == "" {
		return nil, ErrMissingAccessToken
	}
	return c.NewRequest(ctx, method, rawURL, body)
}

func addOptionalQuery(values url.Values, key string, value any) {
	if value == nil {
		return
	}
	switch typed := value.(type) {
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

// NewThreadsClient creates a Threads OAuth client.
func NewThreadsClient(clientID, clientSecret, redirectURL string, options ...option) (*threadsClient, error) {
	if clientID == "" {
		return nil, ErrMissingClientID
	}
	if clientSecret == "" {
		return nil, ErrMissingClientSecret
	}
	if redirectURL == "" {
		return nil, ErrMissingRedirectURL
	}

	client := &threadsClient{
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

var _ platform.Client = (*threadsClient)(nil)
