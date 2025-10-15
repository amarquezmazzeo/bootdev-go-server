package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}

func handlerValidate(w http.ResponseWriter, r *http.Request) {

	type requestBody struct {
		Body string `json:"body"`
	}

	type resposeError struct {
		Error string `json:"error"`
	}

	type resposeSuccess struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	reqBody := requestBody{}
	err := decoder.Decode(&reqBody)
	if err != nil {
		log.Printf("Error decoding request: %s", err)
		w.WriteHeader(500)
		w.Write([]byte(`{"error": "Something went wrong"}`))
		return
	}
	if len(reqBody.Body) > 140 {
		respBody := resposeError{Error: "Chirp is too long"}
		data, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling response: %s", err)
			w.WriteHeader(500)
			w.Write([]byte(`{"error": "Something went wrong"}`))
			return
		}
		log.Printf("Chirp body is over 140 characters: %s", reqBody.Body)
		w.WriteHeader(400)
		w.Write(data)
		return
	}

	respBody := resposeSuccess{Valid: true}
	data, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling response: %s", err)
		w.WriteHeader(500)
		w.Write([]byte(`{"error": "Something went wrong"}`))
		return
	}
	log.Printf("Chirp body is under 140 characters: %s", reqBody.Body)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
