package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/G-Villarinho/fast-feet-api/repositories"
	"github.com/G-Villarinho/fast-feet-api/request"
	"github.com/G-Villarinho/fast-feet-api/services/email"
	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(ctx context.Context, payload models.CreateOrderPayload) (*models.CreateOrderResponse, error)
	PickUpOrder(ctx context.Context, orderID uuid.UUID) (*models.PickUpOrderResponse, error)
	DeliverOrder(ctx context.Context, orderID uuid.UUID, payload models.DeliverOrderPayload) error
	GetOrders(ctx context.Context, paginations *models.Pagination) (*models.PaginatedResponse[*models.OrderResponse], error)
	GetOrder(ctx context.Context, orderID uuid.UUID) (*models.OrderDetailsResponse, error)
}

type orderService struct {
	i  *di.Injector
	ef *email.EmailFactory
	es email.EmailService
	fs FileService
	or repositories.OrderRepository
	rr repositories.RecipientRepository
	ur repositories.UserRepository
}

func NewOrderService(i *di.Injector) (OrderService, error) {
	ef := email.NewEmailFactory()

	es, err := di.Invoke[email.EmailService](i)
	if err != nil {
		return nil, fmt.Errorf("invoke email service: %w", err)
	}

	fs, err := di.Invoke[FileService](i)
	if err != nil {
		return nil, fmt.Errorf("invoke file service: %w", err)
	}

	or, err := di.Invoke[repositories.OrderRepository](i)
	if err != nil {
		return nil, fmt.Errorf("invoke order repository: %w", err)
	}

	rr, err := di.Invoke[repositories.RecipientRepository](i)
	if err != nil {
		return nil, fmt.Errorf("invoke recipient repository: %w", err)
	}

	ur, err := di.Invoke[repositories.UserRepository](i)
	if err != nil {
		return nil, fmt.Errorf("invoke user repository: %w", err)
	}

	return &orderService{
		i:  i,
		ef: ef,
		es: es,
		fs: fs,
		or: or,
		rr: rr,
		ur: ur,
	}, nil
}

func (o *orderService) CreateOrder(ctx context.Context, payload models.CreateOrderPayload) (*models.CreateOrderResponse, error) {
	userID, found := request.UserID(ctx)
	if !found {
		return nil, models.ErrUserNotFoundInContext
	}

	user, err := o.ur.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if models.Cannot(user.Role, models.Create, models.Orders) {
		return nil, models.ErrInsufficientPermission
	}

	recipient, err := o.rr.GetRecipientByID(ctx, payload.RecipientID)
	if err != nil {
		return nil, fmt.Errorf("get recipient by id %q: %w", payload.RecipientID, err)
	}

	if recipient == nil {
		return nil, models.ErrRecipientNotFound
	}

	order := payload.ToOrder()

	if err := o.or.CreateOrder(ctx, *order); err != nil {
		return nil, fmt.Errorf("create order: %w", err)
	}

	return &models.CreateOrderResponse{
		OrderID: order.ID,
	}, nil
}

func (o *orderService) PickUpOrder(ctx context.Context, orderID uuid.UUID) (*models.PickUpOrderResponse, error) {
	userID, found := request.UserID(ctx)
	if !found {
		return nil, models.ErrUserNotFoundInContext
	}

	user, err := o.ur.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if models.Cannot(user.Role, models.UpdateStatus, models.Orders) {
		return nil, models.ErrInsufficientPermission
	}

	order, err := o.or.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("get order by id %q: %w", orderID, err)
	}

	if order == nil {
		return nil, models.ErrOrderNotFound
	}

	if order.Status != models.Waiting {
		return nil, models.ErrCannotTransitionToDelivered
	}

	order.DeliverymanID = &userID
	order.PicknUpAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	order.Status = models.PicknUp

	if err := o.or.UpdateOrder(ctx, *order); err != nil {
		return nil, fmt.Errorf("update order %q status: %w", orderID, err)
	}

	go func() {
		sendEmailPayload := o.ef.CreatePickUpSendEmail(order.Recipient.Email, "Pedido em rota de entrega", order.Recipient.FullName, order.TrackingCode.String())
		o.es.SendEmail(ctx, sendEmailPayload)
	}()

	return &models.PickUpOrderResponse{
		PicknUpAt: order.PicknUpAt.Time,
	}, nil
}

func (o *orderService) DeliverOrder(ctx context.Context, orderID uuid.UUID, payload models.DeliverOrderPayload) error {
	userID, found := request.UserID(ctx)
	if !found {
		return models.ErrUserNotFoundInContext
	}

	user, err := o.ur.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if models.Cannot(user.Role, models.UpdateStatus, models.Orders) {
		return models.ErrInsufficientPermission
	}

	if err := o.fs.ValidateImage(ctx, payload.OrderImage); err != nil {
		return err
	}

	order, err := o.or.GetOrderByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("get order by id %q: %w", orderID, err)
	}

	if order == nil {
		return models.ErrOrderNotFound
	}

	if order.Status != models.PicknUp {
		return models.ErrCannotTransitionToDelivered
	}

	order.DeliveryAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	order.Status = models.Done

	// TODO: Notification recipient
	// TODO: Persist imagem in cloudflare or anothe image store client

	if err := o.or.UpdateOrder(ctx, *order); err != nil {
		return fmt.Errorf("update order %q status: %w", orderID, err)
	}

	return nil
}

func (o *orderService) GetOrders(ctx context.Context, pagination *models.Pagination) (*models.PaginatedResponse[*models.OrderResponse], error) {
	userID, found := request.UserID(ctx)
	if !found {
		return nil, models.ErrUserNotFoundInContext
	}

	user, err := o.ur.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if models.Cannot(user.Role, models.Read, models.Deliveries) {
		return nil, models.ErrInsufficientPermission
	}

	var deliveryManID *uuid.UUID
	if user.Role == models.DeliveryMan {
		deliveryManID = &userID
	}

	paginatedOrders, err := o.or.GetOrdersPagedList(ctx, deliveryManID, pagination)
	if err != nil {
		return nil, fmt.Errorf("get paginated orders: %w", err)
	}

	paginatedOrdersResponse := models.MapPaginatedResult(paginatedOrders, func(order models.Order) *models.OrderResponse {
		return order.ToOrderResponse()
	})

	return paginatedOrdersResponse, nil
}

func (o *orderService) GetOrder(ctx context.Context, orderID uuid.UUID) (*models.OrderDetailsResponse, error) {
	userID, found := request.UserID(ctx)
	if !found {
		return nil, models.ErrUserNotFoundInContext
	}

	user, err := o.ur.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id %q: %w", userID, err)
	}

	if models.Cannot(user.Role, models.Read, models.Deliveries) {
		return nil, models.ErrInsufficientPermission
	}

	order, err := o.or.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("get order by id %q: %w", orderID, err)
	}

	if order == nil {
		return nil, models.ErrOrderNotFound
	}

	if user.Role == models.DeliveryMan && order.DeliverymanID != nil && *order.DeliverymanID != userID {
		return nil, models.ErrNotAssignedToOrder
	}

	return order.ToOrderDetailsResponse(), nil
}
