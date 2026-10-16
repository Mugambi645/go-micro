package main

import (
	"net/http"
	"encoding/json"
)



func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	// Implementation for broker handler
	payload := jsonResponse{
		Error: false,
		Message: "Broker hit",

	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}