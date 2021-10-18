package bot

import (
	"fmt"
	"github.com/AlexSafatli/Garrus/chat"
	"github.com/AlexSafatli/Garrus/sound"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func searchSounds(query string) (possibilities []string) {
	query = strings.ToLower(query)
	closest := sound.GetLibrary().GetClosestMatchingSoundID(query)
	if len(closest) > 0 {
		possibilities = append(possibilities, closest)
	}
	for _, name := range sound.GetLibrary().GetSoundNames() {
		if strings.Contains(name, query) && name != closest {
			possibilities = append(possibilities, name)
		}
	}
	return
}

func getSoundsForCategoryAsMessageStrings(cat string) ([]string, int) {
	mb := chat.MessageBuilder{}
	var total int
	for _, s := range sound.GetLibrary().SoundMap {
		if s.ContainsCategory(cat) {
			_ = mb.Write("`?" + s.ID + "` ")
			total++
		}
	}
	if total == 0 {
		return []string{}, 0
	}
	return mb.GetMessageStrings(), total
}

func sendSoundsForCategoryForMessageCommand(s *discordgo.Session, channelID string, c string) {
	msgs, total := getSoundsForCategoryAsMessageStrings(c)
	title := fmt.Sprintf("%s (%d)", c, total)
	if total == 0 {
		chat.SendEmbedMessage(s, channelID, title, "No sounds found for this category.", map[string]string{})
	} else {
		for _, msg := range msgs {
			chat.SendEmbedMessage(s, channelID, title, msg, map[string]string{})
		}
	}
}

func sendSoundsForCategoryForSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate, c string) {
	msgs, total := getSoundsForCategoryAsMessageStrings(c)
	if total == 0 {
		chat.SendSimpleInteractionEmbedForAction(s, i, c, "No sounds found for this category.", nil)
	} else {
		chat.SendSimpleInteractionEmbedsForAction(s, i, fmt.Sprintf("%s (%d)", c, total), msgs, nil)
	}
}

func getMainChannelIDForGuild(b *Bot, guildID string) string {
	var id string
	if id, ok := b.mainGuildChannelIDs[guildID]; ok {
		return id
	}
	var lowestPos = -1
	gc, err := b.Session.GuildChannels(guildID)
	if err != nil {
		return id
	}
	for _, c := range gc {
		if c.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		if lowestPos == -1 || c.Position < lowestPos {
			lowestPos = c.Position
			id = c.ID
		}
	}
	if len(id) > 0 {
		b.mainGuildChannelIDs[guildID] = id // cache the channel ID
	}
	return id
}

func authenticate(b *Bot, guildID, userID string) bool {
	m, err := b.Session.GuildMember(guildID, userID)
	if err != nil {
		log.Println("Could not authenticate user", userID)
		return false
	}
	for _, roleID := range m.Roles {
		role, err := b.State.Role(guildID, roleID)
		if err != nil {
			return false
		}
		if role.Permissions&discordgo.PermissionAdministrator != 0 {
			return true
		}
	}
	return false
}
