package x

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

const (
	oauthAuthorizeURL = "https://x.com/i/oauth2/authorize"
	oauthTokenURL     = BaseURL + "/oauth2/token"
	oauthRevokeURL    = BaseURL + "/oauth2/revoke"
)

// oAuth provides access to OAuth endpoints.
type oAuth struct {
	client *XClient
}

// NewOAuth creates an OAuth endpoint group using client.
func NewOAuth(client *XClient) *oAuth {
	return &oAuth{client: client}
}

type (
	RequestAuthorizeQuery struct {
		ResponseType        string  `json:"response_type"`
		ClientID            string  `json:"client_id"`
		RedirectUri         string  `json:"redirect_uri"`
		Scopes              []Scope `json:"scope"`
		State               string  `json:"state"`
		CodeChallenge       string  `json:"code_challenge"`
		CodeChallengeMethod string  `json:"code_challenge_method"`
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
		GrantType    string `json:"grant_type"`
		Code         string `json:"code"`
		RedirectUri  string `json:"redirect_uri"`
		ClientID     string `json:"client_id"`
		CodeVerifier string `json:"code_verifier"`
	}

	RequestExchangeCode struct {
		Headers RequestExchangeCodeHeaders `json:"headers"`
		Body    RequestExchangeCodeBody    `json:"body"`
	}

	ResponseExchangeCodeSuccess struct {
		TokenType    string  `json:"token_type"`
		ExpiresIn    int     `json:"expires_in"`
		AccessToken  string  `json:"access_token"`
		Scope        string  `json:"scope"`
		RefreshToken *string `json:"refresh_token"`
	}

	ResponseExchangeCodeError struct {
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}

	ResponseExchangeCode struct {
		Success ResponseExchangeCodeSuccess `json:"success"`
		Error   ResponseExchangeCodeError   `json:"error"`
	}

	RequestRefreshTokenHeaders struct {
		ContentType any `json:"Content-Type"`
	}

	RequestRefreshTokenBody struct {
		GrantType    string `json:"grant_type"`
		RefreshToken string `json:"refresh_token"`
		ClientID     string `json:"client_id"`
	}

	RequestRefreshToken struct {
		Headers RequestRefreshTokenHeaders `json:"headers"`
		Body    RequestRefreshTokenBody    `json:"body"`
	}

	ResponseRefreshTokenSuccess struct {
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		AccessToken  string `json:"access_token"`
		Scope        string `json:"scope"`
		RefreshToken string `json:"refresh_token"`
	}

	ResponseRefreshToken struct {
		Success ResponseRefreshTokenSuccess `json:"success"`
	}

	RequestRevokeTokenHeaders struct {
		ContentType any `json:"Content-Type"`
	}

	RequestRevokeTokenBody struct {
		Token    string `json:"token"`
		ClientID string `json:"client_id"`
	}

	RequestRevokeToken struct {
		Headers RequestRevokeTokenHeaders `json:"headers"`
		Body    RequestRevokeTokenBody    `json:"body"`
	}

	ResponseRevokeTokenSuccess struct {
		Revoked bool `json:"revoked"`
	}

	ResponseRevokeToken struct {
		Success ResponseRevokeTokenSuccess `json:"success"`
	}

	oauthTokenResponse struct {
		TokenType    string  `json:"token_type"`
		ExpiresIn    int     `json:"expires_in"`
		AccessToken  string  `json:"access_token"`
		Scope        string  `json:"scope"`
		RefreshToken *string `json:"refresh_token"`
	}
)

