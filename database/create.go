package database

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
	"github.com/jackc/pgx/v5" 

	"blogging-platform/database/init"
)

func (s *PostgresStore) CreatePost(post *Post) error {
	query := `
	INSERT INTO posts (title, content, category, tags)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at, updated_at`

	now := time.Now().UTC()
	post.CreatedAt = now
	post.UpdatedAt = now
	
	return s.db.QueryRow(
		context.Background(),
		query,
		post.Title, 
		post.Content,
		post.Category,
		post.Tags,
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
}
