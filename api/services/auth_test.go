package services

import (
	"context"
	"errors"
	"testing"

	"github.com/G-Villarinho/fast-feet-api/mocks"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_Login(t *testing.T) {
	t.Run("WhenUserNotFound_ShouldReturnErrInvalidCredentials", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		mockTokenService := new(mocks.TokenService)
		mockSecureService := new(mocks.SecureService)

		service := authService{
			ur: mockRepo,
			ts: mockTokenService,
			ss: mockSecureService,
		}

		payload := models.LoginPayload{
			CPF:      "12345678900",
			Password: "password",
		}

		mockRepo.On("GetUserByCPF", mock.Anything, payload.CPF).
			Return(nil, nil)

		response, err := service.Login(context.Background(), payload)

		assert.ErrorIs(t, err, models.ErrInvalidCredentials)
		assert.Nil(t, response)
		mockRepo.AssertExpectations(t)
	})

	t.Run("WhenUserIsBlocked_ShouldReturnErrUserBlocked", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		mockTokenService := new(mocks.TokenService)
		mockSecureService := new(mocks.SecureService)

		service := authService{
			ur: mockRepo,
			ts: mockTokenService,
			ss: mockSecureService,
		}

		payload := models.LoginPayload{
			CPF:      "12345678900",
			Password: "password",
		}

		user := &models.User{
			Status: models.BlockedStatus,
		}

		mockRepo.On("GetUserByCPF", mock.Anything, payload.CPF).
			Return(user, nil)

		response, err := service.Login(context.Background(), payload)

		assert.ErrorIs(t, err, models.ErrUserBlocked)
		assert.Nil(t, response)
		mockRepo.AssertExpectations(t)
	})

	t.Run("WhenPasswordIsIncorrect_ShouldReturnErrInvalidCredentials", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		mockTokenService := new(mocks.TokenService)
		mockSecureService := new(mocks.SecureService)

		service := authService{
			ur: mockRepo,
			ts: mockTokenService,
			ss: mockSecureService,
		}

		payload := models.LoginPayload{
			CPF:      "12345678900",
			Password: "wrongpassword",
		}

		user := &models.User{
			PasswordHash: "$2a$10$somehashedpassword",
		}

		mockRepo.On("GetUserByCPF", mock.Anything, payload.CPF).
			Return(user, nil)

		mockSecureService.On("CheckPassword", mock.Anything, user.PasswordHash, payload.Password).
			Return(errors.New("incorrect password"))

		response, err := service.Login(context.Background(), payload)

		assert.ErrorIs(t, err, models.ErrInvalidCredentials)
		assert.Nil(t, response)
		mockRepo.AssertExpectations(t)
		mockSecureService.AssertExpectations(t)
	})

	t.Run("WhenLoginIsSuccessful_ShouldReturnToken", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		mockTokenService := new(mocks.TokenService)
		mockSecureService := new(mocks.SecureService)

		service := authService{
			ur: mockRepo,
			ts: mockTokenService,
			ss: mockSecureService,
		}

		payload := models.LoginPayload{
			CPF:      "12345678900",
			Password: "password",
		}

		user := &models.User{
			BaseModel: models.BaseModel{
				ID: uuid.New(),
			},
			FullName:     "some-name",
			PasswordHash: "$2a$10$somehashedpassword",
			Role:         models.Admin,
		}

		mockRepo.On("GetUserByCPF", mock.Anything, payload.CPF).
			Return(user, nil)

		mockSecureService.On("CheckPassword", mock.Anything, user.PasswordHash, payload.Password).
			Return(nil)

		mockTokenService.On("CreateToken", mock.Anything, mock.Anything).
			Return("some-jwt-token", nil)

		response, err := service.Login(context.Background(), payload)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, "some-jwt-token", response.Token)
		mockRepo.AssertExpectations(t)
		mockSecureService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})
}
