package x

// ReplySetting controls who may reply to a post.
type ReplySetting string

const (
	ReplySettingFollowing      ReplySetting = "following"
	ReplySettingMentionedUsers ReplySetting = "mentionedUsers"
	ReplySettingSubscribers    ReplySetting = "subscribers"
	ReplySettingVerified       ReplySetting = "verified"
)

// ContentType identifies a request content type accepted by an X endpoint.
type ContentType string

const (
	ContentTypeJSON ContentType = "application/json"
)

// TokenTypeHint identifies the token type being revoked.
type TokenTypeHint string

const (
	TokenTypeHintAccessToken  TokenTypeHint = "access_token"
	TokenTypeHintRefreshToken TokenTypeHint = "refresh_token"
)

// UserField selects an optional field on an X user resource.
type UserField string

const (
	UserFieldAffiliation       UserField = "affiliation"
	UserFieldConfirmedEmail    UserField = "confirmed_email"
	UserFieldConnectionStatus  UserField = "connection_status"
	UserFieldCreatedAt         UserField = "created_at"
	UserFieldDescription       UserField = "description"
	UserFieldEntities          UserField = "entities"
	UserFieldID                UserField = "id"
	UserFieldIdentityVerified  UserField = "is_identity_verified"
	UserFieldLocation          UserField = "location"
	UserFieldMostRecentTweetID UserField = "most_recent_tweet_id"
	UserFieldName              UserField = "name"
	UserFieldParody            UserField = "parody"
	UserFieldPinnedTweetID     UserField = "pinned_tweet_id"
	UserFieldProfileBannerURL  UserField = "profile_banner_url"
	UserFieldProfileImageURL   UserField = "profile_image_url"
	UserFieldProtected         UserField = "protected"
	UserFieldPublicMetrics     UserField = "public_metrics"
	UserFieldReceivesYourDM    UserField = "receives_your_dm"
	UserFieldSubscription      UserField = "subscription"
	UserFieldSubscriptionType  UserField = "subscription_type"
	UserFieldURL               UserField = "url"
	UserFieldUsername          UserField = "username"
	UserFieldVerified          UserField = "verified"
	UserFieldVerifiedType      UserField = "verified_type"
	UserFieldWithheld          UserField = "withheld"
)

// TweetField selects an optional field on an X post resource.
type TweetField string

const (
	TweetFieldArticle                TweetField = "article"
	TweetFieldAttachments            TweetField = "attachments"
	TweetFieldAuthorID               TweetField = "author_id"
	TweetFieldCardURI                TweetField = "card_uri"
	TweetFieldCommunityID            TweetField = "community_id"
	TweetFieldContextAnnotations     TweetField = "context_annotations"
	TweetFieldConversationID         TweetField = "conversation_id"
	TweetFieldCreatedAt              TweetField = "created_at"
	TweetFieldDisplayTextRange       TweetField = "display_text_range"
	TweetFieldEditControls           TweetField = "edit_controls"
	TweetFieldEditHistoryTweetIDs    TweetField = "edit_history_tweet_ids"
	TweetFieldEntities               TweetField = "entities"
	TweetFieldGeo                    TweetField = "geo"
	TweetFieldID                     TweetField = "id"
	TweetFieldInReplyToUserID        TweetField = "in_reply_to_user_id"
	TweetFieldLanguage               TweetField = "lang"
	TweetFieldMatchedMediaNotes      TweetField = "matched_media_notes"
	TweetFieldMediaMetadata          TweetField = "media_metadata"
	TweetFieldNonPublicMetrics       TweetField = "non_public_metrics"
	TweetFieldNoteRequestSuggestions TweetField = "note_request_suggestions"
	TweetFieldNoteTweet              TweetField = "note_tweet"
	TweetFieldOrganicMetrics         TweetField = "organic_metrics"
	TweetFieldPossiblySensitive      TweetField = "possibly_sensitive"
	TweetFieldPromotedMetrics        TweetField = "promoted_metrics"
	TweetFieldPublicMetrics          TweetField = "public_metrics"
	TweetFieldReferencedTweets       TweetField = "referenced_tweets"
	TweetFieldReplySettings          TweetField = "reply_settings"
	TweetFieldSource                 TweetField = "source"
	TweetFieldText                   TweetField = "text"
	TweetFieldWithheld               TweetField = "withheld"
)

// Expansion selects related resources to include with an X post response.
type Expansion string

const (
	ExpansionArticleCoverMedia           Expansion = "article.cover_media"
	ExpansionArticleMediaEntities        Expansion = "article.media_entities"
	ExpansionAttachmentsMediaKeys        Expansion = "attachments.media_keys"
	ExpansionAttachmentsMediaSourceTweet Expansion = "attachments.media_source_tweet"
	ExpansionAttachmentsPollIDs          Expansion = "attachments.poll_ids"
	ExpansionAuthorID                    Expansion = "author_id"
	ExpansionEditHistoryTweetIDs         Expansion = "edit_history_tweet_ids"
	ExpansionEntitiesMentionsUsername    Expansion = "entities.mentions.username"
	ExpansionGeoPlaceID                  Expansion = "geo.place_id"
	ExpansionInReplyToUserID             Expansion = "in_reply_to_user_id"
	ExpansionReferencedTweetID           Expansion = "referenced_tweets.id"
	ExpansionReferencedTweetAuthorID     Expansion = "referenced_tweets.id.author_id"
)

// MediaType is the type of media returned by X.
type MediaType string

const (
	MediaTypeAnimatedGIF MediaType = "animated_gif"
	MediaTypePhoto       MediaType = "photo"
	MediaTypeVideo       MediaType = "video"
)

// TweetReferenceType describes a relationship to another post.
type TweetReferenceType string

const (
	TweetReferenceRetweeted TweetReferenceType = "retweeted"
	TweetReferenceQuoted    TweetReferenceType = "quoted"
	TweetReferenceRepliedTo TweetReferenceType = "replied_to"
)
