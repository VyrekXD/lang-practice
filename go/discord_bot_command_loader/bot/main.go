package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/vyrekxd/bot/core/commands"
	"github.com/vyrekxd/bot/core/config"
	"github.com/vyrekxd/bot/core/utils"
)

func main() {
	client, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panicf("An error ocurred when instancing the client: %v", err)
	}

	defer client.Close()

	client.Identify.Intents = discordgo.Intent(38409)

	client.AddHandler(ready)
	client.AddHandler(interactionCreate)

	err = client.Open()
	if err != nil {
		log.Panicf("Error opening connection: %v", err)
		return
	}

	utils.AddCommands(
		client,
		utils.RawCommands(commands.Commands),
		config.TestGuildId,
	)

	log.Println("Bot running!  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if cmd, ok := commands.Commands[i.ApplicationCommandData().Name]; ok {
		cmd.Run(s, i)
	}
}

func ready(_ *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged as: %v", r.User.Username)
}
