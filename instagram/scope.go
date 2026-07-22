package instagram

// Scope is an OAuth permission supported by instagram.
type Scope string

const (
	// ScopeInstagramBusinessBasic grants the instagram_business_basic permission.
	ScopeInstagramBusinessBasic Scope = "instagram_business_basic"
	// ScopeInstagramBusinessContentPublish grants the instagram_business_content_publish permission.
	ScopeInstagramBusinessContentPublish Scope = "instagram_business_content_publish"
	// ScopeInstagramBusinessManageComments grants the instagram_business_manage_comments permission.
	ScopeInstagramBusinessManageComments Scope = "instagram_business_manage_comments"
	// ScopeInstagramBusinessManageMessages grants the instagram_business_manage_messages permission.
	ScopeInstagramBusinessManageMessages Scope = "instagram_business_manage_messages"
)
