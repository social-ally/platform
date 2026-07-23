// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package instagram

import (
	"context"
	"net/http"
	"net/url"
)

// Media provides access to media endpoints.
type media struct {
	client *instagramClient
}

// NewMedia creates a Media endpoint group using client.
func NewMedia(client *instagramClient) *media {
	return &media{client: client}
}

type (
	RequestListMediaPath struct {
		IgUserID string `json:"ig_user_id"`
	}

	RequestListMediaQuery struct {
		Fields any     `json:"fields"`
		Limit  int     `json:"limit"`
		After  *string `json:"after"`
	}

	RequestListMedia struct {
		Path  RequestListMediaPath  `json:"path"`
		Query RequestListMediaQuery `json:"query"`
	}

	ResponseListMediaSuccessDataItem struct {
		ID        string    `json:"id"`
		Caption   string    `json:"caption"`
		MediaType MediaType `json:"media_type"`
		MediaURL  string    `json:"media_url"`
		Permalink string    `json:"permalink"`
		Timestamp any       `json:"timestamp"`
	}

	ResponseListMediaSuccessPagingCursors struct {
		Before string `json:"before"`
		After  string `json:"after"`
	}

	ResponseListMediaSuccessPaging struct {
		Cursors ResponseListMediaSuccessPagingCursors `json:"cursors"`
		Next    *string                               `json:"next"`
	}

	ResponseListMediaSuccess struct {
		Data   []ResponseListMediaSuccessDataItem `json:"data"`
		Paging ResponseListMediaSuccessPaging     `json:"paging"`
	}

	ResponseListMedia struct {
		Success ResponseListMediaSuccess `json:"success"`
	}
)

// ListMedia calls GET https://graph.instagram.com/v24.0/{ig_user_id}/media.
func (s *media) ListMedia(ctx context.Context, request *RequestListMedia) (*ResponseListMedia, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.IgUserID == "" {
		return nil, ErrMissingID
	}
	query := url.Values{"fields": {stringValue(request.Query.Fields)}}
	if request.Query.Limit != 0 {
		query.Set("limit", stringValue(request.Query.Limit))
	}
	if request.Query.After != nil {
		query.Set("after", *request.Query.After)
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, APIBaseURL+"/"+url.PathEscape(request.Path.IgUserID)+"/media?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseListMedia)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
