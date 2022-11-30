package main

import (
	"log"

	"github.com/beintil/goregisterjwt/controllers"
	"github.com/beintil/goregisterjwt/database"
	"github.com/beintil/goregisterjwt/home"
	"github.com/beintil/goregisterjwt/setting"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DBConnect("root:password@tcp(localhost:3306)/go_register_jwt?parseTime=true")

	app := fiber.New()

	app.Post("/api/token", controllers.NewToken)
	app.Post("/api/user/register", controllers.RegisterUsers)
	openHome := app.Group("/api/home").Use(setting.Authorization)
	{
		openHome.Get("/myhome", home.Home)
	}

	log.Fatal(app.Listen(":8080"))
}
