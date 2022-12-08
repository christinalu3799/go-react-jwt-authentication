package database

import (
	"github.com/christinalu3799/go-react-jwt-authentication/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:03071999cl!@/go-react-jwt-authentication"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the db!")
	}

	connection.AutoMigrate(&models.User{})
}
