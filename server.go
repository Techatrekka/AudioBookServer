package main

import (
	"server/configs"
	"server/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()
	//database.SignInWithEmailPassword("test@email.com", "test")

	router := gin.Default() // this is a max upload of 256 MIBs
	router.MaxMultipartMemory = 256 << 20

	endpoints.TapeEndpoints(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
