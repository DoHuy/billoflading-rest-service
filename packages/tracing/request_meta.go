package tracing

import (
	"context"
	"vtp-apis/packages/constants"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"vtp-apis/packages/tracing/trcontext"
)

const (
	keyRequestMeta     = "request_meta"
	keyURI             = "uri"
	keyMethod          = "method"
	keyRequestID       = "request_id"
	keyCommitHashShort = "commit_short"
)

type RequestMeta struct {
	URI       string
	Method    string
	RequestID string
}

func (m *RequestMeta) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	if m == nil {
		return nil
	}
	encoder.AddString(keyURI, m.URI)
	encoder.AddString(keyMethod, m.Method)
	encoder.AddString(keyRequestID, m.RequestID)
	return nil
}

func ContextWithRequestMeta(ctx context.Context, r *RequestMeta) context.Context {
	return trcontext.WithRequestMeta(ctx, r)
}

func RequestMetaFromContext(ctx context.Context) *RequestMeta {
	if ctx == nil {
		return nil
	}
	r, _ := trcontext.RequestMetaFromContext(ctx).(*RequestMeta)
	return r
}

func ExtractRequestMetaZapFields(ctx context.Context) []zap.Field {
	if ctx == nil {
		return nil
	}
	requestMeta := RequestMetaFromContext(ctx)
	if requestMeta == nil {
		return nil
	}
	return []zap.Field{zap.Object(keyRequestMeta, requestMeta)}
}

func ExtractServiceMetaZapFields(ctx context.Context) []zap.Field {
	_ = ctx
	return []zap.Field{zap.String(keyCommitHashShort, constants.CommitHashShort)}
}
