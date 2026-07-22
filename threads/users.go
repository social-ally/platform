// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package threads

import (
	"context"
	"net/http"
	"net/url"
)

// Users provides access to users endpoints.
type Users struct {
	client *ThreadsClient
}

// NewUsers creates a Users endpoint group using client.
func NewUsers(client *ThreadsClient) *Users {
	return &Users{client: client}
}

type (
	RequestGetMeQuery struct {
		Fields any `json:"fields"`
	}

	RequestGetMe struct {
		Query RequestGetMeQuery `json:"query"`
	}

	ResponseGetMeSuccess struct {
		ID                       string `json:"id"`
		Username                 string `json:"username"`
		Name                     string `json:"name"`
		ThreadsProfilePictureURL string `json:"threads_profile_picture_url"`
		ThreadsBiography         string `json:"threads_biography"`
	}

	ResponseGetMe struct {
		Success ResponseGetMeSuccess `json:"success"`
	}
)

// GetMe calls GET https://graph.threads.net/v1.0/me.
func (s *Users) GetMe(ctx context.Context, request *RequestGetMe) (*ResponseGetMe, error) {
	q := url.Values{}
	if request != nil {
		addOptionalQuery(q, "fields", request.Query.Fields)
	}
	raw := BaseURL + "/v1.0/me"
	if q.Encode() != "" {
		raw += "?" + q.Encode()
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodGet, raw, nil)
	if e != nil {
		return nil, e
	}
	var data ResponseGetMeSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseGetMe{Success: data}, nil
}
