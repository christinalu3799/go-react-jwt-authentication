package models

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
	Id     uint   `json:"checking_id"`
	Number string `json:"number"`
	UserID uint   `json:"user_id"`
}
