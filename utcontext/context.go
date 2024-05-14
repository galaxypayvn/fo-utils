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
	userID, ok := ctx.Value(uthttp.HeaderUserID).(string)
	if userID == "" || !ok {
		return "", ErrNotFound
	}

	return userID, nil
}

func GetRequestIDFromContext(ctx context.Context) (string, error) {
	requestID, ok := ctx.Value(uthttp.HeaderRequestID).(string)
	if requestID == "" || !ok {
		return "", ErrNotFound
	}

	return requestID, nil
}

func GetLocaleFromContext(ctx context.Context) (string, error) {
	locale, ok := ctx.Value(uthttp.HeaderLocale).(string)
	if locale == "" || !ok {
		return "", ErrNotFound
	}

	return locale, nil
}
