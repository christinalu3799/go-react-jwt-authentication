package models

import (
	"gorm.io/gorm"
)

// Belongs to association ---
// `Checking` belongs to `User`
// Here, `UserID` is the foreign key in Checking
type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

type Checking struct {
	gorm.Model
	Number string `json:"number"`
	UserID uint
}
