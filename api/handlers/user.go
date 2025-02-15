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
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateAdmin(ectx echo.Context) error
	GetUser(ectx echo.Context) error
}

type userHandler struct {
	i  *di.Injector
	us services.UserService
}

func NewUserHandler(i *di.Injector) (UserHandler, error) {
	us, err := di.Invoke[services.UserService](i)
	if err != nil {
		return nil, fmt.Errorf("invoke user service: %w", err)
	}

	return &userHandler{
		i:  i,
		us: us,
	}, nil
}

func (u *userHandler) CreateAdmin(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "user"),
		slog.String("func", "CreateAdmin"),
	)

	var payload models.CreateUserPayload
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

	if err := u.us.CreateAdmin(ectx.Request().Context(), payload); err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrEmailAlreadyExists) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusConflict, "Um usuário com o mesmo e-mail já está cadastrado.")
		}

		if errors.Is(err, models.ErrInsufficientPermission) {
			return responses.ForbiddenPermissionAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrCPFAlreadyExists) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusConflict, "Um usuário com o mesmo CPF já está cadastrado.")
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.NoContent(http.StatusCreated)
}

func (u *userHandler) GetUser(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "user"),
		slog.String("func", "GetUser"),
	)

	response, err := u.us.GetUser(ectx.Request().Context())
	if err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrUserNotFoundInContext) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		if errors.Is(err, models.ErrUserNotFound) {
			responses.AccessDeniedAPIErrorResponse(ectx)
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	return ectx.JSON(http.StatusOK, response)
}
