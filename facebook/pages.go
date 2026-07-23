package facebook

import (
	"context"
	"net/http"
	"net/url"
)

// Pages provides access to pages endpoints.
type pages struct {
	client *FacebookClient
}

// NewPages creates a Pages endpoint group using client.
func NewPages(client *FacebookClient) *pages {
	return &pages{client: client}
}

type (
	RequestListManagedPagesQuery struct {
		Fields []PageField `json:"fields"`
	}

	RequestListManagedPages struct {
		Query RequestListManagedPagesQuery `json:"query"`
	}

	ResponseListManagedPagesSuccessDataItem struct {
		ID                 string `json:"id"`
		Name               string `json:"name"`
		AccessToken        any    `json:"access_token"`
		Category           string `json:"category"`
		CategoryList       []any  `json:"category_list"`
		Tasks              []any  `json:"tasks"`
		About              string `json:"about"`
		Description        string `json:"description"`
		FanCount           int    `json:"fan_count"`
		FollowersCount     int    `json:"followers_count"`
		Website            string `json:"website"`
		Username           string `json:"username"`
		Picture            any    `json:"picture"`
		VerificationStatus string `json:"verification_status"`
	}

	ResponseListManagedPagesSuccessPaging struct {
		Cursors  any    `json:"cursors"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
	}

	ResponseListManagedPagesSuccess struct {
		Data   []ResponseListManagedPagesSuccessDataItem `json:"data"`
		Paging ResponseListManagedPagesSuccessPaging     `json:"paging"`
	}

	ResponseListManagedPages struct {
		Success ResponseListManagedPagesSuccess `json:"success"`
	}

	RequestGetPagePath struct {
		PageID string `json:"page_id"`
	}

	RequestGetPageQuery struct {
		Fields []PageField `json:"fields"`
	}

	RequestGetPage struct {
		Path  RequestGetPagePath  `json:"path"`
		Query RequestGetPageQuery `json:"query"`
	}

	ResponseGetPageSuccess struct {
		ID                 string `json:"id"`
		Name               string `json:"name"`
		FanCount           int    `json:"fan_count"`
		FollowersCount     int    `json:"followers_count"`
		AccessToken        string `json:"access_token"`
		Category           string `json:"category"`
		CategoryList       []any  `json:"category_list"`
		Tasks              []any  `json:"tasks"`
		About              string `json:"about"`
		Description        string `json:"description"`
		Website            string `json:"website"`
		Username           string `json:"username"`
		Picture            any    `json:"picture"`
		VerificationStatus string `json:"verification_status"`
	}

	ResponseGetPage struct {
		Success ResponseGetPageSuccess `json:"success"`
	}
)

// ListManagedPages calls GET https://graph.facebook.com/v24.0/me/accounts.
func (s *pages) ListManagedPages(ctx context.Context, request *RequestListManagedPages) (*ResponseListManagedPages, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := url.Values{"fields": {stringValue(request.Query.Fields)}}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, APIBaseURL+"/me/accounts?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseListManagedPages)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// GetPage calls GET https://graph.facebook.com/v24.0/{page_id}.
func (s *pages) GetPage(ctx context.Context, request *RequestGetPage) (*ResponseGetPage, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Path.PageID == "" {
		return nil, ErrMissingID
	}
	query := url.Values{"fields": {stringValue(request.Query.Fields)}}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, APIBaseURL+"/"+url.PathEscape(request.Path.PageID)+"?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseGetPage)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
