package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": true,
		"time":   time.Now(),
	}

	js, err := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(js)

	if err != nil {
		data = envelope{
			"error": err,
		}
		js, err = json.Marshal(data)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(js)
	}
}
