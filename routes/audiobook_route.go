package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func audioBookRoute(app *fiber.App) {
	//All routes related to users comes here
	app.Post("/GetUser", controllers.RetrieveAudioBookList)
}
