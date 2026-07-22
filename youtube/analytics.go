// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package youtube

import (
	"context"
	"net/http"
	"net/url"
)

// Analytics provides access to analytics endpoints.
type Analytics struct {
	client *YouTubeClient
}

// NewAnalytics creates a Analytics endpoint group using client.
func NewAnalytics(client *YouTubeClient) *Analytics {
	return &Analytics{client: client}
}

type (
	RequestQueryChannelAnalyticsQuery struct {
		Ids        any     `json:"ids"`
		StartDate  any     `json:"startDate"`
		EndDate    any     `json:"endDate"`
		Metrics    any     `json:"metrics"`
		Dimensions any     `json:"dimensions"`
		Filters    *string `json:"filters"`
		Sort       *string `json:"sort"`
	}

	RequestQueryChannelAnalytics struct {
		Query RequestQueryChannelAnalyticsQuery `json:"query"`
	}

	ResponseQueryChannelAnalyticsSuccessColumnHeadersItem struct {
		Name       string `json:"name"`
		ColumnType any    `json:"columnType"`
		DataType   any    `json:"dataType"`
	}

	ResponseQueryChannelAnalyticsSuccess struct {
		ColumnHeaders []ResponseQueryChannelAnalyticsSuccessColumnHeadersItem `json:"columnHeaders"`
		Rows          [][]any                                                 `json:"rows"`
	}

	ResponseQueryChannelAnalytics struct {
		Success ResponseQueryChannelAnalyticsSuccess `json:"success"`
	}
)

// QueryChannelAnalytics calls GET https://youtubeanalytics.googleapis.com/v2/reports.
func (s *Analytics) QueryChannelAnalytics(ctx context.Context, request *RequestQueryChannelAnalytics) (*ResponseQueryChannelAnalytics, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := url.Values{"ids": {stringValue(request.Query.Ids)}, "startDate": {stringValue(request.Query.StartDate)}, "endDate": {stringValue(request.Query.EndDate)}, "metrics": {stringValue(request.Query.Metrics)}, "dimensions": {stringValue(request.Query.Dimensions)}}
	if request.Query.Filters != nil {
		query.Set("filters", *request.Query.Filters)
	}
	if request.Query.Sort != nil {
		query.Set("sort", *request.Query.Sort)
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, AnalyticsBaseURL+"/reports?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseQueryChannelAnalytics)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
