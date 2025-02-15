package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/G-Villarinho/fast-feet-api/config"
	"github.com/G-Villarinho/fast-feet-api/mocks"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const validLoginPayload = `{
	"cpf": "322.784.410-94",
	"password": "validPassword123"
}`

func TestAuthHandler_Login(t *testing.T) {

	t.Run("WhenCannotBindPayload_ShouldReturnError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		handler := &authHandler{}

		err := handler.Login(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Contains(t, rec.Body.String(), "Ops! Não conseguimos entender os dados enviados")
	})

	t.Run("WhenValidationFails_ShouldReturnValidationError", func(t *testing.T) {
		invalidPayload := `{
			"cpf": "322.784.410-94",
			"password": "short"
		}`

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(invalidPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		handler := &authHandler{}

		err := handler.Login(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Contains(t, rec.Body.String(), "Por favor, insira um valor com no mínimo 8 caracteres.")
	})

	t.Run("WhenInvalidCredentials_ShouldReturnConflict", func(t *testing.T) {
		mockAuthService := new(mocks.AuthService)
		handler := &authHandler{as: mockAuthService}

		mockAuthService.On("Login", mock.Anything, mock.Anything).Return(nil, models.ErrInvalidCredentials)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(validLoginPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		err := handler.Login(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.Contains(t, rec.Body.String(), "Credenciais inválidas. Verifique seu CPF e senha e tente novamente.")
		mockAuthService.AssertExpectations(t)
	})

	t.Run("WhenUserBlocked_ShouldReturnForbidden", func(t *testing.T) {
		mockAuthService := new(mocks.AuthService)
		handler := &authHandler{as: mockAuthService}

		mockAuthService.On("Login", mock.Anything, mock.Anything).Return(nil, models.ErrUserBlocked)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(validLoginPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		err := handler.Login(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)
		assert.Contains(t, rec.Body.String(), "Sua conta está temporariamente bloqueada. Entre em contato com o suporte para mais informações.")
		mockAuthService.AssertExpectations(t)
	})

	t.Run("WhenLoginSucceeds_ShouldReturnTokenInCookie", func(t *testing.T) {
		mockAuthService := new(mocks.AuthService)
		handler := &authHandler{as: mockAuthService}

		mockAuthService.On("Login", mock.Anything, mock.Anything).Return(&models.LoginResponse{Token: "valid-token"}, nil)

		config.Env.Session.CookieName = "fast-feet.token"

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(validLoginPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		err := handler.Login(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		cookies := rec.Result().Cookies()
		var tokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == config.Env.Session.CookieName {
				tokenCookie = cookie
				break
			}
		}

		assert.NotNil(t, tokenCookie)
		assert.Equal(t, "valid-token", tokenCookie.Value)

		mockAuthService.AssertExpectations(t)
	})

	t.Run("WhenLoginFails_ShouldReturnInternalServerError", func(t *testing.T) {
		mockAuthService := new(mocks.AuthService)
		handler := &authHandler{as: mockAuthService}

		mockAuthService.On("Login", mock.Anything, mock.Anything).Return(nil, errors.New("failed to authenticate"))

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(validLoginPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		err := handler.Login(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockAuthService.AssertExpectations(t)
	})
}
