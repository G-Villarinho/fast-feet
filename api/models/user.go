package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/G-Villarinho/fast-feet-api/utils"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound          = errors.New("user not found in the database")
	ErrEmailAlreadyExists    = errors.New("user with same e-mail already exists")
	ErrCPFAlreadyExists      = errors.New("user with same CPF already exists")
	ErrUserNotFoundInContext = errors.New("user not found in the context")
	ErrUserBlocked           = errors.New("user is blocked")
)

type Status string

const (
	ActiveStatus  Status = "ACTIVE"
	BlockedStatus Status = "BLOCKED"
)

type User struct {
	BaseModel
	FullName     string       `gorm:"not null"`
	CPF          string       `gorm:"not null;unique"`
	Email        string       `gorm:"not null;unique"`
	PasswordHash string       `gorm:"not null"`
	Status       Status       `gorm:"not null;default:'ACTIVE';index"`
	Role         Role         `gorm:"not null;index"`
	BlockedAt    sql.NullTime `gorm:"default:null"`

	DeliverymanOrders []Order `gorm:"foreignKey:DeliverymanID;references:ID"`
}

type CreateUserPayload struct {
	FullName string `json:"fullName" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	CPF      string `json:"cpf" validate:"required,cpf"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
	Email    string    `json:"email"`
	Role     Role      `json:"role"`
}

func (cup *CreateUserPayload) ToUser(passwordHash string, role Role) *User {
	return &User{
		BaseModel: BaseModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		FullName:     cup.FullName,
		CPF:          utils.RemoveCPFFormat(cup.CPF),
		Email:        cup.Email,
		PasswordHash: passwordHash,
		Role:         role,
	}
}

func (u *User) ToUserResponse() *UserResponse {
	return &UserResponse{
		ID:       u.ID,
		FullName: u.FullName,
		Email:    u.Email,
		Role:     u.Role,
	}
}
