// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package facebook

import (
	"context"
	"net/http"
	"net/url"
)

// OAuth provides access to oauth endpoints.
type oAuth struct {
	client *FacebookClient
}

// NewOAuth creates a OAuth endpoint group using client.
func NewOAuth(client *FacebookClient) *oAuth {
	return &oAuth{client: client}
}

type (
	RequestAuthorizeQuery struct {
		ClientID     string  `json:"client_id"`
		RedirectUri  string  `json:"redirect_uri"`
		State        string  `json:"state"`
		Scopes       []Scope `json:"scope"`
		ResponseType string  `json:"response_type"`
	}

	RequestAuthorize struct {
		Query RequestAuthorizeQuery `json:"query"`
	}

	ResponseAuthorizeSuccessRedirectQuery struct {
		Code  string `json:"code"`
		State string `json:"state"`
	}

	ResponseAuthorizeErrorRedirectQuery struct {
		Error            string `json:"error"`
		ErrorReason      string `json:"error_reason"`
		ErrorDescription string `json:"error_description"`
	}

	ResponseAuthorize struct {
		URL                  string                                `json:"url"`
		SuccessRedirectQuery ResponseAuthorizeSuccessRedirectQuery `json:"success_redirect_query"`
		ErrorRedirectQuery   ResponseAuthorizeErrorRedirectQuery   `json:"error_redirect_query"`
	}

	RequestExchangeCodeQuery struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RedirectUri  string `json:"redirect_uri"`
		Code         string `json:"code"`
	}

	RequestExchangeCode struct {
		Query RequestExchangeCodeQuery `json:"query"`
	}

	ResponseExchangeCodeSuccess struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	ResponseExchangeCode struct {
		Success ResponseExchangeCodeSuccess `json:"success"`
	}

	RequestExchangeLongLivedUserTokenQuery struct {
		GrantType       string `json:"grant_type"`
		ClientID        string `json:"client_id"`
		ClientSecret    string `json:"client_secret"`
		FbExchangeToken string `json:"fb_exchange_token"`
	}

	RequestExchangeLongLivedUserToken struct {
		Query RequestExchangeLongLivedUserTokenQuery `json:"query"`
	}

	ResponseExchangeLongLivedUserTokenSuccess struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	ResponseExchangeLongLivedUserToken struct {
		Success ResponseExchangeLongLivedUserTokenSuccess `json:"success"`
	}

	RequestDebugTokenQuery struct {
		InputToken  string `json:"input_token"`
		AccessToken string `json:"access_token"`
	}

	RequestDebugToken struct {
		Query RequestDebugTokenQuery `json:"query"`
	}

	ResponseDebugTokenSuccessData struct {
		AppID       string   `json:"app_id"`
		Type        string   `json:"type"`
		Application string   `json:"application"`
		ExpiresAt   int      `json:"expires_at"`
		IsValid     bool     `json:"is_valid"`
		Scopes      []string `json:"scopes"`
		UserID      string   `json:"user_id"`
	}

	ResponseDebugTokenSuccess struct {
		Data ResponseDebugTokenSuccessData `json:"data"`
	}

	ResponseDebugToken struct {
		Success ResponseDebugTokenSuccess `json:"success"`
	}
)

// Authorize calls GET https://www.facebook.com/v24.0/dialog/oauth.
func (s *oAuth) Authorize(ctx context.Context, request *RequestAuthorize) (*ResponseAuthorize, error) {
	if s.client == nil {
		return nil, ErrNilHTTPClient
	}
	query := RequestAuthorizeQuery{}
	if request != nil {
		query = request.Query
	}
	if query.ClientID == "" {
		query.ClientID = s.client.clientID
	}
	if query.RedirectUri == "" {
		query.RedirectUri = s.client.redirectURL
	}
	if len(query.Scopes) == 0 {
		query.Scopes = s.client.scopes
	}
	if len(query.Scopes) == 0 {
		return nil, ErrMissingScopes
	}
	if query.ResponseType == "" {
		query.ResponseType = "code"
	}
	values := url.Values{"client_id": {query.ClientID}, "redirect_uri": {query.RedirectUri}, "scope": {scopeValue(query.Scopes)}, "response_type": {stringValue(query.ResponseType)}}
	if query.State != "" {
		values.Set("state", query.State)
	}
	return &ResponseAuthorize{URL: AuthorizationBaseURL + "/dialog/oauth?" + values.Encode()}, nil
}

// ExchangeCode calls GET https://graph.facebook.com/v24.0/oauth/access_token.
func (s *oAuth) ExchangeCode(ctx context.Context, request *RequestExchangeCode) (*ResponseExchangeCode, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := request.Query
	if query.Code == "" {
		return nil, ErrMissingCode
	}
	if query.ClientID == "" {
		query.ClientID = s.client.clientID
	}
	if query.ClientSecret == "" {
		query.ClientSecret = s.client.clientSecret
	}
	if query.RedirectUri == "" {
		query.RedirectUri = s.client.redirectURL
	}
	values := url.Values{"client_id": {query.ClientID}, "client_secret": {query.ClientSecret}, "redirect_uri": {query.RedirectUri}, "code": {query.Code}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodGet, APIBaseURL+"/oauth/access_token?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseExchangeCode)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// ExchangeLongLivedUserToken calls GET https://graph.facebook.com/v24.0/oauth/access_token.
func (s *oAuth) ExchangeLongLivedUserToken(ctx context.Context, request *RequestExchangeLongLivedUserToken) (*ResponseExchangeLongLivedUserToken, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := request.Query
	if query.FbExchangeToken == "" {
		return nil, ErrMissingAccessToken
	}
	if query.ClientID == "" {
		query.ClientID = s.client.clientID
	}
	if query.ClientSecret == "" {
		query.ClientSecret = s.client.clientSecret
	}
	if query.GrantType == "" {
		query.GrantType = "fb_exchange_token"
	}
	values := url.Values{"grant_type": {stringValue(query.GrantType)}, "client_id": {query.ClientID}, "client_secret": {query.ClientSecret}, "fb_exchange_token": {stringValue(query.FbExchangeToken)}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodGet, APIBaseURL+"/oauth/access_token?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseExchangeLongLivedUserToken)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// DebugToken calls GET https://graph.facebook.com/debug_token.
func (s *oAuth) DebugToken(ctx context.Context, request *RequestDebugToken) (*ResponseDebugToken, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := request.Query
	if query.InputToken == "" || query.AccessToken == "" {
		return nil, ErrMissingAccessToken
	}
	values := url.Values{"input_token": {stringValue(query.InputToken)}, "access_token": {stringValue(query.AccessToken)}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodGet, BaseURL+"/debug_token?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseDebugToken)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
