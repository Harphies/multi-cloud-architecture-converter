package main

import (
	"fmt"
	"net/http"
)

func (app *application) heathcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Healthcheck")
}
