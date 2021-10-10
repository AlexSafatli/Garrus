package chat

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

const (
	// https://github.com/izy521/discord.io/blob/master/docs/colors.md
	discordColorBlack      = 0
	discordColorAqua       = 1752220
	discordColorGreen      = 3066993
	discordColorBlue       = 3447003
	discordColorPurple     = 10181046
	discordColorGold       = 15844367
	discordColorOrange     = 15105570
	discordColorRed        = 15158332
	discordColorGrey       = 9807270
	discordColorDarkerGrey = 8359053
	discordColorNavy       = 3426654
	discordColorDarkAqua   = 1146986
	discordColorDarkGreen  = 2067276
	discordColorDarkBlue   = 2123412
	discordColorDarkPurple = 7419530
	discordColorDarkGold   = 12745742
	discordColorDarkOrange = 11027200
	discordColorDarkRed    = 10038562
	discordColorDarkGrey   = 9936031
	discordColorLightGrey  = 12370112
	discordColorDarkNavy   = 2899536

	botMessageSearchLimit = 50
)

// DeleteReceivedMessage takes a created message and deletes it if it is not private
func DeleteReceivedMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.GuildID != "" {
		err := s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			log.Printf("Could not delete message from %s", m.Author.Username)
		}
	}
}

// DeleteBotMessages deletes all bot messages found in the given channel
func DeleteBotMessages(s *discordgo.Session, channelID, aroundID string) {
	msgs, err := s.ChannelMessages(channelID, botMessageSearchLimit, "", "", aroundID)
	if err != nil {
		log.Printf("Could not find bot messages around message ID %s in channel %s", aroundID, channelID)
		return
	}
	var botMessageIDs []string
	for _, m := range msgs {
		if m.Author.ID == s.State.User.ID {
			botMessageIDs = append(botMessageIDs, m.ID)
		}
	}
	if len(botMessageIDs) > 0 {
		if err = s.ChannelMessagesBulkDelete(channelID, botMessageIDs); err != nil {
			log.Printf("Could not bulk delete all found %d bot messages in channel %s => %s", len(botMessageIDs), channelID, err)
		}
	}
}

// DeleteInteractionResponse deletes a response sent from an interaction
func DeleteInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate) {
	go func() {
		time.Sleep(time.Second * 10)
		_ = s.InteractionResponseDelete(s.State.User.ID, i.Interaction)
	}()
}

// SendInteractionResponseForAction sends a response to an interaction
func SendInteractionResponseForAction(s *discordgo.Session, i *discordgo.InteractionCreate, payload string, err error) {
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
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: payload,
			},
		})
	}
	if errResponse != nil {
		_, _ = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong with sending a response",
		})
	}
	DeleteInteractionResponse(s, i)
}

// SendSimpleInteractionEmbedForAction sends an embed response for an interaction with no fields
func SendSimpleInteractionEmbedForAction(s *discordgo.Session, i *discordgo.InteractionCreate, title, description string, err error) {
	SendInteractionEmbedForAction(s, i, title, description, map[string]string{}, err)
}

// SendInteractionEmbedForAction sends an embed response to an interaction
func SendInteractionEmbedForAction(s *discordgo.Session, i *discordgo.InteractionCreate, title, description string, fields map[string]string, err error) {
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
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		})
	}
	if errResponse != nil {
		_, _ = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong with sending a response",
		})
	}
	defer DeleteInteractionResponse(s, i)
	if err != nil {
		return
	}
	SendEmbedInteractionResponse(s, i, title, description, fields)
}

// SendInteractionAckForAction sends a deferred interaction acknowledgment for an action
func SendInteractionAckForAction(s *discordgo.Session, i *discordgo.InteractionCreate, err error) error {
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
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		})
	}
	if errResponse != nil {
		_, _ = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong with sending a response",
		})
	}
	return errResponse
}

// SendSimpleMessageResponseForAction sends a response for a message command with no fields
func SendSimpleMessageResponseForAction(s *discordgo.Session, channelID, title, payload string, err error) {
	SendMessageResponseForAction(s, channelID, title, payload, map[string]string{}, err)
}

// SendMessageResponseForAction sends a response for a message command
func SendMessageResponseForAction(s *discordgo.Session, channelID, title, payload string, fields map[string]string, err error) {
	var msg *discordgo.Message
	if err != nil {
		msg = SendErrorEmbedMessage(s, channelID, title, err)
	} else {
		msg = SendEmbedMessage(s, channelID, title, payload, fields)
	}
	time.Sleep(time.Second * 10)
	_ = s.ChannelMessageDelete(channelID, msg.ID)
}
