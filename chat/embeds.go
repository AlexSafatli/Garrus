package chat

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func SendEmbedMessage(s *discordgo.Session, channelId string, title string, description string, fields map[string]string) *discordgo.Message {
	embed := &discordgo.MessageEmbed{Title: title, Description: description, Color: discordColorNavy}
	if fields != nil {
		for _, embedField := range makeMessageEmbedFieldSlice(fields) {
			embed.Fields = append(embed.Fields, embedField)
		}
	}
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: Version.Name + " " + Version.Version + " – " + Version.Developer,
	}
	msg, err := s.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
}

func makeMessageEmbedFieldSlice(vals map[string]string) (arr []*discordgo.MessageEmbedField) {
	for k, v := range vals {
		arr = append(arr, &discordgo.MessageEmbedField{Name: k, Value: v})
	}
	return
}
