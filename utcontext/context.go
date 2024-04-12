package utcontext

import (
	"context"
	"errors"

	"code.finan.cc/finan-one-be/fo-utils/net/uthttp"
)

var (
	ErrNotFound = errors.New("value not found in context")
)

func GetUserIDFromContext(ctx context.Context) (string, error) {
	userID := ctx.Value(uthttp.HeaderUserID).(string)
	if userID == "" {
		return "", ErrNotFound
	}

	return userID, nil
}

func GetRequestIDFromContext(ctx context.Context) (string, error) {
	requestID := ctx.Value(uthttp.HeaderRequestID).(string)
	if requestID == "" {
		return "", ErrNotFound
	}

	return requestID, nil
}

func GetLocaleFromContext(ctx context.Context) (string, error) {
	locale := ctx.Value(uthttp.HeaderLocale).(string)
	if locale == "" {
		return "", ErrNotFound
	}

	return locale, nil
}