// Authorize calls GET https://x.com/i/oauth2/authorize.
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
	if query.CodeChallengeMethod == "" {
		query.CodeChallengeMethod = "S256"
	}
	if query.CodeChallenge == "" {
		return nil, ErrMissingCodeChallenge
	}

	authorizeURL, err := url.Parse(oauthAuthorizeURL)
	if err != nil {
		return nil, err
	}
	values := authorizeURL.Query()
	values.Set("response_type", query.ResponseType)
	values.Set("client_id", query.ClientID)
	values.Set("redirect_uri", query.RedirectUri)
	values.Set("scope", joinScopes(query.Scopes))
	values.Set("state", query.State)
	values.Set("code_challenge", query.CodeChallenge)
	values.Set("code_challenge_method", query.CodeChallengeMethod)
	authorizeURL.RawQuery = values.Encode()

	return &ResponseAuthorize{URL: authorizeURL.String()}, nil
}

// ExchangeCode calls POST https://api.x.com/2/oauth2/token.
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
	if body.CodeVerifier == "" {
		return nil, ErrMissingCodeVerifier
	}
	if body.ClientID == "" {
		body.ClientID = s.client.clientID
	}
	if body.RedirectUri == "" {
		body.RedirectUri = s.client.redirectURL
	}
	if body.GrantType == "" {
		body.GrantType = "authorization_code"
	}

	values := url.Values{
		"grant_type":    {body.GrantType},
		"code":          {body.Code},
		"redirect_uri":  {body.RedirectUri},
		"code_verifier": {body.CodeVerifier},
	}
	if !s.client.confidential {
		values.Set("client_id", body.ClientID)
	}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodPost, oauthTokenURL, values)
	if err != nil {
		return nil, err
	}
	if s.client.confidential {
		httpRequest.SetBasicAuth(body.ClientID, s.client.clientSecret)
	}
	var token oauthTokenResponse
	if err := s.client.Do(httpRequest, &token); err != nil {
		return nil, err
	}
	return &ResponseExchangeCode{Success: ResponseExchangeCodeSuccess{TokenType: token.TokenType, ExpiresIn: token.ExpiresIn, AccessToken: token.AccessToken, Scope: token.Scope, RefreshToken: token.RefreshToken}}, nil
}

// RefreshToken calls POST https://api.x.com/2/oauth2/token.
func (s *oAuth) RefreshToken(ctx context.Context, request *RequestRefreshToken) (*ResponseRefreshToken, error) {
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
	if body.ClientID == "" {
		body.ClientID = s.client.clientID
	}
	if body.GrantType == "" {
		body.GrantType = "refresh_token"
	}

	values := url.Values{"grant_type": {body.GrantType}, "refresh_token": {body.RefreshToken}}
	if !s.client.confidential {
		values.Set("client_id", body.ClientID)
	}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodPost, oauthTokenURL, values)
	if err != nil {
		return nil, err
	}
	if s.client.confidential {
		httpRequest.SetBasicAuth(body.ClientID, s.client.clientSecret)
	}
	var token oauthTokenResponse
	if err := s.client.Do(httpRequest, &token); err != nil {
		return nil, err
	}
	refreshToken := ""
	if token.RefreshToken != nil {
		refreshToken = *token.RefreshToken
	}
	return &ResponseRefreshToken{Success: ResponseRefreshTokenSuccess{TokenType: token.TokenType, ExpiresIn: token.ExpiresIn, AccessToken: token.AccessToken, Scope: token.Scope, RefreshToken: refreshToken}}, nil
}

// RevokeToken calls POST https://api.x.com/2/oauth2/revoke.
func (s *oAuth) RevokeToken(ctx context.Context, request *RequestRevokeToken) (*ResponseRevokeToken, error) {
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
	if body.ClientID == "" {
		body.ClientID = s.client.clientID
	}

	values := url.Values{"token": {body.Token}, "client_id": {body.ClientID}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodPost, oauthRevokeURL, values)
	if err != nil {
		return nil, err
	}
	if err := s.client.Do(httpRequest, nil); err != nil {
		return nil, err
	}
	return &ResponseRevokeToken{Success: ResponseRevokeTokenSuccess{Revoked: true}}, nil
}

func joinScopes(scopes []Scope) string {
	values := make([]string, len(scopes))
	for index, scope := range scopes {
		values[index] = string(scope)
	}
	return strings.Join(values, " ")
}
