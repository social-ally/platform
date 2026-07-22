package tiktok

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

const oauthAuthorizeURL = "https://www.tiktok.com/v2/auth/authorize/"

// OAuth provides access to oauth endpoints.
type OAuth struct {
	client *TikTokClient
}

// NewOAuth creates a OAuth endpoint group using client.
func NewOAuth(client *TikTokClient) *OAuth {
	return &OAuth{client: client}
}

type (
	RequestAuthorizeQuery struct {
		ClientKey           string  `json:"client_key"`
		Scopes              []Scope `json:"scope"`
		ResponseType        string  `json:"response_type"`
		RedirectUri         string  `json:"redirect_uri"`
		State               string  `json:"state"`
		CodeChallenge       *string `json:"code_challenge"`
		CodeChallengeMethod *string `json:"code_challenge_method"`
	}

	RequestAuthorize struct {
		Query RequestAuthorizeQuery `json:"query"`
	}

	ResponseAuthorizeSuccessRedirectQuery struct {
		Code   string `json:"code"`
		Scopes string `json:"scopes"`
		State  string `json:"state"`
	}

	ResponseAuthorizeErrorRedirectQuery struct {
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
		State            string `json:"state"`
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
		ClientKey    string  `json:"client_key"`
		ClientSecret string  `json:"client_secret"`
		Code         string  `json:"code"`
		GrantType    string  `json:"grant_type"`
		RedirectUri  string  `json:"redirect_uri"`
		CodeVerifier *string `json:"code_verifier"`
	}

	RequestExchangeCode struct {
		Headers RequestExchangeCodeHeaders `json:"headers"`
		Body    RequestExchangeCodeBody    `json:"body"`
	}

	ResponseExchangeCodeSuccess struct {
		AccessToken      string `json:"access_token"`
		ExpiresIn        int    `json:"expires_in"`
		OpenID           string `json:"open_id"`
		RefreshExpiresIn int    `json:"refresh_expires_in"`
		RefreshToken     string `json:"refresh_token"`
		Scope            string `json:"scope"`
		TokenType        string `json:"token_type"`
	}

	ResponseExchangeCodeError struct {
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
		LogID            string `json:"log_id"`
	}

	ResponseExchangeCode struct {
		Success ResponseExchangeCodeSuccess `json:"success"`
		Error   ResponseExchangeCodeError   `json:"error"`
	}

	RequestRefreshTokenHeaders struct {
		ContentType any `json:"Content-Type"`
	}

	RequestRefreshTokenBody struct {
		ClientKey    string `json:"client_key"`
		ClientSecret string `json:"client_secret"`
		GrantType    string `json:"grant_type"`
		RefreshToken string `json:"refresh_token"`
	}

	RequestRefreshToken struct {
		Headers RequestRefreshTokenHeaders `json:"headers"`
		Body    RequestRefreshTokenBody    `json:"body"`
	}

	ResponseRefreshTokenSuccess struct {
		AccessToken      string `json:"access_token"`
		ExpiresIn        int    `json:"expires_in"`
		OpenID           string `json:"open_id"`
		RefreshExpiresIn int    `json:"refresh_expires_in"`
		RefreshToken     string `json:"refresh_token"`
		Scope            string `json:"scope"`
		TokenType        string `json:"token_type"`
	}

	ResponseRefreshToken struct {
		Success ResponseRefreshTokenSuccess `json:"success"`
	}

	RequestRevokeTokenHeaders struct {
		ContentType any `json:"Content-Type"`
	}

	RequestRevokeTokenBody struct {
		ClientKey    string `json:"client_key"`
		ClientSecret string `json:"client_secret"`
		Token        string `json:"token"`
	}

	RequestRevokeToken struct {
		Headers RequestRevokeTokenHeaders `json:"headers"`
		Body    RequestRevokeTokenBody    `json:"body"`
	}

	ResponseRevokeTokenSuccess struct {
		HttpStatus int `json:"http_status"`
	}

	ResponseRevokeToken struct {
		Success ResponseRevokeTokenSuccess `json:"success"`
	}
)

