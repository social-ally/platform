// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package tiktok

import (
	"context"
	"net/http"
	"net/url"
)

// Videos provides access to videos endpoints.
type videos struct {
	client *TikTokClient
}

// NewVideos creates a Videos endpoint group using client.
func NewVideos(client *TikTokClient) *videos {
	return &videos{client: client}
}

type (
	RequestListUserVideosQuery struct {
		Fields any `json:"fields"`
	}

	RequestListUserVideosBody struct {
		MaxCount int  `json:"max_count"`
		Cursor   *int `json:"cursor"`
	}

	RequestListUserVideos struct {
		Query RequestListUserVideosQuery `json:"query"`
		Body  RequestListUserVideosBody  `json:"body"`
	}

	ResponseListUserVideosSuccessDataVideosItem struct {
		ID               string `json:"id"`
		CreateTime       int    `json:"create_time"`
		CoverImageURL    string `json:"cover_image_url"`
		ShareURL         string `json:"share_url"`
		VideoDescription string `json:"video_description"`
		Duration         int    `json:"duration"`
		LikeCount        int    `json:"like_count"`
		CommentCount     int    `json:"comment_count"`
		ShareCount       int    `json:"share_count"`
		ViewCount        int    `json:"view_count"`
	}

	ResponseListUserVideosSuccessData struct {
		Videos  []ResponseListUserVideosSuccessDataVideosItem `json:"videos"`
		Cursor  int                                           `json:"cursor"`
		HasMore bool                                          `json:"has_more"`
	}

	ResponseListUserVideosSuccessError struct {
		Code    any    `json:"code"`
		Message any    `json:"message"`
		LogID   string `json:"log_id"`
	}

	ResponseListUserVideosSuccess struct {
		Data  ResponseListUserVideosSuccessData  `json:"data"`
		Error ResponseListUserVideosSuccessError `json:"error"`
	}

	ResponseListUserVideos struct {
		Success ResponseListUserVideosSuccess `json:"success"`
	}

	RequestQueryVideosQuery struct {
		Fields any `json:"fields"`
	}

	RequestQueryVideosBodyFilters struct {
		VideoIds []string `json:"video_ids"`
	}

	RequestQueryVideosBody struct {
		Filters RequestQueryVideosBodyFilters `json:"filters"`
	}

	RequestQueryVideos struct {
		Query RequestQueryVideosQuery `json:"query"`
		Body  RequestQueryVideosBody  `json:"body"`
	}

	ResponseQueryVideosSuccessDataVideosItem struct {
		ID string `json:"id"`
	}

	ResponseQueryVideosSuccessData struct {
		Videos []ResponseQueryVideosSuccessDataVideosItem `json:"videos"`
	}

	ResponseQueryVideosSuccessError struct {
		Code    any    `json:"code"`
		Message any    `json:"message"`
		LogID   string `json:"log_id"`
	}

	ResponseQueryVideosSuccess struct {
		Data  ResponseQueryVideosSuccessData  `json:"data"`
		Error ResponseQueryVideosSuccessError `json:"error"`
	}

	ResponseQueryVideos struct {
		Success ResponseQueryVideosSuccess `json:"success"`
	}
)

// ListUserVideos calls POST https://open.tiktokapis.com/v2/video/list/.
func (s *videos) ListUserVideos(ctx context.Context, request *RequestListUserVideos) (*ResponseListUserVideos, error) {
	if request == nil {
		return nil, ErrNilOAuthRequest
	}
	q := url.Values{}
	addOptionalQuery(q, "fields", request.Query.Fields)
	raw := BaseURL + "/v2/video/list/"
	if q.Encode() != "" {
		raw += "?" + q.Encode()
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodPost, raw, request.Body)
	if e != nil {
		return nil, e
	}
	var data ResponseListUserVideosSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseListUserVideos{Success: data}, nil
}

// QueryVideos calls POST https://open.tiktokapis.com/v2/video/query/.
func (s *videos) QueryVideos(ctx context.Context, request *RequestQueryVideos) (*ResponseQueryVideos, error) {
	if request == nil {
		return nil, ErrNilOAuthRequest
	}
	q := url.Values{}
	addOptionalQuery(q, "fields", request.Query.Fields)
	raw := BaseURL + "/v2/video/query/"
	if q.Encode() != "" {
		raw += "?" + q.Encode()
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodPost, raw, request.Body)
	if e != nil {
		return nil, e
	}
	var data ResponseQueryVideosSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseQueryVideos{Success: data}, nil
}
