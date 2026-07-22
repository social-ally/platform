package facebook

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) { return fn(request) }

func TestTransferVideoChunkUsesPageIDAndFormBody(t *testing.T) {
	client, err := NewFacebookClient("client", "secret", "https://example.com/callback", WithScopes(ScopePagesManagePosts), WithAccessToken("token"), WithHTTPClient(&http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if got, want := request.URL.EscapedPath(), "/v24.0/page%2Fid/videos"; got != want {
			t.Errorf("path = %q, want %q", got, want)
		}
		if got, want := request.Header.Get("Content-Type"), "application/x-www-form-urlencoded"; got != want {
			t.Errorf("content type = %q, want %q", got, want)
		}
		return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Body: io.NopCloser(strings.NewReader(`{"start_offset":"1","end_offset":"2"}`)), Header: make(http.Header)}, nil
	})}))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := NewPublishing(client).TransferVideoChunk(context.Background(), &RequestTransferVideoChunk{Path: RequestStartVideoUploadPath{PageID: "page/id"}, Body: RequestTransferVideoChunkBody{UploadSessionID: "session", StartOffset: "0"}}); err != nil {
		t.Fatal(err)
	}
}
