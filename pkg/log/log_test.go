/**
@author: yeebing
@date: 2022/9/25
**/

package log

import (
	"testing"

	"go.uber.org/zap"
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
