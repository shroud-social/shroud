package guild

import (
	"time"

	"github.com/google/uuid"
)

type AccessType string

const (
	AccessInviteOnly  AccessType = "INVITE_ONLY"
	AccessApplication AccessType = "APPLICATION"
	AccessOpen        AccessType = "OPEN"
)

type VerificationLevel string

const (
	VerificationLevelNone   VerificationLevel = "NONE"
	VerificationLevelLow    VerificationLevel = "LOW"
	VerificationLevelMedium VerificationLevel = "MEDIUM"
	VerificationLevelHigh   VerificationLevel = "HIGH"
)

type QuestionType string

const (
	QuestionTypeShort          QuestionType = "SHORT"
	QuestionTypeLong           QuestionType = "LONG"
	QuestionTypeMultipleChoice QuestionType = "MULTIPLE_CHOICE"
)

// Guild represents a community and its associated configuration
type Guild struct {
	ID      uuid.UUID `json:"id"`
	RealmID uuid.UUID `json:"realm_id"`
	OwnerID uuid.UUID `json:"owner_id"`

	CreationTime time.Time `json:"creation_time"`

	NSFW bool `json:"nsfw"`

	// Identity
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Icon             string   `json:"icon"`
	Banner           string   `json:"banner"`
	InviteBackground string   `json:"invite_background"`
	Tags             []string `json:"tags"`
	HTMLWidget       bool     `json:"html_widget"`

	badge GuildBadge `json:"badge"`

	Tier guild.Tier `json:"tier"`

	// Engagement
	WelcomeMessage   string `json:"welcome_message"` // if blank, random
	SupportMessage   string `json:"support_message"` // if blank, random
	UserFeed         bool   `json:"user_feed"`
	NotificationsAll bool   `json:"notifications"` // All Messages or just mentions

	// Realm Admin Notifications
	RealmGeneralNotificationChannel    uuid.UUID `json:"realm_general_notification_channel"`
	RealmModerationNotificationChannel uuid.UUID `json:"realm_moderation_notification_channel"`

	// Access
	AccessType           AccessType             `json:"access_type"` // Invite Only, Application, Open
	ApplicationQuestions []*ApplicationQuestion `json:"application_questions"`

	// Guild Security
	SuspiciousActivityNotifications uuid.UUID         `json:"suspicious_activity_notifications"`
	VerificationLevel               VerificationLevel `json:"verification_level"`
	AllowDMsFromSuspicious          bool              `json:"allow_dm_from_suspicious"`
	AllowDMsFromUnknown             bool              `json:"allow_dm_from_unknown"`
	RequireModerator2FA             bool              `json:"require_moderator_2fa"`

	// AutoMod
	AutoModRules  []*AutoModRule `json:"auto_mod_rules"`
	ContentFilter bool           `json:"content_filter"`
}

// GuildBadge represents a custom badge which members can apply to their profile globally
type GuildBadge struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// ApplicationQuestion represents a question a prospective user might answer to join a guild
type ApplicationQuestion struct {
	Type     QuestionType `json:"type"`
	Question string       `json:"text"`
	Optional bool         `json:"optional"`

	// Type: Multiple Choice
	Options     []string `json:"options"`
	AnswerLimit int      `json:"answer_limit"`
}

// AutoModRule represents a rule which messages or usernames will be checked against
type AutoModRule struct {
	WordMatch    string `json:"word_match"`
	RegexPattern string `json:"regex_pattern"`
	ExcludeWords string `json:"exceptions"`

	BlockEvent   bool      `json:"block_event"`
	AlertChannel uuid.UUID `json:"send_alert"` // If blank, no alert

	ExcludeRoles []string `json:"exclude_roles"`

	ApplyToNames     bool `json:"apply_to_names"`
	ApplyToNicknames bool `json:"apply_to_nicknames"`
	ApplyToChats     bool `json:"apply_to_chats"`
}
