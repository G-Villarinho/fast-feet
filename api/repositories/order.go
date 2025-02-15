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

type OrderRepository interface {
	CreateOrder(ctx context.Context, order models.Order) error
	GetOrderByID(ctx context.Context, ID uuid.UUID) (*models.Order, error)
	GetOrderByTrackingCode(ctx context.Context, trackingCode uuid.UUID) (*models.Order, error)
	DeleteOrder(ctx context.Context, ID uuid.UUID) error
	UpdateOrder(ctx context.Context, order models.Order) error
	GetOrdersPagedList(ctx context.Context, deliveryManID *uuid.UUID, pagination *models.Pagination) (*models.PaginatedResponse[models.Order], error)
}

type orderRepository struct {
	i  *di.Injector
	DB *gorm.DB
}

func NewOrderRepository(i *di.Injector) (OrderRepository, error) {
	DB, err := di.Invoke[*gorm.DB](i)
	if err != nil {
		return nil, fmt.Errorf("invoke DB: %w", err)
	}

	return &orderRepository{
		i:  i,
		DB: DB,
	}, nil
}

func (o *orderRepository) CreateOrder(ctx context.Context, order models.Order) error {
	if err := o.DB.
		WithContext(ctx).
		Create(&order).Error; err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) GetOrderByID(ctx context.Context, ID uuid.UUID) (*models.Order, error) {
	var order models.Order

	if err := o.DB.
		WithContext(ctx).
		Where("id = ?", ID).
		First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}

	return &order, nil
}

func (o *orderRepository) DeleteOrder(ctx context.Context, ID uuid.UUID) error {
	if err := o.DB.
		WithContext(ctx).
		Where("id = ?", ID).
		Delete(&models.Order{}).Error; err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) UpdateOrder(ctx context.Context, order models.Order) error {
	if err := o.DB.
		WithContext(ctx).
		Save(&order).
		Error; err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) GetOrderByTrackingCode(ctx context.Context, trackingCode uuid.UUID) (*models.Order, error) {
	var order models.Order

	if err := o.DB.
		WithContext(ctx).
		Where("tracking_code = ?", trackingCode).
		First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}

	return &order, nil
}

func (o *orderRepository) GetOrdersPagedList(ctx context.Context, deliveryManID *uuid.UUID, pagination *models.Pagination) (*models.PaginatedResponse[models.Order], error) {
	query := o.DB.WithContext(ctx).
		Model(&models.Order{})

	if deliveryManID != nil {
		query = query.Where("status = ? OR deliveryman_id = ?", models.Waiting, deliveryManID)
	}

	orders, err := paginate[models.Order](query, pagination, &models.Order{})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return orders, nil
}
