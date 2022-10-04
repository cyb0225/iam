/**
@author: yeebing
@date: 2022/9/25
**/

package log

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

// TestLog test log into files and log into console.
func TestLog(t *testing.T) {
	opts := Option{
		Level:     "debug",
		AccessLog: "../../tmp/log/access.log",
		ErrorLog:  "../../tmp/log/error.log",
		Console:   true,
	}

	t.Run("test unknown level", func(t *testing.T) {
		errOpts := opts
		errOpts.Level = "unknown"
		_, err := New(errOpts)
		if err == nil {
			t.Fatalf("got nil want err, given %v", errOpts.Level)
		}
	})

	logger, err := New(opts)
	if err != nil {
		t.Fatalf("got an unexpected error: %v", err)
	}

	t.Run("test log", func(t *testing.T) {
		logger.Debug("msg", zap.String("key", "value"))
		logger.Info("msg", zap.String("key", "value"))
		logger.Warn("msg", zap.String("key", "value"))
		logger.Error("msg", zap.String("key", "value"))
	})

	t.Run("test info log with json writer", func(t *testing.T) {
		infoOpts := opts
		infoOpts.Level = "info"
		logger, err := New(infoOpts)
		if err != nil {
			t.Fatalf("got an unexpected error: %v", err)
		}
		logger.Debug("info level", zap.String("key", "value"))
		logger.Info("info level", zap.String("key", "value"))
	})
}

func TestWithRequestID(t *testing.T) {
	opts := Option{
		Level:     "debug",
		AccessLog: "../../tmp/log/access.log",
		ErrorLog:  "../../tmp/log/error.log",
		Console:   true,
	}

	_, err := New(opts)
	assert.Equal(t, nil, err)

	t.Run("not set request ID", func(t *testing.T) {
		WithRequestID(context.TODO()).Info("test", zap.String("str", "Str"))
	})

	t.Run("success", func(t *testing.T) {
		ctx := context.WithValue(context.TODO(), KeyRequestID, "1")
		WithRequestID(ctx).Info("test", zap.String("str", "Str"))
	})

}
