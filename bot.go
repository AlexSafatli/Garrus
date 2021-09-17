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
	Start time.Time
	*discordgo.Session
}

// NewBot creates a new Bot
func NewBot(token string) (b *Bot, err error) {
	discord, err := discordgo.New(token)
	if err != nil {
		return
	}
	b = &Bot{Start: time.Now(), Session: discord}
	b.routeHandlers()
	return
}

func (b *Bot) routeHandlers() {
	b.AddHandler(chat.NewHandler(b.Session, ".about", commands.AboutCommand))
	b.AddHandler(chat.DiceRollHandler)
}

// Self returns the User struct associated with the bot user
func (b *Bot) Self() (*discordgo.User, error) {
	return b.User("@me")
}

func (b *Bot) String() string {
	return fmt.Sprintf("Bot[%s] - started at %v", b.Token, b.Start)
}
