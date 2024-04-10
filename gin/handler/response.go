package handler

import (
	messagecode "code.finan.cc/finan-one-be/fo-utils/config/messagecode"
	"code.finan.cc/finan-one-be/fo-utils/net/uthttp"
	"code.finan.cc/finan-one-be/fo-utils/valid"

	"github.com/gin-gonic/gin"
	"gitlab.com/goxp/cloud0/ginext"
)

type response struct {
	Message   Message        `json:"message"`
	Code      int            `json:"code,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
	Data      any            `json:"data,omitempty"`
	Meta      map[string]any `json:"meta,omitempty"`
}

type Message struct {
	Content string `json:"content,omitempty"`
	Params  []any  `json:"params,omitempty"`
}

type ResponseHandler struct {
	messClient *messagecode.Client
}

func NewResponse(messClient *messagecode.Client) *ResponseHandler {
	return &ResponseHandler{
		messClient: messClient,
	}
}

func (h *ResponseHandler) NewResponseWithMessageCode(c *gin.Context, messageCode int, data any, meta map[string]any, params ...any) *ginext.Response {
	requestID := c.GetString(ginext.RequestIDName)

	locale := uthttp.GetLocaleFromHeader(c.Request)
	res := &response{
		Message: Message{
			Content: h.messClient.GetMessage(locale, messageCode),
			Params:  params,
		},
		Code:      messageCode,
		RequestID: requestID,
		Meta:      meta,
	}

	switch {
	case valid.IsSlice(data), data == nil:
		res.Data = data
	default:
		res.Data = []any{data}
	}

	return &ginext.Response{
		Code: h.messClient.GetHTTPCode(locale, messageCode),
		Body: res,
	}
}

func GeneralBadRequestResponse(err error) (*ginext.Response, error) {
	return nil, messagecode.Error{
		Code:  messagecode.GeneralBadRequestCode,
		Cause: err,
	}
}
