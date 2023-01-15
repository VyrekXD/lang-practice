package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	cmds "github.com/vyrekxd/discord_bot/pkg/commands"
	"github.com/vyrekxd/discord_bot/pkg/config"
)

var (
	s        *discordgo.Session
	commands = []*discordgo.ApplicationCommand{
		cmds.PingData,
		cmds.YobData,
	}
)

func init() {
	var err error
	s, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Error instancing the bot: ", err)
		return
	}

	s.Identify.Intents = discordgo.IntentsGuildMessages

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	s.AddHandler(messageCreate)

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.ApplicationCommandData().Name {
		case cmds.PingData.Name:
			{
				cmds.PingRun(s, i)
			}
		case cmds.YobData.Name:
			{
				cmds.YobRun(s, i)
			}

		}
	})
}

func main() {
	if err := s.Open(); err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, config.GuildId, v)

		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}

		registeredCommands[i] = cmd
	}

	defer s.Close()

	fmt.Println("\nBot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("Removing commands...")

	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, config.GuildId, v.ID)

		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	fmt.Println("\nGracefully shutting down.")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	} else if m.Author.Bot {
		return
	} else if !strings.HasPrefix(m.Content, config.Prefix) {
		return
	}

	args := strings.Fields(strings.TrimPrefix(m.Content, config.Prefix))

	if len(args) == 0 {
		return
	}

	commandName := args[0]
	args = args[1:]

	if commandName == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")

		return
	} else if commandName == "args" {
		argsString := strings.Join(args, ", ")

		if len(argsString) == 0 {
			s.ChannelMessageSend(m.ChannelID, "No args.")
			return
		}

		s.ChannelMessageSend(m.ChannelID, argsString)

		return
	} else if commandName == "yob" {
		if len(args) == 0 {
			s.ChannelMessageSend(m.ChannelID, "No year specified.")
			return
		}

		year, err := strconv.Atoi(args[0])

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Argument wasnt a number, err: "+err.Error())
			return
		}

		actualYear := time.Now().Year()
		age := actualYear - year

		s.ChannelMessageSend(m.ChannelID, "Your age is: "+strconv.Itoa(age))

		return
	}
}
