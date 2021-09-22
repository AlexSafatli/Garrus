package commands

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/bot"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	entranceTitle = "Set Entrance"
)

// SetEntranceMessageCommand sets an entrance for a user
func SetEntranceMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	var soundID, content, pmsg string
	i := strings.Index(m.Content, " ")
	if i > 0 {
		content = m.Content[i+1:]
	} else {
		err := sound.DeleteEntranceForUser(m.Author.ID, bot.Db)
		chat.SendSimpleMessageResponseForAction(s, m.ChannelID, entranceTitle, "Cleared your entrance.", err)
		return
	}
	args := strings.Split(content, ",")
	if len(args) >= 1 {
		soundID = args[0]
	} else if len(args) >= 2 {
		pmsg = args[1]
	}
	err := sound.SetEntranceForUser(m.Author.ID, soundID, pmsg, bot.Db)
	chat.SendSimpleMessageResponseForAction(s, m.ChannelID, entranceTitle, fmt.Sprintf("Set your entrance to `%s`.", soundID), err)
}

// SetEntranceSlashCommand sets an entrance for a user
func SetEntranceSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var soundID, pmsg string
	if len(i.ApplicationCommandData().Options) == 0 {
		err := sound.DeleteEntranceForUser(i.User.ID, bot.Db)
		chat.SendInteractionResponseForAction(s, i, "Cleared your entrance.", err)
		return
	}
	args := i.ApplicationCommandData().Options
	if len(args) >= 1 {
		soundID = args[0].StringValue()
	} else if len(args) >= 2 {
		pmsg = args[1].StringValue()
	}
	err := sound.SetEntranceForUser(i.User.ID, soundID, pmsg, bot.Db)
	chat.SendInteractionResponseForAction(s, i, fmt.Sprintf("Set your entrance to `%s`.", soundID), err)
}
