// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package facebook

import (
	"context"
	"net/http"
	"net/url"
)

// Publishing provides access to publishing endpoints.
type Publishing struct {
	client *FacebookClient
}

// NewPublishing creates a Publishing endpoint group using client.
func NewPublishing(client *FacebookClient) *Publishing {
	return &Publishing{client: client}
}

type (
	RequestCreateFeedPostPath struct {
		PageID string `json:"page_id"`
	}

	RequestCreateFeedPostBodyAttachedMediaItem struct {
		MediaFbid string `json:"media_fbid"`
	}

	RequestCreateFeedPostBody struct {
		Message              *string                                      `json:"message"`
		Link                 any                                          `json:"link"`
		Published            *bool                                        `json:"published"`
		ScheduledPublishTime any                                          `json:"scheduled_publish_time"`
		BackdatedTime        any                                          `json:"backdated_time"`
		AttachedMedia        []RequestCreateFeedPostBodyAttachedMediaItem `json:"attached_media"`
		AccessToken          any                                          `json:"access_token"`
	}

	RequestCreateFeedPost struct {
		Path RequestCreateFeedPostPath `json:"path"`
		Body RequestCreateFeedPostBody `json:"body"`
	}

	ResponseCreateFeedPostSuccess struct {
		ID any `json:"id"`
	}

	ResponseCreateFeedPost struct {
		Success ResponseCreateFeedPostSuccess `json:"success"`
	}

	RequestUploadPhotoPath struct {
		PageID string `json:"page_id"`
	}

	RequestUploadPhotoBody struct {
		URL                  any     `json:"url"`
		Source               any     `json:"source"`
		Caption              *string `json:"caption"`
		Published            *bool   `json:"published"`
		ScheduledPublishTime any     `json:"scheduled_publish_time"`
		AccessToken          any     `json:"access_token"`
	}

	RequestUploadPhoto struct {
		Path RequestUploadPhotoPath `json:"path"`
		Body RequestUploadPhotoBody `json:"body"`
	}

	ResponseUploadPhotoSuccess struct {
		ID     any     `json:"id"`
		PostID *string `json:"post_id"`
	}

	ResponseUploadPhoto struct {
		Success ResponseUploadPhotoSuccess `json:"success"`
	}

	RequestStartVideoUploadPath struct {
		PageID string `json:"page_id"`
	}

	RequestStartVideoUploadBody struct {
		UploadPhase any `json:"upload_phase"`
		FileSize    int `json:"file_size"`
		AccessToken any `json:"access_token"`
	}

	RequestStartVideoUpload struct {
		Path RequestStartVideoUploadPath `json:"path"`
		Body RequestStartVideoUploadBody `json:"body"`
	}

	ResponseStartVideoUploadSuccess struct {
		VideoID         string `json:"video_id"`
		UploadSessionID string `json:"upload_session_id"`
		StartOffset     string `json:"start_offset"`
		EndOffset       string `json:"end_offset"`
	}

	ResponseStartVideoUpload struct {
		Success ResponseStartVideoUploadSuccess `json:"success"`
	}

	RequestTransferVideoChunkBody struct {
		UploadPhase     any    `json:"upload_phase"`
		UploadSessionID string `json:"upload_session_id"`
		StartOffset     string `json:"start_offset"`
		VideoFileChunk  any    `json:"video_file_chunk"`
	}

	RequestTransferVideoChunk struct {
		Path RequestStartVideoUploadPath   `json:"path"`
		Body RequestTransferVideoChunkBody `json:"body"`
	}

	ResponseTransferVideoChunkSuccess struct {
		StartOffset string `json:"start_offset"`
		EndOffset   string `json:"end_offset"`
	}

	ResponseTransferVideoChunk struct {
		Success ResponseTransferVideoChunkSuccess `json:"success"`
	}

	RequestFinishVideoUploadBody struct {
		UploadPhase          any     `json:"upload_phase"`
		UploadSessionID      string  `json:"upload_session_id"`
		Title                *string `json:"title"`
		Description          *string `json:"description"`
		Published            *bool   `json:"published"`
		ScheduledPublishTime any     `json:"scheduled_publish_time"`
	}

	RequestFinishVideoUpload struct {
		Path RequestStartVideoUploadPath  `json:"path"`
		Body RequestFinishVideoUploadBody `json:"body"`
	}

	ResponseFinishVideoUploadSuccess struct {
		Success bool `json:"success"`
	}

	ResponseFinishVideoUpload struct {
		Success ResponseFinishVideoUploadSuccess `json:"success"`
	}

	RequestDeletePostPath struct {
		PostID string `json:"post_id"`
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

// CreateFeedPost calls POST https://graph.facebook.com/v24.0/{page_id}/feed.
func (s *Publishing) CreateFeedPost(ctx context.Context, request *RequestCreateFeedPost) (*ResponseCreateFeedPost, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.PageID == "" {
		return nil, ErrMissingID
	}
	body, err := formValues(request.Body)
	if err != nil {
		return nil, err
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, APIBaseURL+"/"+url.PathEscape(request.Path.PageID)+"/feed", body)
	if err != nil {
		return nil, err
	}
	response := new(ResponseCreateFeedPost)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// UploadPhoto calls POST https://graph.facebook.com/v24.0/{page_id}/photos.
func (s *Publishing) UploadPhoto(ctx context.Context, request *RequestUploadPhoto) (*ResponseUploadPhoto, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.PageID == "" {
		return nil, ErrMissingID
	}
	body, err := formValues(request.Body)
	if err != nil {
		return nil, err
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, APIBaseURL+"/"+url.PathEscape(request.Path.PageID)+"/photos", body)
	if err != nil {
		return nil, err
	}
	response := new(ResponseUploadPhoto)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// StartVideoUpload calls POST https://graph-video.facebook.com/v24.0/{page_id}/videos.
func (s *Publishing) StartVideoUpload(ctx context.Context, request *RequestStartVideoUpload) (*ResponseStartVideoUpload, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.PageID == "" {
		return nil, ErrMissingID
	}
	body, err := formValues(request.Body)
	if err != nil {
		return nil, err
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, VideoBaseURL+"/"+url.PathEscape(request.Path.PageID)+"/videos", body)
	if err != nil {
		return nil, err
	}
	response := new(ResponseStartVideoUpload)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// TransferVideoChunk calls POST https://graph-video.facebook.com/v24.0/{page_id}/videos.
func (s *Publishing) TransferVideoChunk(ctx context.Context, request *RequestTransferVideoChunk) (*ResponseTransferVideoChunk, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.PageID == "" {
		return nil, ErrMissingID
	}
	body, err := formValues(request.Body)
	if err != nil {
		return nil, err
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, VideoBaseURL+"/"+url.PathEscape(request.Path.PageID)+"/videos", body)
	if err != nil {
		return nil, err
	}
	response := new(ResponseTransferVideoChunk)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// FinishVideoUpload calls POST https://graph-video.facebook.com/v24.0/{page_id}/videos.
func (s *Publishing) FinishVideoUpload(ctx context.Context, request *RequestFinishVideoUpload) (*ResponseFinishVideoUpload, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.PageID == "" {
		return nil, ErrMissingID
	}
	body, err := formValues(request.Body)
	if err != nil {
		return nil, err
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, VideoBaseURL+"/"+url.PathEscape(request.Path.PageID)+"/videos", body)
	if err != nil {
		return nil, err
	}
	response := new(ResponseFinishVideoUpload)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// DeletePost calls DELETE https://graph.facebook.com/v24.0/{post_id}.
func (s *Publishing) DeletePost(ctx context.Context, request *RequestDeletePost) (*ResponseDeletePost, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.PostID == "" {
		return nil, ErrMissingID
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodDelete, APIBaseURL+"/"+url.PathEscape(request.Path.PostID), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseDeletePost)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
