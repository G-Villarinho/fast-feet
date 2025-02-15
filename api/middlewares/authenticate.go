package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/G-Villarinho/fast-feet-api/config"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/G-Villarinho/fast-feet-api/request"
	"github.com/G-Villarinho/fast-feet-api/responses"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		cookie, err := ectx.Cookie(config.Env.Session.CookieName)
		if err != nil {
			return responses.AccessDeniedAPIErrorResponse(ectx)
		}

		claims, err := validateToken(cookie.Value)
		if err != nil {
			removeCookie(ectx)
			return responses.AccessDeniedAPIErrorResponse(ectx)
		}

		ctx := ectx.Request().Context()
		ctx = request.WithUserID(ctx, claims.UserID)
		ctx = request.WithToken(ctx, cookie.Value)

		ectx.SetRequest(ectx.Request().WithContext(ctx))

		return next(ectx)
	}
}

func validateToken(tokenString string) (*models.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.TokenClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(config.Env.Session.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func removeCookie(ectx echo.Context) {
	ectx.SetCookie(&http.Cookie{
		Name:     config.Env.Session.CookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})
}
