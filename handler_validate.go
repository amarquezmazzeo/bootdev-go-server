package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidate(w http.ResponseWriter, r *http.Request) {

	const maxChirpLength = 140

	type requestParameters struct {
		Body string `json:"body"`
	}

	type resposeSuccess struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := requestParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding request", err)
		return
	}
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}
	respondWithJson(w, http.StatusOK, resposeSuccess{Valid: true})

}
