package database

import (
	"fmt"
	"github.com/christinalu3799/go-react-jwt-authentication/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// create a GLOBAL variable for our database
var DB *gorm.DB

func Connect() {
	// load all env variables into a map
	envMap, mapErr := godotenv.Read("./.env")
	if mapErr != nil {
		fmt.Printf("Error loading .env file into map[sting]string\n")
		os.Exit(1)
	}
	// this is where we are setting up our database connection
	// dsn = URL string of database passed into the .Open() method
	dsn := envMap["dbUser"] + ":" + envMap["dbPass"] + "@" + envMap["tcp"] + "/" + envMap["dbName"]
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// error handling
	if err != nil {
		panic("could not connect to the db!")
	}

	DB = connection

	// create users and checkings tables in our database
	connection.AutoMigrate(&models.User{}, &models.Checking{})
}
