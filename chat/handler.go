package chat

import (
	"github.com/bwmarrin/discordgo"
	"log"
)
import "strings"

func NewHandler(s *discordgo.Session, command string, commandFunction func(*discordgo.Session, *discordgo.MessageCreate)) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.HasPrefix(m.Content, command) {
			commandFunction(s, m)
			if m.GuildID != "" {
				err := s.ChannelMessageDelete(m.ChannelID, m.ID)
				if err != nil {
					log.Printf("Could not remove message from %s (%s, %s)", m.Author.Username, m.ChannelID, m.ID)
				}
			}
		}
	}
}
