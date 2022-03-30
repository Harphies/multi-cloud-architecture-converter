package main

import (
	"encoding/json"
	"net/http"
)

// define the envoper
type envelope map[string]interface{}

// write JSON
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Append new line to make it easier to view in terminal applications
	js = append(js, '\n')

	// loop through the header
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Add content type as application/json, then write the status code and the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
