package chat

import (
	"github.com/bwmarrin/discordgo"
)

type SlashCommand struct {
	Command  *discordgo.ApplicationCommand
	Function func(*discordgo.Session, *discordgo.InteractionCreate)
}

// NewSlashCommandRouteHandler is a generic handler function that calls a corresponding command function by name for slash commands
func NewSlashCommandRouteHandler(s *discordgo.Session, cmds map[string]*SlashCommand) func(*discordgo.Session, *discordgo.InteractionCreate) {
	if s == nil {
		return nil
	}
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if c, ok := cmds[i.ApplicationCommandData().Name]; ok {
			c.Function(s, i)
		}
	}
}
