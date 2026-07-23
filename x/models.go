package x

// OAuthToken is an OAuth token response from X.
type OAuthToken struct {
	TokenType    string  `json:"token_type"`
	ExpiresIn    int     `json:"expires_in"`
	AccessToken  string  `json:"access_token"`
	Scope        string  `json:"scope"`
	RefreshToken *string `json:"refresh_token"`
}

// OAuthError is an OAuth error response from X.
type OAuthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorURI         string `json:"error_uri"`
	State            string `json:"state"`
}

// Problem is the common RFC 7807-style error response returned by X.
type Problem struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Status   int    `json:"status"`
	Instance string `json:"instance"`
	Errors   []any  `json:"errors"`
}

// PublicMetrics contains public X account and post counters.
type PublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	LikeCount      int `json:"like_count"`
	ListedCount    int `json:"listed_count"`
	MediaCount     int `json:"media_count"`
	TweetCount     int `json:"tweet_count"`
}

// UserEntities contains entities attached to an X user profile.
type UserEntities struct {
	Description any `json:"description"`
	URL         any `json:"url"`
}

// User is an X user resource. Optional API fields may be omitted by X.
type User struct {
	ID                 string         `json:"id"`
	Name               string         `json:"name"`
	Username           string         `json:"username"`
	CreatedAt          string         `json:"created_at"`
	Description        string         `json:"description"`
	Entities           *UserEntities  `json:"entities"`
	Location           string         `json:"location"`
	ProfileImageURL    string         `json:"profile_image_url"`
	ProfileBannerURL   string         `json:"profile_banner_url"`
	Protected          bool           `json:"protected"`
	PublicMetrics      *PublicMetrics `json:"public_metrics"`
	URL                string         `json:"url"`
	Verified           bool           `json:"verified"`
	VerifiedType       string         `json:"verified_type"`
	PinnedTweetID      string         `json:"pinned_tweet_id"`
	MostRecentTweetID  string         `json:"most_recent_tweet_id"`
	ReceivesYourDM     bool           `json:"receives_your_dm"`
	SubscriptionType   string         `json:"subscription_type"`
	ConnectionStatus   []string       `json:"connection_status"`
	ConfirmedEmail     bool           `json:"confirmed_email"`
	IsIdentityVerified bool           `json:"is_identity_verified"`
	Parody             bool           `json:"parody"`
	Affiliation        any            `json:"affiliation"`
	Subscription       any            `json:"subscription"`
	Withheld           any            `json:"withheld"`
}

// TweetMetrics contains metrics included on an X post.
type TweetMetrics struct {
	BookmarkCount     int `json:"bookmark_count"`
	ImpressionCount   int `json:"impression_count"`
	LikeCount         int `json:"like_count"`
	QuoteCount        int `json:"quote_count"`
	ReplyCount        int `json:"reply_count"`
	RetweetCount      int `json:"retweet_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
}

// Tweet is an X post resource.
type Tweet struct {
	ID                     string           `json:"id"`
	Text                   string           `json:"text"`
	AuthorID               string           `json:"author_id"`
	ConversationID         string           `json:"conversation_id"`
	CreatedAt              string           `json:"created_at"`
	InReplyToUserID        string           `json:"in_reply_to_user_id"`
	Lang                   string           `json:"lang"`
	PossiblySensitive      bool             `json:"possibly_sensitive"`
	ReplySettings          ReplySetting     `json:"reply_settings"`
	Source                 string           `json:"source"`
	Attachments            any              `json:"attachments"`
	ContextAnnotations     []any            `json:"context_annotations"`
	EditControls           any              `json:"edit_controls"`
	EditHistoryTweetIDs    []string         `json:"edit_history_tweet_ids"`
	Entities               any              `json:"entities"`
	Geo                    any              `json:"geo"`
	NoteTweet              any              `json:"note_tweet"`
	OrganicMetrics         *TweetMetrics    `json:"organic_metrics"`
	PublicMetrics          *TweetMetrics    `json:"public_metrics"`
	NonPublicMetrics       *TweetMetrics    `json:"non_public_metrics"`
	PromotedMetrics        *TweetMetrics    `json:"promoted_metrics"`
	ReferencedTweets       []TweetReference `json:"referenced_tweets"`
	Withheld               any              `json:"withheld"`
	Article                any              `json:"article"`
	CardURI                *string          `json:"card_uri"`
	CommunityID            *string          `json:"community_id"`
	DisplayTextRange       []int            `json:"display_text_range"`
	MatchedMediaNotes      []any            `json:"matched_media_notes"`
	MediaMetadata          any              `json:"media_metadata"`
	NoteRequestSuggestions []any            `json:"note_request_suggestions"`
}

// TweetReference identifies another post related to a post.
type TweetReference struct {
	Type TweetReferenceType `json:"type"`
	ID   string             `json:"id"`
}

// Media is a media resource included by X.
type Media struct {
	MediaKey         string    `json:"media_key"`
	Type             MediaType `json:"type"`
	URL              string    `json:"url"`
	PreviewImageURL  string    `json:"preview_image_url"`
	DurationMS       int       `json:"duration_ms"`
	Height           int       `json:"height"`
	Width            int       `json:"width"`
	AltText          string    `json:"alt_text"`
	Variants         []any     `json:"variants"`
	PublicMetrics    any       `json:"public_metrics"`
	NonPublicMetrics any       `json:"non_public_metrics"`
	OrganicMetrics   any       `json:"organic_metrics"`
	PromotedMetrics  any       `json:"promoted_metrics"`
}

// Includes contains related resources returned alongside a post or user.
type Includes struct {
	Media  []Media `json:"media"`
	Places []any   `json:"places"`
	Polls  []any   `json:"polls"`
	Tweets []Tweet `json:"tweets"`
	Users  []User  `json:"users"`
}

// Meta describes an X collection response.
type Meta struct {
	NewestID      string `json:"newest_id"`
	NextToken     string `json:"next_token"`
	OldestID      string `json:"oldest_id"`
	PreviousToken string `json:"previous_token"`
	ResultCount   int    `json:"result_count"`
}
