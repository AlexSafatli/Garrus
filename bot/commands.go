package bot

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/rpg"
	"github.com/AlexSafatli/Garrus/rpg/dice"
	"github.com/bwmarrin/discordgo"
)

// AboutCommand shows an about embed for the bot
func AboutCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	chat.SendInteractionRawEmbedForAction(s, i, chat.GetRawAboutEmbedMessage(s), nil)
}

// RollDiceCommand rolls an RPG dice query string
func RollDiceCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var formula string
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		formula = args[0].StringValue()
	}
	d, err := rpg.Roll(dice.Formula(formula))
	if err != nil {
		chat.SendErrorInteractionEmbedForAction(s, i, "Roll Dice", err)
	} else {
		chat.SendSimpleInteractionEmbedsForAction(s, i, "Roll Dice",
			[]string{fmt.Sprintf("You rolled `%s` and got %d.", formula, d.Result)},
			err)
	}
}