// Authorize calls GET https://www.tiktok.com/v2/auth/authorize/.
func (s *OAuth) Authorize(ctx context.Context, request *RequestAuthorize) (*ResponseAuthorize, error) {
	if s.client == nil {
		return nil, ErrNilOAuthClient
	}
	query := RequestAuthorizeQuery{}
	if request != nil {
		query = request.Query
	}
	if query.ClientKey == "" {
		query.ClientKey = s.client.clientID
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
	if s.client.pkce && (query.CodeChallenge == nil || *query.CodeChallenge == "") {
		return nil, ErrMissingCodeChallenge
	}
	if query.CodeChallenge != nil && query.CodeChallengeMethod == nil {
		method := "S256"
		query.CodeChallengeMethod = &method
	}
	authorizeURL, err := url.Parse(oauthAuthorizeURL)
	if err != nil {
		return nil, err
	}
	values := authorizeURL.Query()
	values.Set("client_key", query.ClientKey)
	values.Set("scope", joinScopes(query.Scopes))
	values.Set("response_type", query.ResponseType)
	values.Set("redirect_uri", query.RedirectUri)
	if query.State != "" {
		values.Set("state", query.State)
	}
	if query.CodeChallenge != nil {
		values.Set("code_challenge", *query.CodeChallenge)
	}
	if query.CodeChallengeMethod != nil {
		values.Set("code_challenge_method", *query.CodeChallengeMethod)
	}
	authorizeURL.RawQuery = values.Encode()
	return &ResponseAuthorize{URL: authorizeURL.String()}, nil
}

// ExchangeCode calls POST https://open.tiktokapis.com/v2/oauth/token/.
func (s *OAuth) ExchangeCode(ctx context.Context, request *RequestExchangeCode) (*ResponseExchangeCode, error) {
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
	if s.client.pkce && (body.CodeVerifier == nil || *body.CodeVerifier == "") {
		return nil, ErrMissingCodeVerifier
	}
	if body.ClientKey == "" {
		body.ClientKey = s.client.clientID
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
	values := url.Values{"client_key": {body.ClientKey}, "client_secret": {body.ClientSecret}, "code": {body.Code}, "grant_type": {body.GrantType}, "redirect_uri": {body.RedirectUri}}
	if body.CodeVerifier != nil {
		values.Set("code_verifier", *body.CodeVerifier)
	}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodPost, BaseURL+"/v2/oauth/token/", values)
	if err != nil {
		return nil, err
	}
	var raw ResponseExchangeCodeSuccess
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseExchangeCode{Success: raw}, nil
}

// RefreshToken calls POST https://open.tiktokapis.com/v2/oauth/token/.
func (s *OAuth) RefreshToken(ctx context.Context, request *RequestRefreshToken) (*ResponseRefreshToken, error) {
	if s.client == nil {
		return nil, ErrNilOAuthClient
	}
	if request == nil {
		return nil, ErrNilOAuthRequest
	}
	body := request.Body
	if body.RefreshToken == "" {
		return nil, ErrMissingToken
	}
	if body.ClientKey == "" {
		body.ClientKey = s.client.clientID
	}
	if body.ClientSecret == "" {
		body.ClientSecret = s.client.clientSecret
	}
	if body.GrantType == "" {
		body.GrantType = "refresh_token"
	}
	values := url.Values{"client_key": {body.ClientKey}, "client_secret": {body.ClientSecret}, "grant_type": {body.GrantType}, "refresh_token": {body.RefreshToken}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodPost, BaseURL+"/v2/oauth/token/", values)
	if err != nil {
		return nil, err
	}
	var raw ResponseRefreshTokenSuccess
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseRefreshToken{Success: raw}, nil
}

// RevokeToken calls POST https://open.tiktokapis.com/v2/oauth/revoke/.
func (s *OAuth) RevokeToken(ctx context.Context, request *RequestRevokeToken) (*ResponseRevokeToken, error) {
	if s.client == nil {
		return nil, ErrNilOAuthClient
	}
	if request == nil {
		return nil, ErrNilOAuthRequest
	}
	body := request.Body
	if body.Token == "" {
		return nil, ErrMissingToken
	}
	if body.ClientKey == "" {
		body.ClientKey = s.client.clientID
	}
	if body.ClientSecret == "" {
		body.ClientSecret = s.client.clientSecret
	}
	values := url.Values{"client_key": {body.ClientKey}, "client_secret": {body.ClientSecret}, "token": {body.Token}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodPost, BaseURL+"/v2/oauth/revoke/", values)
	if err != nil {
		return nil, err
	}
	if err := s.client.Do(httpRequest, nil); err != nil {
		return nil, err
	}
	return &ResponseRevokeToken{Success: ResponseRevokeTokenSuccess{HttpStatus: http.StatusOK}}, nil
}

func joinScopes(scopes []Scope) string {
	values := make([]string, len(scopes))
	for index, scope := range scopes {
		values[index] = string(scope)
	}
	return strings.Join(values, ",")
}
