// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package youtube

import (
	"context"
	"net/http"
	"net/url"
)

// Playlists provides access to playlists endpoints.
type playlists struct {
	client *youTubeClient
}

// NewPlaylists creates a Playlists endpoint group using client.
func NewPlaylists(client *youTubeClient) *playlists {
	return &playlists{client: client}
}

type (
	RequestCreatePlaylistQuery struct {
		Part any `json:"part"`
	}

	RequestCreatePlaylistBodySnippet struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	RequestCreatePlaylistBodyStatus struct {
		PrivacyStatus PrivacyStatus `json:"privacyStatus"`
	}

	RequestCreatePlaylistBody struct {
		Snippet RequestCreatePlaylistBodySnippet `json:"snippet"`
		Status  RequestCreatePlaylistBodyStatus  `json:"status"`
	}

	RequestCreatePlaylist struct {
		Query RequestCreatePlaylistQuery `json:"query"`
		Body  RequestCreatePlaylistBody  `json:"body"`
	}

	ResponseCreatePlaylistSuccessSnippet struct {
	}

	ResponseCreatePlaylistSuccessStatus struct {
	}

	ResponseCreatePlaylistSuccess struct {
		ID      string                               `json:"id"`
		Snippet ResponseCreatePlaylistSuccessSnippet `json:"snippet"`
		Status  ResponseCreatePlaylistSuccessStatus  `json:"status"`
	}

	ResponseCreatePlaylist struct {
		Success ResponseCreatePlaylistSuccess `json:"success"`
	}

	RequestAddVideoToPlaylistQuery struct {
		Part any `json:"part"`
	}

	RequestAddVideoToPlaylistBodySnippetResourceId struct {
		Kind    any    `json:"kind"`
		VideoId string `json:"videoId"`
	}

	RequestAddVideoToPlaylistBodySnippet struct {
		PlaylistId string                                         `json:"playlistId"`
		ResourceId RequestAddVideoToPlaylistBodySnippetResourceId `json:"resourceId"`
		Position   *int                                           `json:"position"`
	}

	RequestAddVideoToPlaylistBody struct {
		Snippet RequestAddVideoToPlaylistBodySnippet `json:"snippet"`
	}

	RequestAddVideoToPlaylist struct {
		Query RequestAddVideoToPlaylistQuery `json:"query"`
		Body  RequestAddVideoToPlaylistBody  `json:"body"`
	}

	ResponseAddVideoToPlaylistSuccessSnippet struct {
	}

	ResponseAddVideoToPlaylistSuccess struct {
		ID      string                                   `json:"id"`
		Snippet ResponseAddVideoToPlaylistSuccessSnippet `json:"snippet"`
	}

	ResponseAddVideoToPlaylist struct {
		Success ResponseAddVideoToPlaylistSuccess `json:"success"`
	}
)

// CreatePlaylist calls POST https://www.googleapis.com/youtube/v3/playlists.
func (s *playlists) CreatePlaylist(ctx context.Context, request *RequestCreatePlaylist) (*ResponseCreatePlaylist, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := url.Values{"part": {stringValue(request.Query.Part)}}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, BaseURL+"/playlists?"+query.Encode(), request.Body)
	if err != nil {
		return nil, err
	}
	response := new(ResponseCreatePlaylist)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// AddVideoToPlaylist calls POST https://www.googleapis.com/youtube/v3/playlistItems.
func (s *playlists) AddVideoToPlaylist(ctx context.Context, request *RequestAddVideoToPlaylist) (*ResponseAddVideoToPlaylist, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := url.Values{"part": {stringValue(request.Query.Part)}}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, BaseURL+"/playlistItems?"+query.Encode(), request.Body)
	if err != nil {
		return nil, err
	}
	response := new(ResponseAddVideoToPlaylist)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
