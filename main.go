package main

import(
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5" 
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatalf("Failed to initialize a database store: %v", err)
	}

	handler := NewPostHandler(store)
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/posts", func(r chi.Router){
		r.Get("/", handler.HandleGetAllPosts)

		r.Post("/", handler.HndleCreatePost)

		r.Route("/{id}", func(r chi.Router){
			r.Get("/", handler.HandlerGetPostByID)
			
			r.Put("/", handler.HandleUpdatePost)

			r.Delete("/, handler.HandleDeletePost")
		})
	})

	port := ":8080"
	log.Printf("Starting Blog API server on http://localhost%s", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}