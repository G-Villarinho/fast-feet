package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/G-Villarinho/fast-feet-api/mocks"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const validCreateUserPayload = `{
	"email": "admin@example.com",
	"cpf": "322.784.410-94",
	"fullName": "Admin User"
}`

func TestUserHandler_CreateAdmin(t *testing.T) {
	t.Run("WhenCannotBindPayload_ShouldReturnError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/admins", nil)
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		handler := &userHandler{}

		err := handler.CreateAdmin(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Contains(t, rec.Body.String(), "Ops! Não conseguimos entender os dados enviados")
	})

	t.Run("WhenEmailAlreadyExists_ShouldReturnConflict", func(t *testing.T) {
		mockUserService := new(mocks.UserService)
		handler := &userHandler{us: mockUserService}

		mockUserService.On("CreateAdmin", mock.Anything, mock.Anything).Return(models.ErrEmailAlreadyExists)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/admins", strings.NewReader(validCreateUserPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		err := handler.CreateAdmin(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.Contains(t, rec.Body.String(), "Um usuário com o mesmo e-mail já está cadastrado.")
		mockUserService.AssertExpectations(t)
	})

	t.Run("WhenCPFAlreadyExists_ShouldReturnConflict", func(t *testing.T) {
		mockUserService := new(mocks.UserService)
		handler := &userHandler{us: mockUserService}

		mockUserService.On("CreateAdmin", mock.Anything, mock.Anything).Return(models.ErrCPFAlreadyExists)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/admin", strings.NewReader(validCreateUserPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		err := handler.CreateAdmin(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.Contains(t, rec.Body.String(), "Um usuário com o mesmo CPF já está cadastrado.")
		mockUserService.AssertExpectations(t)
	})

	t.Run("WhenEmailIsInvalid_ShouldReturnValidationError", func(t *testing.T) {
		invalidEmailPayload := `{
			"email": "invalid-email",
			"cpf": "322.784.410-94",
			"fullName": "Admin User"
		}`

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/admin", strings.NewReader(invalidEmailPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		handler := &userHandler{}

		err := handler.CreateAdmin(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Contains(t, rec.Body.String(), "O formato do e-mail está inválido.")
	})

	t.Run("WhenCPFIsInvalid_ShouldReturnValidationError", func(t *testing.T) {
		invalidCPFPayload := `{
			"email": "admin@example.com",
			"cpf": "32248441094", 
			"fullName": "Admin User"
		}`

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/admin", strings.NewReader(invalidCPFPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		handler := &userHandler{}

		err := handler.CreateAdmin(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Contains(t, rec.Body.String(), "O formato do CPF está inválido.")
	})

	t.Run("WhenUserIsValid_ShouldCreateSuccessfully", func(t *testing.T) {
		mockUserService := new(mocks.UserService)
		handler := &userHandler{us: mockUserService}

		mockUserService.On("CreateAdmin", mock.Anything, mock.Anything).Return(nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/admin", strings.NewReader(validCreateUserPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		err := handler.CreateAdmin(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUserService.AssertExpectations(t)
	})

	t.Run("WhenCreateAdminFails_ShouldReturnInternalServerError", func(t *testing.T) {
		mockUserService := new(mocks.UserService)
		handler := &userHandler{us: mockUserService}

		mockUserService.On("CreateAdmin", mock.Anything, mock.Anything).Return(errors.New("failed to create admin"))

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/admin", strings.NewReader(validCreateUserPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		err := handler.CreateAdmin(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUserService.AssertExpectations(t)
	})
}
