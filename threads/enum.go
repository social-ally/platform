package threads

// MediaType identifies the content type of a Threads post container.
type MediaType string

const (
	MediaTypeText     MediaType = "TEXT"
	MediaTypeImage    MediaType = "IMAGE"
	MediaTypeVideo    MediaType = "VIDEO"
	MediaTypeCarousel MediaType = "CAROUSEL"
)

// ContainerStatus is the processing state of a Threads post container.
type ContainerStatus string

const (
	ContainerStatusInProgress ContainerStatus = "IN_PROGRESS"
	ContainerStatusFinished   ContainerStatus = "FINISHED"
	ContainerStatusError      ContainerStatus = "ERROR"
	ContainerStatusExpired    ContainerStatus = "EXPIRED"
)
