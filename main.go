package main

import (
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
	"time"

	"log"

	"os"

	"./config"
)

var (
	// TODO use config library
	discordToken = config.String("discord-token", "default.token")
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	err := config.Parse("discord.toml")
	if err != nil {
		log.Fatalln("Could not parse discord toml")
	}

	discord, _ := NewBot("Bot " + *discordToken)
	user, err := discord.Self()
	if err != nil {
		log.Fatalln("Could not get user info for bot")
	}
	log.SetPrefix(user.Username + " – ")

	if err = discord.Open(); err != nil {
		log.Fatalln("Could not open websocket")
	}
	log.Println("Loaded bot with token", *discordToken)

	// Wait here until CTRL-C or another term signal
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sig

	_ = discord.Close()
	fmt.Println("Aborted!")
}
