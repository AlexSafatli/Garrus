package chat

import (
	"github.com/bwmarrin/discordgo"
	"log"
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

	discordCustomColorRoseGold      = 15648450
	discordCustomColorWarmPink      = 16734337
	discordCustomColorMilkWhite     = 16643563
	discordCustomColorWisteria      = 10976706
	discordCustomColorPaleGoldenrod = 16114856

	botMessageSearchLimit = 50
)

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
