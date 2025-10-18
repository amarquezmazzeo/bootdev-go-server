package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerValidate(w http.ResponseWriter, r *http.Request) {

	const maxChirpLength = 140

	type requestParameters struct {
		Body string `json:"body"`
	}

	type resposeSuccess struct {
		CleanedBody string `json:"cleaned_body"`
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
	respondWithJson(w, http.StatusOK, resposeSuccess{
		CleanedBody: cleanString(params.Body),
	})

}

func cleanString(body string) string {
	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}

	words := strings.Split(body, " ")
	for i := range words {
		if _, exists := badWords[strings.ToLower(words[i])]; exists {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}
