package bot

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	entranceTitle        = "Set Entrance"
	playSoundTitle       = "Play Sound"
	playRandomSoundTitle = "Play Random Sound"
	searchSoundsTitle    = "Search Sounds"
	listCategoriesTitle  = "Categories"
	listSoundsTitle      = "List Sounds"
)

// AboutMessageCommand takes a created message and returns an About embed message
func AboutMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	chat.SendRawEmbedMessage(s, m.ChannelID, chat.GetRawAboutEmbedMessage(s))
}

// AboutSlashCommand returns an About embed message for a slash command
func AboutSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	chat.SendInteractionRawEmbedForAction(s, i, chat.GetRawAboutEmbedMessage(s), nil)
}

// SetEntranceMessageCommand sets an entrance for the invoking user
func SetEntranceMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	var soundID, content, pmsg string
	db := LoadDatabase()
	defer db.Close()
	i := strings.Index(m.Content, " ")
	if i > 0 {
		content = m.Content[i+1:]
	} else {
		err := sound.DeleteEntranceForUser(m.Author.ID, db)
		chat.SendSimpleMessageResponseForAction(s, m.ChannelID, entranceTitle, "Cleared your entrance.", err)
		return
	}
	args := strings.Split(content, ",")
	if len(args) >= 1 {
		soundID = args[0]
	} else if len(args) >= 2 {
		pmsg = args[1]
	}
	err := sound.SetEntranceForUser(m.Author.ID, soundID, pmsg, db)
	chat.SendSimpleMessageResponseForAction(s, m.ChannelID, entranceTitle, fmt.Sprintf("Set your entrance to `%s`.", soundID), err)
}

// SetEntranceSlashCommand sets an entrance for the invoking user
func SetEntranceSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var soundID, pmsg string
	db := LoadDatabase()
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
		err := sound.SetEntranceForUser(i.Member.User.ID, soundID, pmsg, db)
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
			chat.SendSimpleEmbedMessage(s, m.ChannelID, searchSoundsTitle, msg)
		}
	} else {
		chat.SendWarningEmbedMessage(s, m.ChannelID, searchSoundsTitle, "Could not find any sounds for query `"+query+"`.")
	}
}

// SearchSoundsSlashCommand returns the collection of sounds matching a query
func SearchSoundsSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var query string
	var possibilities []string
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		query = args[0].StringValue()
	}
	possibilities = searchSounds(query)
	mb := chat.MessageBuilder{}
	_ = mb.Write(fmt.Sprintf("Found **%d** possible sounds for query `%s`%s%s.\n\n", len(possibilities), query, chat.Separator, i.Member.Mention()))
	if len(possibilities) > 0 {
		for _, k := range possibilities {
			_ = mb.Write("`?" + k + "` ")
		}
		chat.SendSimpleInteractionEmbedsForAction(s, i, searchSoundsTitle, mb.GetMessageStrings(), nil)
	} else {
		chat.SendWarningInteractionEmbedForAction(s, i, searchSoundsTitle, "Could not find any sounds for query `"+query+"`.", nil)
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
	ok, category := sound.GetLibrary().Category(query)
	if !ok {
		chat.SendWarningEmbedMessage(s, m.ChannelID, listSoundsTitle, "Could not find any category with name "+query)
		return
	}
	sendSoundsForCategoryForMessageCommand(s, m.ChannelID, category)
}

// ListSoundsSlashCommand returns a collection of sounds over multiple messages
func ListSoundsSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var query string
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		query = args[0].StringValue()
	}
	ok, category := sound.GetLibrary().Category(query)
	if !ok {
		chat.SendWarningInteractionEmbedForAction(s, i, listSoundsTitle, "Could not find any category with name "+query, nil)
		return
	}
	sendSoundsForCategoryForSlashCommand(s, i, category)
}

// ListCategoriesMessageCommand returns the list of categories
func ListCategoriesMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	chat.SendSimpleEmbedMessage(s, m.ChannelID, listCategoriesTitle, chat.SliceToMessageString(sound.GetLibrary().Categories))
}

// ListCategoriesSlashCommand returns the list of categories
func ListCategoriesSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	chat.SendSimpleInteractionEmbedForAction(s, i, listCategoriesTitle, chat.SliceToMessageString(sound.GetLibrary().Categories), nil)
}

