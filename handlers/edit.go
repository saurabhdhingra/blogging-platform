package handlers


import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5" // Recommended router for clean REST handling
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5" /

	"blogging-platform/handlers/init"
	"blogging-platform/utils/validation"
	"blogging-platform/utils/respond"
	"blogging-platform/utils/parse"
	"blogging-platform/models/post"
	"blogging-platform/models/postRequest"
)

func (h *PostHandler) HandleUpdatePost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid post ID format.")
		return
	}

	var req PostRequest
	if err := parseRequest(r, &req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	// Validation
	if errors := validatePostRequest(req); len(errors) > 0 {
		respondWithError(w, http.StatusBadRequest, strings.Join(errors, ", "))
		return
	}

	// Database operation
	updatedPost, err := h.store.UpdatePost(id, &req)
	if err != nil {
		log.Printf("DB Error updating post: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to update blog post.")
		return
	}

	// Not Found: 404
	if updatedPost == nil {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("Blog post with ID %d not found.", id))
		return
	}

	// Success: 200 OK
	respondWithJSON(w, http.StatusOK, updatedPost)
}