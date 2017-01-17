package bot

import (
	"fmt"
	"log"
	"regexp"

	"github.com/AlexSafatli/DiscordSwissArmyKnife/rpg"
	"github.com/bwmarrin/discordgo"
)

var dieStrMatcher = regexp.MustCompile(`(\[\[[^\[^\]]*\]\])`)

// DiceRollHandler takes a created message and edits it with the result of a dice roll (if it matches a `[[...]]` pattern)
func DiceRollHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if dieStrMatcher.MatchString(m.Content) {
		matches := dieStrMatcher.FindAllString(m.Content, -1)
		for _, match := range matches {
			dieStr := match[2 : len(match)-2]
			msgToSend := fmt.Sprintf("**%d** ([`%s`]) *rolled by* %s", rpg.Roll(dieStr), dieStr, m.Author.Username)
			_, err := s.ChannelMessageSend(m.ChannelID, msgToSend)
			if err != nil {
				log.Println("Failed to send message with result", msgToSend, "to channel", m.ChannelID)
			}
		}
	}
}
