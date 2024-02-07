package main

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/bot"
	"log"
	"os/signal"
	"syscall"

	"os"

	"github.com/AlexSafatli/Garrus/config"
)

func main() {
	configValues := config.LoadConfigs()
	log.Println("Loaded config")

	// Set flat file database location
	bot.SetDatabasePath(configValues.DbPath)

	// Load bot
	discord, _ := bot.NewBot("Bot " + configValues.DiscordToken)
	user, err := discord.Self()
	if err != nil {
		log.Fatalln("Could not get user info for bot", err)
	}
	log.SetPrefix(user.Username + " – ")
	if err = discord.Open(); err != nil {
		log.Fatalln("Could not open websocket", err)
	}
	log.Println("Loaded bot with provided token")

	// Register slash commands
	if err = discord.RegisterSlashCommands(); err != nil {
		log.Fatalln("Could not register slash commands", err)
	}

	// Wait here until CTRL-C or another term signal
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sig

	_ = discord.Close()
	fmt.Println("Aborted!")
}
