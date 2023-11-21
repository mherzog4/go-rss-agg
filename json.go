package main 

import (
	"net/http"
	"encoding/json"
	"log"
	
)

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Printf("Error: %v", msg)
	}

	type errReponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errReponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling payload: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}