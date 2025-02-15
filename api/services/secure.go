package services

import (
	"context"
	"crypto/rand"

	"github.com/G-Villarinho/fast-feet-api/di"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockery --name=SecureService --filename=secure_service.go --output=../mocks --outpkg=mocks
type SecureService interface {
	CreatePassword(ctx context.Context) (string, error)
	HashPassword(ctx context.Context, password string) (string, error)
	CheckPassword(ctx context.Context, hashedPassword, password string) error
}

type secureService struct {
	i *di.Injector
}

func NewSecureService(i *di.Injector) (SecureService, error) {
	return &secureService{
		i: i,
	}, nil
}

func (s *secureService) CreatePassword(ctx context.Context) (string, error) {
	password, err := generateRandomPassword(8)
	if err != nil {
		return "", err
	}

	return password, nil
}

func (s *secureService) CheckPassword(ctx context.Context, hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *secureService) HashPassword(ctx context.Context, password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func generateRandomPassword(size int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	for i, v := range b {
		b[i] = letters[int(v)%len(letters)]
	}
	return string(b), nil
}
