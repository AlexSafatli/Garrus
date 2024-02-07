package bot

import (
	"github.com/bwmarrin/discordgo"
)

func getUsersVoiceConnectionsCountMap(s *discordgo.Session, g *discordgo.Guild) (total int, usersFound map[string]int) {
	usersFound = make(map[string]int)
	for _, vs := range g.VoiceStates {
		if vs.UserID != s.State.User.ID {
			usersFound[vs.ChannelID]++
			total++
		}
	}
	return
}

func getUsersVoiceChannelID(s *discordgo.Session, guildID, userID string) (channelID string) {
	g, err := s.State.Guild(guildID)
	if err != nil {
		return
	}
	for _, vs := range g.VoiceStates {
		if vs.UserID == userID {
			channelID = vs.ChannelID
			return
		}
	}
	return
}

func openVoiceConnection(s *discordgo.Session, channelID, guildID string) error {
	existing, ok := s.VoiceConnections[guildID]
	if ok && existing.ChannelID != channelID || !ok {
		if _, err := s.ChannelVoiceJoin(guildID, channelID, false, true); err != nil {
			return err
		}
	}
	return nil
}

func closeVoiceConnectionOrChangeChannelsIfAlone(s *discordgo.Session, guildID string) {
	if s.VoiceConnections[guildID] == nil {
		return // no connection opened
	}
	g, err := s.State.Guild(guildID)
	if err != nil {
		return
	}

	totalUsersFound, usersFound := getUsersVoiceConnectionsCountMap(s, g)
	channelID := s.VoiceConnections[guildID].ChannelID

	if totalUsersFound == 0 {
		if err = s.VoiceConnections[guildID].Disconnect(); err != nil {
			s.VoiceConnections[guildID].Close()
		}
	} else if usersFound[channelID] == 0 {
		var mostUsers int
		var channelIDWithMostUsers string
		for k, v := range usersFound {
			if v > mostUsers {
				channelIDWithMostUsers = k
			}
		}
		if channelIDWithMostUsers != channelID {
			_ = s.VoiceConnections[guildID].ChangeChannel(channelIDWithMostUsers, false, true)
		}
	}
}
