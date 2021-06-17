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
		c.SendStatus(fiber.StatusBadRequest)
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

func Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(err)
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.SendStatus(fiber.ErrNotFound.Code)
		return c.JSON(fiber.Map{"message": "user not found"})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "failed to login"})
	}

	return c.JSON(user)
}
