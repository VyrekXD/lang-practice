package commands

import (
	"github.com/bwmarrin/discordgo"

	"github.com/vyrekxd/bot/core/commands/info"
	"github.com/vyrekxd/bot/core/commands/info/subcommand"
	"github.com/vyrekxd/bot/core/commands/test"
)

type CommandRun func(*discordgo.Session, *discordgo.InteractionCreate)

type Command struct {
	Data *discordgo.ApplicationCommand
	Run  CommandRun
}

var Commands = map[string]*Command{
	info.PingData.Name: {Data: info.PingData, Run: info.PingRun},
	subcommand.PongData.Name: {Data: subcommand.PongData, Run: subcommand.PongRun},
	test.PogData.Name: {Data: test.PogData, Run: test.PogRun},
}