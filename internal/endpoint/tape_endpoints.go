package endpoint

import (
	"clos/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TapeEndpoints(router *gin.Engine) {
	//GET Endpoints
	router.GET("/downloadfile/:fileId", controller.DownloadImage)
	router.GET("/download/:audioId", func(c *gin.Context) {
		audioId := c.Param("audioId")
		controller.DownloadFolder(c, audioId)
	})
	router.GET("/catalog/:type", controller.GetCatalogByType)
	router.GET("/audio/:id", func(c *gin.Context) {
		audioId := c.Param("id")
		c.Header("Content-Type", "application/json")

		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write(controller.ReturnAudioFileData(audioId))
	})
	router.GET("/getListeningHistory/", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write(controller.GetListeningHistory(c))
	})

	//POST Endpoints
	router.POST("/upload/:id", controller.UploadFile)
	router.POST("/uploadListeningHistory/", controller.UploadListeningHistory)
}
