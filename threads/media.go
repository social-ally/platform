// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package threads

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Media provides access to media endpoints.
type Media struct {
	client *ThreadsClient
}

// NewMedia creates a Media endpoint group using client.
func NewMedia(client *ThreadsClient) *Media {
	return &Media{client: client}
}

type (
	RequestListUserThreadsPath struct {
		ThreadsUserID string `json:"threads_user_id"`
	}

	RequestListUserThreadsQuery struct {
		Fields any     `json:"fields"`
		Limit  int     `json:"limit"`
		After  *string `json:"after"`
	}

	RequestListUserThreads struct {
		Path  RequestListUserThreadsPath  `json:"path"`
		Query RequestListUserThreadsQuery `json:"query"`
	}

	ResponseListUserThreadsSuccessDataItem struct {
		ID        string `json:"id"`
		MediaType string `json:"media_type"`
		Text      string `json:"text"`
		Permalink string `json:"permalink"`
	}

	ResponseListUserThreadsSuccessPaging struct {
	}

	ResponseListUserThreadsSuccess struct {
		Data   []ResponseListUserThreadsSuccessDataItem `json:"data"`
		Paging ResponseListUserThreadsSuccessPaging     `json:"paging"`
	}

	ResponseListUserThreads struct {
		Success ResponseListUserThreadsSuccess `json:"success"`
	}
)

// ListUserThreads calls GET https://graph.threads.net/v1.0/{threads_user_id}/threads.
func (s *Media) ListUserThreads(ctx context.Context, request *RequestListUserThreads) (*ResponseListUserThreads, error) {
	if request == nil || request.Path.ThreadsUserID == "" {
		return nil, ErrMissingID
	}
	q := url.Values{}
	addOptionalQuery(q, "fields", request.Query.Fields)
	if request.Query.Limit != 0 {
		q.Set("limit", fmt.Sprint(request.Query.Limit))
	}
	if request.Query.After != nil {
		q.Set("after", *request.Query.After)
	}
	raw := BaseURL + "/v1.0/" + url.PathEscape(request.Path.ThreadsUserID) + "/threads"
	if q.Encode() != "" {
		raw += "?" + q.Encode()
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodGet, raw, nil)
	if e != nil {
		return nil, e
	}
	var data ResponseListUserThreadsSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseListUserThreads{Success: data}, nil
}
