package controllers

import (
	"fmt"
	"strconv"

	"github.com/christinalu3799/go-react-jwt-authentication/database"
	"github.com/christinalu3799/go-react-jwt-authentication/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CreateChecking(c *fiber.Ctx) error {
	// the data we are posting to the server
	var data map[string]string
	// error handling
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	// need to get user id from cookie
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// error handling if user is not logged in
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated. please login to create checking balance.",
		})
	}
	// get the claims from the token
	claims := token.Claims.(*jwt.StandardClaims)
	uint8UserID, err := strconv.ParseUint(claims.Issuer, 10, 32)
	// error handling for converting sting to uint
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// create our checking schema
	checking := models.Checking{
		Number: data["number"],
		UserID: uint(uint8UserID),
	}
	database.DB.Create(&checking)

	return c.JSON(checking)
}

func GetCheckingBalance(c *fiber.Ctx) error {

	// need to get user id from cookie
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// error handling if user is not logged in
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated. please login to view checking balance.",
		})
	}
	// get the claims from the token
	claims := token.Claims.(*jwt.StandardClaims)
	fmt.Println(claims)
	fmt.Println("claims.Issuer = ", claims.Issuer)

	// initialize variable for checking balance
	var checking models.Checking
	// database.DB.Where("user_id = ?", claims.Issuer).Find(&checking)
	database.DB.Where(map[string]interface{}{"user_id": claims.Issuer}).Find(&checking)
	fmt.Println(checking)
	return c.JSON(checking)
}
