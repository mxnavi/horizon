package common

import (
	"context"

	herror "github.com/horizoncd/horizon/core/errors"
	hctx "github.com/horizoncd/horizon/pkg/context"
)

func WithContextJWTTokenString(ctx context.Context, tokenStr string) context.Context {
	return context.WithValue(ctx, hctx.JWTTokenString, tokenStr)
}

func JWTTokenStringFromContext(ctx context.Context) (string, error) {
	str, ok := ctx.Value(hctx.JWTTokenString).(string)
	if !ok {
		return "", herror.ErrFailedToGetJWTToken
	}
	return str, nil
}
