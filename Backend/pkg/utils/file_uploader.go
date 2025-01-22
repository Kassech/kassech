package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// UploadFile handles the file upload and saves it to the specified path
func UploadFile(file *multipart.FileHeader, uploadPath string) (string, error) {
	// Create the upload path if it doesn't exist
	var err = os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Add timestamp to the filename to avoid duplication
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d_%s", timestamp, file.Filename)

	// Create the file on the server
	dst, err := os.Create(filepath.Join(uploadPath, filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Copy the uploaded file to the destination file
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	return filepath.Join(uploadPath, filename), nil
}
func UploadAndAssignPath(file *multipart.FileHeader, directory string, assign func(path string)) error {
	if file != nil {
		path, err := UploadFile(file, directory)
		if err != nil {
			return err
		}
		assign(path)
	}
	return nil
}
