// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package youtube

import (
	"context"
	"net/http"
	"net/url"
)

// OAuth provides access to oauth endpoints.
type oAuth struct {
	client *YouTubeClient
}

// NewOAuth creates a OAuth endpoint group using client.
func NewOAuth(client *YouTubeClient) *oAuth {
	return &oAuth{client: client}
}

type (
	RequestAuthorizeQuery struct {
		ClientID             string  `json:"client_id"`
		RedirectUri          string  `json:"redirect_uri"`
		ResponseType         string  `json:"response_type"`
		Scopes               []Scope `json:"scope"`
		State                string  `json:"state"`
		AccessType           string  `json:"access_type"`
		IncludeGrantedScopes *bool   `json:"include_granted_scopes"`
		Prompt               string  `json:"prompt"`
	}

	RequestAuthorize struct {
		Query RequestAuthorizeQuery `json:"query"`
	}

	ResponseAuthorizeSuccessRedirectQuery struct {
		Code  string `json:"code"`
		State string `json:"state"`
	}

	ResponseAuthorizeErrorRedirectQuery struct {
		Error string `json:"error"`
	}

	ResponseAuthorize struct {
		URL                  string                                `json:"url"`
		SuccessRedirectQuery ResponseAuthorizeSuccessRedirectQuery `json:"success_redirect_query"`
		ErrorRedirectQuery   ResponseAuthorizeErrorRedirectQuery   `json:"error_redirect_query"`
	}

	RequestExchangeCodeHeaders struct {
		ContentType string `json:"Content-Type"`
	}

	RequestExchangeCodeBody struct {
		ClientID     string  `json:"client_id"`
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
		AccessToken  string  `json:"access_token"`
		ExpiresIn    int     `json:"expires_in"`
		RefreshToken *string `json:"refresh_token"`
		Scope        string  `json:"scope"`
		TokenType    string  `json:"token_type"`
		IDToken      *string `json:"id_token"`
	}

	ResponseExchangeCode struct {
		Success ResponseExchangeCodeSuccess `json:"success"`
	}

	RequestRefreshTokenHeaders struct {
		ContentType string `json:"Content-Type"`
	}

	RequestRefreshTokenBody struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RefreshToken string `json:"refresh_token"`
		GrantType    string `json:"grant_type"`
	}

	RequestRefreshToken struct {
		Headers RequestRefreshTokenHeaders `json:"headers"`
		Body    RequestRefreshTokenBody    `json:"body"`
	}

	ResponseRefreshTokenSuccess struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
	}

	ResponseRefreshToken struct {
		Success ResponseRefreshTokenSuccess `json:"success"`
	}

	RequestRevokeTokenHeaders struct {
		ContentType string `json:"Content-Type"`
	}

	RequestRevokeTokenBody struct {
		Token string `json:"token"`
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

// Authorize calls GET https://accounts.google.com/o/oauth2/v2/auth.
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
	if query.AccessType != "" {
		values.Set("access_type", stringValue(query.AccessType))
	}
	if query.IncludeGrantedScopes != nil {
		values.Set("include_granted_scopes", stringValue(query.IncludeGrantedScopes))
	}
	if query.Prompt != "" {
		values.Set("prompt", stringValue(query.Prompt))
	}
	return &ResponseAuthorize{URL: AuthorizationBaseURL + "/auth?" + values.Encode()}, nil
}

// ExchangeCode calls POST https://oauth2.googleapis.com/token.
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
	if body.GrantType == "" {
		body.GrantType = "authorization_code"
	}
	values := url.Values{"client_id": {body.ClientID}, "client_secret": {body.ClientSecret}, "code": {body.Code}, "grant_type": {stringValue(body.GrantType)}, "redirect_uri": {body.RedirectUri}}
	if body.CodeVerifier != nil {
		values.Set("code_verifier", *body.CodeVerifier)
	}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodPost, OAuthBaseURL+"/token", values)
	if err != nil {
		return nil, err
	}
	response := new(ResponseExchangeCode)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// RefreshToken calls POST https://oauth2.googleapis.com/token.
func (s *oAuth) RefreshToken(ctx context.Context, request *RequestRefreshToken) (*ResponseRefreshToken, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	body := request.Body
	if body.RefreshToken == "" {
		return nil, ErrMissingAccessToken
	}
	if body.ClientID == "" {
		body.ClientID = s.client.clientID
	}
	if body.ClientSecret == "" {
		body.ClientSecret = s.client.clientSecret
	}
	if body.GrantType == "" {
		body.GrantType = "refresh_token"
	}
	values := url.Values{"client_id": {body.ClientID}, "client_secret": {body.ClientSecret}, "refresh_token": {body.RefreshToken}, "grant_type": {stringValue(body.GrantType)}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodPost, OAuthBaseURL+"/token", values)
	if err != nil {
		return nil, err
	}
	response := new(ResponseRefreshToken)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// RevokeToken calls POST https://oauth2.googleapis.com/revoke.
func (s *oAuth) RevokeToken(ctx context.Context, request *RequestRevokeToken) (*ResponseRevokeToken, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Body.Token == "" {
		return nil, ErrMissingAccessToken
	}
	values := url.Values{"token": {request.Body.Token}}
	httpRequest, err := s.client.NewRequest(ctx, http.MethodPost, OAuthBaseURL+"/revoke", values)
	if err != nil {
		return nil, err
	}
	if err := s.client.Do(httpRequest, nil); err != nil {
		return nil, err
	}
	return &ResponseRevokeToken{Success: ResponseRevokeTokenSuccess{HttpStatus: http.StatusOK}}, nil
}
