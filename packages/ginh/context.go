package ginh

import (
	"context"

	"github.com/gin-gonic/gin"

	"vtp-apis/packages/tracing"
)

func RequestContextFromGinCtx(ginCtx *gin.Context) context.Context {
	ctx := ginCtx.Request.Context()
	ctx = ContextWithRequestMeta(ctx, ginCtx)
	ctx = ContextWithMetaData(ctx)
	return ctx
}

func ContextWithRequestMeta(ctx context.Context, ginCtx *gin.Context) context.Context {
	reqMeta := getRequestMeta(ginCtx)
	return tracing.ContextWithRequestMeta(ctx, &reqMeta)
}

func getRequestMeta(ginCtx *gin.Context) tracing.RequestMeta {
	return tracing.RequestMeta{
		URI:    ginCtx.Request.RequestURI,
		Method: ginCtx.Request.Method,
	}
}
