package filemanager

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
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

func CreateNewAudioBook() {

}

func ZipFolder(folderPath, zipFileFolder string, zipFileName string) error {
	_, fileExistsErr := os.Stat(zipFileName)
	switch fileExistsErr {
	case nil:
		return nil
	default:
		mkdirErr := os.MkdirAll(zipFileFolder, fs.ModePerm)
		if mkdirErr != nil {
			return mkdirErr
		}

		zipFile, err := os.Create(zipFileName)
		if err != nil {
			return err
		}
		defer zipFile.Close()

		zipWriter := zip.NewWriter(zipFile)
		defer zipWriter.Close()

		return filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			relativePath, err := filepath.Rel(folderPath, path)
			if err != nil {
				return err
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			zipEntryWriter, err := zipWriter.Create(relativePath)
			if err != nil {
				return err
			}

			_, err = io.Copy(zipEntryWriter, file)
			return err
		})
	}
}
