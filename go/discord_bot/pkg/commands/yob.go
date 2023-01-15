package commands

import (
	"github.com/bwmarrin/discordgo"
)

var YobData *discordgo.ApplicationCommand = &discordgo.ApplicationCommand{
	Name:        "yob",
	Description: "How old are you?",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "yob",
			Description: "Year of birth",
			Required:    true,
		},
	},
}

var YobRun = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
