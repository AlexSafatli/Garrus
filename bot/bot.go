package bot

import (
	"fmt"
	"time"

	"github.com/AlexSafatli/Garrus/chat"
	"github.com/bwmarrin/discordgo"
)

// Bot encompasses a DiscordGo Bot
type Bot struct {
	Start                   time.Time
	MessageCommands         []*chat.MessageCommand
	SlashCommands           map[string]*chat.SlashCommand
	lastSentEntranceMessage map[string]string
	mainGuildChannelIDs     map[string]string
	*discordgo.Session
}

// NewBot creates a new Bot
func NewBot(token string) (b *Bot, err error) {
	discord, err := discordgo.New(token)
	if err != nil {
		return
	}
	b = &Bot{Start: time.Now(), Session: discord}
	b.lastSentEntranceMessage = make(map[string]string)
	b.mainGuildChannelIDs = make(map[string]string)
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
		{
			Command:  ".list",
			Function: ListSoundsMessageCommand,
		},
		{
			Command:  ".categories",
			Function: ListCategoriesMessageCommand,
		},
		{
			Command:  "?",
			Function: PlaySoundMessageCommand,
		},
	}
}

func (b *Bot) initSlashCommands() {
	b.SlashCommands = map[string]*chat.SlashCommand{
		"about": {
			Command: &discordgo.ApplicationCommand{
				Name:        "about",
				Description: "About this bot",
			},
			Function: AboutSlashCommand,
		},
		"entrance": {
			Command: &discordgo.ApplicationCommand{
				Name:        "entrance",
				Description: "Set your entrance sound when joining a voice channel (or clear it if an empty string is given)",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "sound",
						Description: "The name of the sound",
						Required:    true,
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
		"searchfor": {
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
		"list": {
			Command: &discordgo.ApplicationCommand{
				Name:        "list",
				Description: "Lists all sound files for a category",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "category",
						Description: "A category by name (case-insensitive)",
						Required:    true,
					},
				},
			},
			Function: ListSoundsSlashCommand,
		},
		"categories": {
			Command: &discordgo.ApplicationCommand{
				Name:        "categories",
				Description: "Lists all categories",
			},
			Function: ListCategoriesSlashCommand,
		},
		"random": {
			Command: &discordgo.ApplicationCommand{
				Name:        "random",
				Description: "Play a random sound file (optionally within a category)",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "category",
						Description: "A category",
						Required:    false,
					},
				},
			},
			Function: PlayRandomSoundSlashCommand,
		},
		"sound": {
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
			Function: PlaySoundSlashCommand,
		},
	}
}

func (b *Bot) routeHandlers() {
	b.AddHandler(chat.NewMessageCommandRouteHandler(b.Session, b.MessageCommands))
	b.AddHandler(chat.NewSlashCommandRouteHandler(b.Session, b.SlashCommands))
	b.AddHandler(OnGuildVoiceJoinHandler(b))
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
