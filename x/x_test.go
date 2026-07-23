package x

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

func TestNewXClientValidation(t *testing.T) {
	_, err := NewXClient("", "secret", "https://example.com/callback", WithScopes(ScopeTweetRead))
	if !errors.Is(err, ErrMissingClientID) {
		t.Fatalf("NewXClient() error = %v, want ErrMissingClientID", err)
	}

	_, err = NewXClient("client", "secret", "https://example.com/callback")
	if !errors.Is(err, ErrMissingScopes) {
		t.Fatalf("NewXClient() error = %v, want ErrMissingScopes", err)
	}

	_, err = NewXClient("client", "", "https://example.com/callback", WithScopes(ScopeTweetRead), WithConfidentialClient())
	if !errors.Is(err, ErrMissingClientSecret) {
		t.Fatalf("NewXClient() error = %v, want ErrMissingClientSecret", err)
	}
}

func TestXClientRequestAndResponse(t *testing.T) {
	httpClient := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if got, want := r.Header.Get("Content-Type"), "application/x-www-form-urlencoded"; got != want {
			t.Errorf("Content-Type = %q, want %q", got, want)
		}
		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}
		if got, want := r.Form.Get("code"), "code-value"; got != want {
			t.Errorf("code = %q, want %q", got, want)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(`{"ok":true}`)),
			Request:    r,
		}, nil
	})}

	client, err := NewXClient(
		"client",
		"secret",
		"https://example.com/callback",
		WithScopes(ScopeTweetRead),
		WithHTTPClient(httpClient),
	)
	if err != nil {
		t.Fatal(err)
	}

	request, err := client.NewRequest(context.Background(), http.MethodPost, "https://api.example.test", url.Values{"code": {"code-value"}})
	if err != nil {
		t.Fatal(err)
	}
	var response struct {
		OK bool `json:"ok"`
	}
	if err := client.Do(request, &response); err != nil {
		t.Fatal(err)
	}
	if !response.OK {
		t.Fatal("response was not decoded")
	}
}

func TestWithAccessTokenReturnsAuthenticatedCopy(t *testing.T) {
	client, err := NewXClient("client", "", "https://example.com/callback", WithScopes(ScopeUsersRead))
	if err != nil {
		t.Fatal(err)
	}
	authenticated, err := client.WithAccessToken("access-token")
	if err != nil {
		t.Fatal(err)
	}
	request, err := authenticated.NewRequest(context.Background(), http.MethodGet, "https://api.example.test", nil)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := request.Header.Get("Authorization"), "Bearer access-token"; got != want {
		t.Errorf("Authorization = %q, want %q", got, want)
	}
	originalRequest, err := client.NewRequest(context.Background(), http.MethodGet, "https://api.example.test", nil)
	if err != nil {
		t.Fatal(err)
	}
	if got := originalRequest.Header.Get("Authorization"); got != "" {
		t.Errorf("original Authorization = %q, want empty", got)
	}
}

func TestOAuthAuthorizeUsesClientDefaults(t *testing.T) {
	client, err := NewXClient("client", "secret", "https://example.com/callback", WithScopes(ScopeTweetRead, ScopeUsersRead))
	if err != nil {
		t.Fatal(err)
	}

	response, err := NewOAuth(client).Authorize(context.Background(), &RequestAuthorize{Query: RequestAuthorizeQuery{
		State:         "state-value",
		CodeChallenge: "challenge-value",
	}})
	if err != nil {
		t.Fatal(err)
	}
	authorizeURL, err := url.Parse(response.URL)
	if err != nil {
		t.Fatal(err)
	}
	query := authorizeURL.Query()
	if got, want := query.Get("client_id"), "client"; got != want {
		t.Errorf("client_id = %q, want %q", got, want)
	}
	if got, want := query.Get("scope"), "tweet.read users.read"; got != want {
		t.Errorf("scope = %q, want %q", got, want)
	}
	if got, want := query.Get("code_challenge_method"), "S256"; got != want {
		t.Errorf("code_challenge_method = %q, want %q", got, want)
	}
}

