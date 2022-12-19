package main

import (
	"fmt"
	"os"

	"github.com/christinalu3799/go-react-jwt-authentication/database"
	"github.com/christinalu3799/go-react-jwt-authentication/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// connect to database
	database.Connect()

	app := fiber.New()

	// user cors to allow frontend to make requests to backend
	app.Use(cors.New(cors.Config{
		// this allows our frontend to receive the cookie that server sends
		// also allows the frontend to send the cookie back
		AllowCredentials: true,
	}))

	routes.Setup(app)

	fmt.Println("HOME: ", os.Getenv("HOME"))
	shell, ok := os.LookupEnv("SHELL")
	if !ok {
		fmt.Println("the env var SHELL is not set")
	} else {
		fmt.Println("SHELL: ", shell)
	}
	app.Listen(":8000")
}
