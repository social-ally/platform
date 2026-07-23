// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package tiktok

import (
	"context"
	"net/http"
	"net/url"
)

// Users provides access to users endpoints.
type users struct {
	client *TikTokClient
}

// NewUsers creates a Users endpoint group using client.
func NewUsers(client *TikTokClient) *users {
	return &users{client: client}
}

type (
	RequestQueryUserInfoQuery struct {
		Fields any `json:"fields"`
	}

	RequestQueryUserInfo struct {
		Query RequestQueryUserInfoQuery `json:"query"`
	}

	ResponseQueryUserInfoSuccessDataUser struct {
		OpenID          string `json:"open_id"`
		UnionID         string `json:"union_id"`
		AvatarURL       string `json:"avatar_url"`
		DisplayName     string `json:"display_name"`
		BioDescription  string `json:"bio_description"`
		ProfileDeepLink string `json:"profile_deep_link"`
		IsVerified      bool   `json:"is_verified"`
		FollowerCount   int    `json:"follower_count"`
		FollowingCount  int    `json:"following_count"`
		LikesCount      int    `json:"likes_count"`
		VideoCount      int    `json:"video_count"`
	}

	ResponseQueryUserInfoSuccessData struct {
		User ResponseQueryUserInfoSuccessDataUser `json:"user"`
	}

	ResponseQueryUserInfoSuccessError struct {
		Code    any    `json:"code"`
		Message any    `json:"message"`
		LogID   string `json:"log_id"`
	}

	ResponseQueryUserInfoSuccess struct {
		Data  ResponseQueryUserInfoSuccessData  `json:"data"`
		Error ResponseQueryUserInfoSuccessError `json:"error"`
	}

	ResponseQueryUserInfo struct {
		Success ResponseQueryUserInfoSuccess `json:"success"`
	}
)

// QueryUserInfo calls GET https://open.tiktokapis.com/v2/user/info/.
func (s *users) QueryUserInfo(ctx context.Context, request *RequestQueryUserInfo) (*ResponseQueryUserInfo, error) {
	q := url.Values{}
	if request != nil {
		addOptionalQuery(q, "fields", request.Query.Fields)
	}
	raw := BaseURL + "/v2/user/info/"
	if q.Encode() != "" {
		raw += "?" + q.Encode()
	}
	r, e := s.client.authenticatedRequest(ctx, http.MethodGet, raw, nil)
	if e != nil {
		return nil, e
	}
	var data ResponseQueryUserInfoSuccess
	if e = s.client.Do(r, &data); e != nil {
		return nil, e
	}
	return &ResponseQueryUserInfo{Success: data}, nil
}
