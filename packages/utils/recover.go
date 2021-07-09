package u

import (
	"context"
	"fmt"
	"vtp-apis/packages/logger/zapl"

	"go.uber.org/zap"
)

var (
	defaultPanicHandler = func(e interface{}) {
		zapl.Error(context.TODO(), "recovered from panic", nil, zap.String("error", fmt.Sprint(e)))
	}
	panicHandler = defaultPanicHandler
)

func ReplacePanicHandler(f func(interface{})) {
	panicHandler = f
}

func WithRecover(f func()) {
	defer func() {
		if panicHandler != nil {
			if e := recover(); e != nil {
				panicHandler(e)
			}
		}
	}()
	f()
}

func WithRecoverError(f func() error) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("recovered from panic: %v", e)
		}
	}()
	return f()
}
