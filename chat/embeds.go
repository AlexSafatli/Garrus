package chat

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/version"
	"log"

	"github.com/bwmarrin/discordgo"
)

func makeEmbed(title string, description string, fields map[string]string, color int) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{Title: title, Description: description, Color: color}
	if fields != nil {
		for _, embedField := range makeMessageEmbedFieldSlice(fields) {
			embed.Fields = append(embed.Fields, embedField)
		}
	}
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: version.Version.Name + " " + version.Version.Version + "+" + version.Version.GitCommit + Separator + version.Version.Developer,
	}
	return embed
}

func makeWarningEmbed(title string, description string, fields map[string]string) *discordgo.MessageEmbed {
	return makeEmbed(title, description, fields, discordColorOrange)
}

func makeErrorEmbed(title string, err error) *discordgo.MessageEmbed {
	return makeEmbed(title, err.Error(), map[string]string{}, discordColorRed)
}

func makeMessageEmbedFieldSlice(vals map[string]string) (arr []*discordgo.MessageEmbedField) {
	for k, v := range vals {
		arr = append(arr, &discordgo.MessageEmbedField{Name: k, Value: v})
	}
	return
}

func SendEmbedMessage(s *discordgo.Session, channelId string, title string, description string, fields map[string]string, color int) *discordgo.Message {
	embed := makeEmbed(title, description, fields, color)
	msg, err := s.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
}

func SendSimpleEmbedMessage(s *discordgo.Session, channelId string, title string, description string) *discordgo.Message {
	return SendEmbedMessage(s, channelId, title, description, map[string]string{}, discordCustomColorMilkWhite)
}

func SendWarningEmbedMessage(s *discordgo.Session, channelId, title, warning string) *discordgo.Message {
	embed := makeWarningEmbed(title, warning, map[string]string{})
	msg, err := s.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
}

func SendErrorEmbedMessage(s *discordgo.Session, channelId string, title string, err error) *discordgo.Message {
	embed := makeErrorEmbed(title, err)
	msg, err := s.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
}

func SendRawEmbedMessage(s *discordgo.Session, channelId string, e *discordgo.MessageEmbed) *discordgo.Message {
	msg, err := s.ChannelMessageSendEmbed(channelId, e)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
}

func GetRawAboutEmbedMessage(s *discordgo.Session) *discordgo.MessageEmbed {
	var desc = fmt.Sprintf("My name is Garrus Vakarian. I am a bot created by %s.", version.Version.Developer)
	e := makeEmbed("My Name Is "+version.Version.Name, desc, map[string]string{
		"Usage":             "Start typing with a `/` to see my commands.",
		RandomString(Whats): "I automate and organize things.",
	}, discordCustomColorMilkWhite)
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: s.State.User.AvatarURL("2048"),
	}
	return e
}
