package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

const (
	HeaderRequestID  = "x-request-id"
	HeaderUserID     = "x-user-id"
	HeaderBusinessID = "x-business-id"
)

func GetIDFromHeader(req *http.Request, headerKey string) (res uuid.UUID, err error) {
	headerValue := req.Header.Get(headerKey)
	if headerValue == "" {
		return uuid.Nil, errors.New("empty header value")
	}

	idString := strings.Split(headerValue, "|")[0]
	res, err = uuid.Parse(idString)
	if err != nil {
		return uuid.Nil, err
	}

	return res, nil
}

// GetBusinessIDFromHeader
// retrieves the business ID from the request header
func GetBusinessIDFromHeader(req *http.Request) (uuid.UUID, error) {
	return GetIDFromHeader(req, HeaderBusinessID)
}

// GetUserIDFromHeader
// retrieves the user ID from the request header
func GetUserIDFromHeader(req *http.Request) (uuid.UUID, error) {
	return GetIDFromHeader(req, HeaderUserID)
}

// GetRequestIDFromHeader
// retrieves the request ID from the request header
func GetRequestIDFromHeader(req *http.Request) (uuid.UUID, error) {
	return GetIDFromHeader(req, HeaderRequestID)
}
