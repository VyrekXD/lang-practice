package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyrekxd/domain_email_verifier/pkg/controllers"
)

var RegisterRoutes = func(app *fiber.App) {
	app.Get("/domain", controllers.GetDomain)
	app.Get("/", controllers.GetIndex)
}
