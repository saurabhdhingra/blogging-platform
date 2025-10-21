package utils

import (
	"net/http"
	"encoding/json"

	"blogging-platform/models/errorResponse"
)

func respondWithJSON(w http.ReponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
		}
	}
}


func respondWithError(w. http.ResponseWritter, code int, message string){
	respondWithJSON(w, code, ErrorResponse(Error: message))
}