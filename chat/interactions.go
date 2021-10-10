package chat

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func sendInteractionEmbedHeader(s *discordgo.Session, i *discordgo.InteractionCreate, respType discordgo.InteractionResponseType, respData *discordgo.InteractionResponseData, err error) {
	var errResponse error
	if err != nil {
		errResponse = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Something went wrong: " + err.Error(),
			},
		})
	} else {
		errResponse = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: respType,
			Data: respData,
		})
	}
	if errResponse != nil {
		_, _ = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong with sending a response",
		})
	}
}

// DeleteInteractionResponse deletes a response sent from an interaction
func DeleteInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate) {
	go func() {
		time.Sleep(time.Second * 10)
		_ = s.InteractionResponseDelete(s.State.User.ID, i.Interaction)
	}()
}

// EditInteractionResponse edits the original response of an interaction
func EditInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	_, _ = s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
		Content: content,
	})
}

// SendInteractionResponseForAction sends a response to an interaction
func SendInteractionResponseForAction(s *discordgo.Session, i *discordgo.InteractionCreate, payload string, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Content: payload}, err)
	DeleteInteractionResponse(s, i)
}

// SendSimpleInteractionEmbedForAction sends an embed response for an interaction with no fields
func SendSimpleInteractionEmbedForAction(s *discordgo.Session, i *discordgo.InteractionCreate, title, description string, err error) {
	SendInteractionEmbedForAction(s, i, title, description, map[string]string{}, err)
}

// SendInteractionEmbedForAction sends an embed response to an interaction
func SendInteractionEmbedForAction(s *discordgo.Session, i *discordgo.InteractionCreate, title, description string, fields map[string]string, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{
			makeEmbed(title, description, fields),
		}}, err)
	DeleteInteractionResponse(s, i)
}

// SendSimpleInteractionEmbedsForAction sends multiple embed responses for an interaction
func SendSimpleInteractionEmbedsForAction(s *discordgo.Session, i *discordgo.InteractionCreate, title string, descriptions []string, err error) {
	var embeds []*discordgo.MessageEmbed
	for i := range descriptions {
		embeds = append(embeds, makeEmbed(title, descriptions[i], map[string]string{}))
	}
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Embeds: embeds}, err)
	DeleteInteractionResponse(s, i)
}

// SendInteractionRawEmbedForAction sends a raw embed response to an interaction
func SendInteractionRawEmbedForAction(s *discordgo.Session, i *discordgo.InteractionCreate, embed *discordgo.MessageEmbed, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{embed}}, err)
	DeleteInteractionResponse(s, i)
}

// SendInteractionRawEmbedsForAction sends a raw embed response to an interaction
func SendInteractionRawEmbedsForAction(s *discordgo.Session, i *discordgo.InteractionCreate, embeds []*discordgo.MessageEmbed, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Embeds: embeds}, err)
	DeleteInteractionResponse(s, i)
}

// SendWarningInteractionEmbedForAction sends a warning embed response for an interaction with no fields
func SendWarningInteractionEmbedForAction(s *discordgo.Session, i *discordgo.InteractionCreate, title, description string, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{
			makeWarningEmbed(title, description, map[string]string{}),
		}}, err)
	DeleteInteractionResponse(s, i)
}

// SendInteractionAckForAction sends a deferred interaction acknowledgment for an action
func SendInteractionAckForAction(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseDeferredChannelMessageWithSource,
		&discordgo.InteractionResponseData{}, err)
}
