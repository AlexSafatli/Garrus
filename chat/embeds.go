package chat

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func SendEmbedMessage(s *discordgo.Session, channelId string, title string, description string, fields map[string]string) *discordgo.Message {
	embed := makeEmbed(title, description, fields)
	msg, err := s.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
}

func SendWelcomeEmbedMessage(s *discordgo.Session, channelId string, user *discordgo.User, soundInfo string) *discordgo.Message {
	var title, desc string // randomWelcome + "**" + user.GetName() + "**"
	if len(soundInfo) > 0 {
		desc = soundInfo + separator + user.Mention()
	} else {
		desc = user.Mention()
	}
	// TODO add image avatar to embed
	return nil
}

func makeEmbed(title string, description string, fields map[string]string) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{Title: title, Description: description, Color: discordColorNavy}
	if fields != nil {
		for _, embedField := range makeMessageEmbedFieldSlice(fields) {
			embed.Fields = append(embed.Fields, embedField)
		}
	}
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: Version.Name + " " + Version.Version + separator + Version.Developer,
	}
	return embed
}

func makeMessageEmbedFieldSlice(vals map[string]string) (arr []*discordgo.MessageEmbedField) {
	for k, v := range vals {
		arr = append(arr, &discordgo.MessageEmbedField{Name: k, Value: v})
	}
	return
}
