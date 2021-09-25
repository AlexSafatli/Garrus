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
	entranceTitle     = "Set Entrance"
	playSoundTitle    = "Play Sound"
	searchSoundsTitle = "Search Sounds"
)

func searchSounds(query string) (possibilities []string) {
	closest := sound.GetLibrary().GetClosestMatchingSoundID(query)
	if len(closest) > 0 {
		possibilities = append(possibilities, closest)
	}
	for _, name := range sound.GetLibrary().GetSoundNames() {
		if strings.Contains(name, query) && name != closest {
			possibilities = append(possibilities, closest)
		}
	}
	return
}

// AboutMessageCommand takes a created message and returns an About embed message
func AboutMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Printf("About handler running for message from %s", m.Author.Username)
}

// AboutSlashCommand returns an About embed message for a slash command
func AboutSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Printf("About handler running for message from %s", i.Member.User.Username)
}

// SetEntranceMessageCommand sets an entrance for a user
func SetEntranceMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	var soundID, content, pmsg string
	i := strings.Index(m.Content, " ")
	if i > 0 {
		content = m.Content[i+1:]
	} else {
		err := sound.DeleteEntranceForUser(m.Author.ID, Db)
		chat.SendSimpleMessageResponseForAction(s, m.ChannelID, entranceTitle, "Cleared your entrance.", err)
		return
	}
	args := strings.Split(content, ",")
	if len(args) >= 1 {
		soundID = args[0]
	} else if len(args) >= 2 {
		pmsg = args[1]
	}
	err := sound.SetEntranceForUser(m.Author.ID, soundID, pmsg, Db)
	chat.SendSimpleMessageResponseForAction(s, m.ChannelID, entranceTitle, fmt.Sprintf("Set your entrance to `%s`.", soundID), err)
}

// SetEntranceSlashCommand sets an entrance for a user
func SetEntranceSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var soundID, pmsg string
	if len(i.ApplicationCommandData().Options) == 0 {
		err := sound.DeleteEntranceForUser(i.User.ID, Db)
		chat.SendInteractionResponseForAction(s, i, "Cleared your entrance.", err)
		return
	}
	args := i.ApplicationCommandData().Options
	if len(args) >= 1 {
		soundID = args[0].StringValue()
	} else if len(args) >= 2 {
		pmsg = args[1].StringValue()
	}
	err := sound.SetEntranceForUser(i.User.ID, soundID, pmsg, Db)
	chat.SendInteractionResponseForAction(s, i, fmt.Sprintf("Set your entrance to `%s`.", soundID), err)
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

// PlaySoundMessageCommand plays a sound in a voice channel
func PlaySoundMessageCommand(s *discordgo.Session, vc *discordgo.VoiceConnection, m *discordgo.MessageCreate) {
	var query = m.Content[1:] // assumes '?' (or one character) start prefix
	var closest = sound.GetLibrary().GetClosestMatchingSoundID(query)
	if closest == query {
		file := sound.GetLibrary().SoundMap[query]
		err := sound.PlayDCA(file.FilePath, vc)
		if err != nil {
			return
		}
		file.NumberPlays++
		if err = sound.GetLibrary().SetSoundData(file, Db); err != nil {
			log.Fatalln("When updating sound =>" + err.Error())
		}
	} else {
		chat.SendWarningEmbedMessage(s, m.ChannelID, playSoundTitle, "Could not find a sound by name `"+query+"`. Did you mean `"+closest+"`?")
	}
}

// PlaySoundSlashCommand plays a sound in a voice channel
func PlaySoundSlashCommand(s *discordgo.Session, vc *discordgo.VoiceConnection, i *discordgo.InteractionCreate) {
	var query string
	if err := chat.SendInteractionAckForAction(s, i, nil); err != nil {
		return
	}
	defer chat.DeleteInteractionResponse(s, i)
	args := i.ApplicationCommandData().Options
	if len(args) == 1 {
		query = args[0].StringValue()
	} else {
		chat.SendErrorEmbedInteractionResponse(s, i, playSoundTitle, errors.New("need a sound to play"))
		return
	}
	var closest = sound.GetLibrary().GetClosestMatchingSoundID(query)
	if closest == query {
		file := sound.GetLibrary().SoundMap[query]
		err := sound.PlayDCA(file.FilePath, vc)
		if err != nil {
			return
		}
		file.NumberPlays++
		if err = sound.GetLibrary().SetSoundData(file, Db); err != nil {
			log.Fatalln("When updating sound =>" + err.Error())
		}
		chat.SendEmbedInteractionResponse(s, i, playSoundTitle, "Played sound `"+query+"`.", map[string]string{})
	} else {
		chat.SendWarningEmbedInteractionResponse(s, i, playSoundTitle, "Could not find a sound by name `"+query+"`. Did you mean `"+closest+"`?")
	}
}
