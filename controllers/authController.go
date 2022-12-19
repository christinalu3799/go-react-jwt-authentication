package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/christinalu3799/go-react-jwt-authentication/database"
	"github.com/christinalu3799/go-react-jwt-authentication/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey string = "secret"

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
	// then, we are creating + inserting the user we created into DB
	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// want to get user associated w email
	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	// if we haven't found the user based off email
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// if we found the user, then we need to compare the password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	// when we get the right email and password, we want to return a JWT token
	// creating the claims, which are statements about an entity (typically, the user) and additional data

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		// issuer is our user, need to convert user id back to string
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	// here, we are signing our token to make sure that we are who we say we are
	// in other words,signing our JWTs with a secret lets us know whether the content has been tampered with
	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not log in",
		})
	}

	// store our token in cookies
	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
		// store cookie on client-side
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(user)
}

func User(c *fiber.Ctx) error {
	// retrieve the cookie from the client
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// get the claims from the token
	claims := token.Claims.(*jwt.StandardClaims)

	// with the claims data, we want to retrieve the logged in user
	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)
	fmt.Println(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// to logout, we need to remove the cookie that we just created in the login function
	// here, we are encrypting a cookie values
	cookie := fiber.Cookie{
		Name:  "jwt",
		Value: "",
		// set the expiry time to 1 hour ago to "remove" the cookie
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "user has successfully logged out.",
	})
}
