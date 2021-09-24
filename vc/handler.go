package vc

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/bot"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"github.com/bwmarrin/discordgo"
	"log"
)

// OnGuildVoiceJoinHandler is a very specific use-case handler function that controls follow and entrance behavior
func OnGuildVoiceJoinHandler(b *bot.Bot) func(*discordgo.Session, *discordgo.VoiceStateUpdate) {
	return func(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
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
		u, err := b.Session.User(vs.UserID)
		if err != nil {
			return
		}
		if vs.BeforeUpdate.ChannelID != vs.ChannelID {
			existing, ok := b.VoiceConnections[vs.GuildID]
			if ok && existing.ChannelID != vs.ChannelID || !ok {
				vc, err := s.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, true)
				if err != nil {
					return
				}
				b.VoiceConnections[vs.GuildID] = vc
			}
			entrance := sound.GetEntranceForUser(vs.UserID)
			if entrance != nil {
				var soundMap = sound.GetSounds()
				var file = (*soundMap)[entrance.SoundID]
				var soundInfo string
				err = sound.PlayDCA(file.FilePath, b.VoiceConnections[vs.GuildID])
				if err != nil {
					return
				}
				file.NumberPlays++
				if err = sound.GetLibrary().SetSoundData(file, bot.Db); err != nil {
					log.Fatalln("When updating sound =>" + err.Error())
				}
				soundInfo = fmt.Sprintf("Played `%s` from **%s** (**%d** plays)", file.ID, file.Categories[0], file.NumberPlays)
				chat.SendWelcomeEmbedMessage(b.Session, vs.ChannelID, u, soundInfo)
			}
		}
	}
}
