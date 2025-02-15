package services

import (
	"context"
	"time"

	"github.com/G-Villarinho/fast-feet-api/config"
	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockery --name=TokenService --filename=token_service.go --output=../mocks --outpkg=mocks
type TokenService interface {
	CreateToken(ctx context.Context, payload models.TokenPayload) (string, error)
}

type tokenService struct {
	i *di.Injector
}

func NewTokenService(i *di.Injector) (TokenService, error) {

	return &tokenService{
		i: i,
	}, nil
}

func (t *tokenService) CreateToken(ctx context.Context, payload models.TokenPayload) (string, error) {
	claims := models.TokenClaims{
		UserID: payload.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.Env.Session.TokenExp))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env.Session.JWTSecret))
}
