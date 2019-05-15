package chat

import (
	"fmt"
	"log"
	"regexp"

	"../rpg"
	"github.com/bwmarrin/discordgo"
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

var dieStrMatcher = regexp.MustCompile(`(\[\[[^\[^\]]*]])`)

// DiceRollHandler takes a created message and returns dice roll results (if a roll matches a `[[...]]` pattern)
func DiceRollHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if dieStrMatcher.MatchString(m.Content) {
		log.Printf("Found dice roll(s) to handle in %s", m.Content)
		matches := dieStrMatcher.FindAllString(m.Content, -1)
		for _, match := range matches {
			dieStr := match[2 : len(match)-2]
			msgToSend := fmt.Sprintf("**%d** ([`%s`]) *was rolled by* %s", rpg.Roll(dieStr), dieStr, m.Author.Username)
			_ = SendEmbedMessage(s, m.ChannelID, fmt.Sprintf("%s is Rolling Dice", m.Author.Username), msgToSend, nil)
		}
		DeleteReceivedMessage(s, m)
	}
}

// DeleteReceivedMessage takes a created message and deletes it if it is not private
func DeleteReceivedMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.GuildID != "" {
		err := s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			log.Printf("Could not delete message from %s", m.Author.Username)
		}
	}
}
