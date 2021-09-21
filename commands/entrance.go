package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

// SetEntranceMessageCommand sets an entrance for a user
func SetEntranceMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Printf("About handler running for message from %s", m.Author.Username)
}

// SetEntranceSlashCommand sets an entrance for a user
func SetEntranceSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Printf("About handler running for message from %s", i.Member.User.Username)
}
