package main

import (
	"fmt"
	"server/configs"
	"server/endpoints" // Ensure this is imported

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()

	router := gin.Default()
	router.MaxMultipartMemory = 256 << 20 // this is a max upload of 256 MIBs

	endpoints.TapeEndpoints(router)      // Register existing endpoints
	endpoints.AudiobookEndpoints(router) // Register audiobook endpoints

	// Print all registered routes for debugging
	fmt.Println("Routes registered:")
	for _, route := range router.Routes() {
		fmt.Println(route.Method, route.Path)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.Run(":8080") // Start the server
}
