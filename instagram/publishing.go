// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package instagram

import (
	"context"
	"net/http"
	"net/url"
)

// Publishing provides access to publishing endpoints.
type publishing struct {
	client *InstagramClient
}

// NewPublishing creates a Publishing endpoint group using client.
func NewPublishing(client *InstagramClient) *publishing {
	return &publishing{client: client}
}

type (
	RequestCreateMediaContainerPath struct {
		IgUserID string `json:"ig_user_id"`
	}

	RequestCreateMediaContainerBody struct {
		ImageURL       any                 `json:"image_url"`
		VideoURL       any                 `json:"video_url"`
		MediaType      PublishingMediaType `json:"media_type"`
		Caption        *string             `json:"caption"`
		ShareToFeed    *bool               `json:"share_to_feed"`
		CoverURL       any                 `json:"cover_url"`
		ThumbOffset    *int                `json:"thumb_offset"`
		Children       []any               `json:"children"`
		IsCarouselItem *bool               `json:"is_carousel_item"`
		IsAiGenerated  *bool               `json:"is_ai_generated"`
		AccessToken    string              `json:"access_token"`
	}

	RequestCreateMediaContainer struct {
		Path RequestCreateMediaContainerPath `json:"path"`
		Body RequestCreateMediaContainerBody `json:"body"`
	}

	ResponseCreateMediaContainerSuccess struct {
		ID any `json:"id"`
	}

	ResponseCreateMediaContainer struct {
		Success ResponseCreateMediaContainerSuccess `json:"success"`
	}

	RequestGetContainerStatusPath struct {
		CreationID string `json:"creation_id"`
	}

	RequestGetContainerStatusQuery struct {
		Fields      any    `json:"fields"`
		AccessToken string `json:"access_token"`
	}

	RequestGetContainerStatus struct {
		Path  RequestGetContainerStatusPath  `json:"path"`
		Query RequestGetContainerStatusQuery `json:"query"`
	}

	ResponseGetContainerStatusSuccess struct {
		ID         string          `json:"id"`
		StatusCode ContainerStatus `json:"status_code"`
		Status     string          `json:"status"`
	}

	ResponseGetContainerStatus struct {
		Success ResponseGetContainerStatusSuccess `json:"success"`
	}

	RequestPublishMediaPath struct {
		IgUserID string `json:"ig_user_id"`
	}

	RequestPublishMediaBody struct {
		CreationID  string `json:"creation_id"`
		AccessToken string `json:"access_token"`
	}

	RequestPublishMedia struct {
		Path RequestPublishMediaPath `json:"path"`
		Body RequestPublishMediaBody `json:"body"`
	}

	ResponsePublishMediaSuccess struct {
		ID any `json:"id"`
	}

	ResponsePublishMedia struct {
		Success ResponsePublishMediaSuccess `json:"success"`
	}

	RequestGetPublishingLimitPath struct {
		IgUserID string `json:"ig_user_id"`
	}

	RequestGetPublishingLimitQuery struct {
		Fields      any    `json:"fields"`
		AccessToken string `json:"access_token"`
	}

	RequestGetPublishingLimit struct {
		Path  RequestGetPublishingLimitPath  `json:"path"`
		Query RequestGetPublishingLimitQuery `json:"query"`
	}

	ResponseGetPublishingLimitSuccessDataItemConfig struct {
		QuotaTotal    int `json:"quota_total"`
		QuotaDuration int `json:"quota_duration"`
	}

	ResponseGetPublishingLimitSuccessDataItem struct {
		Config     ResponseGetPublishingLimitSuccessDataItemConfig `json:"config"`
		QuotaUsage int                                             `json:"quota_usage"`
	}

	ResponseGetPublishingLimitSuccess struct {
		Data []ResponseGetPublishingLimitSuccessDataItem `json:"data"`
	}

	ResponseGetPublishingLimit struct {
		Success ResponseGetPublishingLimitSuccess `json:"success"`
	}
)

// CreateMediaContainer calls POST https://graph.instagram.com/v24.0/{ig_user_id}/media.
func (s *publishing) CreateMediaContainer(ctx context.Context, request *RequestCreateMediaContainer) (*ResponseCreateMediaContainer, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.IgUserID == "" {
		return nil, ErrMissingID
	}
	body, err := formValues(request.Body)
	if err != nil {
		return nil, err
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, APIBaseURL+"/"+url.PathEscape(request.Path.IgUserID)+"/media", body)
	if err != nil {
		return nil, err
	}
	response := new(ResponseCreateMediaContainer)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// GetContainerStatus calls GET https://graph.instagram.com/v24.0/{creation_id}.
func (s *publishing) GetContainerStatus(ctx context.Context, request *RequestGetContainerStatus) (*ResponseGetContainerStatus, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.CreationID == "" {
		return nil, ErrMissingID
	}
	query := url.Values{"fields": {stringValue(request.Query.Fields)}}
	if request.Query.AccessToken != "" {
		query.Set("access_token", request.Query.AccessToken)
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, APIBaseURL+"/"+url.PathEscape(request.Path.CreationID)+"?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseGetContainerStatus)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// PublishMedia calls POST https://graph.instagram.com/v24.0/{ig_user_id}/media_publish.
func (s *publishing) PublishMedia(ctx context.Context, request *RequestPublishMedia) (*ResponsePublishMedia, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.IgUserID == "" {
		return nil, ErrMissingID
	}
	body, err := formValues(request.Body)
	if err != nil {
		return nil, err
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, APIBaseURL+"/"+url.PathEscape(request.Path.IgUserID)+"/media_publish", body)
	if err != nil {
		return nil, err
	}
	response := new(ResponsePublishMedia)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// GetPublishingLimit calls GET https://graph.instagram.com/v24.0/{ig_user_id}/content_publishing_limit.
func (s *publishing) GetPublishingLimit(ctx context.Context, request *RequestGetPublishingLimit) (*ResponseGetPublishingLimit, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.IgUserID == "" {
		return nil, ErrMissingID
	}
	query := url.Values{"fields": {stringValue(request.Query.Fields)}}
	if request.Query.AccessToken != "" {
		query.Set("access_token", request.Query.AccessToken)
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, APIBaseURL+"/"+url.PathEscape(request.Path.IgUserID)+"/content_publishing_limit?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseGetPublishingLimit)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