// PlaySoundMessageCommand plays a sound in a voice channel
func PlaySoundMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	targetChannelID := getUsersVoiceChannelID(s, m.GuildID, m.Author.ID)
	if s.VoiceConnections[m.GuildID] == nil || targetChannelID != s.VoiceConnections[m.GuildID].ChannelID {
		// Not in a voice channel; join
		if err := openVoiceConnection(s, targetChannelID, m.GuildID); err != nil {
			chat.SendErrorEmbedMessage(s, m.ChannelID, playSoundTitle, err)
			return
		}
	}
	var query = m.Content[1:] // assumes '?' (or one character) start prefix
	var library = sound.GetLibrary()
	if library.Contains(query) {
		db := LoadDatabase()
		defer db.Close()
		file := library.SoundMap[query]
		playSoundWithSave(file, s.VoiceConnections[m.GuildID], db)
	} else {
		msg := "Could not find a sound by name `" + query + "`"
		closestMatch := library.GetClosestMatchingSoundID(query)
		if len(closestMatch) > 0 {
			msg += " Did you mean `" + closestMatch + "`?"
		}
		msg += chat.Separator + m.Author.Mention()
		chat.SendWarningEmbedMessage(s, m.ChannelID, playSoundTitle, msg)
	}
}

// PlaySoundSlashCommand plays a sound in a voice channel
func PlaySoundSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var query string
	targetChannelID := getUsersVoiceChannelID(s, i.GuildID, i.Member.User.ID)
	if s.VoiceConnections[i.GuildID] == nil || targetChannelID != s.VoiceConnections[i.GuildID].ChannelID {
		// Not in a voice channel; join
		if err := openVoiceConnection(s, targetChannelID, i.GuildID); err != nil {
			chat.SendErrorInteractionEmbedForAction(s, i, playSoundTitle, err)
			return
		}
	}
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		query = args[0].StringValue()
	}
	var library = sound.GetLibrary()
	if library.Contains(query) {
		db := LoadDatabase()
		defer db.Close()
		file := library.SoundMap[query]
		playSoundWithSave(file, s.VoiceConnections[i.GuildID], db)
		chat.SendInteractionResponseForAction(s, i, "Played sound `"+query+"`.", nil)
	} else {
		msg := "Could not find a sound by name `" + query + "`"
		closestMatch := library.GetClosestMatchingSoundID(query)
		if len(closestMatch) > 0 {
			msg += " Did you mean `" + closestMatch + "`?"
		}
		msg += chat.Separator + i.Member.Mention()
		chat.SendWarningInteractionEmbedForAction(s, i, playSoundTitle, msg, nil)
	}
}

// PlayRandomSoundSlashCommand plays a random sound in a voice channel
func PlayRandomSoundSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var query, category string
	targetChannelID := getUsersVoiceChannelID(s, i.GuildID, i.Member.User.ID)
	if s.VoiceConnections[i.GuildID] == nil || targetChannelID != s.VoiceConnections[i.GuildID].ChannelID {
		// Not in a voice channel; join
		if err := openVoiceConnection(s, targetChannelID, i.GuildID); err != nil {
			chat.SendErrorInteractionEmbedForAction(s, i, playRandomSoundTitle, err)
			return
		}
	}
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		query = args[0].StringValue()
	}
	var library = sound.GetLibrary()
	if len(query) > 0 {
		var ok bool
		ok, category = library.Category(query)
		if !ok {
			chat.SendWarningInteractionEmbedForAction(s, i, playRandomSoundTitle, "Could not find any category with name "+query, nil)
			return
		}
	}
	db := LoadDatabase()
	defer db.Close()
	var file *sound.File
	if len(category) > 0 {
		file = library.GetRandomSoundForCategory(category)
	} else {
		file = library.GetRandomSound()
	}
	playSoundWithSave(file, s.VoiceConnections[i.GuildID], db)
	soundInfo := fmt.Sprintf("Played random sound `%s` from **%s** (**%d** plays)", file.ID, file.Categories[0], file.NumberPlays)
	msg := soundInfo + chat.Separator + i.Member.Mention()
	chat.SendSimpleInteractionEmbedForAction(s, i, playRandomSoundTitle, msg, nil)
}
