package filemanager

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// Struct to hold file details
type FileInfo struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// Function to list files with detailed information (name and size)
func ListAudioBookFilesDetailed() ([]FileInfo, error) {
	files, err := os.ReadDir("./Audiobooks")
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	var fileInfos []FileInfo
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return nil, fmt.Errorf("error getting file info: %w", err)
		}
		fileInfos = append(fileInfos, FileInfo{
			Name: info.Name(),
			Size: info.Size(),
		})
	}
	return fileInfos, nil
}

// Function to list files (returns a slice of DirEntry)
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

// Function to retrieve a specific audiobook file
func RetrieveAudioBook(Name string) []byte {
	file, err := os.ReadFile("./Audiobooks/" + Name)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return file
}

// Placeholder function to create a new audiobook (currently empty)
func CreateNewAudioBook() {
	// Implementation for creating new audiobooks can go here
}

// Function to zip a folder and save the zip file to a specified location
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
