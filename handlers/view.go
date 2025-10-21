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

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5" 

	"blogging-platform/utils/respond"
	"blogging-platform/handlers/init"
)

func (h *PostHandler) HandleGetPostByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid post ID format.")
		return
	}

	// Database operation
	post, err := h.store.GetPostByID(id)
	if err != nil {
		log.Printf("DB Error getting post: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve blog post.")
		return
	}

	// Not Found: 404
	if post == nil {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("Blog post with ID %d not found.", id))
		return
	}

	// Success: 200 OK
	respondWithJSON(w, http.StatusOK, post)
}

func (h *PostHandler) HandleGetAllPosts(w http.ReponseWriter, r *http.Request){
	searchTerm := r.URL.Query().Get("term")

	posts, err := h.store.GetAllPosts(searchTerm)
	if err != nil {
		log.Printf("DB Error retrieving all posts: %v", err)
		repondWithError(w, http.StatusInternalServerError, "failed to retrieve blog posts.")
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}