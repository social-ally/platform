package x

import (
	"context"
	"net/http"
	"net/url"
)

// Analytics provides access to analytics endpoints.
type Analytics struct {
	client *XClient
}

// NewAnalytics creates a Analytics endpoint group using client.
func NewAnalytics(client *XClient) *Analytics {
	return &Analytics{client: client}
}

type (
	RequestGetPostMetricsPath struct {
		ID string `json:"id"`
	}

	RequestGetPostMetricsQuery struct {
		TweetFields any `json:"tweet.fields"`
	}

	RequestGetPostMetrics struct {
		Path  IDPath                     `json:"path"`
		Query RequestGetPostMetricsQuery `json:"query"`
	}

	ResponseGetPostMetricsSuccessDataPublicMetrics struct {
		RetweetCount    int `json:"retweet_count"`
		ReplyCount      int `json:"reply_count"`
		LikeCount       int `json:"like_count"`
		QuoteCount      int `json:"quote_count"`
		BookmarkCount   int `json:"bookmark_count"`
		ImpressionCount int `json:"impression_count"`
	}

	ResponseGetPostMetricsSuccessData struct {
		ID            string                                         `json:"id"`
		PublicMetrics ResponseGetPostMetricsSuccessDataPublicMetrics `json:"public_metrics"`
	}

	ResponseGetPostMetricsSuccess struct {
		Data ResponseGetPostMetricsSuccessData `json:"data"`
	}

	ResponseGetPostMetrics struct {
		Success ResponseGetPostMetricsSuccess `json:"success"`
	}
)

// GetPostMetrics calls GET https://api.x.com/2/tweets/{id}.
func (s *Analytics) GetPostMetrics(ctx context.Context, request *RequestGetPostMetrics) (*ResponseGetPostMetrics, error) {
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
	var raw struct {
		Data ResponseGetPostMetricsSuccessData `json:"data"`
	}
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseGetPostMetrics{Success: ResponseGetPostMetricsSuccess{Data: raw.Data}}, nil
}
