package trcontext

import "context"

type userMetaKey struct{}
type requestMetaKey struct{}
type appMetaKey struct{}
type kafkaMetaKey struct{}

func WithUserMeta(ctx context.Context, userMeta interface{}) context.Context {
	return context.WithValue(ctx, userMetaKey{}, userMeta)
}

func UserMetaFromContext(ctx context.Context) interface{} {
	if ctx == nil {
		return nil
	}
	return ctx.Value(userMetaKey{})
}

func WithRequestMeta(ctx context.Context, requestMeta interface{}) context.Context {
	return context.WithValue(ctx, requestMetaKey{}, requestMeta)
}

func RequestMetaFromContext(ctx context.Context) interface{} {
	return ctx.Value(requestMetaKey{})
}

func WithAppMeta(ctx context.Context, appMeta interface{}) context.Context {
	return context.WithValue(ctx, appMetaKey{}, appMeta)
}

func AppMetaFromContext(ctx context.Context) interface{} {
	return ctx.Value(appMetaKey{})
}

func WithKafkaMeta(ctx context.Context, kafkaMeta interface{}) context.Context {
	return context.WithValue(ctx, kafkaMetaKey{}, kafkaMeta)
}

func KafkaMetaFromContext(ctx context.Context) interface{} {
	return ctx.Value(kafkaMetaKey{})
}
