package threads

// MediaType identifies the content type of a Threads post container.
type MediaType string

const (
	MediaTypeText     MediaType = "TEXT"
	MediaTypeImage    MediaType = "IMAGE"
	MediaTypeVideo    MediaType = "VIDEO"
	MediaTypeCarousel MediaType = "CAROUSEL"
)

// UserField selects a field returned by the Threads user endpoint.
type UserField string

const (
	UserFieldID                UserField = "id"
	UserFieldUsername          UserField = "username"
	UserFieldName              UserField = "name"
	UserFieldProfilePictureURL UserField = "threads_profile_picture_url"
	UserFieldBiography         UserField = "threads_biography"
)

// ThreadField selects a field returned by the Threads listing endpoint.
type ThreadField string

const (
	ThreadFieldID               ThreadField = "id"
	ThreadFieldMediaProductType ThreadField = "media_product_type"
	ThreadFieldMediaType        ThreadField = "media_type"
	ThreadFieldMediaURL         ThreadField = "media_url"
	ThreadFieldPermalink        ThreadField = "permalink"
	ThreadFieldOwner            ThreadField = "owner"
	ThreadFieldUsername         ThreadField = "username"
	ThreadFieldText             ThreadField = "text"
	ThreadFieldTimestamp        ThreadField = "timestamp"
	ThreadFieldShortcode        ThreadField = "shortcode"
	ThreadFieldThumbnailURL     ThreadField = "thumbnail_url"
	ThreadFieldChildren         ThreadField = "children"
	ThreadFieldQuotePost        ThreadField = "is_quote_post"
	ThreadFieldQuotedPost       ThreadField = "quoted_post"
	ThreadFieldRepostedPost     ThreadField = "reposted_post"
	ThreadFieldReplyAudience    ThreadField = "reply_audience"
)

// ContainerStatus is the processing state of a Threads post container.
type ContainerStatus string

const (
	ContainerStatusInProgress ContainerStatus = "IN_PROGRESS"
	ContainerStatusFinished   ContainerStatus = "FINISHED"
	ContainerStatusError      ContainerStatus = "ERROR"
	ContainerStatusExpired    ContainerStatus = "EXPIRED"
)
