package instagram

// AccountType identifies the type of an Instagram account.
type AccountType string

const (
	AccountTypeBusiness     AccountType = "BUSINESS"
	AccountTypeMediaCreator AccountType = "MEDIA_CREATOR"
)

// UserField selects a field returned by the Instagram user endpoint.
type UserField string

const (
	UserFieldID                UserField = "id"
	UserFieldUserID            UserField = "user_id"
	UserFieldUsername          UserField = "username"
	UserFieldName              UserField = "name"
	UserFieldAccountType       UserField = "account_type"
	UserFieldProfilePictureURL UserField = "profile_picture_url"
	UserFieldBiography         UserField = "biography"
	UserFieldWebsite           UserField = "website"
	UserFieldFollowersCount    UserField = "followers_count"
	UserFieldFollowsCount      UserField = "follows_count"
	UserFieldMediaCount        UserField = "media_count"
)

// MediaField selects a field returned by an Instagram media endpoint.
type MediaField string

const (
	MediaFieldID             MediaField = "id"
	MediaFieldCaption        MediaField = "caption"
	MediaFieldCommentsCount  MediaField = "comments_count"
	MediaFieldLikeCount      MediaField = "like_count"
	MediaFieldProductType    MediaField = "media_product_type"
	MediaFieldType           MediaField = "media_type"
	MediaFieldURL            MediaField = "media_url"
	MediaFieldOwner          MediaField = "owner"
	MediaFieldPermalink      MediaField = "permalink"
	MediaFieldShortcode      MediaField = "shortcode"
	MediaFieldThumbnailURL   MediaField = "thumbnail_url"
	MediaFieldTimestamp      MediaField = "timestamp"
	MediaFieldUsername       MediaField = "username"
	MediaFieldChildren       MediaField = "children"
	MediaFieldCommentEnabled MediaField = "is_comment_enabled"
)

// MediaType identifies the type of media returned by Instagram.
type MediaType string

const (
	MediaTypeImage         MediaType = "IMAGE"
	MediaTypeVideo         MediaType = "VIDEO"
	MediaTypeCarouselAlbum MediaType = "CAROUSEL_ALBUM"
)

// PublishingMediaType identifies the type of media to create.
type PublishingMediaType string

const (
	PublishingMediaTypeImage    PublishingMediaType = "IMAGE"
	PublishingMediaTypeReels    PublishingMediaType = "REELS"
	PublishingMediaTypeStories  PublishingMediaType = "STORIES"
	PublishingMediaTypeCarousel PublishingMediaType = "CAROUSEL"
)

// ContainerStatus is the processing state of an Instagram media container.
type ContainerStatus string

const (
	ContainerStatusExpired    ContainerStatus = "EXPIRED"
	ContainerStatusError      ContainerStatus = "ERROR"
	ContainerStatusFinished   ContainerStatus = "FINISHED"
	ContainerStatusInProgress ContainerStatus = "IN_PROGRESS"
	ContainerStatusPublished  ContainerStatus = "PUBLISHED"
)

// InsightPeriod is the aggregation period for Instagram insights.
type InsightPeriod string

const (
	InsightPeriodDay      InsightPeriod = "day"
	InsightPeriodLifetime InsightPeriod = "lifetime"
)
