package main

import (
	"math/rand"
	"time"

	"log"

	"os"

	"./config"
)

var (
	discordToken = config.String("discord-token", "default.token")
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	err := config.Parse("discord.toml")
	if err != nil {
		log.Fatalln("Could not parse discord toml")
		os.Exit(1)
	}
	discord, _ := NewBot("Bot " + *discordToken)
	user, err := discord.Self()
	if err != nil {
		log.Fatalln("Could not get user info for bot")
		os.Exit(1)
	}
	log.SetPrefix(user.Username + " – ")
	err = discord.Open()
	if err != nil {
		log.Fatalln("Could not open websocket")
		os.Exit(1)
	}
	log.Println("Loaded bot with token", *discordToken)
	for {
	}
}
