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
	"blogging-platform/utils/respond"
)

func (h *PostHandler) HandleDeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid post ID format.")
		return
	}

	// Database operation
	if err := h.store.DeletePost(id); err != nil {
		if err == pgx.ErrNoRows {
			// Not Found: 404
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("Blog post with ID %d not found.", id))
			return
		}
		log.Printf("DB Error deleting post: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to delete blog post.")
		return
	}

	// Success: 204 No Content
	respondWithJSON(w, http.StatusNoContent, nil)
}