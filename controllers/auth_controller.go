package controllers

import (
	"strconv"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/reecerose/go-jwt-auth/database"
	"github.com/reecerose/go-jwt-auth/models"
	"github.com/reecerose/go-jwt-auth/utils"
	"golang.org/x/crypto/bcrypt"
)

// Register a new user
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

// Login an existing user
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

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(utils.SECRET_KEY))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "failed to login"})
	}

	return c.JSON(fiber.Map{"token": token})
}
