package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"server/database"
	"server/filemanager"
	"server/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type TapeResponse struct {
	TapeId      string `json:"tape_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Synopsis    string `json:"synopsis"`
	IsAudiobook string `json:"is_audiobook"`
	Tags        string `json:"tags"`
}

func ReturnCatalog(section string) []byte {
	result := database.SelectAll("Tape")
	var tapes []models.Tape
	err := json.Unmarshal(result, &tapes)
	if err != nil {
		print(err)
	}
	if len(tapes) == 0 {
		panic("Empty array returned from DB")
	}
	var TapeResponses []TapeResponse
	for i := 0; i < len(tapes); i++ {
		newTapeResponse := TapeResponse{
			TapeId:      strconv.Itoa(tapes[i].TapeId),
			Title:       tapes[i].Title,
			Author:      tapes[i].Author,
			Synopsis:    tapes[i].Synopsis,
			IsAudiobook: strconv.FormatBool(tapes[i].IsAudiobook),
			Tags:        strings.Trim(strings.Join(strings.Split(fmt.Sprint(tapes[i].Tags), " "), ","), "[]"),
		}
		TapeResponses = append(TapeResponses, newTapeResponse)
	}

	jsonString, _ := json.Marshal(TapeResponses)
	return jsonString
}

func ReturnAudioFileData(id string) []byte {
	result := database.SelectById("Tape", id)
	var value []models.Tape
	err := json.Unmarshal(result, &value)
	if err != nil {
		print(err)
	}
	if len(value) == 0 {
		panic("Empty array returned from DB")
	}
	jsonString, _ := json.Marshal(value[0])
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

func UploadFile(c *gin.Context) {
	id := c.Param("id")

	form, err1 := c.MultipartForm()
	if err1 != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err1.Error())
		return
	}
	files := form.File["files"]

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
}

func DownloadImage(c *gin.Context) {
	fileName := c.Param("fileId")
	filePath := "./Audiobooks/" + fileName + "/image.png"

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	c.Data(http.StatusOK, "image/png", fileBytes)
}

func GetCatalogByType(c *gin.Context) {
	sectionType := c.Param("type")
	c.Header("Content-Type", "application/json")

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(ReturnCatalog(sectionType))
}

func UploadListeningHistory(c *gin.Context) {
	var req models.ListeningHistory
	c.BindJSON(&req)
	database.UploadObjectToTable("ListeningHistory", req)
}

func GetListeningHistory(c *gin.Context) []byte {
	var req []models.ListeningHistory
	userId := c.Query("user_id")
	tapeId := c.Query("tape_id")
	result := database.SelectByCompositeId("ListeningHistory", "user", "tape", userId, tapeId)
	err := json.Unmarshal(result, &req)
	if err != nil {
		print("Unmarshalling failed")
	}
	return result
	// print(req[len(req)-1].CurrentChapter)
}
