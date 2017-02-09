package main

import (
	"math/rand"
	"time"

	"log"

	"os"

	"github.com/AlexSafatli/DiscordSwissArmyKnife/bot"
	"github.com/AlexSafatli/DiscordSwissArmyKnife/config"
)

var (
	discordToken = config.String("discord-token", "default.token")
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	config.Parse("discord.toml")
	discord, _ := bot.NewBot("Bot " + *discordToken)
	user, err := discord.Self()
	if err != nil {
		log.Println("Could not get user info for bot")
		os.Exit(1)
	}
	log.SetPrefix(user.Username)
	discord.AddHandler(bot.DiceRollHandler)
	err = discord.Open()
	if err != nil {
		log.Println("Could not open websocket")
		os.Exit(1)
	}
	log.Println("Loaded bot with token", *discordToken)
	for {
	}
}
