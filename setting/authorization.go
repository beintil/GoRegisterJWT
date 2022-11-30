package setting

import (
	"github.com/gofiber/fiber/v2"
)

func Authorization(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	//проверяем, содержит ли запрос хоть какие либо данные
	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "missing token",
		})
	}

	// Высылаем токен в func ChecToken где она проверяются
	err := CheckToken(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Next()
}
