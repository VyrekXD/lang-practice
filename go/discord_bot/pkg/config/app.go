package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token   string
	Prefix  string = "go!"
	GuildId string = "973427352560365658"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token = os.Getenv("BOT_TOKEN")
}
