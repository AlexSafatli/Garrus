package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// AboutMessageCommand takes a created message and returns an About embed message
func AboutMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Printf("About handler running for message from %s", m.Author.Username)
}

// AboutSlashCommand returns an About embed message for a slash command
func AboutSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Printf("About handler running for message from %s", i.Member.User.Username)
}
