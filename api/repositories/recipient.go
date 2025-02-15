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

//go:generate mockery --name=RecipientRepository --filename=recipient_repository.go --output=../mocks --outpkg=mocks
type RecipientRepository interface {
	CreateRecipient(ctx context.Context, recipient models.Recipient) error
	GetRecipientByID(ctx context.Context, ID uuid.UUID) (*models.Recipient, error)
	GetRecipientByEmail(ctx context.Context, email string) (*models.Recipient, error)
	UpdateRecipient(ctx context.Context, recipient models.Recipient) error
	DeleteRecipient(ctx context.Context, ID uuid.UUID) error
	GetRecipientLitePagedList(ctx context.Context, pagination *models.RecipientBasicInfoPagination) (*models.PaginatedResponse[models.Recipient], error)
}

type recipientRepository struct {
	i  *di.Injector
	DB *gorm.DB
}

func NewRecipientRepository(i *di.Injector) (RecipientRepository, error) {
	DB, err := di.Invoke[*gorm.DB](i)
	if err != nil {
		return nil, fmt.Errorf("invoke DB: %w", err)
	}

	return &recipientRepository{
		i:  i,
		DB: DB,
	}, nil
}

func (r *recipientRepository) CreateRecipient(ctx context.Context, recipient models.Recipient) error {
	if err := r.DB.
		WithContext(ctx).
		Create(&recipient).Error; err != nil {
		return err
	}

	return nil
}

func (r *recipientRepository) GetRecipientByID(ctx context.Context, ID uuid.UUID) (*models.Recipient, error) {
	var recipient models.Recipient

	if err := r.DB.
		WithContext(ctx).
		Where("id = ?", ID).
		First(&recipient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}

	return &recipient, nil
}

func (r *recipientRepository) GetRecipientByEmail(ctx context.Context, email string) (*models.Recipient, error) {
	var recipient models.Recipient

	if err := r.DB.
		WithContext(ctx).
		Where("email = ?", email).
		First(&recipient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}

	return &recipient, nil
}

func (r *recipientRepository) UpdateRecipient(ctx context.Context, recipient models.Recipient) error {
	if err := r.DB.
		WithContext(ctx).
		Save(&recipient).Error; err != nil {
		return err
	}

	return nil
}

func (r *recipientRepository) DeleteRecipient(ctx context.Context, ID uuid.UUID) error {
	if err := r.DB.
		WithContext(ctx).
		Where("id = ?", ID).
		Delete(&models.Recipient{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *recipientRepository) GetRecipientLitePagedList(ctx context.Context, pagination *models.RecipientBasicInfoPagination) (*models.PaginatedResponse[models.Recipient], error) {
	query := r.DB.WithContext(ctx).
		Model(&models.Recipient{})

	if pagination.Q != nil {
		query = query.Where("full_name LIKE ? OR email LIKE ?", fmt.Sprintf("%%%s%%", *pagination.Q), fmt.Sprintf("%%%s%%", *pagination.Q))
	}

	recipients, err := paginate[models.Recipient](query, &pagination.Pagination, &models.Recipient{})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return recipients, nil
}
