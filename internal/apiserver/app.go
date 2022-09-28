/**
@author: yeebing
@date: 2022/9/24
**/

package apiserver

import (
	"github.com/cyb0225/iam/pkg/app"
	"log"
	"net/http"
	"time"

	"github.com/cyb0225/iam/pkg/shutdown"
)

// NewApp create a command project frame.
func NewApp(name, basename string) *app.App {
	app := app.NewApp(
		name,
		basename,
	)
	return app
}

// run start apiserver Server.
func run() {
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
