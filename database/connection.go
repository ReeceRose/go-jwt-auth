package database

import (
	"github.com/reecerose/go-jwt-auth/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is used to access DB in other files
var DB *gorm.DB

// Connect is used for initializing the database and opening the connection. Should be called on program start
func Connect() {
	connection, err := gorm.Open(sqlite.Open("jwt.db"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
