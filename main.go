package main

import (
	"fmt"
	"github.com/christinalu3799/go-react-jwt-authentication/database"
	"github.com/christinalu3799/go-react-jwt-authentication/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {
	// connect to database
	database.Connect()

	app := fiber.New()

	// use cors to allow client to send requests to backend
	app.Use(cors.New(cors.Config{
		// this allows our frontend to receive/send back the cookie to/from the server
		AllowCredentials: true,
	}))

	routes.Setup(app)

	// get port
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}
	fmt.Printf("Listening on port %s\n\n", port)
	app.Listen(port)
}
