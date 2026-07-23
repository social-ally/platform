// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package youtube

import (
	"context"
	"net/http"
	"net/url"
)

// Channels provides access to channels endpoints.
type channels struct {
	client *YouTubeClient
}

// NewChannels creates a Channels endpoint group using client.
func NewChannels(client *YouTubeClient) *channels {
	return &channels{client: client}
}

type (
	RequestGetMyChannelsQuery struct {
		Part []ChannelPart `json:"part"`
		Mine bool          `json:"mine"`
	}

	RequestGetMyChannels struct {
		Query RequestGetMyChannelsQuery `json:"query"`
	}

	ResponseGetMyChannelsSuccessItemsItemSnippetThumbnails struct {
		Default  Thumbnail `json:"default"`
		Medium   Thumbnail `json:"medium"`
		High     Thumbnail `json:"high"`
		Standard Thumbnail `json:"standard"`
		Maxres   Thumbnail `json:"maxres"`
	}

	// Thumbnail is a YouTube channel or video image at a specific resolution.
	Thumbnail struct {
		URL    string `json:"url"`
		Width  *int   `json:"width"`
		Height *int   `json:"height"`
	}

	ResponseGetMyChannelsSuccessItemsItemSnippet struct {
		Title       string                                                 `json:"title"`
		Description string                                                 `json:"description"`
		Thumbnails  ResponseGetMyChannelsSuccessItemsItemSnippetThumbnails `json:"thumbnails"`
	}

	ResponseGetMyChannelsSuccessItemsItemStatistics struct {
		ViewCount             string `json:"viewCount"`
		SubscriberCount       string `json:"subscriberCount"`
		HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
		VideoCount            string `json:"videoCount"`
	}

	ResponseGetMyChannelsSuccessItemsItem struct {
		ID         string                                          `json:"id"`
		Snippet    ResponseGetMyChannelsSuccessItemsItemSnippet    `json:"snippet"`
		Statistics ResponseGetMyChannelsSuccessItemsItemStatistics `json:"statistics"`
	}

	ResponseGetMyChannelsSuccess struct {
		Kind  string                                  `json:"kind"`
		Items []ResponseGetMyChannelsSuccessItemsItem `json:"items"`
	}

	ResponseGetMyChannels struct {
		Success ResponseGetMyChannelsSuccess `json:"success"`
	}
)

// GetMyChannels calls GET https://www.googleapis.com/youtube/v3/channels.
func (s *channels) GetMyChannels(ctx context.Context, request *RequestGetMyChannels) (*ResponseGetMyChannels, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := url.Values{"part": {stringValue(request.Query.Part)}, "mine": {stringValue(request.Query.Mine)}}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, BaseURL+"/channels?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseGetMyChannels)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
