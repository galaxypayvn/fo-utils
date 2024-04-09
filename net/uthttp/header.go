package uthttp

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
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

type UriParse struct {
	ID []string `json:"id" uri:"id"`
}

func ParseIDFromUri(c *gin.Context) *uuid.UUID {
	tID := UriParse{}
	if err := c.ShouldBindUri(&tID); err != nil {
		_ = c.Error(err)
		return nil
	}
	if len(tID.ID) == 0 {
		_ = c.Error(fmt.Errorf("error: Empty when parse ID from URI"))
		return nil
	}
	if id, err := uuid.Parse(tID.ID[0]); err != nil {
		_ = c.Error(err)
		return nil
	} else {
		return &id
	}
}
