// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package tiktok

import (
	"context"
	"fmt"
	"net/http"
)

// Publishing provides access to publishing endpoints.
type publishing struct {
	client *TikTokClient
}

// NewPublishing creates a Publishing endpoint group using client.
func NewPublishing(client *TikTokClient) *publishing {
	return &publishing{client: client}
}

type (
	RequestQueryCreatorInfoBody struct {
	}

	RequestQueryCreatorInfo struct {
		Body RequestQueryCreatorInfoBody `json:"body"`
	}

	ResponseQueryCreatorInfoSuccessData struct {
		CreatorAvatarURL        string         `json:"creator_avatar_url"`
		CreatorUsername         string         `json:"creator_username"`
		CreatorNickname         string         `json:"creator_nickname"`
		PrivacyLevelOptions     []PrivacyLevel `json:"privacy_level_options"`
		CommentDisabled         bool           `json:"comment_disabled"`
		DuetDisabled            bool           `json:"duet_disabled"`
		StitchDisabled          bool           `json:"stitch_disabled"`
		MaxVideoPostDurationSec int            `json:"max_video_post_duration_sec"`
	}

	ResponseQueryCreatorInfoSuccessError struct {
		Code    any    `json:"code"`
		Message any    `json:"message"`
		LogID   string `json:"log_id"`
	}

	ResponseQueryCreatorInfoSuccess struct {
		Data  ResponseQueryCreatorInfoSuccessData  `json:"data"`
		Error ResponseQueryCreatorInfoSuccessError `json:"error"`
	}

	ResponseQueryCreatorInfo struct {
		Success ResponseQueryCreatorInfoSuccess `json:"success"`
	}

	RequestDirectPostVideoInitBodyPostInfo struct {
		Title                 string       `json:"title"`
		PrivacyLevel          PrivacyLevel `json:"privacy_level"`
		DisableDuet           bool         `json:"disable_duet"`
		DisableComment        bool         `json:"disable_comment"`
		DisableStitch         bool         `json:"disable_stitch"`
		VideoCoverTimestampMs *int         `json:"video_cover_timestamp_ms"`
		BrandContentToggle    *bool        `json:"brand_content_toggle"`
		BrandOrganicToggle    *bool        `json:"brand_organic_toggle"`
	}

	RequestDirectPostVideoInitBodySourceInfo struct {
		Source          Source `json:"source"`
		VideoSize       *int   `json:"video_size"`
		ChunkSize       *int   `json:"chunk_size"`
		TotalChunkCount *int   `json:"total_chunk_count"`
		VideoURL        string `json:"video_url"`
	}

	RequestDirectPostVideoInitBody struct {
		PostInfo   RequestDirectPostVideoInitBodyPostInfo   `json:"post_info"`
		SourceInfo RequestDirectPostVideoInitBodySourceInfo `json:"source_info"`
	}

	RequestDirectPostVideoInit struct {
		Body RequestDirectPostVideoInitBody `json:"body"`
	}

	ResponseDirectPostVideoInitSuccessData struct {
		PublishID string  `json:"publish_id"`
		UploadURL *string `json:"upload_url"`
	}

	ResponseDirectPostVideoInitSuccessError struct {
		Code    any    `json:"code"`
		Message any    `json:"message"`
		LogID   string `json:"log_id"`
	}

	ResponseDirectPostVideoInitSuccess struct {
		Data  ResponseDirectPostVideoInitSuccessData  `json:"data"`
		Error ResponseDirectPostVideoInitSuccessError `json:"error"`
	}

	ResponseDirectPostVideoInit struct {
		Success ResponseDirectPostVideoInitSuccess `json:"success"`
	}

	RequestUploadVideoChunkHeaders struct {
		ContentType   VideoContentType `json:"Content-Type"`
		ContentLength int              `json:"Content-Length"`
		ContentRange  string           `json:"Content-Range"`
	}

	RequestUploadVideoChunk struct {
		UploadURL string                         `json:"-"`
		Headers   RequestUploadVideoChunkHeaders `json:"headers"`
		Body      []byte                         `json:"body"`
	}

	ResponseUploadVideoChunkSuccess struct {
		HttpStatus int `json:"http_status"`
	}

	ResponseUploadVideoChunk struct {
		Success ResponseUploadVideoChunkSuccess `json:"success"`
	}

	RequestUploadVideoDraftInitBodySourceInfo struct {
		Source          Source `json:"source"`
		VideoSize       *int   `json:"video_size"`
		ChunkSize       *int   `json:"chunk_size"`
		TotalChunkCount *int   `json:"total_chunk_count"`
		VideoURL        string `json:"video_url"`
	}

	RequestUploadVideoDraftInitBody struct {
		SourceInfo RequestUploadVideoDraftInitBodySourceInfo `json:"source_info"`
	}

	RequestUploadVideoDraftInit struct {
		Body RequestUploadVideoDraftInitBody `json:"body"`
	}

	ResponseUploadVideoDraftInitSuccessData struct {
		PublishID string  `json:"publish_id"`
		UploadURL *string `json:"upload_url"`
	}

	ResponseUploadVideoDraftInitSuccessError struct {
		Code    any    `json:"code"`
		Message any    `json:"message"`
		LogID   string `json:"log_id"`
	}

	ResponseUploadVideoDraftInitSuccess struct {
		Data  ResponseUploadVideoDraftInitSuccessData  `json:"data"`
		Error ResponseUploadVideoDraftInitSuccessError `json:"error"`
	}

	ResponseUploadVideoDraftInit struct {
		Success ResponseUploadVideoDraftInitSuccess `json:"success"`
	}

	RequestDirectPostPhotoInitBodyPostInfo struct {
		Title          string       `json:"title"`
		Description    string       `json:"description"`
		DisableComment bool         `json:"disable_comment"`
		PrivacyLevel   PrivacyLevel `json:"privacy_level"`
		AutoAddMusic   bool         `json:"auto_add_music"`
	}

	RequestDirectPostPhotoInitBodySourceInfo struct {
		Source          Source   `json:"source"`
		PhotoCoverIndex int      `json:"photo_cover_index"`
		PhotoImages     []string `json:"photo_images"`
	}

	RequestDirectPostPhotoInitBody struct {
		PostInfo   RequestDirectPostPhotoInitBodyPostInfo   `json:"post_info"`
		SourceInfo RequestDirectPostPhotoInitBodySourceInfo `json:"source_info"`
		PostMode   any                                      `json:"post_mode"`
		MediaType  any                                      `json:"media_type"`
	}

	RequestDirectPostPhotoInit struct {
		Body RequestDirectPostPhotoInitBody `json:"body"`
	}

	ResponseDirectPostPhotoInitSuccessData struct {
		PublishID string `json:"publish_id"`
	}

	ResponseDirectPostPhotoInitSuccessError struct {
		Code    any    `json:"code"`
		Message any    `json:"message"`
		LogID   string `json:"log_id"`
	}

	ResponseDirectPostPhotoInitSuccess struct {
		Data  ResponseDirectPostPhotoInitSuccessData  `json:"data"`
		Error ResponseDirectPostPhotoInitSuccessError `json:"error"`
	}

	ResponseDirectPostPhotoInit struct {
		Success ResponseDirectPostPhotoInitSuccess `json:"success"`
	}

	RequestGetPublishStatusBody struct {
		PublishID string `json:"publish_id"`
	}

	RequestGetPublishStatus struct {
		Body RequestGetPublishStatusBody `json:"body"`
	}

	ResponseGetPublishStatusSuccessData struct {
		Status                   PublishStatus `json:"status"`
		FailReason               *string       `json:"fail_reason"`
		PublicalyAvailablePostID []string      `json:"publicaly_available_post_id"`
		UploadedBytes            *int          `json:"uploaded_bytes"`
	}

	ResponseGetPublishStatusSuccessError struct {
		Code    any    `json:"code"`
		Message any    `json:"message"`
		LogID   string `json:"log_id"`
	}

	ResponseGetPublishStatusSuccess struct {
		Data  ResponseGetPublishStatusSuccessData  `json:"data"`
		Error ResponseGetPublishStatusSuccessError `json:"error"`
	}

	ResponseGetPublishStatus struct {
		Success ResponseGetPublishStatusSuccess `json:"success"`
	}
)

