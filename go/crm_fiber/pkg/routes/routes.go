package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyrekxd/crm_fiber/pkg/controllers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("api")

	v1 := api.Group("v1")

	v1.Get("/lead", controllers.GetLeads)
	v1.Get("/lead/:id", controllers.GetLead)
	v1.Post("/lead/", controllers.CreateLead)
	v1.Delete("/lead/:id", controllers.DeleteLead)
}
