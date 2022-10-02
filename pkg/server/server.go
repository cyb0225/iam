/**
@author: yeebing
@date: 2022/10/1
**/

package server

import (
	"github.com/cyb0225/iam/pkg/shutdown"
	"log"
	"net/http"
	"time"
)

type Option struct {
	Mode            string        `yaml:"mode"` // run mode, debug or release
	Port            string        `yaml:"port"`
	ReadTimeout     time.Duration `yaml:"readTimeout"`  // second
	WriteTimeout    time.Duration `yaml:"writeTimeout"` // second
	MaxHeaderBytes  int           `yaml:"maxHeaderBytes"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout"` // second
}

// Run start http server.
func Run(opts Option, handler http.Handler) {

	s := &http.Server{
		Addr:           ":" + opts.Port,
		Handler:        handler,
		ReadTimeout:    opts.ReadTimeout * time.Second,
		WriteTimeout:   opts.WriteTimeout * time.Second,
		MaxHeaderBytes: opts.MaxHeaderBytes,
	}

	// start http server
	go func() {
		log.Printf("Start http server %s...\n", opts.Port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe err: %v", err)
		}
	}()

	shutdown.WithTimeout(s, opts.ShutdownTimeout*time.Second)
}
