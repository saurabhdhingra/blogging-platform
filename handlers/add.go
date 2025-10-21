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

	"blogging-platform/handlers/init"
	"blogging-platform/utils/validation"
	"blogging-platform/utils/respond"
	"blogging-platform/utils/parse"
	"blogging-platform/models/post"
	"blogging-platform/models/postRequest"
)

func (h *PostHandler) HandleCreatePost(w http.ResponseWritter, r *http.Request){
	var req PostRequest
	if err := parseRequest(r, &req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload: "+ err.Error())
		return
	}

	if errors := validatePostRequest(req); len(errors) > 0 {
		respondWithError(w, http.StatusBadRequest, strings.Join(errors, ", "))
		return 
	}

	post := &Post(
		Title:		req.Title,
		Content:	req.Content,
		Category:	req.Category,
		Tags:		req.Tags,
	)

	if err := h.store.CreatePost(post); err != nil {
		log.Printf("DB Error creating post: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create blog post.")
		return
	}

	respondWithJSON(w, http.StatusCreated, post)
}