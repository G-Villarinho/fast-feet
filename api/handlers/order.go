package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/G-Villarinho/fast-feet-api/responses"
	"github.com/G-Villarinho/fast-feet-api/services"
	"github.com/G-Villarinho/fast-feet-api/validators"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type OrderHandler interface {
	CreateOrder(ectx echo.Context) error
	PickUpOrder(ectx echo.Context) error
	DeliverOrder(ectx echo.Context) error
	GetOrders(ectx echo.Context) error
	GetOrder(ectx echo.Context) error
}

type orderHandler struct {
	i  *di.Injector
	os services.OrderService
}

func NewOrderHandler(i *di.Injector) (OrderHandler, error) {
	os, err := di.Invoke[services.OrderService](i)
	if err != nil {
		return nil, fmt.Errorf("invoke order service: %w", err)
	}

	return &orderHandler{
		i:  i,
		os: os,
	}, nil
}

func (o *orderHandler) CreateOrder(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "order"),
		slog.String("func", "CreateOrder"),
	)

	var payload models.CreateOrderPayload
	if err := jsoniter.NewDecoder(ectx.Request().Body).Decode(&payload); err != nil {
		log.Warn("Error to decode JSON payload", slog.String("error", err.Error()))
		return responses.CannotBindPayloadAPIErrorResponse(ectx)
	}

	if validationErrors := validators.ValidateStruct(&payload); validationErrors != nil {
		if msg, exists := validationErrors["validation_setup"]; exists {
			log.Warn("Error in validation setup", slog.String("message", msg))
			return responses.CannotBindPayloadAPIErrorResponse(ectx)
		}

		return responses.NewValidationErrorResponse(ectx, validationErrors)
	}

	response, err := o.os.CreateOrder(ectx.Request().Context(), payload)
	if err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrInsufficientPermission) {
			return responses.ForbiddenPermissionAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrRecipientNotFound) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusNotFound, "Não foi encontrado um destinário para criar a encomenda.")
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.JSON(http.StatusCreated, response)
}

func (o *orderHandler) PickUpOrder(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "order"),
		slog.String("func", "PickUpOrder"),
	)

	orderID, err := uuid.Parse(ectx.Param("orderId"))
	if err != nil {
		log.Warn(err.Error())
		return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "Parâmetro de busca de encomenda inválido.")
	}

	response, err := o.os.PickUpOrder(ectx.Request().Context(), orderID)
	if err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrInsufficientPermission) {
			return responses.ForbiddenPermissionAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrOrderNotFound) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusNotFound, "Não foi encontrado um pedido para retirar.")
		}

		if errors.Is(err, models.ErrCannotTransitionToPicknUp) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "Não é possível entregar a encomenda que já foi entregue ou que foi retirada por outro entregador.")
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.JSON(http.StatusOK, response)
}

func (o *orderHandler) DeliverOrder(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "order"),
		slog.String("func", "DeliverOrder"),
	)

	orderID, err := uuid.Parse(ectx.Param("orderId"))
	if err != nil {
		log.Warn(err.Error())
		return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "Parâmetro de busca de encomenda inválido.")
	}

	image, err := ectx.FormFile("image")
	if err != nil {
		log.Warn(err.Error())
		return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "É necessário enviar uma imagem para terminar a entrega.")
	}

	payload := models.DeliverOrderPayload{
		OrderImage: image,
	}

	if err := o.os.DeliverOrder(ectx.Request().Context(), orderID, payload); err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrInsufficientPermission) {
			return responses.ForbiddenPermissionAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrOrderNotFound) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusNotFound, "Não foi encontrado um pedido para relizar a entrega.")
		}

		if errors.Is(err, models.ErrCannotTransitionToDelivered) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "Não é possível entregar a encomenda sem antes retira-la.")
		}

		if errors.Is(err, models.ErrImageTooLarge) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "A imagem é muito grande. O tamanho máximo permitido é 5MB.")
		}

		if errors.Is(err, models.ErrInvalidImageFormat) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "Formato de imagem inválido. Por favor, envie uma imagem com formato válido (JPG, JPEG, PNG).")
		}

		if errors.Is(err, models.ErrImageCorrupted) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "A imagem está corrompida ou tem um formato inválido.")
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.NoContent(http.StatusOK)
}

func (o *orderHandler) GetOrders(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "order"),
		slog.String("func", "GetOrders"),
	)

	pagination := models.NewPagination(ectx.QueryParam("page"), ectx.QueryParam("limit"))

	response, err := o.os.GetOrders(ectx.Request().Context(), pagination)
	if err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrInsufficientPermission) {
			return responses.ForbiddenPermissionAPIErrorResponse(ectx)
		}

		responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.JSON(http.StatusOK, response)
}

func (o *orderHandler) GetOrder(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "order"),
		slog.String("func", "GetOrder"),
	)

	orderID, err := uuid.Parse(ectx.Param("orderId"))
	if err != nil {
		log.Warn(err.Error())
		return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "Parâmetro de busca de encomenda inválido.")
	}

	response, err := o.os.GetOrder(ectx.Request().Context(), orderID)
	if err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrInsufficientPermission) {
			return responses.ForbiddenPermissionAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrOrderNotFound) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusNotFound, "Não foi encontrado nenhuma entrega.")
		}

		if errors.Is(err, models.ErrNotAssignedToOrder) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusConflict, "Você não está atribuído a esta entrega.")
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.JSON(http.StatusOK, response)
}
