package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"server/filemanager"
	"server/models"

	"github.com/gin-gonic/gin"
)

// this is unfinished
func ReturnCatalog(section string) []byte {
	audiobook := models.Audiobook{AudioFile: "awd", Title: "awd", Author: "awdawd", Synopsis: "awda", Id: "awd", IconLocation: "awdawd"}
	var books [1]models.Audiobook
	books[0] = audiobook
	jsonString, _ := json.Marshal(books)
	return jsonString
}

func ReturnAudioFileData(id string) []byte {
	audiobook := models.Audiobook{AudioFile: "awd", Title: "awd", Author: "awdawd", Synopsis: "awda", Id: "awd", IconLocation: "awdawd"}
	jsonString, _ := json.Marshal(audiobook)
	return jsonString
}

func DownloadFolder(c *gin.Context, audioId string) {
	folderPath := "./Audiobooks/" + audioId + "/"
	zipFileFolder := "./AudiobooksZipped/" + audioId + "/"
	zipFileName := "./AudiobooksZipped/" + audioId + "/files.zip"

	err := filemanager.ZipFolder(folderPath, zipFileFolder, zipFileName)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating zip file: %v", err)
		return
	}
	// we can talk about this if you would like but I think we should reuse the zip files
	// and leave the original files there as a backup
	//defer os.Remove(zipFileName)

	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+zipFileName)
	c.Writer.Header().Set("Content-Type", "application/zip")

	file, err := os.Open(zipFileName)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error opening zip file: %v", err)
		return
	}
	defer file.Close()

	io.Copy(c.Writer, file)
}