// QueryCreatorInfo calls POST https://open.tiktokapis.com/v2/post/publish/creator_info/query/.
func (s *publishing) QueryCreatorInfo(ctx context.Context, request *RequestQueryCreatorInfo) (*ResponseQueryCreatorInfo, error) {
	if request == nil {
		return nil, ErrNilOAuthRequest
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodPost, BaseURL+"/v2/post/publish/creator_info/query/", request.Body)
	if e != nil {
		return nil, e
	}
	var data ResponseQueryCreatorInfoSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseQueryCreatorInfo{Success: data}, nil
}

// DirectPostVideoInit calls POST https://open.tiktokapis.com/v2/post/publish/video/init/.
func (s *publishing) DirectPostVideoInit(ctx context.Context, request *RequestDirectPostVideoInit) (*ResponseDirectPostVideoInit, error) {
	if request == nil {
		return nil, ErrNilOAuthRequest
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodPost, BaseURL+"/v2/post/publish/video/init/", request.Body)
	if e != nil {
		return nil, e
	}
	var data ResponseDirectPostVideoInitSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseDirectPostVideoInit{Success: data}, nil
}

// UploadVideoChunk calls PUT {upload_url_returned_by_init}.
func (s *publishing) UploadVideoChunk(ctx context.Context, request *RequestUploadVideoChunk) (*ResponseUploadVideoChunk, error) {
	if request == nil || request.UploadURL == "" {
		return nil, ErrMissingToken
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodPut, request.UploadURL, request.Body)
	if e != nil {
		return nil, e
	}
	if request.Headers.ContentType != "" {
		r.Header.Set("Content-Type", fmt.Sprint(request.Headers.ContentType))
	}
	if request.Headers.ContentLength != 0 {
		r.ContentLength = int64(request.Headers.ContentLength)
	}
	if request.Headers.ContentRange != "" {
		r.Header.Set("Content-Range", fmt.Sprint(request.Headers.ContentRange))
	}
	if e = s.client.Do(r, nil); e != nil {
		return nil, e
	}
	return &ResponseUploadVideoChunk{Success: ResponseUploadVideoChunkSuccess{HttpStatus: http.StatusOK}}, nil
}

