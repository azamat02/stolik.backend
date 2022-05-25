package routes

import (
	"github.com/gofiber/fiber/v2"
	"stolik.online/controllers"
)

func Setup(app *fiber.App) {
	//User routes
	app.Post("api/users", controllers.CheckPhone)
}
