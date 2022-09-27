/**
@author: yeebing
@date: 2022/9/25
**/

package db

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	log := LogOption{
		LogLevel:                  4,
		SlowThreshold:             time.Millisecond * 200, // 200ms
		LogFile:                   "stdout",
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	}

	opts := Option{
		Host:                  "127.0.0.1",
		Port:                  "3306",
		Username:              "root",
		Password:              "123456",
		Database:              "iam_test",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: 1 * time.Minute,
		LogOpt:                log,
	}

	t.Run("normal test", func(t *testing.T) {
		_, err := New(opts)
		if err != nil {
			t.Fatalf("unexpected err or check mysql server's config: %v", err)
		}
	})

	t.Run("test log in logfile", func(t *testing.T) {
		opts.LogOpt.LogFile = "db.log"
		_, err := New(opts)
		if err != nil {
			t.Fatalf("unexpected err or check mysql server's config: %v", err)
		}
	})

}
