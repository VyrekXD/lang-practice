package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vyrekxd/domain_email_verifier/pkg/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Domain Email Checker v0.0.1",
	})

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen("127.0.0.1:8080"))

}
