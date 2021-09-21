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
)

const (
	separator = " - "
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
