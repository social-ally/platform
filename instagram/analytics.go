// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package instagram

import (
	"context"
	"net/http"
	"net/url"
)

// Analytics provides access to analytics endpoints.
type analytics struct {
	client *InstagramClient
}

// NewAnalytics creates a Analytics endpoint group using client.
func NewAnalytics(client *InstagramClient) *analytics {
	return &analytics{client: client}
}

type (
	RequestGetMediaInsightsPath struct {
		MediaID string `json:"media_id"`
	}

	RequestGetMediaInsightsQuery struct {
		Metric any           `json:"metric"`
		Period InsightPeriod `json:"period"`
	}

	RequestGetMediaInsights struct {
		Path  RequestGetMediaInsightsPath  `json:"path"`
		Query RequestGetMediaInsightsQuery `json:"query"`
	}

	ResponseGetMediaInsightsSuccessDataItemValuesItem struct {
		Value   float64 `json:"value"`
		EndTime any     `json:"end_time"`
	}

	ResponseGetMediaInsightsSuccessDataItem struct {
		Name        string                                              `json:"name"`
		Period      string                                              `json:"period"`
		Values      []ResponseGetMediaInsightsSuccessDataItemValuesItem `json:"values"`
		Title       string                                              `json:"title"`
		Description string                                              `json:"description"`
		ID          string                                              `json:"id"`
	}

	ResponseGetMediaInsightsSuccess struct {
		Data []ResponseGetMediaInsightsSuccessDataItem `json:"data"`
	}

	ResponseGetMediaInsights struct {
		Success ResponseGetMediaInsightsSuccess `json:"success"`
	}
)

// GetMediaInsights calls GET https://graph.instagram.com/v24.0/{media_id}/insights.
func (s *analytics) GetMediaInsights(ctx context.Context, request *RequestGetMediaInsights) (*ResponseGetMediaInsights, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.MediaID == "" {
		return nil, ErrMissingID
	}
	query := url.Values{"metric": {stringValue(request.Query.Metric)}, "period": {stringValue(request.Query.Period)}}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, APIBaseURL+"/"+url.PathEscape(request.Path.MediaID)+"/insights?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseGetMediaInsights)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
