package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//go:generate mockery --name=UserRepository --filename=user_repository.go --output=../mocks --outpkg=mocks
type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByID(ctx context.Context, ID uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByCPF(ctx context.Context, CPF string) (*models.User, error)
	DeleteUser(ctx context.Context, ID uuid.UUID) error
}

type userRepository struct {
	i  *di.Injector
	DB *gorm.DB
}

func NewUserRepository(i *di.Injector) (UserRepository, error) {
	DB, err := di.Invoke[*gorm.DB](i)
	if err != nil {
		return nil, fmt.Errorf("invoke DB: %w", err)
	}

	return &userRepository{
		i:  i,
		DB: DB,
	}, nil
}

func (u *userRepository) CreateUser(ctx context.Context, user models.User) error {
	if err := u.DB.
		WithContext(ctx).
		Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) GetUserByID(ctx context.Context, ID uuid.UUID) (*models.User, error) {
	var user models.User

	if err := u.DB.
		WithContext(ctx).
		Where("id = ?", ID).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}

	return &user, nil
}

func (u *userRepository) GetUserByCPF(ctx context.Context, CPF string) (*models.User, error) {
	var user models.User

	if err := u.DB.
		WithContext(ctx).
		Where("cpf = ?", CPF).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}

	return &user, nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	if err := u.DB.
		WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}

	return &user, nil
}

func (u *userRepository) DeleteUser(ctx context.Context, ID uuid.UUID) error {
	if err := u.DB.
		WithContext(ctx).
		Where("id = ?", ID).
		Delete(&models.User{}).Error; err != nil {
		return err
	}

	return nil
}
