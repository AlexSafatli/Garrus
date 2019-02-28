package chat

import (
	"github.com/bwmarrin/discordgo"
)
import "strings"

func NewHandler(s *discordgo.Session, command string, commandFunction func(*discordgo.Session, *discordgo.MessageCreate)) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.HasPrefix(m.Content, command) {
			commandFunction(s, m)
			DeleteReceivedMessage(s, m)
		}
	}
}
