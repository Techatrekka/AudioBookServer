package controller

import (
	"clos/internal/filemanager"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUploadedAudiobooks handles listing all files in the Audiobooks directory.
func GetUploadedAudiobooks(c *gin.Context) {
	files, err := filemanager.ListAudioBookFilesDetailed()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}
