package services

import (
	"context"
	"errors"
	"testing"

	"github.com/G-Villarinho/fast-feet-api/mocks"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/G-Villarinho/fast-feet-api/request"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_CreateAdmin(t *testing.T) {
	t.Run("WhenEmailAlreadyExists_ShouldReturnErrEmailAlreadyExists", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		secureServiceMock := new(mocks.SecureService)

		service := userService{
			ur: mockRepo,
			ss: secureServiceMock,
		}

		payload := models.CreateUserPayload{
			Email: "existing@example.com",
			CPF:   "12345678900",
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockRepo.On("GetUserByID", mock.Anything, userID).
			Return(&models.User{Role: models.Admin}, nil)

		mockRepo.On("GetUserByEmail", mock.Anything, payload.Email).
			Return(&models.User{}, nil)

		err := service.CreateAdmin(ctx, payload)

		assert.ErrorIs(t, err, models.ErrEmailAlreadyExists)
		mockRepo.AssertExpectations(t)
	})

	t.Run("WhenCPFAlreadyExists_ShouldReturnErrCPFAlreadyExists", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		secureServiceMock := new(mocks.SecureService)

		service := userService{
			ur: mockRepo,
			ss: secureServiceMock,
		}

		payload := models.CreateUserPayload{
			Email: "new@example.com",
			CPF:   "12345678900",
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockRepo.On("GetUserByID", mock.Anything, userID).
			Return(&models.User{Role: models.Admin}, nil)

		mockRepo.On("GetUserByEmail", mock.Anything, payload.Email).
			Return(nil, nil)

		mockRepo.On("GetUserByCPF", mock.Anything, payload.CPF).
			Return(&models.User{}, nil)

		err := service.CreateAdmin(ctx, payload)

		assert.ErrorIs(t, err, models.ErrCPFAlreadyExists)
		mockRepo.AssertExpectations(t)
	})

	t.Run("WhenUserIsValid_ShouldCreateSuccessfullyAndReturnNil", func(t *testing.T) {
		userRepoMock := new(mocks.UserRepository)
		secureServiceMock := new(mocks.SecureService)

		service := userService{
			ur: userRepoMock,
			ss: secureServiceMock,
		}

		payload := models.CreateUserPayload{
			Email: "new@example.com",
			CPF:   "12345678900",
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		userRepoMock.On("GetUserByID", mock.Anything, userID).
			Return(&models.User{Role: models.Admin}, nil)

		userRepoMock.On("GetUserByEmail", mock.Anything, payload.Email).
			Return(nil, nil)

		userRepoMock.On("GetUserByCPF", mock.Anything, payload.CPF).
			Return(nil, nil)

		secureServiceMock.On("CreatePassword", mock.Anything).
			Return("password", nil)

		secureServiceMock.On("HashPassword", mock.Anything, mock.Anything).
			Return("$2y$10$hash", nil)

		userRepoMock.On("CreateUser", mock.Anything, mock.Anything).
			Return(nil)

		err := service.CreateAdmin(ctx, payload)

		assert.NoError(t, err)
		userRepoMock.AssertExpectations(t)
	})

	t.Run("WhenCreateUserFails_ShouldReturnError", func(t *testing.T) {
		userRepoMock := new(mocks.UserRepository)
		secureServiceMock := new(mocks.SecureService)

		service := userService{
			ur: userRepoMock,
			ss: secureServiceMock,
		}

		payload := models.CreateUserPayload{
			Email: "new@example.com",
			CPF:   "12345678900",
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		userRepoMock.On("GetUserByID", mock.Anything, userID).
			Return(&models.User{Role: models.Admin}, nil)

		userRepoMock.On("GetUserByEmail", mock.Anything, payload.Email).
			Return(nil, nil)

		userRepoMock.On("GetUserByCPF", mock.Anything, payload.CPF).
			Return(nil, nil)

		secureServiceMock.On("CreatePassword", mock.Anything).
			Return("password", nil)

		secureServiceMock.On("HashPassword", mock.Anything, mock.Anything).
			Return("$2y$10$hash", nil)

		userRepoMock.On("CreateUser", mock.Anything, mock.Anything).
			Return(errors.New("failed to create user"))

		err := service.CreateAdmin(ctx, payload)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create user")
		userRepoMock.AssertExpectations(t)
	})

	t.Run("WhenUserHasNoPermission_ShouldReturnErrInsufficientPermission", func(t *testing.T) {
		userRepoMock := new(mocks.UserRepository)
		secureServiceMock := new(mocks.SecureService)

		service := userService{
			ur: userRepoMock,
			ss: secureServiceMock,
		}

		payload := models.CreateUserPayload{
			Email: "new@example.com",
			CPF:   "12345678900",
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		userRepoMock.On("GetUserByID", mock.Anything, userID).
			Return(&models.User{Role: models.DeliveryMan}, nil)

		err := service.CreateAdmin(ctx, payload)

		assert.ErrorIs(t, err, models.ErrInsufficientPermission)
		userRepoMock.AssertExpectations(t)
	})
}
