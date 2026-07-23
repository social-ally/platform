// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package facebook

import (
	"context"
	"net/http"
	"net/url"
)

// Analytics provides access to analytics endpoints.
type analytics struct {
	client *facebookClient
}

// NewAnalytics creates a Analytics endpoint group using client.
func NewAnalytics(client *facebookClient) *analytics {
	return &analytics{client: client}
}

type (
	RequestGetPageInsightsPath struct {
		PageID string `json:"page_id"`
	}

	RequestGetPageInsightsQuery struct {
		Metric any           `json:"metric"`
		Period InsightPeriod `json:"period"`
		Since  any           `json:"since"`
		Until  any           `json:"until"`
	}

	RequestGetPageInsights struct {
		Path  RequestGetPageInsightsPath  `json:"path"`
		Query RequestGetPageInsightsQuery `json:"query"`
	}

	ResponseGetPageInsightsSuccessDataItemValuesItem struct {
		Value   any `json:"value"`
		EndTime any `json:"end_time"`
	}

	ResponseGetPageInsightsSuccessDataItem struct {
		Name        string                                             `json:"name"`
		Period      string                                             `json:"period"`
		Values      []ResponseGetPageInsightsSuccessDataItemValuesItem `json:"values"`
		Title       string                                             `json:"title"`
		Description string                                             `json:"description"`
		ID          string                                             `json:"id"`
	}

	ResponseGetPageInsightsSuccess struct {
		Data []ResponseGetPageInsightsSuccessDataItem `json:"data"`
	}

	ResponseGetPageInsights struct {
		Success ResponseGetPageInsightsSuccess `json:"success"`
	}
)

// GetPageInsights calls GET https://graph.facebook.com/v24.0/{page_id}/insights.
func (s *analytics) GetPageInsights(ctx context.Context, request *RequestGetPageInsights) (*ResponseGetPageInsights, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.PageID == "" {
		return nil, ErrMissingID
	}
	query := url.Values{"metric": {stringValue(request.Query.Metric)}, "period": {stringValue(request.Query.Period)}, "since": {stringValue(request.Query.Since)}, "until": {stringValue(request.Query.Until)}}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, APIBaseURL+"/"+url.PathEscape(request.Path.PageID)+"/insights?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseGetPageInsights)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
