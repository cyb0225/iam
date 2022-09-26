/**
@author: yeebing
@date: 2022/9/25
**/

package db

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Option struct {
	Host                  string        `yaml:"host"`
	Port                  string        `yaml:"port"`
	Username              string        `yaml:"username"`
	Password              string        `yaml:"password"`
	Database              string        `yaml:"database"`
	MaxIdleConnections    int           `yaml:"maxIdleConnections"`
	MaxOpenConnections    int           `yaml:"maxOpenConnections"`
	MaxConnectionLifeTime time.Duration `yaml:"maxConnectionLifeTime"`
	LogOpt                LogOption     `yaml:"logOpt"`
}

type LogOption struct {
	LogLevel                  int           `yaml:"logLevel"` // Silent 1 、Error 2、Warn 3、Info 4
	LogFile                   string        `yaml:"logFile"`  // if it is stdout/stderr then log to the stdout/stderr.
	SlowThreshold             time.Duration `yaml:"slowThreshold"`
	IgnoreRecordNotFoundError bool          `yaml:"ignoreRecordNotFoundError"` // ignored ErrRecordNotFound
	Colorful                  bool          `yaml:"colorful"`                  // log with color
}

// New Open initialize db session based on opts
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
	switch opts.LogOpt.LogFile {
	case "stdout":
		logWriter = os.Stdout
	case "stderr":
		logWriter = os.Stderr
	default:

	}

	newLogger := logger.New(
		log.New(logWriter, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             opts.LogOpt.SlowThreshold,
			LogLevel:                  logger.LogLevel(opts.LogOpt.LogLevel),
			IgnoreRecordNotFoundError: opts.LogOpt.IgnoreRecordNotFoundError,
			Colorful:                  opts.LogOpt.Colorful,
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
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	DB = db

	return DB, nil
}
