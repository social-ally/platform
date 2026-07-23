package facebook

// InsightPeriod is the aggregation period for Page insights.
type InsightPeriod string

const (
	InsightPeriodDay      InsightPeriod = "day"
	InsightPeriodWeek     InsightPeriod = "week"
	InsightPeriodDays28   InsightPeriod = "days_28"
	InsightPeriodLifetime InsightPeriod = "lifetime"
)
