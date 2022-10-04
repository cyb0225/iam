/**
@author: yeebing
@date: 2022/9/25
**/

package log

import (
	"context"
	"go.uber.org/zap"
)

var (
	KeyRequestID = "requestID"
)

// WithRequestID set requestID into log fields
// and return the log instance for chaining calls.
func WithRequestID(ctx context.Context) *zap.Logger {
	var lg *zap.Logger
	if Logger == nil || ctx == nil {
		lg = zap.NewExample()
	} else {
		lg = clone(Logger)
	}

	if val := ctx.Value(KeyRequestID); val != nil {
		requestID := val.(string)
		lg = lg.With(zap.String(KeyRequestID, requestID))
	}

	return lg
}

func clone(logger *zap.Logger) *zap.Logger {
	copy := *logger

	return &copy
}
