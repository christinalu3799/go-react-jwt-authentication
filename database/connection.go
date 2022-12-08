package database

import (
	"github.com/christinalu3799/go-react-jwt-authentication/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// create a GLOBAL variables (CAPS)
var DB *gorm.DB

func Connect() {
	// this is where we are setting up our database connection
	connection, err := gorm.Open(mysql.Open("root:03071999cl!@/go-react-jwt-authentication"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the db!")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
