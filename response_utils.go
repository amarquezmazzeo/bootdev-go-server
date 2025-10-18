package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with %d error: %s", code, message)
	}

	type resposeError struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, resposeError{Error: message})
}

func respondWithJson(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		log.Printf("Error marshalling response: %s", err)
		w.WriteHeader(500)
		w.Write([]byte(`{"error": "Something went wrong"}`))
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}
