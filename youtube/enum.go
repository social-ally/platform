package youtube

// UploadType identifies the YouTube video upload protocol.
type UploadType string

const (
	UploadTypeResumable UploadType = "resumable"
	UploadTypeMultipart UploadType = "multipart"
)

// ChannelPart selects a resource part on a YouTube channel.
type ChannelPart string

const (
	ChannelPartAuditDetails        ChannelPart = "auditDetails"
	ChannelPartBrandingSettings    ChannelPart = "brandingSettings"
	ChannelPartContentDetails      ChannelPart = "contentDetails"
	ChannelPartContentOwnerDetails ChannelPart = "contentOwnerDetails"
	ChannelPartID                  ChannelPart = "id"
	ChannelPartLocalizations       ChannelPart = "localizations"
	ChannelPartSnippet             ChannelPart = "snippet"
	ChannelPartStatistics          ChannelPart = "statistics"
	ChannelPartStatus              ChannelPart = "status"
	ChannelPartTopicDetails        ChannelPart = "topicDetails"
)

// VideoPart selects a resource part on a YouTube video.
type VideoPart string

const (
	VideoPartContentDetails              VideoPart = "contentDetails"
	VideoPartFileDetails                 VideoPart = "fileDetails"
	VideoPartID                          VideoPart = "id"
	VideoPartLiveStreamingDetails        VideoPart = "liveStreamingDetails"
	VideoPartLocalizations               VideoPart = "localizations"
	VideoPartPaidProductPlacementDetails VideoPart = "paidProductPlacementDetails"
	VideoPartPlayer                      VideoPart = "player"
	VideoPartProcessingDetails           VideoPart = "processingDetails"
	VideoPartRecordingDetails            VideoPart = "recordingDetails"
	VideoPartSnippet                     VideoPart = "snippet"
	VideoPartStatistics                  VideoPart = "statistics"
	VideoPartStatus                      VideoPart = "status"
	VideoPartSuggestions                 VideoPart = "suggestions"
	VideoPartTopicDetails                VideoPart = "topicDetails"
)

// PlaylistPart selects a resource part on a YouTube playlist.
type PlaylistPart string

const (
	PlaylistPartContentDetails PlaylistPart = "contentDetails"
	PlaylistPartID             PlaylistPart = "id"
	PlaylistPartLocalizations  PlaylistPart = "localizations"
	PlaylistPartPlayer         PlaylistPart = "player"
	PlaylistPartSnippet        PlaylistPart = "snippet"
	PlaylistPartStatus         PlaylistPart = "status"
)

// PlaylistItemPart selects a resource part on a YouTube playlist item.
type PlaylistItemPart string

const (
	PlaylistItemPartContentDetails PlaylistItemPart = "contentDetails"
	PlaylistItemPartID             PlaylistItemPart = "id"
	PlaylistItemPartSnippet        PlaylistItemPart = "snippet"
	PlaylistItemPartStatus         PlaylistItemPart = "status"
)

// MetadataContentType identifies the JSON content type used to initialize a YouTube upload.
type MetadataContentType string

const MetadataContentTypeJSON MetadataContentType = "application/json"

// VideoUploadContentType identifies a media content type accepted by YouTube uploads.
type VideoUploadContentType string

const VideoUploadContentTypeVideo VideoUploadContentType = "video/*"

// PrivacyStatus controls visibility of a YouTube resource.
type PrivacyStatus string

const (
	PrivacyStatusPrivate  PrivacyStatus = "private"
	PrivacyStatusPublic   PrivacyStatus = "public"
	PrivacyStatusUnlisted PrivacyStatus = "unlisted"
)

// AnalyticsDimension identifies a YouTube Analytics grouping dimension.
type AnalyticsDimension string

const (
	AnalyticsDimensionDay     AnalyticsDimension = "day"
	AnalyticsDimensionVideo   AnalyticsDimension = "video"
	AnalyticsDimensionCountry AnalyticsDimension = "country"
)

// AnalyticsColumnType identifies the kind of Analytics result column.
type AnalyticsColumnType string

const (
	AnalyticsColumnTypeDimension AnalyticsColumnType = "DIMENSION"
	AnalyticsColumnTypeMetric    AnalyticsColumnType = "METRIC"
)

// AnalyticsDataType identifies the value type of an Analytics result column.
type AnalyticsDataType string

const (
	AnalyticsDataTypeString  AnalyticsDataType = "STRING"
	AnalyticsDataTypeInteger AnalyticsDataType = "INTEGER"
	AnalyticsDataTypeFloat   AnalyticsDataType = "FLOAT"
)
