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

func TestCreateRecipient(t *testing.T) {
	t.Run("WhenUserNotFoundInContext_ShouldReturnErrUserNotFoundInContext", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		ctx := context.Background()

		payload := models.CreateRecipientPayload{
			FullName:     "John Doe",
			Email:        "john@example.com",
			State:        "State",
			City:         "City",
			Neighborhood: "Neighborhood",
			Address:      "123 Address",
			Zipcode:      123456,
		}

		resp, err := service.CreateRecipient(ctx, payload)

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrUserNotFoundInContext)
	})

	t.Run("WhenUserHasInsufficientPermission_ShouldReturnErrInsufficientPermission", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.DeliveryMan}, nil)

		payload := models.CreateRecipientPayload{
			FullName:     "John Doe",
			Email:        "john@example.com",
			State:        "State",
			City:         "City",
			Neighborhood: "Neighborhood",
			Address:      "123 Address",
			Zipcode:      123456,
		}

		resp, err := service.CreateRecipient(ctx, payload)

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrInsufficientPermission)
	})

	t.Run("WhenEmailAlreadyExists_ShouldReturnErrEmailAlreadyExists", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		payload := models.CreateRecipientPayload{
			FullName:     "John Doe",
			Email:        "john@example.com",
			State:        "State",
			City:         "City",
			Neighborhood: "Neighborhood",
			Address:      "123 Address",
			Zipcode:      123456,
		}

		mockRepo.On("GetRecipientByEmail", ctx, payload.Email).
			Return(&models.Recipient{}, nil)

		resp, err := service.CreateRecipient(ctx, payload)

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrEmailAlreadyExists)
	})

	t.Run("WhenRecipientCreatedSuccessfully_ShouldReturnCreateRecipientResponse", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		payload := models.CreateRecipientPayload{
			FullName:     "John Doe",
			Email:        "john@example.com",
			State:        "State",
			City:         "City",
			Neighborhood: "Neighborhood",
			Address:      "123 Address",
			Zipcode:      123456,
		}

		mockRepo.On("GetRecipientByEmail", ctx, payload.Email).
			Return(nil, nil)

		mockRepo.On("CreateRecipient", ctx, mock.Anything).
			Return(nil)

		resp, err := service.CreateRecipient(ctx, payload)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
}

func TestGetRecipient(t *testing.T) {
	t.Run("WhenUserNotFoundInContext_ShouldReturnErrUserNotFoundInContext", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		ctx := context.Background()

		recipientID := uuid.New()

		resp, err := service.GetRecipient(ctx, recipientID)

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrUserNotFoundInContext)
	})

	t.Run("WhenUserHasInsufficientPermission_ShouldReturnErrInsufficientPermission", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.DeliveryMan}, nil)

		recipientID := uuid.New()

		resp, err := service.GetRecipient(ctx, recipientID)

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrInsufficientPermission)
	})

	t.Run("WhenRecipientNotFound_ShouldReturnErrRecipientNotFound", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		recipientID := uuid.New()

		mockRepo.On("GetRecipientByID", ctx, recipientID).
			Return(nil, nil)

		resp, err := service.GetRecipient(ctx, recipientID)

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrRecipientNotFound)
	})

	t.Run("WhenRecipientFound_ShouldReturnRecipientResponse", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		recipientID := uuid.New()

		recipient := &models.Recipient{
			BaseModel: models.BaseModel{
				ID: recipientID,
			},
			FullName:     "John Doe",
			Email:        "john@example.com",
			State:        "State",
			City:         "City",
			Neighborhood: "Neighborhood",
			Address:      "123 Address",
			Zipcode:      123456,
		}

		mockRepo.On("GetRecipientByID", ctx, recipientID).
			Return(recipient, nil)

		resp, err := service.GetRecipient(ctx, recipientID)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, recipientID, resp.ID)
		assert.Equal(t, "John Doe", resp.FullName)
		assert.Equal(t, recipient.Email, resp.Email)
	})

	t.Run("WhenErrorFetchingRecipient_ShouldReturnError", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		recipientID := uuid.New()

		mockRepo.On("GetRecipientByID", ctx, recipientID).
			Return(nil, errors.New("failed to fetch recipient"))

		resp, err := service.GetRecipient(ctx, recipientID)

		assert.Nil(t, resp)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to fetch recipient")
	})
}

