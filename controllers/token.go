package controllers

import (
	"net/http"

	"github.com/beintil/goregisterjwt/database"
	"github.com/beintil/goregisterjwt/setting"
	"github.com/beintil/goregisterjwt/user"
	"github.com/gofiber/fiber/v2"
)

// Определяем структуру с тело запроса в виде password and email
type TokenStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewToken(c *fiber.Ctx) error {

	var (
		user  user.User
		token TokenStruct
	)

	if err := c.BodyParser(&token); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//Проверяем существование email in datadase
	email := database.DBCon.Where("email = ?", token.Email).First(&user)
	if email.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": email.Error.Error(),
		})
	}

	//проверяем совпадает ли пароль прикрепленный к email с тем который ввели
	password := user.ComparesPassword(token.Password)
	if password != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "password doen't match",
		})
	}

	// Генерируем новый токен
	tokenGenerate, err := setting.GenerateJWT(user.Email, user.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"token": tokenGenerate,
	})
}
