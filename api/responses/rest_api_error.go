package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Status  int    `json:"status"`
	Details string `json:"details"`
}

func AccessDeniedAPIErrorResponse(ctx echo.Context) error {
	return ctx.JSON(http.StatusUnauthorized, ErrorResponse{
		Code:    "UNAUTHORIZED",
		Status:  http.StatusUnauthorized,
		Details: "Você precisa estar autenticado para acessar este recurso.",
	})
}

func InternalServerAPIErrorResponse(ctx echo.Context) error {
	return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
		Status:  http.StatusInternalServerError,
		Details: "Pedimos desculpas! Ocorreu um erro inesperado em nossa aplicação. Se o problema persistir, entre em contato com o suporte..",
	})
}

func ForbiddenPermissionAPIErrorResponse(ctx echo.Context) error {
	return ctx.JSON(http.StatusForbidden, ErrorResponse{
		Code:    "FORBIDDEN",
		Status:  http.StatusForbidden,
		Details: "Você não tem permissão para acessar este recurso.",
	})
}

func CannotBindPayloadAPIErrorResponse(ctx echo.Context) error {
	errorResponse := ErrorResponse{
		Status:  http.StatusUnprocessableEntity,
		Details: "Ops! Não conseguimos entender os dados enviados. Por favor, confira o formato e tente novamente.",
	}
	return ctx.JSON(http.StatusUnprocessableEntity, errorResponse)
}

func NewCustomAPIErrorResponse(ctx echo.Context, statusCode int, details string) error {
	errorResponse := ErrorResponse{
		Status:  statusCode,
		Details: details,
	}

	return ctx.JSON(statusCode, errorResponse)
}
