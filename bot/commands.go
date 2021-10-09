package bot

import (
	"errors"
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	entranceTitle        = "Set Entrance"
	playSoundTitle       = "Play Sound"
	playRandomSoundTitle = "Play Random Sound"
	searchSoundsTitle    = "Search Sounds"
	listSoundsTitle      = "List Sounds"
)

// AboutMessageCommand takes a created message and returns an About embed message
func AboutMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	chat.SendAboutEmbedMessage(s, m.ChannelID)
}

// AboutSlashCommand returns an About embed message for a slash command
func AboutSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if err := chat.SendInteractionAckForAction(s, i, nil); err != nil {
		return
	}
	defer chat.DeleteInteractionResponse(s, i)
	chat.SendRawEmbedInteractionResponse(s, i, chat.GetRawEmbedMessage(s))
}

// SetEntranceMessageCommand sets an entrance for a user
func SetEntranceMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	var soundID, content, pmsg string
	db, err := LoadDatabase()
	if err != nil {
		log.Fatalln("Could not load database", err)
	}
	defer db.Close()
	i := strings.Index(m.Content, " ")
	if i > 0 {
		content = m.Content[i+1:]
	} else {
		err = sound.DeleteEntranceForUser(m.Author.ID, db)
		chat.SendSimpleMessageResponseForAction(s, m.ChannelID, entranceTitle, "Cleared your entrance.", err)
		return
	}
	args := strings.Split(content, ",")
	if len(args) >= 1 {
		soundID = args[0]
	} else if len(args) >= 2 {
		pmsg = args[1]
	}
	err = sound.SetEntranceForUser(m.Author.ID, soundID, pmsg, db)
	chat.SendSimpleMessageResponseForAction(s, m.ChannelID, entranceTitle, fmt.Sprintf("Set your entrance to `%s`.", soundID), err)
}

// SetEntranceSlashCommand sets an entrance for a user
func SetEntranceSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var soundID, pmsg string
	db, err := LoadDatabase()
	if err != nil {
		log.Fatalln("Could not load database", err)
	}
	defer db.Close()
	args := i.ApplicationCommandData().Options
	if len(args) >= 1 {
		soundID = args[0].StringValue()
	} else if len(args) >= 2 {
		pmsg = args[1].StringValue()
	}
	if strings.TrimSpace(soundID) == "" {
		err := sound.DeleteEntranceForUser(i.Member.User.ID, db)
		chat.SendInteractionResponseForAction(s, i, "Cleared your entrance.", err)
	} else {
		err = sound.SetEntranceForUser(i.Member.User.ID, soundID, pmsg, db)
		chat.SendInteractionResponseForAction(s, i, fmt.Sprintf("Set your entrance to `%s`.", soundID), err)
	}
}

// SearchSoundsMessageCommand returns the collection of sounds matching a query
func SearchSoundsMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	var query string
	var possibilities []string
	i := strings.Index(m.Content, " ")
	if i > 0 {
		query = m.Content[i+1:]
	}
	possibilities = searchSounds(query)
	mb := chat.NewMessageBuilder()
	_ = mb.Write(fmt.Sprintf("Found **%d** possible sounds for query `%s`%s%s.\n\n", len(possibilities), query, chat.Separator, m.Author.Mention()))
	if len(possibilities) > 0 {
		for _, k := range possibilities {
			_ = mb.Write("`?" + k + "` ")
		}
		for _, msg := range mb.GetMessageStrings() {
			chat.SendEmbedMessage(s, m.ChannelID, searchSoundsTitle, msg, map[string]string{})
		}
	} else {
		chat.SendWarningEmbedMessage(s, m.ChannelID, searchSoundsTitle, "Could not find any sounds for query `"+query+"`.")
	}
}

