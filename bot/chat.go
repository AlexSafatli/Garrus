package bot

import (
	"fmt"
	"regexp"

	"github.com/AlexSafatli/DiscordSwissArmyKnife/rpg"
	"github.com/bwmarrin/discordgo"
)

var dieStrMatcher = regexp.MustCompile(`(\[\[[^\[^\]]*\]\])`)

// DiceRollHandler takes a created message and edits it with the result of a dice roll (if it matches a `[[xdy]]` pattern)
func DiceRollHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if dieStrMatcher.MatchString(m.Content) {
		editedStr := dieStrMatcher.ReplaceAllStringFunc(m.Content, func(s string) string {
			dieStr := s[2 : len(s)-2]
			return fmt.Sprintf("**%d** ([`%s`])", rpg.Roll(dieStr), dieStr)
		})
		err := s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			// TODO Add log message for error.
		}
		msgToSend := fmt.Sprintf("**%s** sent message:\n%s", m.Author.Username, editedStr)
		_, err = s.ChannelMessageSend(m.ChannelID, msgToSend)
	}
}
