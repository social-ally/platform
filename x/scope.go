package x

// Scope is an OAuth permission supported by x.
type Scope string

const (
	// ScopeBlockRead grants the block.read permission.
	ScopeBlockRead Scope = "block.read"
	// ScopeBlockWrite grants the block.write permission.
	ScopeBlockWrite Scope = "block.write"
	// ScopeBookmarkRead grants the bookmark.read permission.
	ScopeBookmarkRead Scope = "bookmark.read"
	// ScopeBookmarkWrite grants the bookmark.write permission.
	ScopeBookmarkWrite Scope = "bookmark.write"
	// ScopeFollowsRead grants the follows.read permission.
	ScopeFollowsRead Scope = "follows.read"
	// ScopeFollowsWrite grants the follows.write permission.
	ScopeFollowsWrite Scope = "follows.write"
	// ScopeLikeRead grants the like.read permission.
	ScopeLikeRead Scope = "like.read"
	// ScopeLikeWrite grants the like.write permission.
	ScopeLikeWrite Scope = "like.write"
	// ScopeListRead grants the list.read permission.
	ScopeListRead Scope = "list.read"
	// ScopeListWrite grants the list.write permission.
	ScopeListWrite Scope = "list.write"
	// ScopeMuteRead grants the mute.read permission.
	ScopeMuteRead Scope = "mute.read"
	// ScopeMuteWrite grants the mute.write permission.
	ScopeMuteWrite Scope = "mute.write"
	// ScopeOfflineAccess grants the offline.access permission.
	ScopeOfflineAccess Scope = "offline.access"
	// ScopeSpaceRead grants the space.read permission.
	ScopeSpaceRead Scope = "space.read"
	// ScopeTweetRead grants the tweet.read permission.
	ScopeTweetRead Scope = "tweet.read"
	// ScopeTweetWrite grants the tweet.write permission.
	ScopeTweetWrite Scope = "tweet.write"
	// ScopeUsersRead grants the users.read permission.
	ScopeUsersRead Scope = "users.read"
)
