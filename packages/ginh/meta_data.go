package ginh

import (
	"context"
)

type MetaData struct {
	Keys map[string]interface{}
}

type metaDataKey struct{}

func ContextWithMetaData(ctx context.Context) context.Context {
	keys := make(map[string]interface{})
	metaData := MetaData{Keys: keys}
	return context.WithValue(ctx, metaDataKey{}, &metaData)
}
