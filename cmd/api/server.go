package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) serve() error {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  1 * time.Minute,
		Handler:      nil,
	}

	app.logger.Printf("Server is starting at %d", app.config.port)

	err := srv.ListenAndServe()

	if err != nil {
		return err
	}
	return nil
}
