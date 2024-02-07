package bot

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

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

func authenticate(s *discordgo.Session, guildID, userID string) bool {
	m, err := s.GuildMember(guildID, userID)
	if err != nil {
		log.Println("Could not authenticate user", userID)
		return false
	}
	for _, roleID := range m.Roles {
		role, err := s.State.Role(guildID, roleID)
		if err != nil {
			return false
		}
		if role.Permissions&discordgo.PermissionAdministrator != 0 {
			return true
		}
	}
	return false
}
