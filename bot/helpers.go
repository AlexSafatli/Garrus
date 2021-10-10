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
