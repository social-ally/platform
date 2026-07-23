package instagram

// AccountType identifies the type of an Instagram account.
type AccountType string

const (
	AccountTypeBusiness     AccountType = "BUSINESS"
	AccountTypeMediaCreator AccountType = "MEDIA_CREATOR"
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
