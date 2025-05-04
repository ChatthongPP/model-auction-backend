package usecase

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"backend-service/config"
	"backend-service/internal/domain"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

const uploadDir = "./uploads"

func (uc *Usecase) UploadMedia(file *multipart.FileHeader, topic string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	conf := config.GetAppconfig()

	if !isSupportedFileType(file.Filename) {
		return "", domain.ErrUnsupportedFileType
	}

	// Create topic folder
	topicPath := filepath.Join(uploadDir, topic)
	if err := os.MkdirAll(topicPath, os.ModePerm); err != nil {
		return "", err
	}

	// Unique filename using UUID
	ext := filepath.Ext(file.Filename)
	filename := strings.TrimSuffix(file.Filename, ext)
	newFilename := fmt.Sprintf("%s-%s%s", filename, uuid.New().String(), ext)
	dstPath := filepath.Join(topicPath, newFilename)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	// Return public URL
	url := fmt.Sprintf("%s/media/%s/%s", conf.Host, topic, newFilename)
	return url, nil
}

func (uc *Usecase) GetMediaPath(filename string) (string, error) {
	path := filepath.Join(uploadDir, filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", domain.ErrFileNotFound
	}

	return path, nil
}

func isSupportedFileType(filename string) bool {
	lower := strings.ToLower(filename)
	return strings.HasSuffix(lower, ".jpg") ||
		strings.HasSuffix(lower, ".jpeg") ||
		strings.HasSuffix(lower, ".png") ||
		strings.HasSuffix(lower, ".mp4") ||
		strings.HasSuffix(lower, ".mov")
}
