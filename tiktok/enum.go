package tiktok

// CodeChallengeMethod identifies the PKCE challenge transformation.
type CodeChallengeMethod string

const CodeChallengeMethodS256 CodeChallengeMethod = "S256"

// PrivacyLevel controls visibility of a TikTok post.
type PrivacyLevel string

const (
	PrivacyLevelPublicToEveryone    PrivacyLevel = "PUBLIC_TO_EVERYONE"
	PrivacyLevelMutualFollowFriends PrivacyLevel = "MUTUAL_FOLLOW_FRIENDS"
	PrivacyLevelSelfOnly            PrivacyLevel = "SELF_ONLY"
)

// Source identifies how TikTok receives the media.
type Source string

const (
	SourceFileUpload  Source = "FILE_UPLOAD"
	SourcePullFromURL Source = "PULL_FROM_URL"
)

// PublishStatus is the processing state of a TikTok publish request.
type PublishStatus string

const (
	PublishStatusProcessingUpload   PublishStatus = "PROCESSING_UPLOAD"
	PublishStatusProcessingDownload PublishStatus = "PROCESSING_DOWNLOAD"
	PublishStatusSendToUserInbox    PublishStatus = "SEND_TO_USER_INBOX"
	PublishStatusComplete           PublishStatus = "PUBLISH_COMPLETE"
	PublishStatusFailed             PublishStatus = "FAILED"
)
