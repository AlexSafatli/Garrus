package chat

import (
	"github.com/bwmarrin/discordgo"
)
import "strings"

type MessageCommand struct {
	Command  string
	Function func(*discordgo.Session, *discordgo.MessageCreate)
}

type SlashCommand struct {
	Command  *discordgo.ApplicationCommand
	Function func(*discordgo.Session, *discordgo.InteractionCreate)
}

// NewMessageCommandRouteHandler is a generic handler function that calls a corresponding command function by name
func NewMessageCommandRouteHandler(s *discordgo.Session, cmds []*MessageCommand) func(*discordgo.Session, *discordgo.MessageCreate) {
	if s == nil {
		return nil
	}
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		for _, c := range cmds {
			if strings.HasPrefix(m.Content, c.Command) {
				DeleteReceivedMessage(s, m)
				c.Function(s, m)
				return
			}
		}
	}
}

// NewSlashCommandRouteHandler is a generic handler function that calls a corresponding command function by name for slash commands
func NewSlashCommandRouteHandler(s *discordgo.Session, cmds []*SlashCommand) func(*discordgo.Session, *discordgo.InteractionCreate) {
	if s == nil {
		return nil
	}
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		for _, c := range cmds {
			if i.ApplicationCommandData().Name == c.Command.Name {
				c.Function(s, i)
				return
			}
		}
	}
}
