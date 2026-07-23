package x

import (
	"context"
	"net/http"
	"net/url"
)

// users provides access to users endpoints.
type users struct {
	client *XClient
}

// NewUsers creates a users endpoint group using client.
func NewUsers(client *XClient) *users {
	return &users{client: client}
}

type (
	RequestGetAuthenticatedUserQuery struct {
		UserFields any `json:"user.fields"`
	}

	RequestGetAuthenticatedUser struct {
		Query RequestGetAuthenticatedUserQuery `json:"query"`
	}

	ResponseGetAuthenticatedUserSuccessDataPublicMetrics struct {
		FollowersCount int `json:"followers_count"`
		FollowingCount int `json:"following_count"`
		TweetCount     int `json:"tweet_count"`
		ListedCount    int `json:"listed_count"`
	}

	ResponseGetAuthenticatedUserSuccessData struct {
		ID              string                                               `json:"id"`
		Name            string                                               `json:"name"`
		Username        string                                               `json:"username"`
		ProfileImageURL string                                               `json:"profile_image_url"`
		Description     string                                               `json:"description"`
		PublicMetrics   ResponseGetAuthenticatedUserSuccessDataPublicMetrics `json:"public_metrics"`
	}

	ResponseGetAuthenticatedUserSuccess struct {
		Data ResponseGetAuthenticatedUserSuccessData `json:"data"`
	}

	ResponseGetAuthenticatedUser struct {
		Success ResponseGetAuthenticatedUserSuccess `json:"success"`
	}

	RequestGetUserByIDPath struct {
		ID string `json:"id"`
	}

	RequestGetUserByIDQuery struct {
		UserFields any `json:"user.fields"`
	}

	RequestGetUserByID struct {
		Path  IDPath                  `json:"path"`
		Query RequestGetUserByIDQuery `json:"query"`
	}

	ResponseGetUserByIDSuccessData struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	}

	ResponseGetUserByIDSuccess struct {
		Data ResponseGetUserByIDSuccessData `json:"data"`
	}

	ResponseGetUserByID struct {
		Success ResponseGetUserByIDSuccess `json:"success"`
	}
)

// GetAuthenticatedUser calls GET https://api.x.com/2/users/me.
func (s *users) GetAuthenticatedUser(ctx context.Context, request *RequestGetAuthenticatedUser) (*ResponseGetAuthenticatedUser, error) {
	if s.client == nil {
		return nil, ErrMissingAccessToken
	}
	query := url.Values{}
	if request != nil {
		addOptionalQuery(query, "user.fields", request.Query.UserFields)
	}
	rawURL := BaseURL + "/users/me"
	if encoded := query.Encode(); encoded != "" {
		rawURL += "?" + encoded
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	var raw struct {
		Data ResponseGetAuthenticatedUserSuccessData `json:"data"`
	}
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseGetAuthenticatedUser{Success: ResponseGetAuthenticatedUserSuccess{Data: raw.Data}}, nil
}

// GetUserByID calls GET https://api.x.com/2/users/{id}.
func (s *users) GetUserByID(ctx context.Context, request *RequestGetUserByID) (*ResponseGetUserByID, error) {
	if s.client == nil {
		return nil, ErrMissingAccessToken
	}
	if request == nil || request.Path.ID == "" {
		return nil, ErrMissingID
	}
	query := url.Values{}
	addOptionalQuery(query, "user.fields", request.Query.UserFields)
	rawURL := BaseURL + "/users/" + url.PathEscape(request.Path.ID)
	if encoded := query.Encode(); encoded != "" {
		rawURL += "?" + encoded
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	var raw struct {
		Data ResponseGetUserByIDSuccessData `json:"data"`
	}
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseGetUserByID{Success: ResponseGetUserByIDSuccess{Data: raw.Data}}, nil
}
