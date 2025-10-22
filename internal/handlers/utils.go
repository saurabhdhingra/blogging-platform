package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"blogging-platform/models"
)

// respondWithJSON writes a JSON successful response.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
		}
	}
}

// respondWithError writes a JSON error response.
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.ErrorResponse{Error: message})
}

// parseRequest reads and decodes the JSON request body into the target struct.
func parseRequest(r *http.Request, target interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// validatePostRequest performs basic validation on the required fields.
func validatePostRequest(req models.PostRequest) []string {
	var errors []string
	if strings.TrimSpace(req.Title) == "" {
		errors = append(errors, "Title is required.")
	}
	if strings.TrimSpace(req.Content) == "" {
		errors = append(errors, "Content is required.")
	}
	if strings.TrimSpace(req.Category) == "" {
		errors = append(errors, "Category is required.")
	}
	return errors
}