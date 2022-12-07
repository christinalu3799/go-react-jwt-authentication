package controllers

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	// get our data back from the post request
	var data map[string]string // similar to an array with a string as a key and value

	// error handling
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	return c.JSON(data)
}
