package commands

import (
	"github.com/bwmarrin/discordgo"
)

var PingData *discordgo.ApplicationCommand = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Pong!",
}

var PingRun = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
