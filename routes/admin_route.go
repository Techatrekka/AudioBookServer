package routes

import "github.com/gofiber/fiber/v2"

func RouteList(app *fiber.App) {
	audioBookRoute(app)
}
