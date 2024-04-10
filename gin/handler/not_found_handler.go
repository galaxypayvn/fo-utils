package handler

import (
	messagecode "code.finan.cc/finan-one-be/fo-utils/config/messagecode"
	"github.com/gin-gonic/gin"
	"gitlab.com/goxp/cloud0/ginext"
)

func NoRoute(h *ResponseHandler) gin.HandlerFunc {
	return ginext.WrapHandler(
		func(r *ginext.Request) (*ginext.Response, error) {
			return h.NewResponseWithMessageCode(r.GinCtx, messagecode.GeneralNotFoundCode, nil, nil), nil
		},
	)
}

func NoMethod(h *ResponseHandler) gin.HandlerFunc {
	return ginext.WrapHandler(
		func(r *ginext.Request) (*ginext.Response, error) {
			return h.NewResponseWithMessageCode(r.GinCtx, messagecode.GeneralNotFoundCode, nil, nil), nil
		},
	)
}
