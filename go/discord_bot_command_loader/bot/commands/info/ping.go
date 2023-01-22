package info

import "github.com/bwmarrin/discordgo"

var PingData = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Pong!",
}

func PingRun(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
