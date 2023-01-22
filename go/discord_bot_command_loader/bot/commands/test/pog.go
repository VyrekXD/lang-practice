package test

import "github.com/bwmarrin/discordgo"

var PogData = &discordgo.ApplicationCommand{
	Name:        "pog",
	Description: "Poggers",
}

func PogRun(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Poggers!",
		},
	})
}
