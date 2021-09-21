package chat

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/sound"
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
	var entrance *sound.Entrance
	var title, desc string
	entrance = sound.GetEntranceForUser(user.ID)
	if entrance == nil {
		return nil
	}
	title = entrance.PersonalizedMessage + " **" + user.Username + "**"
	if len(soundInfo) > 0 {
		desc = soundInfo + separator + user.Mention()
	} else {
		desc = user.Mention()
	}
	e := makeEmbed(title, desc, map[string]string{
		RandomString(Whats): "I play sounds and automate things. Type `.help` for more.",
	})
	e.Color = ToColor(user.Username)
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:    user.AvatarURL("8"),
		Width:  256,
		Height: 256,
	}
	e.Footer.Text += separator + fmt.Sprintf("%d", e.Color)
	msg, err := s.ChannelMessageSendEmbed(channelId, e)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
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
