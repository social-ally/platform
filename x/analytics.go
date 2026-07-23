package x

import (
	"context"
	"net/http"
	"net/url"
)

// analytics provides access to analytics endpoints.
type analytics struct {
	client *XClient
}

// NewAnalytics creates an analytics endpoint group using client.
func NewAnalytics(client *XClient) *analytics {
	return &analytics{client: client}
}

type (
	RequestGetPostMetricsPath struct {
		ID string `json:"id"`
	}

	RequestGetPostMetricsQuery struct {
		TweetFields []TweetField `json:"tweet.fields"`
	}

	RequestGetPostMetrics struct {
		Path  IDPath                     `json:"path"`
		Query RequestGetPostMetricsQuery `json:"query"`
	}

	ResponseGetPostMetricsSuccessData = Tweet

	ResponseGetPostMetricsSuccess struct {
		Data     ResponseGetPostMetricsSuccessData `json:"data"`
		Errors   []Problem                         `json:"errors"`
		Includes *Includes                         `json:"includes"`
	}

	ResponseGetPostMetrics struct {
		Success ResponseGetPostMetricsSuccess `json:"success"`
	}
)

// GetPostMetrics calls GET https://api.x.com/2/tweets/{id}.
func (s *analytics) GetPostMetrics(ctx context.Context, request *RequestGetPostMetrics) (*ResponseGetPostMetrics, error) {
	if s.client == nil {
		return nil, ErrMissingAccessToken
	}
	if request == nil || request.Path.ID == "" {
		return nil, ErrMissingID
	}
	query := url.Values{}
	addOptionalQuery(query, "tweet.fields", request.Query.TweetFields)
	rawURL := BaseURL + "/tweets/" + url.PathEscape(request.Path.ID)
	if encoded := query.Encode(); encoded != "" {
		rawURL += "?" + encoded
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	var raw ResponseGetPostMetricsSuccess
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseGetPostMetrics{Success: raw}, nil
}
