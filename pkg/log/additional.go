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
	KeyRequestID = "X-Request-ID"
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

	if requestID := ctx.Value(KeyRequestID); requestID != nil {
		lg = lg.With(zap.Any(KeyRequestID, requestID))
	}

	return lg
}

func clone(logger *zap.Logger) *zap.Logger {
	cl := *logger
	return &cl
}
