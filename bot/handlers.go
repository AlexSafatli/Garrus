package bot

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"github.com/bwmarrin/discordgo"
	"log"
)

func openConnection(s *discordgo.Session, channelID, guildID string) error {
	existing, ok := s.VoiceConnections[guildID]
	if ok && existing.ChannelID != channelID || !ok {
		_, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func closeConnectionOrChangeChannelsIfAlone(s *discordgo.Session, guildID string) {
	if s.VoiceConnections[guildID] == nil {
		return
	}
	g, err := s.State.Guild(guildID)
	if err != nil {
		return
	}
	var totalUsersFound int
	var usersFound map[string]int
	usersFound = make(map[string]int)
	for _, vs := range g.VoiceStates {
		if vs.UserID != s.State.User.ID {
			usersFound[vs.ChannelID]++
			totalUsersFound++
		}
	}
	if totalUsersFound == 0 {
		if err = s.VoiceConnections[guildID].Disconnect(); err != nil {
			s.VoiceConnections[guildID].Close()
		}
	} else {
		var mostUsers int
		var channelIDWithMostUsers string
		for k, v := range usersFound {
			if v > mostUsers {
				channelIDWithMostUsers = k
			}
		}
		if channelIDWithMostUsers != s.VoiceConnections[guildID].ChannelID {
			_ = s.VoiceConnections[guildID].ChangeChannel(channelIDWithMostUsers, false, true)
		}
	}
}

func getMainChannelIDForGuild(s *discordgo.Session, guildID string) string {
	var id string
	var lowestPos = -1
	gc, err := s.GuildChannels(guildID)
	if err != nil {
		return id
	}
	for _, c := range gc {
		if c.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		if lowestPos == -1 || c.Position < lowestPos {
			lowestPos = c.Position
			id = c.ID
		}
	}
	return id
}

// OnGuildVoiceJoinHandler is a very specific use-case handler function that controls follow and entrance behavior
func OnGuildVoiceJoinHandler(b *Bot) func(*discordgo.Session, *discordgo.VoiceStateUpdate) {
	return func(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
		if vs.UserID == s.State.User.ID { // move done by bot
			return
		}
		u, err := b.Session.User(vs.UserID)
		if err != nil {
			return
		}
		if len(vs.ChannelID) == 0 {
			defer closeConnectionOrChangeChannelsIfAlone(s, vs.GuildID)
		}
		if len(vs.ChannelID) > 0 && (vs.BeforeUpdate == nil || vs.BeforeUpdate.ChannelID != vs.ChannelID) {
			if err = openConnection(s, vs.ChannelID, vs.GuildID); err != nil {
				return
			}
			entrance := sound.GetEntranceForUser(vs.UserID)
			if entrance != nil {
				db, err := LoadDatabase()
				if err != nil {
					log.Fatalln("Could not load database", err)
				}
				defer db.Close()
				var file = sound.GetLibrary().SoundMap[entrance.SoundID]
				var soundInfo string
				channelID := getMainChannelIDForGuild(s, vs.GuildID)
				soundInfo = fmt.Sprintf("Played `%s` from **%s** (**%d** plays)", file.ID, file.Categories[0], file.NumberPlays)
				if lastMessageID, ok := b.lastSentEntranceMessage[vs.GuildID]; ok {
					chat.DeleteBotMessages(s, channelID, lastMessageID)
				}
				m := chat.SendWelcomeEmbedMessage(b.Session, channelID, u, soundInfo)
				b.lastSentEntranceMessage[vs.GuildID] = m.ID // keep track of the last sent entrance message
				err = sound.PlayDCA(file.FilePath, b.VoiceConnections[vs.GuildID])
				if err != nil {
					return
				}
				file.NumberPlays++
				if err = sound.GetLibrary().SetSoundData(file, db); err != nil {
					log.Fatalln("When updating sound => " + err.Error())
				}
			}
		}
	}
}
