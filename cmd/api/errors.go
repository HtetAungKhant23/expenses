package main

import "net/http"

func (app *application) log(r *http.Request, err error) {
	app.logger.Println(err, "at", r.URL.String())
}

func (app *application) error(w http.ResponseWriter, r *http.Request, status int, msg any) {
	data := envelope{
		"error": msg,
	}

	err := app.writeJson(w, status, data, nil)

	if err != nil {
		app.log(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.log(r, err)
	app.error(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.error(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
}

func (app *application) failedValidation(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.error(w, r, http.StatusUnprocessableEntity, errors)
}
