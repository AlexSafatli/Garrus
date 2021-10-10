package chat

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/sound"
	"github.com/AlexSafatli/Garrus/version"
	"log"

	"github.com/bwmarrin/discordgo"
)

func makeEmbed(title string, description string, fields map[string]string) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{Title: title, Description: description, Color: discordColorNavy}
	if fields != nil {
		for _, embedField := range makeMessageEmbedFieldSlice(fields) {
			embed.Fields = append(embed.Fields, embedField)
		}
	}
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: version.Version.Name + " " + version.Version.Version + Separator + version.Version.Developer,
	}
	return embed
}

func makeWarningEmbed(title string, description string, fields map[string]string) *discordgo.MessageEmbed {
	embed := makeEmbed(title, description, fields)
	embed.Color = discordColorOrange
	return embed
}

func makeErrorEmbed(title string, err error) *discordgo.MessageEmbed {
	embed := makeEmbed(title, err.Error(), map[string]string{})
	embed.Color = discordColorRed
	return embed
}

func makeMessageEmbedFieldSlice(vals map[string]string) (arr []*discordgo.MessageEmbedField) {
	for k, v := range vals {
		arr = append(arr, &discordgo.MessageEmbedField{Name: k, Value: v})
	}
	return
}

func SendEmbedMessage(s *discordgo.Session, channelId string, title string, description string, fields map[string]string) *discordgo.Message {
	embed := makeEmbed(title, description, fields)
	msg, err := s.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
}

func SendSimpleEmbedMessage(s *discordgo.Session, channelId string, title string, description string) *discordgo.Message {
	return SendEmbedMessage(s, channelId, title, description, map[string]string{})
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

func SendWelcomeEmbedMessage(s *discordgo.Session, channelId string, user *discordgo.User, soundInfo string) *discordgo.Message {
	var entrance *sound.Entrance
	var title, desc string
	entrance = sound.GetEntranceForUser(user.ID)
	if entrance == nil {
		return nil
	}
	if len(entrance.PersonalizedMessage) > 0 {
		title = entrance.PersonalizedMessage + " **" + user.Username + "**"
	} else {
		title = "Welcome **" + user.Username + "**!"
	}
	if len(soundInfo) > 0 {
		desc = soundInfo + Separator + user.Mention()
	} else {
		desc = user.Mention()
	}
	e := makeEmbed(title, desc, map[string]string{
		RandomString(Whats): "I play sounds and automate things.",
	})
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: user.AvatarURL("2048"),
	}
	msg, err := s.ChannelMessageSendEmbed(channelId, e)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
}

func GetRawAboutEmbedMessage(s *discordgo.Session) *discordgo.MessageEmbed {
	var desc = fmt.Sprintf("My name is Garrus Vakarian and this is my rectum. I am a bot created by %s.", version.Version.Developer)
	e := makeEmbed("My Name Is "+version.Version.Name, desc, map[string]string{
		"Usage":             "Most of my commands are located in slash commands (start typing with a `/` to see them). Some older commands are still found with the `.` prefix such as `.entrance`.",
		RandomString(Whats): "I play sounds and automate things.",
	})
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: s.State.User.AvatarURL("2048"),
	}
	return e
}
