package models

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidCredentials     = errors.New("invalid credencials: CPF or password wrong")
	ErrTokenNotFoundInContext = errors.New("token not found in the context")
)

type TokenClaims struct {
	UserID uuid.UUID `json:"sub"`
	jwt.RegisteredClaims
}

type TokenPayload struct {
	UserID uuid.UUID `json:"sub"`
}

type LoginPayload struct {
	CPF      string `json:"cpf" validate:"required,cpf"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
