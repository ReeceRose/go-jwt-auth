package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(sqlite.Open("jwt.db"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}
}
