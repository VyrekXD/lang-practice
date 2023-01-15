package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyrekxd/crm_fiber/pkg/models"
)

type Error struct {
	Error        int    `json:"error"`
	ErrorMessage string `json:"error_message"`
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(400)
		return c.JSON(&Error{Error: 400, ErrorMessage: "Needs an ID to get lead"})
	}

	lead := models.GetLead(id)

	c.Status(200)
	return c.JSON(lead)
}

func GetLeads(c *fiber.Ctx) error {
	leads := models.GetAllLeads()

	c.Status(200)
	return c.JSON(leads)
}

func CreateLead(c *fiber.Ctx) error {
	lead := &models.Lead{}

	err := c.BodyParser(lead)
	if err != nil {
		c.Status(500)
		return c.JSON(&Error{Error: 500, ErrorMessage: "An error ocurred when parsing the body"})
	}

	lead.Create()

	c.Status(200)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(400)
		return c.JSON(&Error{Error: 400, ErrorMessage: "Needs an ID to get lead"})
	}

	lead := models.GetLead(id)
	if lead.Name == "" {
		c.Status(500)
		return c.JSON(&Error{Error: 500, ErrorMessage: "No lead found with that ID"})
	}

	models.DeleteLead(id)

	c.Status(200)
	return c.JSON(lead)
}
