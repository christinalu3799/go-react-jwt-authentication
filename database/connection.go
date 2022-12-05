package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:03071999cl!@/go-react-jwt-authentication"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the db!")
	}
}
