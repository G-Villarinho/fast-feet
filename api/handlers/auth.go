package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/G-Villarinho/fast-feet-api/config"
	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/G-Villarinho/fast-feet-api/responses"
	"github.com/G-Villarinho/fast-feet-api/services"
	"github.com/G-Villarinho/fast-feet-api/validators"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Login(ectx echo.Context) error
	Logout(ectx echo.Context) error
}

type authHandler struct {
	i  *di.Injector
	as services.AuthService
}

func NewAuthHandler(i *di.Injector) (AuthHandler, error) {
	as, err := di.Invoke[services.AuthService](i)
	if err != nil {
		return nil, fmt.Errorf("invoke auth service: %w", err)
	}

	return &authHandler{
		i:  i,
		as: as,
	}, nil
}

func (a *authHandler) Login(ectx echo.Context) error {
	log := slog.With(
		slog.String("handler", "auth"),
		slog.String("func", "Login"),
	)

	var payload models.LoginPayload
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

	response, err := a.as.Login(ectx.Request().Context(), payload)
	if err != nil {
		log.Error(err.Error())

		if errors.Is(err, models.ErrInvalidCredentials) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusConflict, "Credenciais inválidas. Verifique seu CPF e senha e tente novamente.")
		}

		if errors.Is(err, models.ErrUserBlocked) {
			return responses.NewCustomAPIErrorResponse(ectx, http.StatusForbidden, "Sua conta está temporariamente bloqueada. Entre em contato com o suporte para mais informações.")
		}

		return responses.InternalServerAPIErrorResponse(ectx)
	}

	ectx.SetCookie(&http.Cookie{
		Name:     config.Env.Session.CookieName,
		Value:    response.Token,
		HttpOnly: true,
		Path:     "/",
	})

	return ectx.NoContent(http.StatusOK)
}

func (a *authHandler) Logout(ectx echo.Context) error {
	ectx.SetCookie(&http.Cookie{
		Name:     config.Env.Session.CookieName,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})

	return ectx.NoContent(http.StatusOK)
}
