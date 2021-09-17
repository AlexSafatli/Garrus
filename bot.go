package main

import (
	"fmt"
	"time"

	"github.com/AlexSafatli/DiscordSwissArmyKnife/chat"
	"github.com/AlexSafatli/DiscordSwissArmyKnife/commands"

	"github.com/bwmarrin/discordgo"
)

// Bot encompasses a DiscordGo Bot
type Bot struct {
	Start           time.Time
	MessageCommands []*chat.MessageCommand
	SlashCommands   []*chat.SlashCommand
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
	if err = b.registerSlashCommands(); err != nil {
		return
	}
	return
}

func (b *Bot) initMessageCommands() {
	b.MessageCommands = []*chat.MessageCommand{
		{
			Command:  ".about",
			Function: commands.AboutMessageCommand,
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
			Function: commands.AboutSlashCommand,
		},
	}
}

func (b *Bot) routeHandlers() {
	b.AddHandler(chat.NewMessageCommandRouteHandler(b.Session, b.MessageCommands))
	b.AddHandler(chat.NewSlashCommandRouteHandler(b.Session, b.SlashCommands))
	b.AddHandler(chat.DiceRollHandler)
}

func (b *Bot) registerSlashCommands() error {
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
