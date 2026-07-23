// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package threads

import (
	"context"
	"net/http"
	"net/url"
)

// Analytics provides access to analytics endpoints.
type analytics struct {
	client *ThreadsClient
}

// NewAnalytics creates a Analytics endpoint group using client.
func NewAnalytics(client *ThreadsClient) *analytics {
	return &analytics{client: client}
}

type (
	RequestGetPostInsightsPath struct {
		ThreadsMediaID string `json:"threads_media_id"`
	}

	RequestGetPostInsightsQuery struct {
		Metric any `json:"metric"`
	}

	RequestGetPostInsights struct {
		Path  RequestGetPostInsightsPath  `json:"path"`
		Query RequestGetPostInsightsQuery `json:"query"`
	}

	ResponseGetPostInsightsSuccessDataItemValuesItem struct {
		Value float64 `json:"value"`
	}

	ResponseGetPostInsightsSuccessDataItem struct {
		Name        string                                             `json:"name"`
		Period      any                                                `json:"period"`
		Values      []ResponseGetPostInsightsSuccessDataItemValuesItem `json:"values"`
		Title       string                                             `json:"title"`
		Description string                                             `json:"description"`
	}

	ResponseGetPostInsightsSuccess struct {
		Data []ResponseGetPostInsightsSuccessDataItem `json:"data"`
	}

	ResponseGetPostInsights struct {
		Success ResponseGetPostInsightsSuccess `json:"success"`
	}
)

// GetPostInsights calls GET https://graph.threads.net/v1.0/{threads_media_id}/insights.
func (s *analytics) GetPostInsights(ctx context.Context, request *RequestGetPostInsights) (*ResponseGetPostInsights, error) {
	if request == nil || request.Path.ThreadsMediaID == "" {
		return nil, ErrMissingID
	}
	q := url.Values{}
	addOptionalQuery(q, "metric", request.Query.Metric)
	raw := BaseURL + "/v1.0/" + url.PathEscape(request.Path.ThreadsMediaID) + "/insights"
	if q.Encode() != "" {
		raw += "?" + q.Encode()
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodGet, raw, nil)
	if e != nil {
		return nil, e
	}
	var data ResponseGetPostInsightsSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseGetPostInsights{Success: data}, nil
}
