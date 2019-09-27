package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// AboutCommand takes a created message and returns an About embed message
func AboutCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Printf("About handler running for message from %s", m.Author.Username)
}
