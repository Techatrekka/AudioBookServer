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

// func RetrieveAudioBookList(context *fiber.Ctx) error {
// 	return context.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "uwu"}})
// }

// func retrieveAudioBook(context *fiber.Ctx) {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			audioFile, err := os.ReadFile("./audio.mp3")

// 			if err != nil {
// 				log.Println(err)
// 				os.Exit(1)
// 			}

// 			w.Header().Set("Access-Control-Allow-Origin", "*")
// 			w.Header().Set("Content-Type", "audio/mpeg")
// 			audioChunks := bytes.NewBuffer(audioFile)

// 			if _, err := audioChunks.WriteTo(w); err != nil {
// 				log.Println(err)
// 			}
// 			break
// 		}
// 	})

// 	var ServerPort = os.Getenv("PORT")

// 	if ServerPort == "" {
// 		ServerPort = "4444"
// 	}

// 	log.Printf("The server is streaming on http://localhost:%s/stream", ServerPort)
// 	log.Fatal(http.ListenAndServe(":"+ServerPort, router), nil)
// }

// func streamAudioBook() {
// 	router := mux.NewRouter()
// 	go alternateStream(router)
// 	router.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			audioFile := filemanager.RetrieveAudioBook("audio.mp3")
// 			// if err != nil {
// 			// 	log.Println(err)
// 			// 	os.Exit(1)
// 			// }

// 			w.Header().Set("Access-Control-Allow-Origin", "*")
// 			w.Header().Set("Content-Type", "audio/mpeg")
// 			audioChunks := bytes.NewBuffer(audioFile)

// 			if _, err := audioChunks.WriteTo(w); err != nil {
// 				log.Println(err)
// 			}
// 			break
// 		}
// 	})

// 	var ServerPort = os.Getenv("PORT")

// 	if ServerPort == "" {
// 		ServerPort = "4444"
// 	}

// 	log.Printf("The server is streaming on http://localhost:%s/stream", ServerPort)
// 	log.Fatal(http.ListenAndServe(":"+ServerPort, router), nil)
// }

// func alternateStream(router *mux.Router) {
// 	router.HandleFunc("/stream2", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			audioFile := filemanager.RetrieveAudioBook("audio.mp3")
// 			// if err != nil {
// 			// 	log.Println(err)
// 			// 	os.Exit(1)
// 			// }

// 			w.Header().Set("Access-Control-Allow-Origin", "*")
// 			w.Header().Set("Content-Type", "audio/mpeg")
// 			audioChunks := bytes.NewBuffer(audioFile)

// 			if _, err := audioChunks.WriteTo(w); err != nil {
// 				log.Println(err)
// 			}
// 			break
// 		}
// 	})

// 	var ServerPort = os.Getenv("PORT")

// 	if ServerPort == "" {
// 		ServerPort = "44444"
// 	}

// 	log.Printf("The server is streaming on http://localhost:%s/stream", ServerPort)
// 	log.Fatal(http.ListenAndServe(":"+ServerPort, router), nil)
// }

// func DownloadFile(w http.ResponseWriter, r *http.Request) {
// 	Openfile, err := os.Open("C://Users/User/Desktop/Download/Download.xlsx") //Open the file to be downloaded later
// 	defer Openfile.Close()                                                    //Close after function return

// 	if err != nil {
// 		http.Error(w, "File not found.", 404) //return 404 if file is not found
// 		return
// 	}

// 	tempBuffer := make([]byte, 512)                       //Create a byte array to read the file later
// 	Openfile.Read(tempBuffer)                             //Read the file into  byte
// 	FileContentType := http.DetectContentType(tempBuffer) //Get file header

// 	FileStat, _ := Openfile.Stat()                     //Get info from file
// 	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

// 	Filename := "demo_download"

// 	//Set the headers
// 	w.Header().Set("Content-Type", FileContentType+";"+Filename)
// 	w.Header().Set("Content-Length", FileSize)

// 	Openfile.Seek(0, 0)  //We read 512 bytes from the file already so we reset the offset back to 0
// 	io.Copy(w, Openfile) //'Copy' the file to the client
// }

// func ReturnAudioBook(w http.ResponseWriter, r *http.Request) {
// 	//stringList := []string{"apple", "banana", "orange", "grape"}
// 	audiobook := Audiobook{AudioFile: "awd", Title: "awd", Author: "awdawd", Synopsis: "awda", Id: "awd", IconLocation: "awdawd"}
// 	var books [1]Audiobook
// 	books[0] = audiobook
// 	// Convert the list to JSON
// 	jsonString, err := json.Marshal(books)
// 	// 	[Audiobook{
// 	// 	audiobook.AudioFile,
// 	// 	audiobook.Title,
// 	// 	audiobook.Author,
// 	// 	audiobook.Synopsis,
// 	// 	audiobook.Id,
// 	// 	audiobook.IconLocation,
// 	// }]
// 	if err != nil {
// 		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the response Content-Type header to application/json
// 	w.Header().Set("Content-Type", "application/json")

// 	// Write the JSON response
// 	_, err = w.Write(jsonString)
// 	if err != nil {
// 		log.Println("Error writing response:", err)
// 	}
// }

// func ReturnList(w http.ResponseWriter, r *http.Request) {
// 	stringList := []string{"apple", "banana", "orange", "grape"}

// 	// Convert the list to JSON
// 	jsonString, err := json.Marshal(stringList)
// 	if err != nil {
// 		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the response Content-Type header to application/json
// 	w.Header().Set("Content-Type", "application/json")

// 	// Write the JSON response
// 	_, err = w.Write(jsonString)
// 	if err != nil {
// 		log.Println("Error writing response:", err)
// 	}
// }

// func ServeFile(w http.ResponseWriter, r *http.Request) {
// 	// Open the file
// 	file, err := os.Open("./Audiobooks/audio.mp3")
// 	if err != nil {
// 		http.Error(w, "File not found", http.StatusNotFound)
// 		return
// 	}
// 	defer file.Close()

// 	// Get file information
// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the Content-Type header based on file extension
// 	contentType := http.DetectContentType([]byte(fileInfo.Name()))
// 	w.Header().Set("Content-Type", contentType)

// 	w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())

// 	// Use ServeContent to efficiently serve the file
// 	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
// }

func ReturnCatalog(section string) []byte {
	// stringList := []string{"apple", "banana", "orange", "grape"}
	// jsonString, _ := json.Marshal(stringList)
	// return jsonString
	audiobook := models.Audiobook{AudioFile: "awd", Title: "awd", Author: "awdawd", Synopsis: "awda", Id: "awd", IconLocation: "awdawd"}
	var books [1]models.Audiobook
	books[0] = audiobook
	// Convert the list to JSON
	jsonString, _ := json.Marshal(books)
	return jsonString
	// 	[Audiobook{
	// 	audiobook.AudioFile,
	// 	audiobook.Title,
	// 	audiobook.Author,
	// 	audiobook.Synopsis,
	// 	audiobook.Id,
	// 	audiobook.IconLocation,
	// }]
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
