package utils

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"strings"
	"time"

	"sushee-backend/config"
	"sushee-backend/httperror/domain"

	"cloud.google.com/go/storage"
	"github.com/rs/zerolog/log"
)

type ClientUploader struct {
	client     *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func NewClientUploader() *ClientUploader {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatal().Msg("Failed to create client")
	}

	return &ClientUploader{
		client:     client,
		projectID:  config.Config.GCSConfig.ProjectID,
		bucketName: config.Config.GCSConfig.Bucket,
		uploadPath: config.Config.GCSConfig.UploadPath,
	}
}

type GCSUploader interface {
	UploadFileFromFileHeader(fileHeader multipart.FileHeader, object string) (string, error)
	DeleteFile(object string) error
}

type gcsUploaderImpl struct {
	clientUploader *ClientUploader
}

type GCSUploaderConfig struct {
	ClientUploader *ClientUploader
}

func NewGCSUploader(c GCSUploaderConfig) GCSUploader {
	return &gcsUploaderImpl{
		clientUploader: c.ClientUploader,
	}
}

func (u *gcsUploaderImpl) checkFileLimit(fileHeader *multipart.FileHeader, maxSize int64, allowedTypes []string) error {
	if fileHeader.Size > maxSize {
		return domain.ErrFileSizeExceedLimit
	}

	fileType := fileHeader.Header.Get("Content-Type")
	for _, allowedType := range allowedTypes {
		if allowedType == fileType {
			return nil
		}
	}
	return domain.ErrFileTypeNotAllowed
}

func (u *gcsUploaderImpl) UploadFileFromFileHeader(fileHeader multipart.FileHeader, object string) (string, error) {
	ctx := context.Background()

	max_size := 2 * 1024 * 1024
	allowed_file_type := []string{"image/jpeg", "image/png", "image/jpg"}
	err := u.checkFileLimit(&fileHeader, int64(max_size), allowed_file_type)
	if err != nil {
		return "", err
	}

	var timeoutSecond int64 = 50
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(timeoutSecond))
	defer cancel()

	fileNameLength := 10
	randString := RandomFileName(fileNameLength)
	object = randString + strings.TrimSpace(strings.ReplaceAll(object, " ", "_"))

	file, err := fileHeader.Open()
	if err != nil {
		return "", domain.ErrUploadFile
	}
	defer file.Close()

	wc := u.clientUploader.client.Bucket(u.clientUploader.bucketName).Object(u.clientUploader.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", domain.ErrUploadFile
	}
	if err := wc.Close(); err != nil {
		return "", domain.ErrUploadFile
	}

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", u.clientUploader.bucketName, u.clientUploader.uploadPath+object), nil
}

func (u *gcsUploaderImpl) DeleteFile(object string) error {
	ctx := context.Background()

	var timeoutSecond int64 = 50
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(timeoutSecond))
	defer cancel()

	fileNameArr := strings.Split(object, "/")
	fileName := fileNameArr[len(fileNameArr)-1]

	if err := u.clientUploader.client.Bucket(u.clientUploader.bucketName).Object(u.clientUploader.uploadPath + fileName).Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", u.clientUploader.uploadPath+fileName, err)
	}

	return nil
}

func RandomFileName(fileNameLength int) string {
	b := make([]byte, fileNameLength)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
