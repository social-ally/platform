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

// UserField selects a field returned by the TikTok user-info endpoint.
type UserField string

const (
	UserFieldOpenID          UserField = "open_id"
	UserFieldUnionID         UserField = "union_id"
	UserFieldAvatarURL       UserField = "avatar_url"
	UserFieldAvatarURL100    UserField = "avatar_url_100"
	UserFieldAvatarLargeURL  UserField = "avatar_large_url"
	UserFieldDisplayName     UserField = "display_name"
	UserFieldBioDescription  UserField = "bio_description"
	UserFieldProfileDeepLink UserField = "profile_deep_link"
	UserFieldVerified        UserField = "is_verified"
	UserFieldFollowerCount   UserField = "follower_count"
	UserFieldFollowingCount  UserField = "following_count"
	UserFieldLikesCount      UserField = "likes_count"
	UserFieldVideoCount      UserField = "video_count"
)

// VideoField selects a field returned by TikTok video endpoints.
type VideoField string

const (
	VideoFieldID            VideoField = "id"
	VideoFieldCreateTime    VideoField = "create_time"
	VideoFieldCoverImageURL VideoField = "cover_image_url"
	VideoFieldShareURL      VideoField = "share_url"
	VideoFieldDescription   VideoField = "video_description"
	VideoFieldDuration      VideoField = "duration"
	VideoFieldHeight        VideoField = "height"
	VideoFieldWidth         VideoField = "width"
	VideoFieldTitle         VideoField = "title"
	VideoFieldEmbedHTML     VideoField = "embed_html"
	VideoFieldEmbedLink     VideoField = "embed_link"
	VideoFieldLikeCount     VideoField = "like_count"
	VideoFieldCommentCount  VideoField = "comment_count"
	VideoFieldShareCount    VideoField = "share_count"
	VideoFieldViewCount     VideoField = "view_count"
)

// VideoContentType is the content type accepted by TikTok video uploads.
type VideoContentType string

const VideoContentTypeMP4 VideoContentType = "video/mp4"

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
