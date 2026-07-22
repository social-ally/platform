package threads

// Scope is an OAuth permission supported by threads.
type Scope string

const (
	// ScopeThreadsBasic grants the threads_basic permission.
	ScopeThreadsBasic Scope = "threads_basic"
	// ScopeThreadsContentPublish grants the threads_content_publish permission.
	ScopeThreadsContentPublish Scope = "threads_content_publish"
	// ScopeThreadsKeywordSearch grants the threads_keyword_search permission.
	ScopeThreadsKeywordSearch Scope = "threads_keyword_search"
	// ScopeThreadsLocationTagging grants the threads_location_tagging permission.
	ScopeThreadsLocationTagging Scope = "threads_location_tagging"
	// ScopeThreadsManageInsights grants the threads_manage_insights permission.
	ScopeThreadsManageInsights Scope = "threads_manage_insights"
	// ScopeThreadsManageReplies grants the threads_manage_replies permission.
	ScopeThreadsManageReplies Scope = "threads_manage_replies"
	// ScopeThreadsReadReplies grants the threads_read_replies permission.
	ScopeThreadsReadReplies Scope = "threads_read_replies"
)
