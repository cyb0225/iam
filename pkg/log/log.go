/**
@author: yeebing
@date: 2022/9/24
**/

package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

type Option struct {
	Level      string `yaml:"level"`
	MaxSize    int    `yaml:"maxSize"`    // MB
	MaxAge     int    `yaml:"maxAge"`     // day
	MaxBackups int    `yaml:"maxBackups"` // the number of files
	Compress   bool   `yaml:"compress"`   // if compress the log
	AccessLog  string `yaml:"accessLog"`
	ErrorLog   string `yaml:"errorLog"`
	Console    bool   `yaml:"console"` // console or json
}

// New create logger instance with options
// it can log into.
func New(opts Option) (*zap.Logger, error) {
	infoWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   opts.AccessLog,
		MaxSize:    opts.MaxSize,
		MaxBackups: opts.MaxBackups,
		MaxAge:     opts.MaxAge,
		Compress:   opts.Compress,
	})
	errorWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   opts.ErrorLog,
		MaxSize:    opts.MaxSize,
		MaxBackups: opts.MaxBackups,
		MaxAge:     opts.MaxAge,
		Compress:   opts.Compress,
	})

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		return nil, err
	}

	encoder := getEncoder(opts.Console)
	var lv zap.LevelEnablerFunc = func(l zapcore.Level) bool {
		return l <= zapcore.InfoLevel && l >= zapLevel
	}

	infoCore := zapcore.NewCore(encoder, infoWriteSyncer, lv)
	errorCore := zapcore.NewCore(encoder, errorWriteSyncer, zapcore.WarnLevel)
	core := zapcore.NewTee(infoCore, errorCore)

	// log into stdout
	if opts.Level == "debug" {
		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		stdCore := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), os.Stdout, zapcore.DebugLevel)
		core = zapcore.NewTee(core, stdCore)
	}

	Logger = zap.New(core, zap.AddCaller())
	Sugar = Logger.Sugar()

	return Logger, nil
}

func getEncoder(console bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if console {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}
