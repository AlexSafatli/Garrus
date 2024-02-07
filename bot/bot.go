package bot

import (
	"fmt"
	"time"

	"github.com/AlexSafatli/Garrus/chat"
	"github.com/bwmarrin/discordgo"
)

// Bot encompasses a DiscordGo Bot
type Bot struct {
	Start               time.Time
	SlashCommands       map[string]*chat.SlashCommand
	mainGuildChannelIDs map[string]string
	*discordgo.Session
}

// NewBot creates a new Bot
func NewBot(token string) (b *Bot, err error) {
	discord, err := discordgo.New(token)
	if err != nil {
		return
	}
	b = &Bot{Start: time.Now(), Session: discord}
	b.mainGuildChannelIDs = make(map[string]string)
	b.initCommands()
	b.routeHandlers()
	return
}

func (b *Bot) initCommands() {
	b.SlashCommands = map[string]*chat.SlashCommand{
		"about": {
			Command: &discordgo.ApplicationCommand{
				Name:        "about",
				Description: "About this bot",
			},
			Function: AboutCommand,
		},
		"roll": {
			Command: &discordgo.ApplicationCommand{
				Name:        "roll",
				Description: "Roll RPG dice",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "formula",
						Description: "A dice formula string",
						Required:    true,
					},
				},
			},
			Function: RollDiceCommand,
		},
	}
}

func (b *Bot) routeHandlers() {
	b.AddHandler(chat.NewSlashCommandRouteHandler(b.Session, b.SlashCommands))
	b.AddHandler(OnGuildChannelCreateHandler(b))
}

func (b *Bot) RegisterSlashCommands() error {
	var s []*discordgo.ApplicationCommand
	for _, c := range b.SlashCommands {
		s = append(s, c.Command)
	}
	_, err := b.Session.ApplicationCommandBulkOverwrite(b.Session.State.User.ID, "", s)
	return err
}

// Self returns the User struct associated with the bot user
func (b *Bot) Self() (*discordgo.User, error) {
	return b.User("@me")
}

func (b *Bot) String() string {
	return fmt.Sprintf("Bot[%s] - started at %v", b.Token, b.Start)
}
