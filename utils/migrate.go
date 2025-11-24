package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func UploadFiles(files []*multipart.FileHeader, uploadPath string) ([]string, error) {
	var paths []string

	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.MkdirAll(uploadPath, os.ModePerm)
	}

	for _, file := range files {
		timestamp := time.Now().UnixNano()
		filename := fmt.Sprintf("%d_%s", timestamp, file.Filename)
		fullPath := filepath.Join(uploadPath, filename)

		if err := saveFile(file, fullPath); err != nil {
			return nil, err
		}
		paths = append(paths, fullPath)
	}

	return paths, nil
}

func saveFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
