package main

import (
	"net/http"
	"server/controllers"
	"server/filemanager"
)

func main() {
	filemanager.ListAudioBookFiles()
	filemanager.RetrieveAudioBook("audio.mp3")
	// app := fiber.New()

	// routes.RouteList(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	// })

	// app.Listen(":80")
	//streamAudioBook()
	// Create a basic HTTP server
	http.HandleFunc("/file", controllers.ServeFile)

	// Start the server
	http.ListenAndServe("localhost:8080", nil)
}
