package models

import (
	"database/sql"
	"errors"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

var (
	ErrOrderNotFound               = errors.New("order not found in database")
	ErrCannotTransitionToDelivered = errors.New("cannot transition to 'Delivered' without passing through 'PicknUp'")
	ErrCannotTransitionToPicknUp   = errors.New("cannot transition to 'PicknUp' unless the order is in 'Waiting' status")
	ErrNotAssignedToOrder          = errors.New("delivery man is not assigned to this order")
)

type OrderStatus string

const (
	Waiting OrderStatus = "WAITING"
	PicknUp OrderStatus = "PICKN_UP"
	Done    OrderStatus = "DONE"
)

type Order struct {
	BaseModel
	Title        string       `gorm:"not null"`
	TrackingCode uuid.UUID    `gorm:"not null"`
	Status       OrderStatus  `gorm:"not null;default:'WAITING';index"`
	IsReturned   bool         `gorm:"not null"`
	PicknUpAt    sql.NullTime `gorm:"default:null"`
	DeliveryAt   sql.NullTime `gorm:"default:null"`

	DeliverymanID *uuid.UUID `gorm:"type:uuid;null;default:null"`
	Deliveryman   User       `gorm:"foreignKey:DeliverymanID;references:ID"`

	RecipientID uuid.UUID `gorm:"type:uuid;not null"`
	Recipient   Recipient `gorm:"foreignKey:RecipientID;references:ID"`
}

type CreateOrderPayload struct {
	Title       string    `json:"title" validate:"required,max=255"`
	RecipientID uuid.UUID `json:"recipientId" validate:"required"`
}

type DeliverOrderPayload struct {
	OrderImage *multipart.FileHeader `json:"title" validate:"required"`
}

type CreateOrderResponse struct {
	OrderID uuid.UUID `json:"orderId"`
}

type OrderResponse struct {
	ID        uuid.UUID   `json:"id"`
	Title     string      `json:"title"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"createdAt"`
}

type PickUpOrderResponse struct {
	PicknUpAt time.Time `json:"picknUpAt"`
}

type DeliverOrderResponse struct {
	DeliveryAt time.Time `json:"deliveryAt"`
}

type OrderDetailsResponse struct {
	ID               uuid.UUID   `json:"id"`
	Status           OrderStatus `json:"status"`
	RecipientName    string      `json:"recipientName"`
	RecipientAddress string      `json:"recipientAddress"`
	RecipientZipcode int         `json:"recipientZipcode"`
	CreatedAt        time.Time   `json:"createdAt"`
	PicknUpAt        *time.Time  `json:"picknUpAt,omitempty"`
	DeliveryAt       *time.Time  `json:"deliveryAt,omitempty"`
}

func (p *CreateOrderPayload) ToOrder() *Order {
	return &Order{
		BaseModel: BaseModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		Title:        p.Title,
		RecipientID:  p.RecipientID,
		IsReturned:   false,
		TrackingCode: uuid.New(),
		Status:       Waiting,
	}
}

func (o *Order) ToOrderResponse() *OrderResponse {
	return &OrderResponse{
		ID:        o.ID,
		Title:     o.Title,
		Status:    o.Status,
		CreatedAt: o.CreatedAt,
	}
}

func (o *Order) ToOrderDetailsResponse() *OrderDetailsResponse {
	var picknUpAt *time.Time
	if o.PicknUpAt.Valid {
		picknUpAt = &o.PicknUpAt.Time
	}

	var deliveryAt *time.Time
	if o.DeliveryAt.Valid {
		deliveryAt = &o.DeliveryAt.Time
	}

	return &OrderDetailsResponse{
		ID:               o.ID,
		Status:           o.Status,
		RecipientName:    o.Recipient.FullName,
		RecipientAddress: o.Recipient.Address,
		RecipientZipcode: o.Recipient.Zipcode,
		CreatedAt:        o.CreatedAt,
		PicknUpAt:        picknUpAt,
		DeliveryAt:       deliveryAt,
	}
}
