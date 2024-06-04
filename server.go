package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// this is a max upload of 256 MIBs
	router.MaxMultipartMemory = 256 << 20
	router.GET("/downloadImage/:fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")
		filePath := "./Images/" + fileName
		c.FileAttachment(filePath, fileName)
	})
	router.GET("/downloadAudio/:fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")
		filePath := "./Audiobooks/" + fileName
		c.FileAttachment(filePath, fileName)
	})
	router.GET("/downloadAudioBook/:audiobookId", func(c *gin.Context) {
		audiobookId := c.Param("audiobookId")
		filePath := "./Audiobooks/" + audiobookId
		c.FileAttachment(filePath, "fileName")

	})

	router.GET("/download/:filename", controllers.DownloadFolder)
	router.GET("/catalog/:type", func(c *gin.Context) {
		sectionType := c.Param("type")
		c.Header("Content-Type", "application/json")

		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write(controllers.ReturnCatalog(sectionType))
	})
	// router.GET("/audiofile/:id", func(c *gin.Context) {
	// 	fileId := c.Param("id")
	// 	c.JSON(200, controllers.ReturnAudioFileData(fileId))
	// })
	router.POST("/upload/:id", func(c *gin.Context) {
		id := c.Param("id")

		form, err1 := c.MultipartForm()
		if err1 != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err1.Error())
			return
		}
		files := form.File["files"]

		print(len(files))

		uploadDir := "./AudioBooks/" + id
		err := os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			c.String(http.StatusInternalServerError, "Could not create upload directory")
			return
		}

		for _, file := range files {
			filePath := filepath.Join(uploadDir, filepath.Base(file.Filename))
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("Could not save file %s", file.Filename))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("Uploaded %d files successfully", len(files)))
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
