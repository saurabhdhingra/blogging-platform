package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"blogging-platform/models"
	"blogging-platform/database"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

// PostHandler encapsulates the store dependency.
type PostHandler struct {
	store *database.PostgresStore
}

// NewPostHandler creates a new handler instance.
func NewPostHandler(store *database.PostgresStore) *PostHandler {
	return &PostHandler{store: store}
}

// HandleCreatePost handles POST /posts
func (h *PostHandler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	var req models.PostRequest
	if err := parseRequest(r, &req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	// Validation
	if errors := validatePostRequest(req); len(errors) > 0 {
		respondWithError(w, http.StatusBadRequest, strings.Join(errors, ", "))
		return
	}

	// Create Post object
	post := &models.Post{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Tags:     req.Tags,
	}

	// Database operation
	if err := h.store.CreatePost(post); err != nil {
		log.Printf("DB Error creating post: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create blog post.")
		return
	}

	// Success: 201 Created
	respondWithJSON(w, http.StatusCreated, post)
}

// HandleGetPostByID handles GET /posts/{id}
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

// HandleGetAllPosts handles GET /posts (with optional ?term= search filter)
func (h *PostHandler) HandleGetAllPosts(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("term")

	// Database operation (filtering handled by store)
	posts, err := h.store.GetAllPosts(searchTerm)
	if err != nil {
		log.Printf("DB Error retrieving all posts: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve blog posts.")
		return
	}

	// Success: 200 OK
	respondWithJSON(w, http.StatusOK, posts)
}

// HandleUpdatePost handles PUT /posts/{id}
func (h *PostHandler) HandleUpdatePost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid post ID format.")
		return
	}

	var req models.PostRequest
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

// HandleDeletePost handles DELETE /posts/{id}
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