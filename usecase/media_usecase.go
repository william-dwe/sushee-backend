package usecase

import (
	"mime/multipart"

	"sushee-backend/httperror/domain"
	"sushee-backend/utils"
)

type MediaUsecase interface {
	UploadFileForBinding(file multipart.FileHeader, object string) (string, error)
	DeleteFile(object string) error
}

type MediaUsecaseConfig struct {
	GCSUploader utils.GCSUploader
}

type mediaUsecaseImpl struct {
	gcsUploader utils.GCSUploader
}

func NewMediaUsecase(c MediaUsecaseConfig) MediaUsecase {
	return &mediaUsecaseImpl{
		gcsUploader: c.GCSUploader,
	}
}

func (u *mediaUsecaseImpl) UploadFileForBinding(file multipart.FileHeader, object string) (string, error) {
	url, err := u.gcsUploader.UploadFileFromFileHeader(file, object)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (u *mediaUsecaseImpl) DeleteFile(object string) error {
	err := u.gcsUploader.DeleteFile(object)
	if err != nil {
		return domain.ErrDeleteFile
	}

	return nil
}
