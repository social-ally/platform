package youtube

// UploadType identifies the YouTube video upload protocol.
type UploadType string

const (
	UploadTypeResumable UploadType = "resumable"
	UploadTypeMultipart UploadType = "multipart"
)

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
