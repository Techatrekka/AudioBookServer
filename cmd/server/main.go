package main

import (
	"clos/internal/configs"
	"clos/internal/endpoint"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	configs.ConnectDB()

	// Initialize Gin router
	router := gin.Default()
	router.MaxMultipartMemory = 256 << 20 // Set max upload size to 256 MiB

	// Register endpoint groups
	endpoint.TapeEndpoints(router)
	endpoint.AudiobookEndpoints(router)

	// Start the server on port 8080
	router.Run(":8080")
}
