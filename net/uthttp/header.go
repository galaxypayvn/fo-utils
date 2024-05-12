package uthttp

import (
	valueobject "code.finan.cc/finan-one-be/fo-utils/model/value-object"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

const (
	HeaderRequestID       = "x-request-id"
	HeaderUserID          = "x-user-id"
	HeaderBusinessID      = "x-business-id"
	HeaderOrgID           = "x-org-id"
	HeaderClientRequestID = "x-client-request-id"
	HeaderDeviceID        = "x-device-id"
	HeaderLocale          = "x-locale"
)

const defaultLocale = "vi"

func GetAuthInfoFromToken(req *http.Request) (res *valueobject.Auth, err error) {
	res = &valueobject.Auth{}
	res.BusinessID, err = GetBusinessIDFromHeader(req)
	if err != nil {
		return nil, err
	}

	res.UserID, err = GetUserIDFromHeader(req)
	if err != nil {
		return nil, err
	}

	res.RequestID, err = GetRequestIDFromHeader(req)
	if err != nil {
		return nil, err
	}

	res.OrgID, err = GetOrgIDFromHeader(req)
	if err != nil {
		return nil, err
	}

	res.Locale = GetLocaleFromHeader(req)

	return res, nil
}

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
func GetBusinessIDFromHeader(req *http.Request) (uint64, error) {
	sBiz, err := GetStringFromHeader(req, HeaderBusinessID)
	if err != nil {
		return 0, err
	}
	// Convert sBiz from string to uint64 format using Atoi
	biz, err := strconv.ParseUint(sBiz, 10, 64)
	if err != nil {
		return 0, err
	}
	return biz, nil
}

func GetOrgIDFromHeader(req *http.Request) (uint64, error) {
	orgIdString, err := GetStringFromHeader(req, HeaderOrgID)
	if err != nil {
		return 0, err
	}

	orgID, err := strconv.ParseUint(orgIdString, 10, 64)
	if err != nil {
		return 0, err
	}
	return orgID, nil
}

// GetUserIDFromHeader
// retrieves the user ID from the request header
func GetUserIDFromHeader(req *http.Request) (uuid.UUID, error) {
	return GetUUIDFromHeader(req, HeaderUserID)
}

// GetRequestIDFromHeader
// retrieves the request ID from the request header
func GetRequestIDFromHeader(req *http.Request) (string, error) {
	return GetStringFromHeader(req, HeaderRequestID)
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

func ParseIntIDFromUri(c *gin.Context) *uint64 {
	tID := UriParse{}
	if err := c.ShouldBindUri(&tID); err != nil || len(tID.ID) == 0 {
		_ = c.Error(err)
		return nil
	}
	if len(tID.ID) > 0 {
		id, err := strconv.ParseUint(tID.ID[0], 10, 64)
		if err != nil {
			_ = c.Error(err)
			return nil
		}
		return &id
	}
	return nil
}
