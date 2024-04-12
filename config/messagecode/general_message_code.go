package messagecode

import (
	"fmt"
)

const (
	GeneralSuccessCreatedCode   = 102001
	GeneralSuccessCode          = 102000
	GeneralInternalErrorCode    = 105000
	GeneralBadRequestCode       = 104000
	GeneralNotFoundCode         = 104004
	GeneralDuplicateRequestCode = 104100
	GeneralUnknownFormatCode    = 105100
)

type Error struct {
	Code   int
	Params []any
	Cause  error
}

type ServiceError struct {
	Code    int
	Message message
	Data    any
	Meta    map[string]any
}

type message struct {
	Content string `json:"content,omitempty"`
	Params  []any  `json:"params,omitempty"`
}

func NewError(code int, cause error, params ...any) error {
	return Error{
		Code:   code,
		Params: params,
		Cause:  cause,
	}
}

func NewUnknownFormatError(err error) error {
	return Error{
		Code:  GeneralUnknownFormatCode,
		Cause: err,
	}
}

func NewServiceError(code int, messageContent string, data any, meta map[string]any, params ...any) error {
	return ServiceError{
		Code: code,
		Message: message{
			Content: messageContent,
			Params:  params,
		},
		Data: data,
		Meta: meta,
	}
}

func (err Error) Error() string {
	errStr := fmt.Sprintf("code=%d params=%v", err.Code, err.Params)
	return fmt.Sprintf("%s: %s", errStr, err.Cause)
}
func (err Error) Unwrap() error {
	return err.Cause
}

func (err ServiceError) Error() string {
	errStr := fmt.Sprintf("code=%d messsage=%s params=%v data=%+v metadata=%+v", err.Code, err.Message.Content, err.Message.Params, err.Data, err.Meta)
	return errStr
}
