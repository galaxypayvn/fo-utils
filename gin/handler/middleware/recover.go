package middleware

import (
	messagecode "code.finan.cc/finan-one-be/fo-utils/config/messagecode"
	"code.finan.cc/finan-one-be/fo-utils/gin/handler"

	"github.com/gin-gonic/gin"
	"gitlab.com/goxp/cloud0/ginext"
)

func Recover(h *handler.ResponseHandler) gin.HandlerFunc {
	return ginext.WrapHandler(
		func(r *ginext.Request) (res *ginext.Response, err error) {
			defer func() {
				err := recover()
				if err != nil {
					r.GinCtx.Header("Connection", "close")
					res = h.NewResponseWithMessageCode(r.GinCtx, messagecode.GeneralInternalErrorCode, nil, nil)
				}
			}()
			c := r.GinCtx
			c.Next()

			return nil, nil
		},
	)
}
