// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package threads

import (
	"context"
	"net/http"
	"net/url"
)

// Publishing provides access to publishing endpoints.
type Publishing struct {
	client *ThreadsClient
}

// NewPublishing creates a Publishing endpoint group using client.
func NewPublishing(client *ThreadsClient) *Publishing {
	return &Publishing{client: client}
}

type (
	RequestCreatePostContainerPath struct {
		ThreadsUserID string `json:"threads_user_id"`
	}

	RequestCreatePostContainerBody struct {
		MediaType      any     `json:"media_type"`
		Text           *string `json:"text"`
		ImageURL       any     `json:"image_url"`
		VideoURL       any     `json:"video_url"`
		Children       []any   `json:"children"`
		IsCarouselItem *bool   `json:"is_carousel_item"`
		ReplyToID      *string `json:"reply_to_id"`
		QuotePostID    *string `json:"quote_post_id"`
		LinkAttachment *string `json:"link_attachment"`
		LocationID     *string `json:"location_id"`
		AccessToken    string  `json:"access_token"`
	}

	RequestCreatePostContainer struct {
		Path RequestCreatePostContainerPath `json:"path"`
		Body RequestCreatePostContainerBody `json:"body"`
	}

	ResponseCreatePostContainerSuccess struct {
		ID any `json:"id"`
	}

	ResponseCreatePostContainer struct {
		Success ResponseCreatePostContainerSuccess `json:"success"`
	}

	RequestGetContainerStatusPath struct {
		CreationID string `json:"creation_id"`
	}

	RequestGetContainerStatusQuery struct {
		Fields any `json:"fields"`
	}

	RequestGetContainerStatus struct {
		Path  RequestGetContainerStatusPath  `json:"path"`
		Query RequestGetContainerStatusQuery `json:"query"`
	}

	ResponseGetContainerStatusSuccess struct {
		ID           string  `json:"id"`
		Status       any     `json:"status"`
		ErrorMessage *string `json:"error_message"`
	}

	ResponseGetContainerStatus struct {
		Success ResponseGetContainerStatusSuccess `json:"success"`
	}

	RequestPublishPostPath struct {
		ThreadsUserID string `json:"threads_user_id"`
	}

	RequestPublishPostBody struct {
		CreationID  string `json:"creation_id"`
		AccessToken string `json:"access_token"`
	}

	RequestPublishPost struct {
		Path RequestPublishPostPath `json:"path"`
		Body RequestPublishPostBody `json:"body"`
	}

	ResponsePublishPostSuccess struct {
		ID any `json:"id"`
	}

	ResponsePublishPost struct {
		Success ResponsePublishPostSuccess `json:"success"`
	}

	RequestDeletePostPath struct {
		ThreadsMediaID string `json:"threads_media_id"`
	}

	RequestDeletePost struct {
		Path RequestDeletePostPath `json:"path"`
	}

	ResponseDeletePostSuccess struct {
		Success bool `json:"success"`
	}

	ResponseDeletePost struct {
		Success ResponseDeletePostSuccess `json:"success"`
	}
)

// CreatePostContainer calls POST https://graph.threads.net/v1.0/{threads_user_id}/threads.
func (s *Publishing) CreatePostContainer(ctx context.Context, request *RequestCreatePostContainer) (*ResponseCreatePostContainer, error) {
	if request == nil || request.Path.ThreadsUserID == "" {
		return nil, ErrMissingID
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodPost, BaseURL+"/v1.0/"+url.PathEscape(request.Path.ThreadsUserID)+"/threads", request.Body)
	if e != nil {
		return nil, e
	}
	var data ResponseCreatePostContainerSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseCreatePostContainer{Success: data}, nil
}

// GetContainerStatus calls GET https://graph.threads.net/v1.0/{creation_id}.
func (s *Publishing) GetContainerStatus(ctx context.Context, request *RequestGetContainerStatus) (*ResponseGetContainerStatus, error) {
	if request == nil || request.Path.CreationID == "" {
		return nil, ErrMissingID
	}
	q := url.Values{}
	addOptionalQuery(q, "fields", request.Query.Fields)
	raw := BaseURL + "/v1.0/" + url.PathEscape(request.Path.CreationID)
	if q.Encode() != "" {
		raw += "?" + q.Encode()
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodGet, raw, nil)
	if e != nil {
		return nil, e
	}
	var data ResponseGetContainerStatusSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseGetContainerStatus{Success: data}, nil
}

// PublishPost calls POST https://graph.threads.net/v1.0/{threads_user_id}/threads_publish.
func (s *Publishing) PublishPost(ctx context.Context, request *RequestPublishPost) (*ResponsePublishPost, error) {
	if request == nil || request.Path.ThreadsUserID == "" {
		return nil, ErrMissingID
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodPost, BaseURL+"/v1.0/"+url.PathEscape(request.Path.ThreadsUserID)+"/threads_publish", request.Body)
	if e != nil {
		return nil, e
	}
	var data ResponsePublishPostSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponsePublishPost{Success: data}, nil
}

// DeletePost calls DELETE https://graph.threads.net/v1.0/{threads_media_id}.
func (s *Publishing) DeletePost(ctx context.Context, request *RequestDeletePost) (*ResponseDeletePost, error) {
	if request == nil || request.Path.ThreadsMediaID == "" {
		return nil, ErrMissingID
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodDelete, BaseURL+"/v1.0/"+url.PathEscape(request.Path.ThreadsMediaID), nil)
	if e != nil {
		return nil, e
	}
	var data ResponseDeletePostSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseDeletePost{Success: data}, nil
}
