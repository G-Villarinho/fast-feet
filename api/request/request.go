package request

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const userIDKey contextKey = "userID"
const tokenKey contextKey = "userToken"

func WithUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func WithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func UserID(ctx context.Context) (uuid.UUID, bool) {
	UserID, ok := ctx.Value(userIDKey).(uuid.UUID)
	return UserID, ok
}

func Token(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(tokenKey).(string)
	return token, ok
}
