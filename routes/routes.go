package routes

import (
	"github.com/christinalu3799/go-react-jwt-authentication/controllers"
	"github.com/gofiber/fiber/v2"
)

// ❗️ * = dereferencing operator; a pointer and defining type of a variable
// ❗️ & = address operator which obtains the pointer (address) of a variable
// https://dev.classmethod.jp/articles/understanding-pointers-in-go/

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
}
