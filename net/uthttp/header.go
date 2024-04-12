package uthttp

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

const (
	HeaderRequestID       = "x-request-id"
	HeaderUserID          = "x-user-id"
	HeaderBusinessID      = "x-business-id"
	HeaderClientRequestID = "x-client-request-id"
	HeaderDeviceID        = "x-device-id"
	HeaderLocale          = "x-locale"
)

const defaultLocale = "en"

func GetStringFromHeader(req *http.Request, headerKey string) (string, error) {
	headerValue := req.Header.Get(headerKey)
	if headerValue == "" {
		return "", fmt.Errorf("empty header value %s", headerKey)
	}

	return headerValue, nil
}

func GetUUIDFromHeader(req *http.Request, headerKey string) (uuid.UUID, error) {
	headerValue := req.Header.Get(headerKey)
	if headerValue == "" {
		return uuid.Nil, fmt.Errorf("empty header value %s", headerKey)
	}

	idString := strings.Split(headerValue, "|")[0]
	res, err := uuid.Parse(idString)
	if err != nil {
		return uuid.Nil, err
	}

	return res, nil
}

// GetBusinessIDFromHeader
// retrieves the business ID from the request header
func GetBusinessIDFromHeader(req *http.Request) (uuid.UUID, error) {
	return GetUUIDFromHeader(req, HeaderBusinessID)
}

// GetUserIDFromHeader
// retrieves the user ID from the request header
func GetUserIDFromHeader(req *http.Request) (uuid.UUID, error) {
	return GetUUIDFromHeader(req, HeaderUserID)
}

// GetRequestIDFromHeader
// retrieves the request ID from the request header
func GetRequestIDFromHeader(req *http.Request) (uuid.UUID, error) {
	return GetUUIDFromHeader(req, HeaderRequestID)
}

func GetClientRequestIDFromHeader(req *http.Request) (string, error) {
	return GetStringFromHeader(req, HeaderClientRequestID)
}

func GetDeviceIDFromHeader(req *http.Request) (string, error) {
	return GetStringFromHeader(req, HeaderDeviceID)
}

func GetLocaleFromHeader(r *http.Request) string {
	locale, err := GetStringFromHeader(r, HeaderLocale)
	if err != nil {
		locale = defaultLocale
	}

	return locale
}

type UriParse struct {
	ID []string `json:"id" uri:"id"`
}

func ParseIDFromUri(c *gin.Context) *uuid.UUID {
	tID := UriParse{}
	if err := c.ShouldBindUri(&tID); err != nil || len(tID.ID) == 0 {
		_ = c.Error(err)
		return nil
	}

	id, err := uuid.Parse(tID.ID[0])
	if err != nil {
		_ = c.Error(err)
		return nil
	}
	return &id
}

func ParseStringIDFromUri(c *gin.Context) *string {
	tID := UriParse{}
	if err := c.ShouldBindUri(&tID); err != nil || len(tID.ID) == 0 {
		_ = c.Error(err)
		return nil
	}
	return &tID.ID[0]
}
