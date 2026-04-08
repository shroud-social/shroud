package v1

import (
	"fmt"
)

const (
	event = ApiVersion1 + ".event"
)

func SubjectMessageCreated(guildId, channelId string) string {
	return fmt.Sprintf("%s.guild.%s.channel.%s.message.created", event, guildId, channelId)
}

func SubjectMessageUpdated(guildId, channelId string) string {
	return fmt.Sprintf("%s.guild.%s.channel.%s.message.updated", event, guildId, channelId)
}

func SubjectMessageDeleted(guildId, channelId string) string {
	return fmt.Sprintf("%s.guild.%s.channel.%s.message.deleted", event, guildId, channelId)
}
