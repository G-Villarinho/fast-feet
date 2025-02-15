package responses

import (
	"net/http"

	"github.com/G-Villarinho/fast-feet-api/validators"
	"github.com/labstack/echo/v4"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Code    string            `json:"code,omitempty"`
	Status  int               `json:"status"`
	Details string            `json:"details"`
	Errors  []ValidationError `json:"errors"`
}

func NewValidationErrorResponse(ctx echo.Context, validationErrors validators.ValidationErrors) error {
	errors := make([]ValidationError, 0, len(validationErrors))
	for field, message := range validationErrors {
		errors = append(errors, ValidationError{
			Field:   field,
			Message: message,
		})
	}

	return ctx.JSON(http.StatusUnprocessableEntity, ValidationErrorResponse{
		Code:    "VALIDATION_ERROR",
		Status:  http.StatusUnprocessableEntity,
		Details: "Os dados fornecidos contêm erros de validação. Corrija os campos indicados e tente novamente.",
		Errors:  errors,
	})
}
