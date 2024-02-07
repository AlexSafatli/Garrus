package bot

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/bwmarrin/discordgo"
)

// AboutCommand shows an about embed for the bot
func AboutCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	chat.SendInteractionRawEmbedForAction(s, i, chat.GetRawAboutEmbedMessage(s), nil)
}

// RollDiceCommand rolls an RPG dice query string
func RollDiceCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var formula string
	var possibilities []string
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		formula = args[0].StringValue()
	}
	result = rollDice(formula)
	mb := chat.MessageBuilder{}
	_ = mb.Write(fmt.Sprintf("Found **%d** possible sounds for query `%s`%s%s.\n\n", len(possibilities), query, chat.Separator, i.Member.Mention()))
	if len(possibilities) > 0 {
		for _, k := range possibilities {
			_ = mb.Write("`?" + k + "` ")
		}
		chat.SendSimpleInteractionEmbedsForAction(s, i, searchSoundsTitle, mb.GetMessageStrings(), nil)
	} else {
		chat.SendWarningInteractionEmbedForAction(s, i, searchSoundsTitle, "Could not find any sounds for query `"+query+"`.", nil)
	}
}
