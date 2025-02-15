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
	"github.com/G-Villarinho/fast-feet-api/utils"
	"github.com/G-Villarinho/fast-feet-api/validators"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type RecipientHandler interface {
	CreateRecipient(ectx echo.Context) error
	GetRecipient(ectx echo.Context) error
	DeleteRecipient(ectx echo.Context) error
	UpdateRecipient(ectx echo.Context) error
	GetRecipientsBasicInfo(ectx echo.Context) error
}

type recipientHandler struct {
	i  *di.Injector
	rs services.RecipientService
}

func NewRecipientHandler(i *di.Injector) (RecipientHandler, error) {
	rs, err := di.Invoke[services.RecipientService](i)
	if err != nil {
		return nil, fmt.Errorf("invoke recipient service: %w", err)
	}

	return &recipientHandler{
		i:  i,
		rs: rs,
	}, nil
}

func (r *recipientHandler) CreateRecipient(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "recipient"),
		slog.String("func", "CreateRecipient"),
	)

	var payload models.CreateRecipientPayload
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

	response, err := r.rs.CreateRecipient(ectx.Request().Context(), payload)
	if err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrInsufficientPermission) {
			return responses.ForbiddenPermissionAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrEmailAlreadyExists) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusConflict, "Um destinatário com o mesmo e-mail já está cadastrado.")
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.JSON(http.StatusCreated, response)
}

func (r *recipientHandler) GetRecipient(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "recipient"),
		slog.String("func", "GetRecipient"),
	)

	recipientID, err := uuid.Parse(ectx.Param("recipientId"))
	if err != nil {
		log.Warn(err.Error())
		return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "Parâmetro de busca de destinatário inválido.")
	}

	response, err := r.rs.GetRecipient(ectx.Request().Context(), recipientID)
	if err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrInsufficientPermission) {
			return responses.ForbiddenPermissionAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrRecipientNotFound) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusNotFound, "Não foi encontrado um destinário com esse parâmetro de busca.")
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.JSON(http.StatusOK, response)
}

func (r *recipientHandler) DeleteRecipient(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "recipient"),
		slog.String("func", "DeleteRecipient"),
	)

	recipientID, err := uuid.Parse(ectx.Param("recipientId"))
	if err != nil {
		log.Warn(err.Error())
		return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "Parâmetro de busca de destinatário inválido.")
	}

	if err := r.rs.DeleteRecipient(ectx.Request().Context(), recipientID); err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrInsufficientPermission) {
			return responses.ForbiddenPermissionAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrRecipientNotFound) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusNotFound, "Não foi encontrado um destinário com esse parâmetro de busca para deletar.")
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.NoContent(http.StatusNoContent)
}

func (r *recipientHandler) UpdateRecipient(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "recipient"),
		slog.String("func", "UpdateRecipient"),
	)

	recipientID, err := uuid.Parse(ectx.Param("recipientId"))
	if err != nil {
		log.Warn(err.Error())
		return responses.NewCustomAPIErrorResponse(ectx, http.StatusBadRequest, "Parâmetro de busca de destinatário inválido.")
	}

	var payload models.UpdateRecipientPayload
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

	response, err := r.rs.UpdateRecipient(ectx.Request().Context(), recipientID, payload)
	if err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrInsufficientPermission) {
			return responses.ForbiddenPermissionAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrRecipientNotFound) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusNotFound, "Não foi encontrado um destinário com esse parâmetro de busca para deletar.")
		}

		if errors.Is(err, models.ErrEmailAlreadyExists) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusConflict, "Um destinatário com o mesmo e-mail já está cadastrado.")
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.JSON(http.StatusOK, response)
}

func (r *recipientHandler) GetRecipientsBasicInfo(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "recipient"),
		slog.String("func", "GetRecipientsBasicInfo"),
	)

	pagination := &models.RecipientBasicInfoPagination{
		Pagination: *models.NewPagination(ectx.QueryParam("page"), ectx.QueryParam("limit")),
		Q:          utils.GetQueryStringPointer(ectx.QueryParam("q")),
	}

	response, err := r.rs.GetRecipientsBasicInfo(ectx.Request().Context(), pagination)
	if err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			return responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrInsufficientPermission) {
			return responses.ForbiddenPermissionAPIErrorResponse(ectx)
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.JSON(http.StatusOK, response)
}
