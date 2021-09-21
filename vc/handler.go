package vc

import (
	"github.com/AlexSafatli/Garrus/bot"
	"github.com/bwmarrin/discordgo"
)

// OnGuildVoiceJoinHandler is a very specific use-case handler function that controls follow and entrance behavior
func OnGuildVoiceJoinHandler(b *bot.Bot) func(*discordgo.Session, *discordgo.VoiceStateUpdate) {
	if s == nil {
		return nil
	}
	return func(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
		c, err := s.Channel(vs.ChannelID)
		if err != nil {
			return
		}
		g, err := s.Guild(vs.GuildID)
		if err != nil {
			return
		}
		if vs.UserID == s.State.User.ID { // move done by bot
			if len(g.VoiceStates) == 1 {
				b.VoiceConnections[vs.GuildID].Close()
			}
			return
		}
		vc, err := s.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, true)
		if err != nil {
			return
		}
		b.VoiceConnections[vs.GuildID] = vc
	}
}
