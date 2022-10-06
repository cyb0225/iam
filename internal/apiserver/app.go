/**
@author: yeebing
@date: 2022/9/24
**/

package apiserver

import (
	"github.com/cyb0225/iam/pkg/cache"
	"github.com/cyb0225/iam/pkg/db"
	"github.com/cyb0225/iam/pkg/email"
	zaplog "github.com/cyb0225/iam/pkg/log"
	"github.com/cyb0225/iam/pkg/server"
	"github.com/spf13/pflag"
	"log"
)

var (
	port       string
	mode       string
	configPath string
)

// Run start apiserver Server.
func Run() {
	// set flags
	SetFlags()
	pflag.Parse()

	// load config
	opts, err := NewOption(configPath)
	if err != nil {
		log.Fatalf("load config failed: %v\n", err)
	}

	if len(mode) != 0 {
		opts.Server.Mode = mode
	}
	if len(port) != 0 {
		opts.Server.Port = port
	}

	// init main relies.
	InitRelies(opts)

	router := InitRouter(opts.Server.Mode)
	server.Run(opts.Server, router)
}

func SetFlags() {
	pflag.StringVar(&port, "port", "", "http server port")
	pflag.StringVar(&mode, "mode", "", "http server start mode( release or debug)")
	pflag.StringVar(&configPath, "config", "", "the filepath of config file")
}

// InitRelies init global relies
func InitRelies(opts *Option) {
	if _, err := db.New(opts.Mysql); err != nil {
		log.Fatalf("init mysql failed: %v", err)
	}

	if _, err := cache.New(); err != nil {
		log.Fatalf("init gocache failed: %v", err)
	}

	if _, err := email.New(opts.Email); err != nil {
		log.Fatalf("init email failed: %v", err)
	}

	if _, err := zaplog.New(opts.Log); err != nil {
		log.Fatalf("init zap log failed: %v", err)
	}
}
