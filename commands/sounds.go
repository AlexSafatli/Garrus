package commands

import (
	"errors"
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	searchSoundsTitle = "Search Sounds"
)

func searchSounds(query string) (possibilities []string) {
	closest := sound.GetLibrary().GetClosestMatchingSoundID(query)
	if len(closest) > 0 {
		possibilities = append(possibilities, closest)
	}
	for _, name := range sound.GetLibrary().GetSoundNames() {
		if strings.Contains(name, query) && name != closest {
			possibilities = append(possibilities, closest)
		}
	}
	return
}

// SearchSoundsMessageCommand returns the collection of sounds matching a query
func SearchSoundsMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	var query string
	var possibilities []string
	i := strings.Index(m.Content, " ")
	if i > 0 {
		query = m.Content[i+1:]
	}
	possibilities = searchSounds(query)
	mb := chat.MessageBuilder{}
	_ = mb.Write(fmt.Sprintf("Found **%d** possible sounds f or query `%s`%s%s.\n\n", len(possibilities), query, chat.Separator, m.Author.Mention()))
	if len(possibilities) > 0 {
		for _, k := range possibilities {
			_ = mb.Write("`?" + k + "` ")
		}
		for _, msg := range mb.GetMessageStrings() {
			chat.SendEmbedMessage(s, m.ChannelID, searchSoundsTitle, msg, map[string]string{})
		}
	} else {
		chat.SendWarningEmbedMessage(s, m.ChannelID, searchSoundsTitle, "Could not find any sounds for query `"+query+"`.")
	}
}

// SearchSoundsSlashCommand returns the collection of sounds matching a query
func SearchSoundsSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var query string
	var possibilities []string
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		query = args[0].StringValue()
	} else {
		chat.SendErrorEmbedMessage(s, i.ChannelID, searchSoundsTitle, errors.New("need a query to search for"))
		return
	}
	if err := chat.SendInteractionAckForAction(s, i, nil); err != nil {
		return
	}
	defer chat.DeleteInteractionResponse(s, i)
	possibilities = searchSounds(query)
	mb := chat.MessageBuilder{}
	_ = mb.Write(fmt.Sprintf("Found **%d** possible sounds f or query `%s`%s%s.\n\n", len(possibilities), query, chat.Separator, i.User.Mention()))
	if len(possibilities) > 0 {
		for _, k := range possibilities {
			_ = mb.Write("`?" + k + "` ")
		}
		for _, msg := range mb.GetMessageStrings() {
			chat.SendEmbedInteractionResponse(s, i, searchSoundsTitle, msg, map[string]string{})
		}
	} else {
		chat.SendWarningEmbedInteractionResponse(s, i, searchSoundsTitle, "Could not find any sounds for query `"+query+"`.")
	}
}
