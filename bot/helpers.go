package bot

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func searchSounds(query string) (possibilities []string) {
	closest := sound.GetLibrary().GetClosestMatchingSoundID(query)
	if len(closest) > 0 {
		possibilities = append(possibilities, closest)
	}
	for _, name := range sound.GetLibrary().GetSoundNames() {
		if strings.Contains(name, query) && name != closest {
			possibilities = append(possibilities, closest)
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
	title := fmt.Sprintf("%s (%d)", c, total)
	if total == 0 {
		chat.SendEmbedInteractionResponse(s, i, title, "No sounds found for this category.", map[string]string{})
	} else {
		for _, msg := range msgs {
			chat.SendEmbedInteractionResponse(s, i, title, msg, map[string]string{})
		}
	}
}
