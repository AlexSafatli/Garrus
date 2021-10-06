package bot

import "github.com/bwmarrin/discordgo"

func openConnection(s *discordgo.Session, channelID, guildID string) error {
	existing, ok := s.VoiceConnections[guildID]
	if ok && existing.ChannelID != channelID || !ok {
		_, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func closeConnectionOrChangeChannelsIfAlone(s *discordgo.Session, guildID string) {
	if s.VoiceConnections[guildID] == nil {
		return
	}
	g, err := s.State.Guild(guildID)
	if err != nil {
		return
	}
	var totalUsersFound int
	var usersFound map[string]int
	usersFound = make(map[string]int)
	for _, vs := range g.VoiceStates {
		if vs.UserID != s.State.User.ID {
			usersFound[vs.ChannelID]++
			totalUsersFound++
		}
	}
	if totalUsersFound == 0 {
		if err = s.VoiceConnections[guildID].Disconnect(); err != nil {
			s.VoiceConnections[guildID].Close()
		}
	} else {
		var mostUsers int
		var channelIDWithMostUsers string
		for k, v := range usersFound {
			if v > mostUsers {
				channelIDWithMostUsers = k
			}
		}
		if channelIDWithMostUsers != s.VoiceConnections[guildID].ChannelID {
			_ = s.VoiceConnections[guildID].ChangeChannel(channelIDWithMostUsers, false, true)
		}
	}
}
