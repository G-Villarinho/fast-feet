package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrRecipientNotFound = errors.New("recipient not found in database")
)

type Recipient struct {
	BaseModel
	FullName     string `gorm:"not null"`
	Email        string `gorm:"not null;unique"`
	State        string `gorm:"not null"`
	City         string `gorm:"not null"`
	Neighborhood string `gorm:"not null"`
	Address      string `gorm:"not null"`
	Zipcode      int    `gorm:"not null"`

	Orders []Order `gorm:"foreignKey:RecipientID;references:ID"`
}

type RecipientBasicInfoPagination struct {
	Pagination
	Q *string `json:"q"`
}

type CreateRecipientPayload struct {
	FullName     string `json:"fullName" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	State        string `json:"state" validate:"required"`
	City         string `json:"city" validate:"required"`
	Neighborhood string `json:"neighborhood" validate:"required"`
	Address      string `json:"address" validate:"required"`
	Zipcode      int    `json:"zipcode" validate:"required"`
}

type UpdateRecipientPayload struct {
	FullName     string `json:"fullName" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	State        string `json:"state" validate:"required"`
	City         string `json:"city" validate:"required"`
	Neighborhood string `json:"neighborhood" validate:"required"`
	Address      string `json:"address" validate:"required"`
	Zipcode      int    `json:"zipcode" validate:"required"`
}

type RecipientResponse struct {
	ID           uuid.UUID `json:"id"`
	FullName     string    `json:"fullName"`
	Email        string    `json:"email"`
	State        string    `json:"state"`
	City         string    `json:"city"`
	Neighborhood string    `json:"neighborhood"`
	Address      string    `json:"address"`
	Zipcode      int       `json:"zipcode"`
	CreatedAt    time.Time `json:"createdAt"`
}

type RecipientBasicInfoResponse struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
	Email    string    `json:"email"`
}

type CreateRecipientResponse struct {
	RecipientID uuid.UUID `json:"recipientId"`
}

func (p CreateRecipientPayload) ToRecipient() *Recipient {
	return &Recipient{
		BaseModel: BaseModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		FullName:     p.FullName,
		Email:        p.Email,
		State:        p.State,
		City:         p.City,
		Neighborhood: p.Neighborhood,
		Address:      p.Address,
		Zipcode:      p.Zipcode,
	}
}

func (r *Recipient) ToRecipientResponse() *RecipientResponse {
	return &RecipientResponse{
		ID:           r.ID,
		FullName:     r.FullName,
		Email:        r.Email,
		State:        r.State,
		City:         r.City,
		Neighborhood: r.Neighborhood,
		Address:      r.Address,
		Zipcode:      r.Zipcode,
		CreatedAt:    r.CreatedAt,
	}
}

func (r *Recipient) ToRecipientBasicInfoResponse() *RecipientBasicInfoResponse {
	return &RecipientBasicInfoResponse{
		ID:       r.ID,
		FullName: r.FullName,
		Email:    r.Email,
	}
}

func (r *Recipient) ApplyUpdates(p *UpdateRecipientPayload) {
	r.FullName = p.FullName
	r.Email = p.Email
	r.State = p.State
	r.City = p.City
	r.Neighborhood = p.Neighborhood
	r.Address = p.Address
	r.Zipcode = p.Zipcode
}
