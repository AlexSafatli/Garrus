package bot

import (
	"fmt"
	"time"

	"github.com/AlexSafatli/Garrus/chat"
	"github.com/bwmarrin/discordgo"
)

// Bot encompasses a DiscordGo Bot
type Bot struct {
	Start            time.Time
	MessageCommands  []*chat.MessageCommand
	SlashCommands    []*chat.SlashCommand
	VoiceConnections map[string]*discordgo.VoiceConnection
	*discordgo.Session
}

// NewBot creates a new Bot
func NewBot(token string) (b *Bot, err error) {
	discord, err := discordgo.New(token)
	if err != nil {
		return
	}
	b = &Bot{Start: time.Now(), Session: discord}
	b.initMessageCommands()
	b.initSlashCommands()
	b.routeHandlers()
	return
}

func (b *Bot) initMessageCommands() {
	b.MessageCommands = []*chat.MessageCommand{
		{
			Command:  ".about",
			Function: AboutMessageCommand,
		},
		{
			Command:  ".entrance",
			Function: SetEntranceMessageCommand,
		},
		{
			Command:  ".search",
			Function: SearchSoundsMessageCommand,
		},
	}
}

func (b *Bot) initSlashCommands() {
	b.SlashCommands = []*chat.SlashCommand{
		{
			Command: &discordgo.ApplicationCommand{
				Name:        "about",
				Description: "About this bot",
			},
			Function: AboutSlashCommand,
		},
		{
			Command: &discordgo.ApplicationCommand{
				Name:        "entrance",
				Description: "Set your entrance sound when joining a voice channel (or clear it if nothing is specified)",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "sound",
						Description: "The name of the sound",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "message",
						Description: "A personalized welcome message",
						Required:    false,
					},
				},
			},
			Function: SetEntranceSlashCommand,
		},
		{
			Command: &discordgo.ApplicationCommand{
				Name:        "searchfor",
				Description: "Lists all sound files matching a query",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "query",
						Description: "A search query",
						Required:    true,
					},
				},
			},
			Function: SearchSoundsSlashCommand,
		},
		{
			Command: &discordgo.ApplicationCommand{
				Name:        "sound",
				Description: "Play a sound file by name",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "sound",
						Description: "A sound file",
						Required:    true,
					},
				},
			},
		},
	}
}

func (b *Bot) routeHandlers() {
	b.AddHandler(chat.NewMessageCommandRouteHandler(b.Session, b.MessageCommands))
	b.AddHandler(chat.NewSlashCommandRouteHandler(b.Session, b.SlashCommands))
	b.AddHandler(OnPlayMessageCommandReceivedHandler(b))
	b.AddHandler(OnPlaySlashCommandReceivedHandler(b))
	b.AddHandler(OnGuildVoiceJoinHandler(b))
}

func (b *Bot) RegisterSlashCommands() error {
	for _, c := range b.SlashCommands {
		_, err := b.Session.ApplicationCommandCreate(b.Session.State.User.ID, "", c.Command)
		if err != nil {
			return err
		}
	}
	return nil
}

// Self returns the User struct associated with the bot user
func (b *Bot) Self() (*discordgo.User, error) {
	return b.User("@me")
}

func (b *Bot) String() string {
	return fmt.Sprintf("Bot[%s] - started at %v", b.Token, b.Start)
}
