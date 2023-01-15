package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conecting to mongo")
	err := mgm.SetDefaultConfig(nil, "database", options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to mongo")

}
