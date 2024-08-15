package handlers

import (
	"fmt"

	"github.com/athena-ctf/jester/config"
	"github.com/bwmarrin/discordgo"
)

func MemberAddHandler(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	channelID := config.DefaultConfig.WelcomeChannelID
	bullet := "⭐ "
	embed := &discordgo.MessageEmbed{
		Title: "Athena CTF!",
		Type:  discordgo.EmbedTypeRich,
		Fields: []*discordgo.MessageEmbedField{
			{
				Value: bullet + fmt.Sprintf("Hi %s, welcome to the server!", m.Member.Mention()),
			},
			{
				Value: bullet + "Head over to <#" + config.DefaultConfig.GenralChannelID + "> and start chatting",
			},
		},
		Color: 0x00FFFF,
	}

	if _, err := s.ChannelMessageSendEmbed(channelID, embed); err != nil {
		s.ChannelMessageSend(channelID, err.Error())
	}
}
