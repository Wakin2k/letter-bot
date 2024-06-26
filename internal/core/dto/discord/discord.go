package discord

import (
	"spot-assistant/internal/core/dto/member"
	"time"
)

// Role is the name of the role that privileged users have with extra permissions.
const PrivilegedRole = "Postman"

// Name of the channel where summaries should be sent.
const SummaryChannel = "letter-summary"

// Name of the channel where commands are expected.
const CommandChannel = "letter"

// ChannelType is the type of a Channel
type ChannelType int

// Block contains known ChannelType values
const (
	ChannelTypeGuildText          ChannelType = 0
	ChannelTypeDM                 ChannelType = 1
	ChannelTypeGuildVoice         ChannelType = 2
	ChannelTypeGroupDM            ChannelType = 3
	ChannelTypeGuildCategory      ChannelType = 4
	ChannelTypeGuildNews          ChannelType = 5
	ChannelTypeGuildStore         ChannelType = 6
	ChannelTypeGuildNewsThread    ChannelType = 10
	ChannelTypeGuildPublicThread  ChannelType = 11
	ChannelTypeGuildPrivateThread ChannelType = 12
	ChannelTypeGuildStageVoice    ChannelType = 13
	ChannelTypeGuildForum         ChannelType = 15
)

type Channel struct {
	ID   string
	Name string
	Type ChannelType
}

// A User stores all data for an individual Discord user.
type User struct {
	// The ID of the user.
	ID string `json:"id"`

	// The user's username.
	Username string `json:"username"`
}

type Message struct {
	// The ID of the message.
	ID string `json:"id"`

	// The ID of the channel in which the message was sent.
	ChannelID string `json:"channel_id"`

	// The ID of the guild in which the message was sent.
	GuildID string `json:"guild_id,omitempty"`

	// The content of the message.
	Content string `json:"content"`

	// The time at which the messsage was sent.
	// CAUTION: this field may be removed in a
	// future API version; it is safer to calculate
	// the creation time via the ID.
	Timestamp time.Time `json:"timestamp"`

	// The time at which the last edit of the message
	// occurred, if it has been edited.
	EditedTimestamp *time.Time `json:"edited_timestamp"`

	// The author of the message. This is not guaranteed to be a
	// valid user (webhook-sent messages do not possess a full author).
	Author *User `json:"author"`

	// The webhook ID of the message, if it was generated by a webhook
	WebhookID string `json:"webhook_id"`

	// Member properties for this message's author,
	// contains only partial information
	Member *member.Member `json:"member"`
}
