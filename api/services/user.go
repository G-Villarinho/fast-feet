package services

import (
	"context"
	"fmt"

	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/G-Villarinho/fast-feet-api/repositories"
	"github.com/G-Villarinho/fast-feet-api/request"
)

//go:generate mockery --name=UserService --filename=user_service.go --output=../mocks --outpkg=mocks
type UserService interface {
	CreateAdmin(ctx context.Context, payload models.CreateUserPayload) error
	CreateDeliveryMan(ctx context.Context, payload models.CreateUserPayload) error
	GetUser(ctx context.Context) (*models.UserResponse, error)
}

type userService struct {
	i  *di.Injector
	ss SecureService
	ur repositories.UserRepository
}

func NewUserService(i *di.Injector) (UserService, error) {
	ss, err := di.Invoke[SecureService](i)
	if err != nil {
		return nil, fmt.Errorf("invoke secure service: %w", err)
	}

	ur, err := di.Invoke[repositories.UserRepository](i)
	if err != nil {
		return nil, fmt.Errorf("invoke user repository: %w", err)
	}

	return &userService{
		i:  i,
		ur: ur,
		ss: ss,
	}, nil
}

func (u *userService) CreateAdmin(ctx context.Context, payload models.CreateUserPayload) error {
	if err := u.createUser(ctx, payload, models.Admin); err != nil {
		return err
	}

	return nil
}

func (u *userService) CreateDeliveryMan(ctx context.Context, payload models.CreateUserPayload) error {
	if err := u.createUser(ctx, payload, models.DeliveryMan); err != nil {
		return err
	}

	return nil
}

func (u *userService) GetUser(ctx context.Context) (*models.UserResponse, error) {
	userID, found := request.UserID(ctx)
	if !found {
		return nil, models.ErrUserNotFoundInContext
	}

	user, err := u.ur.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if user == nil {
		return nil, models.ErrUserNotFound
	}

	return user.ToUserResponse(), nil
}

func (u *userService) createUser(ctx context.Context, payload models.CreateUserPayload, role models.Role) error {
	userID, found := request.UserID(ctx)
	if !found {
		return models.ErrUserNotFoundInContext
	}

	authUser, err := u.ur.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if models.Cannot(authUser.Role, models.Create, models.Users) {
		return models.ErrInsufficientPermission
	}

	userWithSameEmail, err := u.ur.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return fmt.Errorf("get user by email: %w", err)
	}

	if userWithSameEmail != nil {
		return models.ErrEmailAlreadyExists
	}

	userWithSameCPF, err := u.ur.GetUserByCPF(ctx, payload.CPF)
	if err != nil {
		return fmt.Errorf("get user by CPF: %w", err)
	}

	if userWithSameCPF != nil {
		return models.ErrCPFAlreadyExists
	}

	password, err := u.ss.CreatePassword(ctx)
	if err != nil {
		return fmt.Errorf("create password: %w", err)
	}

	passwordHash, err := u.ss.HashPassword(ctx, password)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	user := payload.ToUser(passwordHash, role)

	if err := u.ur.CreateUser(ctx, *user); err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
