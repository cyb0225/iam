/**
@author: yeebing
@date: 2022/9/25
**/

package db

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Option struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	LogFile  string `yaml:"logFile"`
}

var (
	// connect pool
	maxIdleConnections    = 100
	maxOpenConnections    = 100
	maxConnectionLifeTime = time.Hour

	// log
	slowThreshold = 200 * time.Millisecond
	logLevel      = logger.LogLevel(logger.Error)
)

// New Open initialize db session based on opts.
func New(opts Option) (*gorm.DB, error) {
	// if the version of mysql is lower than 8.0, then change to charset from `utf8mb4` to `utf8`
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.Database,
		true,
		"Local")

	var logWriter io.Writer
	switch opts.LogFile {
	case "stdout":
		logWriter = os.Stdout
	case "stderr":
		logWriter = os.Stderr
	default:
		// stored in that file.
		logWriter = &lumberjack.Logger{
			Filename:   opts.LogFile,
			MaxSize:    100,
			MaxAge:     7,
			MaxBackups: 10,
			Compress:   false,
		}
	}

	newLogger := logger.New(
		log.New(logWriter, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             slowThreshold,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// set connect pool

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(maxConnectionLifeTime)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdleConnections)

	DB = db

	return DB, nil
}
