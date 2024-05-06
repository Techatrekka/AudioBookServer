package filemanager

import (
	"fmt"
	"io/fs"
	"os"
)

func ListAudioBookFiles() []fs.DirEntry {
	files, err := os.ReadDir("./AudioBooks")
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	return files
}

func RetrieveAudioBook(Name string) []byte {
	file, err := os.ReadFile("./Audiobooks/" + Name)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return file
}
