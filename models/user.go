package models

// User model for database
type User struct {
	Id       uint
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}
