package routes

import (
	"github.com/christinalu3799/go-react-jwt-authentication/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
