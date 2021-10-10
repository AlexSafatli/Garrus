package bot

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"github.com/boltdb/bolt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func playSound(file *sound.File, vc *discordgo.VoiceConnection, db *bolt.DB) {
	var err error
	go func() {
		err := sound.PlayDCA(file.FilePath, vc)
		if err != nil {
			log.Println("When playing sound " + file.ID + " => " + err.Error())
		}
	}()
	file.NumberPlays++
	if err = sound.GetLibrary().SetSoundData(file, db); err != nil {
		log.Fatalln("When updating sound => " + err.Error())
	}
}

func searchSounds(query string) (possibilities []string) {
	query = strings.ToLower(query)
	closest := sound.GetLibrary().GetClosestMatchingSoundID(query)
	if len(closest) > 0 {
		possibilities = append(possibilities, closest)
	}
	for _, name := range sound.GetLibrary().GetSoundNames() {
		if strings.Contains(name, query) && name != closest {
			possibilities = append(possibilities, name)
		}
	}
	return
}

func getSoundsForCategoryAsMessageStrings(cat string) ([]string, int) {
	mb := chat.MessageBuilder{}
	var total int
	for _, s := range sound.GetLibrary().SoundMap {
		if s.ContainsCategory(cat) {
			_ = mb.Write("`?" + s.ID + "` ")
			total++
		}
	}
	if total == 0 {
		return []string{}, 0
	}
	return mb.GetMessageStrings(), total
}

func sendSoundsForCategoryForMessageCommand(s *discordgo.Session, channelID string, c string) {
	msgs, total := getSoundsForCategoryAsMessageStrings(c)
	title := fmt.Sprintf("%s (%d)", c, total)
	if total == 0 {
		chat.SendEmbedMessage(s, channelID, title, "No sounds found for this category.", map[string]string{})
	} else {
		for _, msg := range msgs {
			chat.SendEmbedMessage(s, channelID, title, msg, map[string]string{})
		}
	}
}

func sendSoundsForCategoryForSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate, c string) {
	msgs, total := getSoundsForCategoryAsMessageStrings(c)
	if total == 0 {
		chat.SendSimpleInteractionEmbedForAction(s, i, c, "No sounds found for this category.", nil)
	} else {
		chat.SendSimpleInteractionEmbedsForAction(s, i, fmt.Sprintf("%s (%d)", c, total), msgs, nil)
	}
}

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

func followOnMove(b *Bot, s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
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
