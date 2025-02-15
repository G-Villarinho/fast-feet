package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/G-Villarinho/fast-feet-api/mocks"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const validCreateRecipientPayload = `{
	"fullName": "John Doe",
	"email": "johndoe@example.com",
	"state": "SP",
	"city": "São Paulo",
	"neighborhood": "Centro",
	"address": "Rua Exemplo, 123",
	"zipcode": 12345678
}`

func TestRecipientHandler_CreateRecipient(t *testing.T) {
	t.Run("WhenCannotBindPayload_ShouldReturnError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/recipients", nil)
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		mockRecipientService := new(mocks.RecipientService)
		handler := &recipientHandler{rs: mockRecipientService}

		err := handler.CreateRecipient(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("WhenValidationFails_ShouldReturnValidationError", func(t *testing.T) {
		invalidPayload := `{"fullName": "", "email": "invalid-email"}`

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/recipients", strings.NewReader(invalidPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		mockRecipientService := new(mocks.RecipientService)
		handler := &recipientHandler{rs: mockRecipientService}

		err := handler.CreateRecipient(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Contains(t, rec.Body.String(), "O formato do e-mail está inválido.")
		assert.Contains(t, rec.Body.String(), "Este campo é obrigatório. Por favor, preencha corretamente.")
	})

	t.Run("WhenEmailAlreadyExists_ShouldReturnConflict", func(t *testing.T) {
		mockRecipientService := new(mocks.RecipientService)
		mockRecipientService.On("CreateRecipient", mock.Anything, mock.Anything).
			Return(nil, models.ErrEmailAlreadyExists)

		handler := &recipientHandler{rs: mockRecipientService}

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/recipients", strings.NewReader(validCreateRecipientPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		err := handler.CreateRecipient(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.Contains(t, rec.Body.String(), "Um destinatário com o mesmo e-mail já está cadastrado.")
		mockRecipientService.AssertExpectations(t)
	})

	t.Run("WhenCreationSucceeds_ShouldReturnCreated", func(t *testing.T) {
		mockRecipientService := new(mocks.RecipientService)
		mockRecipientService.On("CreateRecipient", mock.Anything, mock.Anything).
			Return(&models.CreateRecipientResponse{RecipientID: uuid.New()}, nil)

		handler := &recipientHandler{rs: mockRecipientService}

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/recipients", strings.NewReader(validCreateRecipientPayload))
		rec := httptest.NewRecorder()
		ectx := e.NewContext(req, rec)

		err := handler.CreateRecipient(ectx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		mockRecipientService.AssertExpectations(t)
	})
}
