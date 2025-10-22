package main

import (
	"log"
	"net/http"
	"os"

	"blogging-platform/handlers"
	"blogging-platform/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// The main function initializes the dependencies and starts the HTTP server.
func main() {
	// 1. Initialize PostgreSQL Store
	store, err := database.NewPostgresStore()
	if err != nil {
		log.Fatalf("Failed to initialize database store: %v", err)
	}

	// 2. Initialize Handlers
	postHandler := handlers.NewPostHandler(store)

	// 3. Setup Chi Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Define API routes for /posts
	r.Route("/posts", func(r chi.Router) {
		// GET /posts and GET /posts?term=
		r.Get("/", postHandler.HandleGetAllPosts)
		
		// POST /posts
		r.Post("/", postHandler.HandleCreatePost)

		r.Route("/{id}", func(r chi.Router) {
			// GET /posts/{id}
			r.Get("/", postHandler.HandleGetPostByID)
			
			// PUT /posts/{id}
			r.Put("/", postHandler.HandleUpdatePost)
			
			// DELETE /posts/{id}
			r.Delete("/", postHandler.HandleDeletePost)
		})
	})

	// 4. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Starting Blog API server on http://localhost:%s", port)
	
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}