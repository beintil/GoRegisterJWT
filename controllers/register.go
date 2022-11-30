package controllers

import (
	"net/http"

	"github.com/beintil/goregisterjwt/database"
	"github.com/beintil/goregisterjwt/user"
	"github.com/gofiber/fiber/v2"
)

func RegisterUsers(c *fiber.Ctx) error {
	var user user.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := user.HashPasswod(user.Password); err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	createUser := database.DBCon.Create(&user)
	if createUser.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": createUser.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"password": "*****",
	})
}
