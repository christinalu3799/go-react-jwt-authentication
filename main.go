package main

import (
	"github.com/christinalu3799/go-react-jwt-authentication/database"
	"github.com/christinalu3799/go-react-jwt-authentication/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
