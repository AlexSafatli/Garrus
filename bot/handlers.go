package bot

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"github.com/bwmarrin/discordgo"
	"log"
)

func getMainChannelIDForGuild(b *Bot, guildID string) string {
	var id string
	if id, ok := b.mainGuildChannelIDs[guildID]; ok {
		return id
	}
	var lowestPos = -1
	gc, err := b.Session.GuildChannels(guildID)
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
	if len(id) > 0 {
		b.mainGuildChannelIDs[guildID] = id // cache the channel ID
	}
	return id
}

// OnGuildVoiceJoinHandler is a very specific use-case handler function that controls follow and entrance behavior
func OnGuildVoiceJoinHandler(b *Bot) func(*discordgo.Session, *discordgo.VoiceStateUpdate) {
	return func(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
		var err error
		if vs.UserID == s.State.User.ID { // move done by bot
			return
		}
		if len(vs.ChannelID) == 0 { // empty target voice channel
			defer closeConnectionOrChangeChannelsIfAlone(s, vs.GuildID)
		}
		if len(vs.ChannelID) > 0 && (vs.BeforeUpdate == nil || vs.BeforeUpdate.ChannelID != vs.ChannelID) { // play an entrance
			if err = openConnection(s, vs.ChannelID, vs.GuildID); err != nil {
				return
			}
			entrance := sound.GetEntranceForUser(vs.UserID)

			// If the user has an entrance, play it
			if entrance != nil {

				// Get the file to play
				var file = sound.GetLibrary().SoundMap[entrance.SoundID]

				// Play it in a goroutine
				go func() {
					err := sound.PlayDCA(file.FilePath, b.VoiceConnections[vs.GuildID])
					if err != nil {
						log.Printf("Error while playing entrance for %s -> %v", vs.UserID, err)
					}
				}()

				// Send a welcome message, delete old bot messages
				var soundInfo string
				channelID := getMainChannelIDForGuild(b, vs.GuildID)
				soundInfo = fmt.Sprintf("Played `%s` from **%s** (**%d** plays)", file.ID, file.Categories[0], file.NumberPlays)
				u, err := b.Session.User(vs.UserID)
				if err != nil {
					return
				}
				if lastMessageID, ok := b.lastSentEntranceMessage[vs.GuildID]; ok {
					chat.DeleteBotMessages(s, channelID, lastMessageID)
				}
				m := chat.SendWelcomeEmbedMessage(b.Session, channelID, u, soundInfo)
				b.lastSentEntranceMessage[vs.GuildID] = m.ID // keep track of the last sent entrance message

				// Load database and save changes to database
				db, err := LoadDatabase()
				if err != nil {
					log.Fatalln("Could not load database", err)
				}
				defer db.Close()
				file.NumberPlays++
				if err = sound.GetLibrary().SetSoundData(file, db); err != nil {
					log.Fatalln("When updating sound => " + err.Error())
				}
			}
		}
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
