package youtube

// Scope is an OAuth permission supported by youtube.
type Scope string

const (
	// ScopeEmail grants the email permission.
	ScopeEmail Scope = "email"
	// ScopeYoutube grants the https://www.googleapis.com/auth/youtube permission.
	ScopeYoutube Scope = "https://www.googleapis.com/auth/youtube"
	// ScopeYoutubeForceSSL grants the https://www.googleapis.com/auth/youtube.force-ssl permission.
	ScopeYoutubeForceSSL Scope = "https://www.googleapis.com/auth/youtube.force-ssl"
	// ScopeYoutubeReadonly grants the https://www.googleapis.com/auth/youtube.readonly permission.
	ScopeYoutubeReadonly Scope = "https://www.googleapis.com/auth/youtube.readonly"
	// ScopeYoutubeUpload grants the https://www.googleapis.com/auth/youtube.upload permission.
	ScopeYoutubeUpload Scope = "https://www.googleapis.com/auth/youtube.upload"
	// ScopeYTAnalyticsMonetaryReadonly grants the https://www.googleapis.com/auth/yt-analytics-monetary.readonly permission.
	ScopeYTAnalyticsMonetaryReadonly Scope = "https://www.googleapis.com/auth/yt-analytics-monetary.readonly"
	// ScopeYTAnalyticsReadonly grants the https://www.googleapis.com/auth/yt-analytics.readonly permission.
	ScopeYTAnalyticsReadonly Scope = "https://www.googleapis.com/auth/yt-analytics.readonly"
	// ScopeOpenid grants the openid permission.
	ScopeOpenid Scope = "openid"
	// ScopeProfile grants the profile permission.
	ScopeProfile Scope = "profile"
)
