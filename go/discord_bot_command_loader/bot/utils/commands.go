package utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/vyrekxd/bot/core/commands"
)

func RawCommands(cmds map[string]*commands.Command) []*discordgo.ApplicationCommand {
	rawCommands := []*discordgo.ApplicationCommand{}

	for _, cmd := range cmds {
		rawCommands = append(rawCommands, cmd.Data)
	}

	return rawCommands
}

func AddCommands(s *discordgo.Session, cmds []*discordgo.ApplicationCommand, guildId string) []*discordgo.ApplicationCommand {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(cmds))

	for i, cmd := range cmds {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, guildId, cmd)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", cmd.Name, err)
		}

		registeredCommands[i] = cmd
	}

	return registeredCommands
}
