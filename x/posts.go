package x

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// posts provides access to posts endpoints.
type posts struct {
	client *XClient
}

// NewPosts creates a posts endpoint group using client.
func NewPosts(client *XClient) *posts {
	return &posts{client: client}
}

type (
	RequestCreatePostHeaders struct {
		ContentType any `json:"Content-Type"`
	}

	RequestCreatePostBodyMedia struct {
		MediaIds      []string `json:"media_ids"`
		TaggedUserIds []string `json:"tagged_user_ids"`
	}

	RequestCreatePostBodyReply struct {
		InReplyToTweetID    string   `json:"in_reply_to_tweet_id"`
		ExcludeReplyUserIds []string `json:"exclude_reply_user_ids"`
	}

	RequestCreatePostBodyPoll struct {
		Options         []string `json:"options"`
		DurationMinutes int      `json:"duration_minutes"`
	}

	RequestCreatePostBody struct {
		Text          *string                    `json:"text"`
		Media         RequestCreatePostBodyMedia `json:"media"`
		Reply         RequestCreatePostBodyReply `json:"reply"`
		QuoteTweetID  string                     `json:"quote_tweet_id"`
		Poll          RequestCreatePostBodyPoll  `json:"poll"`
		ReplySettings ReplySetting               `json:"reply_settings"`
	}

	RequestCreatePost struct {
		Headers RequestCreatePostHeaders `json:"headers"`
		Body    RequestCreatePostBody    `json:"body"`
	}

	ResponseCreatePostSuccessData struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}

	ResponseCreatePostSuccess struct {
		Data ResponseCreatePostSuccessData `json:"data"`
	}

	ResponseCreatePost struct {
		Success ResponseCreatePostSuccess `json:"success"`
	}

	RequestDeletePostPath struct {
		ID string `json:"id"`
	}

	RequestDeletePost struct {
		Path IDPath `json:"path"`
	}

	ResponseDeletePostSuccessData struct {
		Deleted bool `json:"deleted"`
	}

	ResponseDeletePostSuccess struct {
		Data ResponseDeletePostSuccessData `json:"data"`
	}

	ResponseDeletePost struct {
		Success ResponseDeletePostSuccess `json:"success"`
	}

	RequestListUserPostsPath struct {
		ID string `json:"id"`
	}

	RequestListUserPostsQuery struct {
		MaxResults      int    `json:"max_results"`
		PaginationToken string `json:"pagination_token"`
		TweetFields     any    `json:"tweet.fields"`
		Expansions      any    `json:"expansions"`
	}

	RequestListUserPosts struct {
		Path  IDPath                    `json:"path"`
		Query RequestListUserPostsQuery `json:"query"`
	}

	ResponseListUserPostsSuccessDataItem struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}

	ResponseListUserPostsSuccessMeta struct {
		ResultCount int     `json:"result_count"`
		NextToken   *string `json:"next_token"`
	}

	ResponseListUserPostsSuccess struct {
		Data []ResponseListUserPostsSuccessDataItem `json:"data"`
		Meta ResponseListUserPostsSuccessMeta       `json:"meta"`
	}

	ResponseListUserPosts struct {
		Success ResponseListUserPostsSuccess `json:"success"`
	}
)

// CreatePost calls POST https://api.x.com/2/tweets.
func (s *posts) CreatePost(ctx context.Context, request *RequestCreatePost) (*ResponseCreatePost, error) {
	if s.client == nil {
		return nil, ErrMissingAccessToken
	}
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	body := make(map[string]any)
	if request.Body.Text != nil {
		body["text"] = *request.Body.Text
	}
	if len(request.Body.Media.MediaIds) != 0 || len(request.Body.Media.TaggedUserIds) != 0 {
		body["media"] = request.Body.Media
	}
	if request.Body.Reply.InReplyToTweetID != "" || len(request.Body.Reply.ExcludeReplyUserIds) != 0 {
		body["reply"] = request.Body.Reply
	}
	if request.Body.QuoteTweetID != "" {
		body["quote_tweet_id"] = request.Body.QuoteTweetID
	}
	if len(request.Body.Poll.Options) != 0 {
		body["poll"] = request.Body.Poll
	}
	if request.Body.ReplySettings != "" {
		body["reply_settings"] = request.Body.ReplySettings
	}
	if len(body) == 0 {
		return nil, ErrMissingPostContent
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, BaseURL+"/tweets", body)
	if err != nil {
		return nil, err
	}
	var raw struct {
		Data ResponseCreatePostSuccessData `json:"data"`
	}
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseCreatePost{Success: ResponseCreatePostSuccess{Data: raw.Data}}, nil
}

// DeletePost calls DELETE https://api.x.com/2/tweets/{id}.
func (s *posts) DeletePost(ctx context.Context, request *RequestDeletePost) (*ResponseDeletePost, error) {
	if s.client == nil {
		return nil, ErrMissingAccessToken
	}
	if request == nil || request.Path.ID == "" {
		return nil, ErrMissingID
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodDelete, BaseURL+"/tweets/"+url.PathEscape(request.Path.ID), nil)
	if err != nil {
		return nil, err
	}
	var raw struct {
		Data ResponseDeletePostSuccessData `json:"data"`
	}
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseDeletePost{Success: ResponseDeletePostSuccess{Data: raw.Data}}, nil
}

// ListUserPosts calls GET https://api.x.com/2/users/{id}/tweets.
func (s *posts) ListUserPosts(ctx context.Context, request *RequestListUserPosts) (*ResponseListUserPosts, error) {
	if s.client == nil {
		return nil, ErrMissingAccessToken
	}
	if request == nil || request.Path.ID == "" {
		return nil, ErrMissingID
	}
	query := url.Values{}
	if request.Query.MaxResults != 0 {
		query.Set("max_results", fmt.Sprint(request.Query.MaxResults))
	}
	addOptionalQuery(query, "pagination_token", request.Query.PaginationToken)
	addOptionalQuery(query, "tweet.fields", request.Query.TweetFields)
	addOptionalQuery(query, "expansions", request.Query.Expansions)
	rawURL := BaseURL + "/users/" + url.PathEscape(request.Path.ID) + "/tweets"
	if encoded := query.Encode(); encoded != "" {
		rawURL += "?" + encoded
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	var raw struct {
		Data []ResponseListUserPostsSuccessDataItem `json:"data"`
		Meta ResponseListUserPostsSuccessMeta       `json:"meta"`
	}
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseListUserPosts{Success: ResponseListUserPostsSuccess{Data: raw.Data, Meta: raw.Meta}}, nil
}
