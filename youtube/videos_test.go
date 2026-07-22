package youtube

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) { return fn(request) }

func TestUploadVideoSendsMultipartMedia(t *testing.T) {
	client, err := NewYouTubeClient("client", "secret", "https://example.com/callback", WithScopes(ScopeYoutubeUpload), WithAccessToken("token"), WithHTTPClient(&http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if got := request.URL.Query().Get("uploadType"); got != "multipart" {
			t.Errorf("uploadType = %q, want multipart", got)
		}
		if !strings.HasPrefix(request.Header.Get("Content-Type"), "multipart/related;") {
			t.Errorf("content type = %q", request.Header.Get("Content-Type"))
		}
		body, err := io.ReadAll(request.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(body), "video bytes") {
			t.Error("media bytes were not included")
		}
		return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Body: io.NopCloser(strings.NewReader(`{"id":"video-id"}`)), Header: make(http.Header)}, nil
	})}))
	if err != nil {
		t.Fatal(err)
	}
	_, err = NewVideos(client).UploadVideo(context.Background(), &RequestUploadVideo{Query: RequestUploadVideoQuery{Part: "snippet,status"}, Body: RequestUploadVideoBody{Snippet: RequestUploadVideoBodySnippet{Title: "title"}}, Media: strings.NewReader("video bytes")})
	if err != nil {
		t.Fatal(err)
	}
}
