package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	// log the error to the standard out
	app.logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	// contextualize the response data
	env := envelope{"error": message}

	// encode error to JSON response
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		// log the error to the console
		app.logError(r, err)
		// write response to internal seerver
		w.WriteHeader(500)
	}
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit excedded"
	app.errorResponse(w, r, http.StatusTooManyRequests, message)
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// not found response
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource cound bot be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
