/**
@author: yeebing
@date: 2022/9/24
**/

package apiserver

import (
	"log"
	"net/http"
	"time"

	"github.com/cyb0225/iam/pkg/shutdown"
)

// Run start apiserver Server.
func Run() {
	// created command.

	// load configs.

	// init router.
	router := InitRouter()
	s := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	// start server server
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	shutdown.WithTimeout(s, time.Second*5)
}