func TestOAuthExchangeCode(t *testing.T) {
	httpClient := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if got, want := request.URL.String(), oauthTokenURL; got != want {
			t.Errorf("URL = %q, want %q", got, want)
		}
		if err := request.ParseForm(); err != nil {
			t.Fatal(err)
		}
		if got, want := request.Form.Get("code"), "authorization-code"; got != want {
			t.Errorf("code = %q, want %q", got, want)
		}
		return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Header: make(http.Header), Body: io.NopCloser(strings.NewReader(`{"token_type":"bearer","expires_in":7200,"access_token":"access-token","scope":"tweet.read","refresh_token":"refresh-token"}`)), Request: request}, nil
	})}
	client, err := NewXClient("client", "secret", "https://example.com/callback", WithScopes(ScopeTweetRead), WithHTTPClient(httpClient))
	if err != nil {
		t.Fatal(err)
	}

	response, err := NewOAuth(client).ExchangeCode(context.Background(), &RequestExchangeCode{Body: RequestExchangeCodeBody{Code: "authorization-code", CodeVerifier: "verifier"}})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := response.Success.AccessToken, "access-token"; got != want {
		t.Errorf("access token = %q, want %q", got, want)
	}
}

func TestOAuthExchangeCodeForConfidentialClient(t *testing.T) {
	httpClient := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		clientID, clientSecret, ok := request.BasicAuth()
		if !ok || clientID != "client" || clientSecret != "secret" {
			t.Errorf("BasicAuth() = (%q, %q, %t), want client credentials", clientID, clientSecret, ok)
		}
		if err := request.ParseForm(); err != nil {
			t.Fatal(err)
		}
		if got := request.Form.Get("client_id"); got != "" {
			t.Errorf("client_id = %q, want omitted for a confidential client", got)
		}
		return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Header: make(http.Header), Body: io.NopCloser(strings.NewReader(`{"token_type":"bearer","expires_in":7200,"access_token":"access-token","scope":"tweet.read"}`)), Request: request}, nil
	})}
	client, err := NewXClient("client", "secret", "https://example.com/callback", WithScopes(ScopeTweetRead), WithConfidentialClient(), WithHTTPClient(httpClient))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := NewOAuth(client).ExchangeCode(context.Background(), &RequestExchangeCode{Body: RequestExchangeCodeBody{Code: "authorization-code", CodeVerifier: "verifier"}}); err != nil {
		t.Fatal(err)
	}
}

func TestUsersGetUserByID(t *testing.T) {
	httpClient := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if got, want := request.Header.Get("Authorization"), "Bearer access-token"; got != want {
			t.Errorf("Authorization = %q, want %q", got, want)
		}
		if got, want := request.URL.Path, "/2/users/user-id"; got != want {
			t.Errorf("path = %q, want %q", got, want)
		}
		return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Header: make(http.Header), Body: io.NopCloser(strings.NewReader(`{"data":{"id":"user-id","name":"Ada","username":"ada"}}`)), Request: request}, nil
	})}
	client, err := NewXClient("client", "secret", "https://example.com/callback", WithScopes(ScopeUsersRead), WithAccessToken("access-token"), WithHTTPClient(httpClient))
	if err != nil {
		t.Fatal(err)
	}

	response, err := NewUsers(client).GetUserByID(context.Background(), &RequestGetUserByID{Path: IDPath{ID: "user-id"}})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := response.Success.Data.Username, "ada"; got != want {
		t.Errorf("username = %q, want %q", got, want)
	}
}

func TestPostsCreatePost(t *testing.T) {
	httpClient := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		var body map[string]any
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if got, want := body["text"], "hello"; got != want {
			t.Errorf("text = %v, want %q", got, want)
		}
		if _, exists := body["media"]; exists {
			t.Error("empty media object was sent")
		}
		return &http.Response{StatusCode: http.StatusCreated, Status: "201 Created", Header: make(http.Header), Body: io.NopCloser(strings.NewReader(`{"data":{"id":"post-id","text":"hello"}}`)), Request: request}, nil
	})}
	client, err := NewXClient("client", "secret", "https://example.com/callback", WithScopes(ScopeTweetWrite), WithAccessToken("access-token"), WithHTTPClient(httpClient))
	if err != nil {
		t.Fatal(err)
	}
	text := "hello"
	response, err := NewPosts(client).CreatePost(context.Background(), &RequestCreatePost{Body: RequestCreatePostBody{Text: &text}})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := response.Success.Data.ID, "post-id"; got != want {
		t.Errorf("post ID = %q, want %q", got, want)
	}
}
