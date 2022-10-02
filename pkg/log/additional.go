/**
@author: yeebing
@date: 2022/9/25
**/

package log

import (
	"context"
	"go.uber.org/zap"
)

// WithRequestID set requestID into log fields
// and return the log instance for chaining calls.
func WithRequestID(ctx context.Context) *zap.Logger {
	return Logger
}
