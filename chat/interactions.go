package chat

import (
	"github.com/bwmarrin/discordgo"
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
		_, _ = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong with sending a response",
		})
	}
}

// EditInteractionResponse edits the original response of an interaction
func EditInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	_, _ = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &content,
	})
}

// SendInteractionResponseForAction sends a response to an interaction
func SendInteractionResponseForAction(s *discordgo.Session, i *discordgo.InteractionCreate, payload string, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Content: payload}, err)
}

// SendSimpleInteractionEmbedForAction sends an embed response for an interaction with no fields
func SendSimpleInteractionEmbedForAction(s *discordgo.Session, i *discordgo.InteractionCreate, title, description string, err error) {
	SendInteractionEmbedForAction(s, i, title, description, map[string]string{}, discordCustomColorMilkWhite, err)
}

// SendInteractionEmbedForAction sends an embed response to an interaction
func SendInteractionEmbedForAction(s *discordgo.Session, i *discordgo.InteractionCreate, title, description string, fields map[string]string, color int, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{
			makeEmbed(title, description, fields, color),
		}}, err)
}

// SendSimpleInteractionEmbedsForAction sends multiple embed responses for an interaction
func SendSimpleInteractionEmbedsForAction(s *discordgo.Session, i *discordgo.InteractionCreate, title string, descriptions []string, err error) {
	var embeds []*discordgo.MessageEmbed
	for i := range descriptions {
		embeds = append(embeds, makeEmbed(title, descriptions[i], map[string]string{}, discordCustomColorMilkWhite))
	}
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Embeds: embeds}, err)
}

// SendInteractionRawEmbedForAction sends a raw embed response to an interaction
func SendInteractionRawEmbedForAction(s *discordgo.Session, i *discordgo.InteractionCreate, embed *discordgo.MessageEmbed, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{embed}}, err)
}

// SendInteractionRawEmbedsForAction sends a raw embed response to an interaction
func SendInteractionRawEmbedsForAction(s *discordgo.Session, i *discordgo.InteractionCreate, embeds []*discordgo.MessageEmbed, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Embeds: embeds}, err)
}

// SendWarningInteractionEmbedForAction sends a warning embed response for an interaction with no fields
func SendWarningInteractionEmbedForAction(s *discordgo.Session, i *discordgo.InteractionCreate, title, description string, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{
			makeWarningEmbed(title, description, map[string]string{}),
		}}, err)
}

// SendErrorInteractionEmbedForAction sends an error embed response for an interaction
func SendErrorInteractionEmbedForAction(s *discordgo.Session, i *discordgo.InteractionCreate, title string, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseChannelMessageWithSource,
		&discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{
			makeErrorEmbed(title, err),
		}}, err)
}

// SendInteractionAckForAction sends a deferred interaction acknowledgment for an action
func SendInteractionAckForAction(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
	sendInteractionEmbedHeader(s, i,
		discordgo.InteractionResponseDeferredChannelMessageWithSource,
		&discordgo.InteractionResponseData{}, err)
}
