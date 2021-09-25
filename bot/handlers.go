package bot

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func openConnection(b *Bot, channelID, guildID string) error {
	existing, ok := b.VoiceConnections[guildID]
	if ok && existing.ChannelID != channelID || !ok {
		vc, err := b.Session.ChannelVoiceJoin(guildID, channelID, false, true)
		if err != nil {
			return err
		}
		b.VoiceConnections[guildID] = vc
	}
	return nil
}

func closeConnectionIfAlone(b *Bot, s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
	//if len(usersFound) == 1 {
	//	b.VoiceConnections[vs.GuildID].Close()
	//}
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
		g, err := s.Guild(vs.GuildID)
		if err != nil {
			return
		}
		defer closeConnectionIfAlone(b, g, vs)
		if vs.UserID == s.State.User.ID { // move done by bot
			return
		}
		u, err := b.Session.User(vs.UserID)
		if err != nil {
			return
		}
		log.Println("Move detected for user", u, "from state", vs.BeforeUpdate, "to channel", vs.ChannelID)
		if len(vs.ChannelID) > 0 && (vs.BeforeUpdate == nil || vs.BeforeUpdate.ChannelID != vs.ChannelID) {
			if err = openConnection(b, vs.ChannelID, vs.GuildID); err != nil {
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
				soundInfo = fmt.Sprintf("Played `%s` from **%s** (**%d** plays)", file.ID, file.Categories[0], file.NumberPlays)
				chat.SendWelcomeEmbedMessage(b.Session, getMainChannelIDForGuild(s, vs.GuildID), u, soundInfo)
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

// OnPlayMessageCommandReceivedHandler is another specific use-case handler function to handle play chat commands that require more than just a Discord session
func OnPlayMessageCommandReceivedHandler(b *Bot) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.HasPrefix(m.Content, "?") && len(m.Content) > 1 {
			if err := openConnection(b, m.ChannelID, m.GuildID); err != nil {
				log.Fatal(err)
				return
			}
			PlaySoundMessageCommand(s, b.VoiceConnections[m.GuildID], m)
			chat.DeleteReceivedMessage(s, m)
		}
	}
}

// OnPlaySlashCommandReceivedHandler is another specific use-case handler function to handle play slash commands that require more than just a Discord session
func OnPlaySlashCommandReceivedHandler(b *Bot) func(*discordgo.Session, *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name == "sound" {
			if err := openConnection(b, i.ChannelID, i.GuildID); err != nil {
				return
			}
			PlaySoundSlashCommand(s, b.VoiceConnections[i.GuildID], i)
		}
	}
}
