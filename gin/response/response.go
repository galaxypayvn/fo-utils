package response

import (
	"errors"
	"reflect"

	messagecode "code.finan.cc/finan-one-be/fo-utils/config/messagecode"
	"code.finan.cc/finan-one-be/fo-utils/net/uthttp"
	"code.finan.cc/finan-one-be/fo-utils/valid"

	"github.com/gin-gonic/gin"
	"gitlab.com/goxp/cloud0/ginext"
)

type Response[T any] struct {
	Message   Message        `json:"message"`
	Code      int            `json:"code,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
	Data      []T            `json:"data,omitempty"`
	Meta      map[string]any `json:"meta,omitempty"`
}

type Message struct {
	Content string `json:"content,omitempty"`
	Params  []any  `json:"params,omitempty"`
	Error   string `json:"error,omitempty"`
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
	locale := uthttp.GetLocaleFromHeader(c.Request)
	messageContent := h.messClient.GetMessage(locale, messageCode)
	return h.newRawResponse(c, messageCode, messageContent, data, meta)
}

func (h *Handler) NewRawResponse(c *gin.Context, messageCode int, messageContent string, data any, meta map[string]any, params ...any) *ginext.Response {
	return h.newRawResponse(c, messageCode, messageContent, data, meta, params...)
}

func (h *Handler) newRawResponse(c *gin.Context, messageCode int, messageContent string, data any, meta map[string]any, params ...any) *ginext.Response {
	requestID := c.GetString(ginext.RequestIDName)

	locale := uthttp.GetLocaleFromHeader(c.Request)
	res := &Response[any]{
		Message: Message{
			Content: messageContent,
			Params:  params,
		},
		Code:      messageCode,
		RequestID: requestID,
		Meta:      meta,
	}

	if len(c.Errors) > 0 {
		res.Message.Error = c.Errors[0].Error()
	}

	switch {
	case valid.IsSlice(data):
		v := reflect.ValueOf(data)
		var out []any
		for i := 0; i < v.Len(); i++ {
			out = append(out, v.Index(i).Interface())
		}
		res.Data = out
	case data == nil:
		res.Data = nil
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

func GeneralUnauthorizedResponse(err error) (*ginext.Response, error) {
	return nil, messagecode.Error{
		Code:  messagecode.GeneralUnauthorizedCode,
		Cause: err,
	}
}

func TranslateToServiceError[T any](resp Response[T]) error {
	if resp.Code == 0 {
		return messagecode.NewUnknownFormatError(errors.New("missing message code"))
	}
	return messagecode.NewServiceError(resp.Code, resp.Message.Content, resp.Data, resp.Meta, resp.Message.Params...)
}
