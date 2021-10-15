package main

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/bot"
	"github.com/AlexSafatli/Garrus/sound"
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
	log.Println("Loaded config")

	// Load flat file database
	bot.SetDatabasePath(configValues.DbPath)
	db := bot.LoadDatabase()
	log.Println("Loaded database from " + configValues.DbPath)

	// Load entrances
	err = sound.LoadEntrances(db)
	if err != nil {
		log.Fatalln("Could not load entrances from database", err)
	}

	// Load sounds and sound data
	if err = sound.LoadSounds(configValues.SoundsPath); err != nil {
		log.Fatalln("Could not load/convert sounds from "+configValues.SoundsPath, err)
	}
	if err = sound.GetLibrary().LoadSoundData(db); err != nil {
		log.Fatalln("Could not load sound data from database", err)
	}
	log.Println("Loaded sounds from " + configValues.SoundsPath)

	// Close database for now
	if err = db.Close(); err != nil {
		log.Fatalln("Could not properly close database", err)
	}

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
