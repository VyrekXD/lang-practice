package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	TestGuildId = ""
	DevGuildId  = ""
)

func init() {
	if os.Getenv("APP_ENV") == "development" {
		err := godotenv.Load(".dev.env")
		if err != nil {
			log.Panicf("An error ocurred when reading .env: %v", err)
		}

		TestGuildId = os.Getenv("DEV_GUILD_ID")
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Panicf("An error ocurred when reading .env: %v", err)
		}

		DevGuildId = os.Getenv("DEV_GUILD_ID")
	}
}
