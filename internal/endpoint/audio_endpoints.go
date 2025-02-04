package endpoint

import (
	"clos/internal/controller"

	"github.com/gin-gonic/gin"
)

// AudiobookEndpoints sets up routes for audiobook-related functionality.
func AudiobookEndpoints(router *gin.Engine) {
	router.GET("/audiobooks", controller.GetUploadedAudiobooks)
}
