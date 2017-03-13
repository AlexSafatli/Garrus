package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) SendEmbedMessage(channelId string, title string, description string, fields []*discordgo.MessageEmbedField) *discordgo.Message {
	embed := &discordgo.MessageEmbed{Title: title, Description: description, Color: discordColorDarkNavy}
	for i := range fields {
		embed.Fields = append(embed.Fields, fields[i])
	}
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: Version.Name + " " + Version.Version + " – " + Version.Developer,
	}
	msg, err := b.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
}

func MakeMessageEmbedFieldSlice(vals map[string]string) (arr []*discordgo.MessageEmbedField) {
	for k, v := range vals {
		arr = append(arr, &discordgo.MessageEmbedField{Name: k, Value: v})
	}
	return
}
