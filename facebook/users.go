package facebook

import (
	"context"
	"net/http"
	"net/url"
)

// Users provides access to users endpoints.
type users struct {
	client *FacebookClient
}

// NewUsers creates a Users endpoint group using client.
func NewUsers(client *FacebookClient) *users {
	return &users{client: client}
}

type (
	RequestGetMeQuery struct {
		Fields []UserField `json:"fields"`
	}

	RequestGetMe struct {
		Query RequestGetMeQuery `json:"query"`
	}

	ResponseGetMeSuccessPicture struct {
		URL          string `json:"url"`
		Height       int    `json:"height"`
		Width        int    `json:"width"`
		IsSilhouette bool   `json:"is_silhouette"`
	}

	ResponseGetMeSuccess struct {
		ID        string                      `json:"id"`
		Name      string                      `json:"name"`
		FirstName string                      `json:"first_name"`
		LastName  string                      `json:"last_name"`
		Email     *string                     `json:"email"`
		Link      string                      `json:"link"`
		Picture   ResponseGetMeSuccessPicture `json:"picture"`
	}

	ResponseGetMe struct {
		Success ResponseGetMeSuccess `json:"success"`
	}
)

// GetMe calls GET https://graph.facebook.com/v24.0/me.
func (s *users) GetMe(ctx context.Context, request *RequestGetMe) (*ResponseGetMe, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := url.Values{"fields": {stringValue(request.Query.Fields)}}
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
