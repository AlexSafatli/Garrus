package bot

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"github.com/boltdb/bolt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func getUsersVoiceConnectionsCountMap(s *discordgo.Session, g *discordgo.Guild) (total int, usersFound map[string]int) {
	usersFound = make(map[string]int)
	for _, vs := range g.VoiceStates {
		if vs.UserID != s.State.User.ID {
			usersFound[vs.ChannelID]++
			total++
		}
	}
	return
}

func getUsersVoiceChannelID(s *discordgo.Session, guildID, userID string) (channelID string) {
	g, err := s.State.Guild(guildID)
	if err != nil {
		return
	}
	for _, vs := range g.VoiceStates {
		if vs.UserID == userID {
			channelID = vs.ChannelID
			return
		}
	}
	return
}

func openVoiceConnection(s *discordgo.Session, channelID, guildID string) error {
	existing, ok := s.VoiceConnections[guildID]
	if ok && existing.ChannelID != channelID || !ok {
		if _, err := s.ChannelVoiceJoin(guildID, channelID, false, true); err != nil {
			return err
		}
	}
	return nil
}

func closeVoiceConnectionOrChangeChannelsIfAlone(s *discordgo.Session, guildID string) {
	if s.VoiceConnections[guildID] == nil {
		return // no connection opened
	}
	g, err := s.State.Guild(guildID)
	if err != nil {
		return
	}

	totalUsersFound, usersFound := getUsersVoiceConnectionsCountMap(s, g)
	channelID := s.VoiceConnections[guildID].ChannelID

	if totalUsersFound == 0 {
		if err = s.VoiceConnections[guildID].Disconnect(); err != nil {
			s.VoiceConnections[guildID].Close()
		}
	} else if usersFound[channelID] == 0 {
		var mostUsers int
		var channelIDWithMostUsers string
		for k, v := range usersFound {
			if v > mostUsers {
				channelIDWithMostUsers = k
			}
		}
		if channelIDWithMostUsers != channelID {
			_ = s.VoiceConnections[guildID].ChangeChannel(channelIDWithMostUsers, false, true)
		}
	}
}

func playSound(file *sound.File, vc *discordgo.VoiceConnection) {
	go func() {
		if err := sound.PlayDCA(file.FilePath, vc); err != nil {
			log.Println("When playing sound " + file.ID + " -> " + err.Error())
		}
	}()
}

func playSoundWithSave(file *sound.File, vc *discordgo.VoiceConnection, db *bolt.DB) {
	playSound(file, vc)
	file.NumberPlays++
	if err := sound.GetLibrary().SetSoundData(file, db); err != nil {
		log.Fatalln("When updating sound => " + err.Error())
	}
}

func followOnMove(b *Bot, s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
	var err error
	vc, ok := s.VoiceConnections[vs.GuildID]
	if vs.UserID == s.State.User.ID { // move done by bot
		return
	}
	if len(vs.ChannelID) == 0 { // empty target voice channel
		defer closeVoiceConnectionOrChangeChannelsIfAlone(s, vs.GuildID)
	}
	if len(vs.ChannelID) > 0 && (vs.BeforeUpdate == nil || vs.BeforeUpdate.ChannelID != vs.ChannelID) {
		if !ok {
			if err = openVoiceConnection(s, vs.ChannelID, vs.GuildID); err != nil {
				log.Printf("Error when joining voice channel %s -> %v", vs.ChannelID, err)
				return
			}
			vc = s.VoiceConnections[vs.GuildID]
		}
		if vc.ChannelID != vs.ChannelID {
			return
		}

		entrance := sound.GetEntranceForUser(vs.UserID)

		// If the user has an entrance, play it
		if entrance != nil {

			// Get the file to play
			var file = sound.GetLibrary().SoundMap[entrance.SoundID]
			if file == nil {
				log.Println("File was not found for entrance", entrance)
				return
			}

			// Play it in a goroutine; do not increment play count
			playSound(file, vc)

			// Send a welcome message, delete old bot messages
			channelID := getMainChannelIDForGuild(b, vs.GuildID)
			soundInfo := fmt.Sprintf("Played `%s` from **%s** (**%d** plays)", file.ID, file.Categories[0], file.NumberPlays)
			u, err := b.Session.User(vs.UserID)
			if err != nil {
				return
			}
			if lastMessageID, ok := b.lastSentEntranceMessage[vs.GuildID]; ok {
				chat.DeleteBotMessages(s, channelID, lastMessageID)
			}
			m := chat.SendWelcomeEmbedMessage(b.Session, channelID, u, soundInfo)
			b.lastSentEntranceMessage[vs.GuildID] = m.ID // keep track of the last sent entrance message
		}
	}
}
