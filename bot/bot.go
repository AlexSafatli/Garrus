package bot

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Bot encompasses a DiscordGo Bot
type Bot struct {
	Start time.Time
	*discordgo.Session
}

// NewBot creates a new Bot
func NewBot(token string) (*Bot, error) {
	discord, err := discordgo.New(token)
	if err != nil {
		return nil, err
	}
	return &Bot{Start: time.Now(), Session: discord}, nil
}

func (b *Bot) String() string {
	return fmt.Sprintf("Bot[%v] - started at %v", b.Token, b.Start)
}
