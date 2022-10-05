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
	Mode string `yaml:"mode"` // run mode, debug or release
	Port string `yaml:"port"`
}

// Run start http server.
func Run(opts Option, handler http.Handler) {

	s := &http.Server{
		Addr:           ":" + opts.Port,
		Handler:        handler,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1500,
	}

	// start http server
	go func() {
		log.Printf("Start http server %s...\n", opts.Port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//log.Fatalf("ListenAndServe err: %v", err)
		}
	}()

	shutdown.WithTimeout(s, 5*time.Second)
}
