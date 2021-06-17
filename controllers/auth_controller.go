package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reecerose/go-jwt-auth/database"
	"github.com/reecerose/go-jwt-auth/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(err)
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