func TestUpdateRecipient(t *testing.T) {
	t.Run("WhenUserNotFoundInContext_ShouldReturnErrUserNotFoundInContext", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		ctx := context.Background()

		recipientID := uuid.New()
		payload := models.UpdateRecipientPayload{
			FullName: "Updated Name",
			Email:    "updated@example.com",
		}

		resp, err := service.UpdateRecipient(ctx, recipientID, payload)

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrUserNotFoundInContext)
	})

	t.Run("WhenUserHasInsufficientPermission_ShouldReturnErrInsufficientPermission", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.DeliveryMan}, nil)

		recipientID := uuid.New()
		payload := models.UpdateRecipientPayload{
			FullName: "Updated Name",
			Email:    "updated@example.com",
		}

		resp, err := service.UpdateRecipient(ctx, recipientID, payload)

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrInsufficientPermission)
	})

	t.Run("WhenRecipientNotFound_ShouldReturnErrRecipientNotFound", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		recipientID := uuid.New()
		payload := models.UpdateRecipientPayload{
			FullName: "Updated Name",
			Email:    "updated@example.com",
		}

		mockRepo.On("GetRecipientByID", ctx, recipientID).
			Return(nil, nil)

		resp, err := service.UpdateRecipient(ctx, recipientID, payload)

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrRecipientNotFound)
	})

	t.Run("WhenEmailAlreadyExists_ShouldReturnErrEmailAlreadyExists", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		recipientID := uuid.New()
		payload := models.UpdateRecipientPayload{
			FullName: "Updated Name",
			Email:    "updated@example.com",
		}

		recipient := &models.Recipient{
			BaseModel: models.BaseModel{
				ID: recipientID,
			},
			FullName:     "John Doe",
			Email:        "john@example.com",
			State:        "State",
			City:         "City",
			Neighborhood: "Neighborhood",
			Address:      "123 Address",
			Zipcode:      123456,
		}

		mockRepo.On("GetRecipientByID", ctx, recipientID).
			Return(recipient, nil)

		mockRepo.On("GetRecipientByEmail", ctx, payload.Email).
			Return(&models.Recipient{}, nil)

		resp, err := service.UpdateRecipient(ctx, recipientID, payload)

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrEmailAlreadyExists)
	})

	t.Run("WhenRecipientUpdatedSuccessfully_ShouldReturnRecipientResponse", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		recipientID := uuid.New()
		payload := models.UpdateRecipientPayload{
			FullName: "Updated Name",
			Email:    "updated@example.com",
		}

		recipient := &models.Recipient{
			BaseModel: models.BaseModel{
				ID: recipientID,
			},
			FullName:     "John Doe",
			Email:        "john@example.com",
			State:        "State",
			City:         "City",
			Neighborhood: "Neighborhood",
			Address:      "123 Address",
			Zipcode:      123456,
		}

		mockRepo.On("GetRecipientByID", ctx, recipientID).
			Return(recipient, nil)

		mockRepo.On("GetRecipientByEmail", ctx, payload.Email).
			Return(nil, nil)

		mockRepo.On("UpdateRecipient", ctx, mock.Anything).
			Return(nil)

		resp, err := service.UpdateRecipient(ctx, recipientID, payload)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, recipientID, resp.ID)
		assert.Equal(t, "Updated Name", resp.FullName)
	})

	t.Run("WhenErrorUpdatingRecipient_ShouldReturnError", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		recipientID := uuid.New()
		payload := models.UpdateRecipientPayload{
			FullName: "Updated Name",
			Email:    "updated@example.com",
		}

		recipient := &models.Recipient{
			BaseModel: models.BaseModel{
				ID: recipientID,
			},
			FullName:     "John Doe",
			Email:        "john@example.com",
			State:        "State",
			City:         "City",
			Neighborhood: "Neighborhood",
			Address:      "123 Address",
			Zipcode:      123456,
		}

		mockRepo.On("GetRecipientByID", ctx, recipientID).
			Return(recipient, nil)

		mockRepo.On("GetRecipientByEmail", ctx, payload.Email).
			Return(nil, nil)

		mockRepo.On("UpdateRecipient", ctx, mock.Anything).
			Return(errors.New("failed to update recipient"))

		resp, err := service.UpdateRecipient(ctx, recipientID, payload)

		assert.Nil(t, resp)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to update recipient")
	})
}

func TestDeleteRecipient(t *testing.T) {
	t.Run("WhenUserNotFoundInContext_ShouldReturnErrUserNotFoundInContext", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		ctx := context.Background()

		recipientID := uuid.New()

		err := service.DeleteRecipient(ctx, recipientID)

		assert.ErrorIs(t, err, models.ErrUserNotFoundInContext)
	})

	t.Run("WhenUserHasInsufficientPermission_ShouldReturnErrInsufficientPermission", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.DeliveryMan}, nil)

		recipientID := uuid.New()

		err := service.DeleteRecipient(ctx, recipientID)

		assert.ErrorIs(t, err, models.ErrInsufficientPermission)
	})

	t.Run("WhenRecipientNotFound_ShouldReturnErrRecipientNotFound", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		recipientID := uuid.New()

		mockRepo.On("GetRecipientByID", ctx, recipientID).
			Return(nil, nil)

		err := service.DeleteRecipient(ctx, recipientID)

		assert.ErrorIs(t, err, models.ErrRecipientNotFound)
	})

	t.Run("WhenErrorDeletingRecipient_ShouldReturnError", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		recipientID := uuid.New()

		recipient := &models.Recipient{
			BaseModel: models.BaseModel{
				ID: recipientID,
			},
		}

		mockRepo.On("GetRecipientByID", ctx, recipientID).
			Return(recipient, nil)

		mockRepo.On("DeleteRecipient", ctx, recipientID).
			Return(errors.New("failed to delete recipient"))

		err := service.DeleteRecipient(ctx, recipientID)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to delete recipient")
	})

	t.Run("WhenRecipientDeletedSuccessfully_ShouldReturnNoError", func(t *testing.T) {
		mockRepo := new(mocks.RecipientRepository)
		mockUserRepo := new(mocks.UserRepository)

		service := recipientService{
			rr: mockRepo,
			ur: mockUserRepo,
		}

		userID := uuid.New()
		ctx := request.WithUserID(context.Background(), userID)

		mockUserRepo.On("GetUserByID", ctx, userID).
			Return(&models.User{Role: models.Admin}, nil)

		recipientID := uuid.New()

		recipient := &models.Recipient{
			BaseModel: models.BaseModel{
				ID: recipientID,
			},
		}

		mockRepo.On("GetRecipientByID", ctx, recipientID).
			Return(recipient, nil)

		mockRepo.On("DeleteRecipient", ctx, recipientID).
			Return(nil)

		err := service.DeleteRecipient(ctx, recipientID)

		assert.NoError(t, err)
	})
}
