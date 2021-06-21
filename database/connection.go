package database

import (
	"github.com/reecerose/go-jwt-auth/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database connection, used to access DB in other files
var DB *gorm.DB

// Connect/Initialize database connection. Should be called on program start
func Connect() {
	connection, err := gorm.Open(sqlite.Open("jwt.db"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
