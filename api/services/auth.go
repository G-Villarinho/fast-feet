package services

import (
	"context"
	"fmt"

	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/G-Villarinho/fast-feet-api/repositories"
	"github.com/G-Villarinho/fast-feet-api/utils"
)

//go:generate mockery --name=AuthService --filename=auth_service.go --output=../mocks --outpkg=mocks
type AuthService interface {
	Login(ctx context.Context, payload models.LoginPayload) (*models.LoginResponse, error)
}

type authService struct {
	i  *di.Injector
	ss SecureService
	ts TokenService
	ur repositories.UserRepository
}

func NewAuthService(i *di.Injector) (AuthService, error) {
	ss, err := di.Invoke[SecureService](i)
	if err != nil {
		return nil, fmt.Errorf("invoke secure service: %w", err)
	}

	ts, err := di.Invoke[TokenService](i)
	if err != nil {
		return nil, fmt.Errorf("invoke token service: %w", err)
	}

	ur, err := di.Invoke[repositories.UserRepository](i)
	if err != nil {
		return nil, fmt.Errorf("invoke user repository: %w", err)
	}

	return &authService{
		i:  i,
		ss: ss,
		ts: ts,
		ur: ur,
	}, nil
}

func (a *authService) Login(ctx context.Context, payload models.LoginPayload) (*models.LoginResponse, error) {
	userFromCPF, err := a.ur.GetUserByCPF(ctx, utils.RemoveCPFFormat(payload.CPF))
	if err != nil {
		return nil, fmt.Errorf("get user by CPF: %w", err)
	}

	if userFromCPF == nil {
		return nil, models.ErrInvalidCredentials
	}

	if userFromCPF.Status == models.BlockedStatus {
		return nil, models.ErrUserBlocked
	}

	if err := a.ss.CheckPassword(ctx, userFromCPF.PasswordHash, payload.Password); err != nil {
		return nil, models.ErrInvalidCredentials
	}

	tokenPayload := models.TokenPayload{
		UserID: userFromCPF.ID,
	}

	token, err := a.ts.CreateToken(ctx, tokenPayload)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
	}, nil
}
