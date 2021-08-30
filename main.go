package main

import (
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
	"time"

	"log"

	"os"

	"github.com/AlexSafatli/Garrus/config"
)

func main() {
	var err error

	rand.Seed(int64(time.Now().Nanosecond()))

	configValues := config.LoadConfigs()
	discord, _ := NewBot("Bot " + configValues.DiscordToken)
	user, err := discord.Self()
	if err != nil {
		log.Fatalln("Could not get user info for bot")
	}
	log.SetPrefix(user.Username + " – ")

	if err = discord.Open(); err != nil {
		log.Fatalln("Could not open websocket")
	}
	log.Println("Loaded bot with provided token")

	// Wait here until CTRL-C or another term signal
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sig

	_ = discord.Close()
	fmt.Println("Aborted!")
}
