package response

import (
	messagecode "code.finan.cc/finan-one-be/fo-utils/config/messagecode"
	"code.finan.cc/finan-one-be/fo-utils/net/uthttp"
	"code.finan.cc/finan-one-be/fo-utils/valid"

	"github.com/gin-gonic/gin"
	"gitlab.com/goxp/cloud0/ginext"
)

type Response struct {
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

type Handler struct {
	messClient *messagecode.Client
}

func NewHandler(messClient *messagecode.Client) *Handler {
	return &Handler{
		messClient: messClient,
	}
}

func (h *Handler) GeneralSuccessResponse(c *gin.Context, params ...any) *ginext.Response {
	return h.NewResponse(c, messagecode.GeneralSuccessCode, nil, nil, params...)
}

func (h *Handler) GeneralSuccessDataResponse(c *gin.Context, data any, params ...any) *ginext.Response {
	return h.NewResponse(c, messagecode.GeneralSuccessCode, data, nil, params...)
}

func (h *Handler) GeneralSuccessDataMetaResponse(c *gin.Context, data any, meta map[string]any, params ...any) *ginext.Response {
	return h.NewResponse(c, messagecode.GeneralSuccessCode, data, meta, params...)
}

func (h *Handler) GeneralSuccessCreatedDataResponse(c *gin.Context, data any, params ...any) *ginext.Response {
	return h.NewResponse(c, messagecode.GeneralSuccessCreatedCode, data, nil, params...)
}

func (h *Handler) NewDataResponse(c *gin.Context, messagecode int, data any, params ...any) *ginext.Response {
	return h.NewResponse(c, messagecode, data, nil, params...)
}

func (h *Handler) NewResponse(c *gin.Context, messageCode int, data any, meta map[string]any, params ...any) *ginext.Response {
	requestID := c.GetString(ginext.RequestIDName)

	locale := uthttp.GetLocaleFromHeader(c.Request)
	res := &Response{
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
