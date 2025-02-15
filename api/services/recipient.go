package services

import (
	"context"
	"fmt"

	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/G-Villarinho/fast-feet-api/repositories"
	"github.com/G-Villarinho/fast-feet-api/request"
	"github.com/google/uuid"
)

//go:generate mockery --name=RecipientService --filename=recipient_service.go --output=../mocks --outpkg=mocks
type RecipientService interface {
	CreateRecipient(ctx context.Context, payload models.CreateRecipientPayload) (*models.CreateRecipientResponse, error)
	GetRecipient(ctx context.Context, recipientID uuid.UUID) (*models.RecipientResponse, error)
	UpdateRecipient(ctx context.Context, recipientID uuid.UUID, payload models.UpdateRecipientPayload) (*models.RecipientResponse, error)
	DeleteRecipient(ctx context.Context, recipientID uuid.UUID) error
	GetRecipientsBasicInfo(ctx context.Context, pagination *models.RecipientBasicInfoPagination) (*models.PaginatedResponse[*models.RecipientBasicInfoResponse], error)
}

type recipientService struct {
	i  *di.Injector
	rr repositories.RecipientRepository
	ur repositories.UserRepository
}

func NewRecipientService(i *di.Injector) (RecipientService, error) {
	rr, err := di.Invoke[repositories.RecipientRepository](i)
	if err != nil {
		return nil, fmt.Errorf("invoke recipient service: %w", err)
	}

	ur, err := di.Invoke[repositories.UserRepository](i)
	if err != nil {
		return nil, fmt.Errorf("invoke user repository: %w", err)
	}

	return &recipientService{
		i:  i,
		rr: rr,
		ur: ur,
	}, nil
}

func (r *recipientService) CreateRecipient(ctx context.Context, payload models.CreateRecipientPayload) (*models.CreateRecipientResponse, error) {
	userID, found := request.UserID(ctx)
	if !found {
		return nil, models.ErrUserNotFoundInContext
	}

	user, err := r.ur.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if models.Cannot(user.Role, models.Create, models.Recipients) {
		return nil, models.ErrInsufficientPermission
	}

	recipientFromEmail, err := r.rr.GetRecipientByEmail(ctx, payload.Email)
	if err != nil {
		return nil, fmt.Errorf("get recipient by email: %w", err)
	}

	if recipientFromEmail != nil {
		return nil, models.ErrEmailAlreadyExists
	}

	recipient := payload.ToRecipient()

	if err := r.rr.CreateRecipient(ctx, *recipient); err != nil {
		return nil, fmt.Errorf("create recipient: %w", err)
	}

	return &models.CreateRecipientResponse{
		RecipientID: recipient.ID,
	}, nil
}

func (r *recipientService) GetRecipient(ctx context.Context, recipientID uuid.UUID) (*models.RecipientResponse, error) {
	userID, found := request.UserID(ctx)
	if !found {
		return nil, models.ErrUserNotFoundInContext
	}

	user, err := r.ur.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if models.Cannot(user.Role, models.Read, models.Recipients) {
		return nil, models.ErrInsufficientPermission
	}

	recipient, err := r.rr.GetRecipientByID(ctx, recipientID)
	if err != nil {
		return nil, fmt.Errorf("get recipient by id %q: %w", recipientID, err)
	}

	if recipient == nil {
		return nil, models.ErrRecipientNotFound
	}

	return recipient.ToRecipientResponse(), nil
}

func (r *recipientService) UpdateRecipient(ctx context.Context, recipientID uuid.UUID, payload models.UpdateRecipientPayload) (*models.RecipientResponse, error) {
	userID, found := request.UserID(ctx)
	if !found {
		return nil, models.ErrUserNotFoundInContext
	}

	user, err := r.ur.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if models.Cannot(user.Role, models.Update, models.Recipients) {
		return nil, models.ErrInsufficientPermission
	}

	recipient, err := r.rr.GetRecipientByID(ctx, recipientID)
	if err != nil {
		return nil, fmt.Errorf("get recipient by id %q: %w", recipientID, err)
	}

	if recipient == nil {
		return nil, models.ErrRecipientNotFound
	}

	if recipient.Email != payload.Email {
		recipientFromEmail, err := r.rr.GetRecipientByEmail(ctx, payload.Email)
		if err != nil {
			return nil, fmt.Errorf("get recipient by email: %w", err)
		}

		if recipientFromEmail != nil {
			return nil, models.ErrEmailAlreadyExists
		}
	}

	recipient.ApplyUpdates(&payload)

	if err := r.rr.UpdateRecipient(ctx, *recipient); err != nil {
		return nil, fmt.Errorf("update recipient %q: %w", recipientID, err)
	}

	return recipient.ToRecipientResponse(), nil
}

func (r *recipientService) DeleteRecipient(ctx context.Context, recipientID uuid.UUID) error {
	userID, found := request.UserID(ctx)
	if !found {
		return models.ErrUserNotFoundInContext
	}

	user, err := r.ur.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if models.Cannot(user.Role, models.Delete, models.Recipients) {
		return models.ErrInsufficientPermission
	}

	recipient, err := r.rr.GetRecipientByID(ctx, recipientID)
	if err != nil {
		return fmt.Errorf("get recipient by id %q: %w", recipientID, err)
	}

	if recipient == nil {
		return models.ErrRecipientNotFound
	}

	if err := r.rr.DeleteRecipient(ctx, recipientID); err != nil {
		return fmt.Errorf("delete recipient %q: %w", recipientID, err)
	}

	return nil
}

func (r *recipientService) GetRecipientsBasicInfo(ctx context.Context, pagination *models.RecipientBasicInfoPagination) (*models.PaginatedResponse[*models.RecipientBasicInfoResponse], error) {
	userID, found := request.UserID(ctx)
	if !found {
		return nil, models.ErrUserNotFoundInContext
	}

	user, err := r.ur.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if models.Cannot(user.Role, models.Read, models.Recipients) {
		return nil, models.ErrInsufficientPermission
	}

	paginatedRecipients, err := r.rr.GetRecipientLitePagedList(ctx, pagination)
	if err != nil {
		return nil, fmt.Errorf("get paginated recipients: %w", err)
	}

	paginatedRecipientBasicInfoResponse := models.MapPaginatedResult(paginatedRecipients, func(recipient models.Recipient) *models.RecipientBasicInfoResponse {
		return recipient.ToRecipientBasicInfoResponse()
	})

	return paginatedRecipientBasicInfoResponse, nil
}
