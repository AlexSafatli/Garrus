package bot

import (
	"fmt"
	"regexp"

	"github.com/AlexSafatli/DiscordSwissArmyKnife/rpg"
	"github.com/bwmarrin/discordgo"
)

var dieStrMatcher = regexp.MustCompile(`\[\[[^\[^\]]*\]\]`)

// DiceRollHandler takes a created message and edits it with the result of a dice roll (if it matches a `[[xdy]]` pattern)
func DiceRollHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if dieStrMatcher.MatchString(m.Content) {
		editedStr := m.Content
		dieStrMatcher.ReplaceAllStringFunc(editedStr, func(s string) string {
			dieStr := s[2 : len(s)-2]
			return fmt.Sprintf("**%d** ([`%s`])", rpg.Roll(dieStr), dieStr)
		})
		_, err := s.ChannelMessageEdit(m.ChannelID, m.ID, editedStr)
		if err != nil {
			// TODO Add log message for error.
		}
	}
}
