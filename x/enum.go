package x

// ReplySetting controls who may reply to a post.
type ReplySetting string

const (
	ReplySettingFollowing      ReplySetting = "following"
	ReplySettingMentionedUsers ReplySetting = "mentionedUsers"
	ReplySettingSubscribers    ReplySetting = "subscribers"
	ReplySettingVerified       ReplySetting = "verified"
)
