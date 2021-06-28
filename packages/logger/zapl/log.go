package zapl

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"vtp-apis/packages/gconfig"
)

var (
	zapStackTraceLogger *zap.Logger
)

func init() {
	zapCfg := zap.NewProductionConfig()
	zapCfg.DisableStacktrace = true
	// zapCfg.DisableCaller = true
	if gconfig.GlobalConfig().DisableSampleZapLog {
		zapCfg.Sampling = nil
	}
	l, err := zapCfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(l)
	zapStackTraceLogger = zap.L().WithOptions(zap.AddStacktrace(zap.DebugLevel))
}

func ReplaceNop() {
	zap.ReplaceGlobals(zap.NewNop())
	zapStackTraceLogger = zap.NewNop()
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = BuildFieldsWithContext(ctx, fields)
	zap.L().Info(msg, fields...)
}

func Error(ctx context.Context, msg string, err error, fields ...zap.Field) {
	fields = BuildFieldsWithContext(ctx, fields)
	fields = BuildFieldsWithError(fields, err)
	var d debugger
	if !errors.As(err, &d) {
		zapStackTraceLogger.Error(msg, fields...)
		return
	}
	zap.L().Error(msg, fields...)
}

func Fatal(ctx context.Context, msg string, err error, fields ...zap.Field) {
	fields = BuildFieldsWithContext(ctx, fields)
	fields = BuildFieldsWithError(fields, err)
	var d debugger
	if !errors.As(err, &d) {
		zapStackTraceLogger.Fatal(msg, fields...)
		return
	}
	zap.L().Fatal(msg, fields...)
}

func Panic(ctx context.Context, msg string, err error, fields ...zap.Field) {
	fields = BuildFieldsWithContext(ctx, fields)
	fields = BuildFieldsWithError(fields, err)
	var d debugger
	if !errors.As(err, &d) {
		zapStackTraceLogger.Panic(msg, fields...)
		return
	}
	zap.L().Panic(msg, fields...)
}

func BuildFieldsWithContext(ctx context.Context, fields []zap.Field) []zap.Field {
	if ctx == nil {
		return fields
	}
	for _, e := range fieldExtractors {
		fields = append(e.ExtractZapFields(ctx), fields...)
	}
	return fields
}

func BuildFieldsWithError(fields []zap.Field, err error) []zap.Field {
	if err == nil {
		return fields
	}
	var d debugger
	if errors.As(err, &d) {
		fields = append(fields, zap.String("root_caller", d.Caller()), zap.String("root_stack", d.StackTrace()))
	}
	fields = append(fields, zap.Error(err))
	return fields
}
