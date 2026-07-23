// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package instagram

import (
	"context"
	"net/http"
	"net/url"
)

// OAuth provides access to oauth endpoints.
type oAuth struct {
	client *InstagramClient
}

// NewOAuth creates a OAuth endpoint group using client.
func NewOAuth(client *InstagramClient) *oAuth {
	return &oAuth{client: client}
}

type (
	RequestAuthorizeQuery struct {
		EnableFbLogin       int     `json:"enable_fb_login"`
		ForceAuthentication int     `json:"force_authentication"`
		ClientID            string  `json:"client_id"`
		RedirectUri         string  `json:"redirect_uri"`
		ResponseType        any     `json:"response_type"`
		Scopes              []Scope `json:"scope"`
		State               string  `json:"state"`
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

	RequestExchangeCodeHeaders struct {
		ContentType any `json:"Content-Type"`
	}

	RequestExchangeCodeBody struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		GrantType    any    `json:"grant_type"`
		RedirectUri  string `json:"redirect_uri"`
		Code         string `json:"code"`
	}

	RequestExchangeCode struct {
		Headers RequestExchangeCodeHeaders `json:"headers"`
		Body    RequestExchangeCodeBody    `json:"body"`
	}

	ResponseExchangeCodeSuccess struct {
		AccessToken string  `json:"access_token"`
		UserID      int     `json:"user_id"`
		Permissions *string `json:"permissions"`
	}

	ResponseExchangeCode struct {
		Success ResponseExchangeCodeSuccess `json:"success"`
	}

	RequestExchangeLongLivedTokenQuery struct {
		GrantType    any    `json:"grant_type"`
		ClientSecret string `json:"client_secret"`
		AccessToken  any    `json:"access_token"`
	}

	RequestExchangeLongLivedToken struct {
		Query RequestExchangeLongLivedTokenQuery `json:"query"`
	}

	ResponseExchangeLongLivedTokenSuccess struct {
		AccessToken string `json:"access_token"`
		TokenType   any    `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	ResponseExchangeLongLivedToken struct {
		Success ResponseExchangeLongLivedTokenSuccess `json:"success"`
	}

	RequestRefreshLongLivedTokenQuery struct {
		GrantType   any    `json:"grant_type"`
		AccessToken string `json:"access_token"`
	}

	RequestRefreshLongLivedToken struct {
		Query RequestRefreshLongLivedTokenQuery `json:"query"`
	}

	ResponseRefreshLongLivedTokenSuccess struct {
		AccessToken string `json:"access_token"`
		TokenType   any    `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	ResponseRefreshLongLivedToken struct {
		Success ResponseRefreshLongLivedTokenSuccess `json:"success"`
	}
)

// Authorize calls GET https://www.instagram.com/oauth/authorize.
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
	if query.ResponseType == nil {
		query.ResponseType = "code"
	}
	values := url.Values{"client_id": {query.ClientID}, "redirect_uri": {query.RedirectUri}, "scope": {scopeValue(query.Scopes)}, "response_type": {stringValue(query.ResponseType)}}
	if query.State != "" {
		values.Set("state", query.State)
	}
	if query.EnableFbLogin != 0 {
		values.Set("enable_fb_login", stringValue(query.EnableFbLogin))
	}
	if query.ForceAuthentication != 0 {
		values.Set("force_authentication", stringValue(query.ForceAuthentication))
	}
	return &ResponseAuthorize{URL: AuthorizationBaseURL + "/oauth/authorize?" + values.Encode()}, nil
}

// ExchangeCode calls POST https://api.instagram.com/oauth/access_token.
func (s *oAuth) ExchangeCode(ctx context.Context, request *RequestExchangeCode) (*ResponseExchangeCode, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	body := request.Body
	if body.Code == "" {
		return nil, ErrMissingCode
	}
	if body.ClientID == "" {
		body.ClientID = s.client.clientID
	}
	if body.ClientSecret == "" {
		body.ClientSecret = s.client.clientSecret
	}
	if body.RedirectUri == "" {
		body.RedirectUri = s.client.redirectURL
	}
	if body.GrantType == nil {
		body.GrantType = "authorization_code"
	}
	values := url.Values{"client_id": {body.ClientID}, "client_secret": {body.ClientSecret}, "grant_type": {stringValue(body.GrantType)}, "redirect_uri": {body.RedirectUri}, "code": {body.Code}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodPost, OAuthBaseURL+"/oauth/access_token", values)
	if err != nil {
		return nil, err
	}
	response := new(ResponseExchangeCode)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// ExchangeLongLivedToken calls GET https://graph.instagram.com/access_token.
func (s *oAuth) ExchangeLongLivedToken(ctx context.Context, request *RequestExchangeLongLivedToken) (*ResponseExchangeLongLivedToken, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := request.Query
	if query.AccessToken == nil {
		return nil, ErrMissingAccessToken
	}
	if query.ClientSecret == "" {
		query.ClientSecret = s.client.clientSecret
	}
	if query.GrantType == nil {
		query.GrantType = "ig_exchange_token"
	}
	values := url.Values{"grant_type": {stringValue(query.GrantType)}, "client_secret": {query.ClientSecret}, "access_token": {stringValue(query.AccessToken)}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodGet, BaseURL+"/access_token?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseExchangeLongLivedToken)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// RefreshLongLivedToken calls GET https://graph.instagram.com/refresh_access_token.
func (s *oAuth) RefreshLongLivedToken(ctx context.Context, request *RequestRefreshLongLivedToken) (*ResponseRefreshLongLivedToken, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := request.Query
	if query.AccessToken == "" {
		return nil, ErrMissingAccessToken
	}
	if query.GrantType == nil {
		query.GrantType = "ig_refresh_token"
	}
	values := url.Values{"grant_type": {stringValue(query.GrantType)}, "access_token": {query.AccessToken}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodGet, BaseURL+"/refresh_access_token?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseRefreshLongLivedToken)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}
