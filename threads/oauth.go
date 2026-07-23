package threads

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

const oauthAuthorizeURL = "https://threads.net/oauth/authorize"

// OAuth provides access to oauth endpoints.
type oAuth struct {
	client *ThreadsClient
}

// NewOAuth creates a OAuth endpoint group using client.
func NewOAuth(client *ThreadsClient) *oAuth {
	return &oAuth{client: client}
}

type (
	RequestAuthorizeQuery struct {
		ClientID     string  `json:"client_id"`
		RedirectUri  string  `json:"redirect_uri"`
		Scopes       []Scope `json:"scope"`
		ResponseType string  `json:"response_type"`
		State        string  `json:"state"`
	}

	RequestAuthorize struct {
		Query RequestAuthorizeQuery `json:"query"`
	}

	ResponseAuthorizeSuccessRedirectQuery struct {
		Code  string `json:"code"`
		State string `json:"state"`
	}

	ResponseAuthorize struct {
		URL                  string                                `json:"url"`
		SuccessRedirectQuery ResponseAuthorizeSuccessRedirectQuery `json:"success_redirect_query"`
	}

	RequestExchangeCodeHeaders struct {
		ContentType string `json:"Content-Type"`
	}

	RequestExchangeCodeBody struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		GrantType    string `json:"grant_type"`
		RedirectUri  string `json:"redirect_uri"`
		Code         string `json:"code"`
	}

	RequestExchangeCode struct {
		Headers RequestExchangeCodeHeaders `json:"headers"`
		Body    RequestExchangeCodeBody    `json:"body"`
	}

	ResponseExchangeCodeSuccess struct {
		AccessToken string `json:"access_token"`
		UserID      int    `json:"user_id"`
	}

	ResponseExchangeCode struct {
		Success ResponseExchangeCodeSuccess `json:"success"`
	}

	RequestExchangeLongLivedTokenQuery struct {
		GrantType    string `json:"grant_type"`
		ClientSecret string `json:"client_secret"`
		AccessToken  string `json:"access_token"`
	}

	RequestExchangeLongLivedToken struct {
		Query RequestExchangeLongLivedTokenQuery `json:"query"`
	}

	ResponseExchangeLongLivedTokenSuccess struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	ResponseExchangeLongLivedToken struct {
		Success ResponseExchangeLongLivedTokenSuccess `json:"success"`
	}

	RequestRefreshLongLivedTokenQuery struct {
		GrantType   string `json:"grant_type"`
		AccessToken string `json:"access_token"`
	}

	RequestRefreshLongLivedToken struct {
		Query RequestRefreshLongLivedTokenQuery `json:"query"`
	}

	ResponseRefreshLongLivedTokenSuccess struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	ResponseRefreshLongLivedToken struct {
		Success ResponseRefreshLongLivedTokenSuccess `json:"success"`
	}
)

// Authorize calls GET https://threads.net/oauth/authorize.
func (s *oAuth) Authorize(ctx context.Context, request *RequestAuthorize) (*ResponseAuthorize, error) {
	if s.client == nil {
		return nil, ErrNilOAuthClient
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
	authorizeURL, err := url.Parse(oauthAuthorizeURL)
	if err != nil {
		return nil, err
	}
	values := authorizeURL.Query()
	values.Set("client_id", query.ClientID)
	values.Set("redirect_uri", query.RedirectUri)
	values.Set("scope", joinScopes(query.Scopes))
	values.Set("response_type", query.ResponseType)
	if query.State != "" {
		values.Set("state", query.State)
	}
	authorizeURL.RawQuery = values.Encode()
	return &ResponseAuthorize{URL: authorizeURL.String()}, nil
}

// ExchangeCode calls POST https://graph.threads.net/oauth/access_token.
func (s *oAuth) ExchangeCode(ctx context.Context, request *RequestExchangeCode) (*ResponseExchangeCode, error) {
	if s.client == nil {
		return nil, ErrNilOAuthClient
	}
	if request == nil {
		return nil, ErrNilOAuthRequest
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
	if body.GrantType == "" {
		body.GrantType = "authorization_code"
	}
	values := url.Values{"client_id": {body.ClientID}, "client_secret": {body.ClientSecret}, "grant_type": {body.GrantType}, "redirect_uri": {body.RedirectUri}, "code": {body.Code}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodPost, BaseURL+"/oauth/access_token", values)
	if err != nil {
		return nil, err
	}
	var raw ResponseExchangeCodeSuccess
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseExchangeCode{Success: raw}, nil
}

// ExchangeLongLivedToken calls GET https://graph.threads.net/access_token.
func (s *oAuth) ExchangeLongLivedToken(ctx context.Context, request *RequestExchangeLongLivedToken) (*ResponseExchangeLongLivedToken, error) {
	if s.client == nil {
		return nil, ErrNilOAuthClient
	}
	if request == nil {
		return nil, ErrNilOAuthRequest
	}
	query := request.Query
	if query.AccessToken == "" {
		return nil, ErrMissingToken
	}
	if query.ClientSecret == "" {
		query.ClientSecret = s.client.clientSecret
	}
	if query.GrantType == "" {
		query.GrantType = "th_exchange_token"
	}
	values := url.Values{"grant_type": {query.GrantType}, "client_secret": {query.ClientSecret}, "access_token": {query.AccessToken}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodGet, BaseURL+"/access_token?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	var raw ResponseExchangeLongLivedTokenSuccess
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseExchangeLongLivedToken{Success: raw}, nil
}

// RefreshLongLivedToken calls GET https://graph.threads.net/refresh_access_token.
func (s *oAuth) RefreshLongLivedToken(ctx context.Context, request *RequestRefreshLongLivedToken) (*ResponseRefreshLongLivedToken, error) {
	if s.client == nil {
		return nil, ErrNilOAuthClient
	}
	if request == nil {
		return nil, ErrNilOAuthRequest
	}
	query := request.Query
	if query.AccessToken == "" {
		return nil, ErrMissingToken
	}
	if query.GrantType == "" {
		query.GrantType = "th_refresh_token"
	}
	values := url.Values{"grant_type": {query.GrantType}, "access_token": {query.AccessToken}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodGet, BaseURL+"/refresh_access_token?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	var raw ResponseRefreshLongLivedTokenSuccess
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseRefreshLongLivedToken{Success: raw}, nil
}

func joinScopes(scopes []Scope) string {
	values := make([]string, len(scopes))
	for index, scope := range scopes {
		values[index] = string(scope)
	}
	return strings.Join(values, ",")
}
