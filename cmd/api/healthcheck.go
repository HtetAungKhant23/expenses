package main

import (
	"net/http"
	"time"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": true,
		"time":   time.Now(),
	}

	err := app.writeJson(w, http.StatusOK, data, nil)

	if err != nil {
		app.serverError(w, r, err)
	}
}
