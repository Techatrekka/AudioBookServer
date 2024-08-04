package endpoints

import (
	"net/http"
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func TapeEndpoints(router *gin.Engine) {
	//GET Endpoints
	router.GET("/downloadfolder/:folderType/filename/:fileName", controllers.DownloadFile)
	router.GET("/download/:audioId", func(c *gin.Context) {
		audioId := c.Param("audioId")
		controllers.DownloadFolder(c, audioId)
	})
	router.GET("/catalog/:type", controllers.GetCatalogByType)
	router.GET("/audio/:id", func(c *gin.Context) {
		audioId := c.Param("id")
		c.Header("Content-Type", "application/json")

		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write(controllers.ReturnAudioFileData(audioId))
	})

	//POST Endpoints
	router.POST("/upload/:id", controllers.UploadFile)
	router.POST("/uploadListeningHistory/", controllers.UploadListeningHistory)
}
