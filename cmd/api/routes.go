package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// return httpHandler
func (app *application) routes() *httprouter.Router {
	// initiate a new httpRouter route instance

	router := httprouter.New()

	// define the routes and each route middlewares
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.heathcheckHandler)

	return router
}
