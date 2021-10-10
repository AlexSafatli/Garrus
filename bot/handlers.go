package bot

import (
	"github.com/bwmarrin/discordgo"
)

// OnGuildVoiceJoinHandler is a very specific use-case handler function that controls follow and entrance behavior
func OnGuildVoiceJoinHandler(b *Bot) func(*discordgo.Session, *discordgo.VoiceStateUpdate) {
	return func(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
		followOnMove(b, s, vs)
	}
}

// OnGuildChannelCreateHandler is another specific use-case handler function that invalidates the bot's channel ID cache if the channel in the cache is deleted
func OnGuildChannelCreateHandler(b *Bot) func(*discordgo.Session, *discordgo.ChannelCreate) {
	return func(s *discordgo.Session, cc *discordgo.ChannelCreate) {
		if len(cc.GuildID) == 0 || cc.Channel == nil {
			return // check if the guild and channels are defined first
		}
		if id, ok := b.mainGuildChannelIDs[cc.GuildID]; ok {
			if id == cc.Channel.ID {
				delete(b.mainGuildChannelIDs, cc.GuildID) // invalidate cache
			} else {
				newMainChannel := getMainChannelIDForGuild(b, cc.GuildID)
				if newMainChannel != id {
					b.mainGuildChannelIDs[cc.GuildID] = newMainChannel // update
				}
			}
		}
	}
}