// UploadVideoDraftInit calls POST https://open.tiktokapis.com/v2/post/publish/inbox/video/init/.
func (s *publishing) UploadVideoDraftInit(ctx context.Context, request *RequestUploadVideoDraftInit) (*ResponseUploadVideoDraftInit, error) {
	if request == nil {
		return nil, ErrNilOAuthRequest
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodPost, BaseURL+"/v2/post/publish/inbox/video/init/", request.Body)
	if e != nil {
		return nil, e
	}
	var data ResponseUploadVideoDraftInitSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseUploadVideoDraftInit{Success: data}, nil
}

// DirectPostPhotoInit calls POST https://open.tiktokapis.com/v2/post/publish/content/init/.
func (s *publishing) DirectPostPhotoInit(ctx context.Context, request *RequestDirectPostPhotoInit) (*ResponseDirectPostPhotoInit, error) {
	if request == nil {
		return nil, ErrNilOAuthRequest
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodPost, BaseURL+"/v2/post/publish/content/init/", request.Body)
	if e != nil {
		return nil, e
	}
	var data ResponseDirectPostPhotoInitSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseDirectPostPhotoInit{Success: data}, nil
}

// GetPublishStatus calls POST https://open.tiktokapis.com/v2/post/publish/status/fetch/.
func (s *publishing) GetPublishStatus(ctx context.Context, request *RequestGetPublishStatus) (*ResponseGetPublishStatus, error) {
	if request == nil {
		return nil, ErrNilOAuthRequest
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodPost, BaseURL+"/v2/post/publish/status/fetch/", request.Body)
	if e != nil {
		return nil, e
	}
	var data ResponseGetPublishStatusSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseGetPublishStatus{Success: data}, nil
}