// SearchSoundsSlashCommand returns the collection of sounds matching a query
func SearchSoundsSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var query string
	var possibilities []string
	if err := chat.SendInteractionAckForAction(s, i, nil); err != nil {
		return
	}
	defer chat.DeleteInteractionResponse(s, i)
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		query = args[0].StringValue()
	} else {
		chat.SendErrorEmbedInteractionResponse(s, i, searchSoundsTitle, errors.New("need a query to search for"))
		return
	}
	possibilities = searchSounds(query)
	mb := chat.MessageBuilder{}
	_ = mb.Write(fmt.Sprintf("Found **%d** possible sounds for query `%s`%s%s.\n\n", len(possibilities), query, chat.Separator, i.User.Mention()))
	if len(possibilities) > 0 {
		for _, k := range possibilities {
			_ = mb.Write("`?" + k + "` ")
		}
		for _, msg := range mb.GetMessageStrings() {
			chat.SendEmbedInteractionResponse(s, i, searchSoundsTitle, msg, map[string]string{})
		}
	} else {
		chat.SendWarningEmbedInteractionResponse(s, i, searchSoundsTitle, "Could not find any sounds for query `"+query+"`.")
	}
}

// ListSoundsMessageCommand returns a collection of sounds over multiple messages
func ListSoundsMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	var query string
	i := strings.Index(m.Content, " ")
	if i > 0 {
		query = m.Content[i+1:]
	}
	if query == "" { // all sounds
		for _, c := range sound.GetLibrary().Categories {
			sendSoundsForCategoryForMessageCommand(s, m.ChannelID, c)
		}
		return
	}
	var catFound bool
	for _, c := range sound.GetLibrary().Categories {
		if strings.ToLower(query) == strings.ToLower(c) {
			sendSoundsForCategoryForMessageCommand(s, m.ChannelID, c)
			catFound = true
		}
	}
	if !catFound {
		chat.SendWarningEmbedMessage(s, m.ChannelID, listSoundsTitle, "Could not find any category with name "+query)
	}
}

// ListSoundsSlashCommand returns a collection of sounds over multiple messages
func ListSoundsSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var query string
	if err := chat.SendInteractionAckForAction(s, i, nil); err != nil {
		return
	}
	defer chat.DeleteInteractionResponse(s, i)
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		query = args[0].StringValue()
	}
	if query == "" { // all sounds
		for _, c := range sound.GetLibrary().Categories {
			sendSoundsForCategoryForSlashCommand(s, i, c)
		}
		return
	}
	var catFound bool
	for _, c := range sound.GetLibrary().Categories {
		if strings.ToLower(query) == strings.ToLower(c) {
			sendSoundsForCategoryForSlashCommand(s, i, c)
			catFound = true
		}
	}
	if !catFound {
		chat.SendWarningEmbedInteractionResponse(s, i, listSoundsTitle, "Could not find any category with name "+query)
	}
}

// PlaySoundMessageCommand plays a sound in a voice channel
func PlaySoundMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.VoiceConnections[m.GuildID] == nil {
		// Not in a voice channel
		chat.SendWarningEmbedMessage(s, m.ChannelID, playSoundTitle, "I am not in a voice channel. Join one first!")
		return
	}
	var query = m.Content[1:] // assumes '?' (or one character) start prefix
	var library = sound.GetLibrary()
	if library.Contains(query) {
		db, err := LoadDatabase()
		if err != nil {
			log.Fatalln("Could not load database", err)
		}
		defer db.Close()
		file := library.SoundMap[query]
		err = sound.PlayDCA(file.FilePath, s.VoiceConnections[m.GuildID])
		if err != nil {
			return
		}
		file.NumberPlays++
		if err = library.SetSoundData(file, db); err != nil {
			log.Fatalln("When updating sound => " + err.Error())
		}
	} else {
		msg := "Could not find a sound by name `" + query + "`."
		closestMatch := library.GetClosestMatchingSoundID(query)
		if len(closestMatch) > 0 {
			msg += " Did you mean `" + closestMatch + "`?"
		}
		msg += " " + m.Author.Mention()
		chat.SendWarningEmbedMessage(s, m.ChannelID, playSoundTitle, msg)
	}
}

