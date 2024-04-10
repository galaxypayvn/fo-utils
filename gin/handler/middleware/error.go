package middleware

import (
	"errors"

	messagecode "code.finan.cc/finan-one-be/fo-utils/config/messagecode"
	"code.finan.cc/finan-one-be/fo-utils/gin/handler"

	"github.com/gin-gonic/gin"
	"gitlab.com/goxp/cloud0/ginext"
)

func HandleError(handler *handler.ResponseHandler) gin.HandlerFunc {
	return ginext.WrapHandler(
		func(r *ginext.Request) (*ginext.Response, error) {
			c := r.GinCtx
			c.Next()
			if len(c.Errors) > 0 {
				err := c.Errors[0]
				c.Errors = nil
				var serviceErr messagecode.Error
				if errors.As(err, &serviceErr) {
					return handler.NewResponseWithMessageCode(r.GinCtx, serviceErr.Code, nil, nil), nil
				} else {
					return handler.NewResponseWithMessageCode(r.GinCtx, messagecode.GeneralInternalErrorCode, nil, nil), nil
				}
			}

			return nil, nil
		},
	)
}
