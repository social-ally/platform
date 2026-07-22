package facebook

// Scope is an OAuth permission supported by facebook.
type Scope string

const (
	// ScopeEmail grants the email permission.
	ScopeEmail Scope = "email"
	// ScopePagesManagePosts grants the pages_manage_posts permission.
	ScopePagesManagePosts Scope = "pages_manage_posts"
	// ScopePagesReadEngagement grants the pages_read_engagement permission.
	ScopePagesReadEngagement Scope = "pages_read_engagement"
	// ScopePagesReadUserContent grants the pages_read_user_content permission.
	ScopePagesReadUserContent Scope = "pages_read_user_content"
	// ScopePagesShowList grants the pages_show_list permission.
	ScopePagesShowList Scope = "pages_show_list"
	// ScopePublicProfile grants the public_profile permission.
	ScopePublicProfile Scope = "public_profile"
	// ScopeReadInsights grants the read_insights permission.
	ScopeReadInsights Scope = "read_insights"
	// ScopeString grants the string permission.
	ScopeString Scope = "string"
)
