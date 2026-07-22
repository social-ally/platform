package tiktok

// Scope is an OAuth permission supported by TikTok.
type Scope string

const (
	ScopeUserInfoBasic           Scope = "user.info.basic"
	ScopeUserInfoProfile         Scope = "user.info.profile"
	ScopeUserInfoStats           Scope = "user.info.stats"
	ScopeVideoList               Scope = "video.list"
	ScopeVideoPublish            Scope = "video.publish"
	ScopeVideoPublishVideoUpload Scope = "video.publish|video.upload"
	ScopeVideoUpload             Scope = "video.upload"
)