// PlaySoundSlashCommand plays a sound in a voice channel
func PlaySoundSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var query string
	if err := chat.SendInteractionAckForAction(s, i, nil); err != nil {
		return
	}
	defer chat.DeleteInteractionResponse(s, i)
	if s.VoiceConnections[i.GuildID] == nil {
		// Not in a voice channel
		chat.SendWarningEmbedInteractionResponse(s, i, playSoundTitle, "I am not in a voice channel. Join one first!")
		return
	}
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		query = args[0].StringValue()
	} else {
		chat.SendErrorEmbedInteractionResponse(s, i, playSoundTitle, errors.New("need a sound to play"))
		return
	}
	var library = sound.GetLibrary()
	if library.Contains(query) {
		db, err := LoadDatabase()
		if err != nil {
			log.Fatalln("Could not load database", err)
		}
		defer db.Close()
		file := library.SoundMap[query]
		err = sound.PlayDCA(file.FilePath, s.VoiceConnections[i.GuildID])
		if err != nil {
			return
		}
		file.NumberPlays++
		if err = library.SetSoundData(file, db); err != nil {
			log.Fatalln("When updating sound => " + err.Error())
		}
		chat.SendEmbedInteractionResponse(s, i, playSoundTitle, "Played sound `"+query+"`.", map[string]string{})
	} else {
		msg := "Could not find a sound by name `" + query + "`."
		closestMatch := library.GetClosestMatchingSoundID(query)
		if len(closestMatch) > 0 {
			msg += " Did you mean `" + closestMatch + "`?"
		}
		msg += " " + i.User.Mention()
		chat.SendWarningEmbedInteractionResponse(s, i, playSoundTitle, msg)
	}
}

// PlayRandomSoundMessageCommand plays a random sound in a voice channel
func PlayRandomSoundMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.VoiceConnections[m.GuildID] == nil {
		// Not in a voice channel
		chat.SendWarningEmbedMessage(s, m.ChannelID, playSoundTitle, "I am not in a voice channel. Join one first!")
		return
	}
	var query, category string
	i := strings.Index(m.Content, " ")
	if i > 0 {
		query = m.Content[i+1:]
	}
	var library = sound.GetLibrary()
	if len(query) > 0 {
		for _, c := range library.Categories {
			if strings.ToLower(query) == strings.ToLower(c) {
				category = c
			}
		}
	}
	if len(category) == 0 && len(query) > 0 {
		chat.SendWarningEmbedMessage(s, m.ChannelID, playRandomSoundTitle, "Could not find any category with name "+query)
		return
	}
	db, err := LoadDatabase()
	if err != nil {
		log.Fatalln("Could not load database", err)
	}
	defer db.Close()
	var file *sound.File
	if len(category) > 0 {
		file = library.GetRandomSoundForCategory(category)
	} else {
		file = library.GetRandomSound()
	}
	err = sound.PlayDCA(file.FilePath, s.VoiceConnections[m.GuildID])
	if err != nil {
		return
	}
	file.NumberPlays++
	if err = library.SetSoundData(file, db); err != nil {
		log.Fatalln("When updating sound => " + err.Error())
	}
}

// PlayRandomSoundSlashCommand plays a random sound in a voice channel
func PlayRandomSoundSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var query, category string
	if err := chat.SendInteractionAckForAction(s, i, nil); err != nil {
		return
	}
	defer chat.DeleteInteractionResponse(s, i)
	if s.VoiceConnections[i.GuildID] == nil {
		// Not in a voice channel
		chat.SendWarningEmbedInteractionResponse(s, i, playSoundTitle, "I am not in a voice channel. Join one first!")
		return
	}
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		query = args[0].StringValue()
	}
	var library = sound.GetLibrary()
	if len(query) > 0 {
		for _, c := range library.Categories {
			if strings.ToLower(query) == strings.ToLower(c) {
				category = c
			}
		}
	}
	if len(category) == 0 && len(query) > 0 {
		chat.SendWarningEmbedInteractionResponse(s, i, playRandomSoundTitle, "Could not find any category with name "+query)
		return
	}
	db, err := LoadDatabase()
	if err != nil {
		log.Fatalln("Could not load database", err)
	}
	defer db.Close()
	var file *sound.File
	if len(category) > 0 {
		file = library.GetRandomSoundForCategory(category)
	} else {
		file = library.GetRandomSound()
	}
	err = sound.PlayDCA(file.FilePath, s.VoiceConnections[i.GuildID])
	if err != nil {
		return
	}
	file.NumberPlays++
	if err = library.SetSoundData(file, db); err != nil {
		log.Fatalln("When updating sound => " + err.Error())
	}
}
