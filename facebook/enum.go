package facebook

// InsightPeriod is the aggregation period for Page insights.
type InsightPeriod string

const (
	InsightPeriodDay      InsightPeriod = "day"
	InsightPeriodWeek     InsightPeriod = "week"
	InsightPeriodDays28   InsightPeriod = "days_28"
	InsightPeriodLifetime InsightPeriod = "lifetime"
)

// UserField selects a field returned by the Facebook user endpoint.
type UserField string

const (
	UserFieldID        UserField = "id"
	UserFieldName      UserField = "name"
	UserFieldFirstName UserField = "first_name"
	UserFieldLastName  UserField = "last_name"
	UserFieldEmail     UserField = "email"
	UserFieldPicture   UserField = "picture"
	UserFieldLink      UserField = "link"
)

// PageField selects a field returned by a Facebook Page endpoint.
type PageField string

const (
	PageFieldID                 PageField = "id"
	PageFieldName               PageField = "name"
	PageFieldAccessToken        PageField = "access_token"
	PageFieldCategory           PageField = "category"
	PageFieldCategoryList       PageField = "category_list"
	PageFieldTasks              PageField = "tasks"
	PageFieldAbout              PageField = "about"
	PageFieldDescription        PageField = "description"
	PageFieldFanCount           PageField = "fan_count"
	PageFieldFollowersCount     PageField = "followers_count"
	PageFieldWebsite            PageField = "website"
	PageFieldUsername           PageField = "username"
	PageFieldPicture            PageField = "picture"
	PageFieldVerificationStatus PageField = "verification_status"
)
