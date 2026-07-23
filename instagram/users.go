// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package instagram

import (
	"context"
	"net/http"
	"net/url"
)

// Users provides access to users endpoints.
type users struct {
	client *InstagramClient
}

// NewUsers creates a Users endpoint group using client.
func NewUsers(client *InstagramClient) *users {
	return &users{client: client}
}

type (
	RequestGetMeQuery struct {
		Fields      []UserField `json:"fields"`
		AccessToken string      `json:"access_token"`
	}

	RequestGetMe struct {
		Query RequestGetMeQuery `json:"query"`
	}

	ResponseGetMeSuccess struct {
		ID                string      `json:"id"`
		UserID            string      `json:"user_id"`
		Username          string      `json:"username"`
		Name              string      `json:"name"`
		AccountType       AccountType `json:"account_type"`
		ProfilePictureURL string      `json:"profile_picture_url"`
		FollowersCount    int         `json:"followers_count"`
		FollowsCount      int         `json:"follows_count"`
		MediaCount        int         `json:"media_count"`
		Biography         string      `json:"biography"`
		Website           string      `json:"website"`
	}

	ResponseGetMe struct {
		Success ResponseGetMeSuccess `json:"success"`
	}
)

// GetMe calls GET https://graph.instagram.com/v24.0/me.
func (s *users) GetMe(ctx context.Context, request *RequestGetMe) (*ResponseGetMe, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := url.Values{"fields": {stringValue(request.Query.Fields)}}
	if request.Query.AccessToken != "" {
		query.Set("access_token", request.Query.AccessToken)
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, APIBaseURL+"/me?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseGetMe)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
