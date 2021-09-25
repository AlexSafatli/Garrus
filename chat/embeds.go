package chat

import (
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

func SendWarningEmbedMessage(s *discordgo.Session, channelId, title, warning string) *discordgo.Message {
	embed := makeEmbed(title, warning, map[string]string{})
	embed.Color = discordColorOrange
	msg, err := s.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
}

func SendErrorEmbedMessage(s *discordgo.Session, channelId string, title string, err error) *discordgo.Message {
	embed := makeEmbed(title, err.Error(), map[string]string{})
	embed.Color = discordColorRed
	msg, err := s.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		log.Println("When sending embed in channel", channelId, "ran into error =>", err)
	}
	return msg
}

func SendEmbedInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate, title string, description string, fields map[string]string) *discordgo.Message {
	embed := makeEmbed(title, description, fields)
	m, _ := s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{embed},
	})
	return m
}

func SendWarningEmbedInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate, title, warning string) *discordgo.Message {
	embed := makeEmbed(title, warning, map[string]string{})
	embed.Color = discordColorOrange
	m, _ := s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{embed},
	})
	return m
}

func SendErrorEmbedInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate, title string, err error) *discordgo.Message {
	embed := makeEmbed(title, err.Error(), map[string]string{})
	embed.Color = discordColorRed
	m, _ := s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{embed},
	})
	return m
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
		RandomString(Whats): "I play sounds and automate things. Type `.help` for more.",
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

func makeEmbed(title string, description string, fields map[string]string) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{Title: title, Description: description, Color: discordColorNavy}
	if fields != nil {
		for _, embedField := range makeMessageEmbedFieldSlice(fields) {
			embed.Fields = append(embed.Fields, embedField)
		}
	}
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: Version.Name + " " + Version.Version + Separator + Version.Developer,
	}
	return embed
}

func makeMessageEmbedFieldSlice(vals map[string]string) (arr []*discordgo.MessageEmbedField) {
	for k, v := range vals {
		arr = append(arr, &discordgo.MessageEmbedField{Name: k, Value: v})
	}
	return
}
