package models

import "errors"

var (
	ErrImageTooLarge      = errors.New("image size exceeds 5MB")
	ErrInvalidImageFormat = errors.New("invalid image format")
	ErrImageCorrupted     = errors.New("image is corrupted or has an invalid format")
	ErrOpenImage          = errors.New("error to open image")
)

const MaxImageSize = 5 * 1024 * 1024

var AllowedImageTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}
