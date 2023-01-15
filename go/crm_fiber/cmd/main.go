package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vyrekxd/crm_fiber/pkg/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "CRM Fiber v0.0.1",
	})

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen("127.0.0.1:8080"))

}
