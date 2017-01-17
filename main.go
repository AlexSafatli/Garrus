package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AlexSafatli/DiscordSwissArmyKnife/bot"
	"github.com/AlexSafatli/DiscordSwissArmyKnife/config"
)

var (
	discordToken = config.String("discord-token", "blah")
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	err := config.Parse("discord.toml")
	discord, err := bot.NewBot("Bot " + *discordToken)
	if err != nil {
		fmt.Printf("Threw error: %v", err)
		return
	}
	discord.AddHandler(bot.DiceRollHandler)
	fmt.Println(discord)
	err = discord.Open()
	fmt.Println("Loaded")
	for {
	}
}
