package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// UploadFile handles the file upload and saves it to the specified path
func UploadFile(r *http.Request, fieldName, uploadPath string) (string, error) {
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		return "", err
	}

	// Retrieve the file from form data
	file, handler, err := r.FormFile(fieldName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create the upload path if it doesn't exist
	err = os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Create the file on the server
	dst, err := os.Create(filepath.Join(uploadPath, handler.Filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return handler.Filename, nil
}
