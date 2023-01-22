package subcommand

import "github.com/bwmarrin/discordgo"

var PongData = &discordgo.ApplicationCommand{
	Name:        "pong",
	Description: "Ping!",
}

func PongRun(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ping!",
		},
	})
}
