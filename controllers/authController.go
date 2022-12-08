package controllers

import (
	"github.com/christinalu3799/go-react-jwt-authentication/database"
	"github.com/christinalu3799/go-react-jwt-authentication/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	// get our data back from the post request
	var data map[string]string // similar to an array with a string as a key and value

	// error handling
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// create the user
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	// 👆🏼 need to convert our password to a byte array
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password, // need to hash the password
	}

	// we are calling the global DB variable from connection.go
	// then, we are creating + inserting the user we created on line 22 into DB
	database.DB.Create(&user)
	return c.JSON(user)
}
