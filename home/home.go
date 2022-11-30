package home

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message":  "Здравствуйте, сейчас Вы на Домашней странице.",
	})
}
