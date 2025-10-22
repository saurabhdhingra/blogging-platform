package models

import "time"

// Post represents the full structure of a blog post stored in the database.
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"` // Maps to TEXT[] in PostgreSQL
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// PostRequest represents the data required to create or update a post (input payload).
type PostRequest struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

// ErrorResponse defines the structure for API error messages.
type ErrorResponse struct {
	Error string `json:"error"`
}
