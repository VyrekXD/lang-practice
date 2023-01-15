package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vyrekxd/domain_email_verifier/pkg/utils"
)

func GetIndex(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/plain")
	c.SendStatus(200)

	return c.SendString("Hello world!")
}

func GetDomain(c *fiber.Ctx) error {
	domain := c.Query("domain")

	if len(domain) == 0 {
		c.SendStatus(400)
		c.Set("Content-Type", "application/json")

		return c.Send([]byte(`{"error": 400, "error_message": "Domain is required"}`))
	}

	domainData, err := utils.CheckDomain(domain)

	if err != nil {
		c.SendStatus(500)
		c.Set("Content-Type", "application/json")

		return c.Send([]byte(fmt.Sprintf("{\"error\": 500, \"error_message\": \"An error ocurred when checking the domain: %v\"}", err)))
	}

	jdata, err := json.Marshal(domainData)

	if err != nil {
		c.SendStatus(500)
		c.Set("Content-Type", "application/json")

		return c.Send([]byte(fmt.Sprintf("{\"error\": 500, \"error_message\": \"An error ocurred when marshaling the struct: %v\"}", err)))
	}

	c.Set("Content-Type", "application/json")
	c.SendStatus(200)

	return c.Send(jdata)
}
