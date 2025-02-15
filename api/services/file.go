package services

import (
	"context"
	"image"
	"mime/multipart"

	_ "image/jpeg"
	_ "image/png"

	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/models"
)

type FileService interface {
	ValidateImage(ctx context.Context, imageFile *multipart.FileHeader) error
}

type fileService struct {
	i *di.Injector
}

func NewFileService(i *di.Injector) (FileService, error) {
	return &fileService{
		i: i,
	}, nil
}

func (f *fileService) ValidateImage(ctx context.Context, imageFile *multipart.FileHeader) error {
	if imageFile.Size > models.MaxImageSize {
		return models.ErrImageTooLarge
	}

	if !models.AllowedImageTypes[imageFile.Header.Get("Content-Type")] {
		return models.ErrInvalidImageFormat
	}

	file, err := imageFile.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	_, _, err = image.Decode(file)
	if err != nil {
		return models.ErrImageCorrupted
	}

	return nil
}
