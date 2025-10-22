package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"blogging-platform/models"

	"github.com/jackc/pgx/v5"
)

// PostgresStore manages the database connection and provides CRUD methods.
type PostgresStore struct {
	db *pgx.Conn
}

// NewPostgresStore establishes the connection and initializes the table.
func NewPostgresStore() (*PostgresStore, error) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// Default local connection string for development
		connStr = "postgresql://your_username:your_password@localhost:5432/postgres"
		log.Printf("WARNING: Using default DB connection string. Set DATABASE_URL env var.")
	}

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	store := &PostgresStore{db: conn}
	if err := store.InitTable(); err != nil {
		return nil, fmt.Errorf("failed to initialize posts table: %w", err)
	}

	return store, nil
}

// InitTable creates the 'posts' table if it does not exist.
func (s *PostgresStore) InitTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		category VARCHAR(100) NOT NULL,
		tags TEXT[] NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := s.db.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error creating posts table: %w", err)
	}
	log.Println("PostgreSQL table 'posts' initialized successfully.")
	return nil
}

// CreatePost inserts a new post into the database.
func (s *PostgresStore) CreatePost(post *models.Post) error {
	query := `
	INSERT INTO posts (title, content, category, tags)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at`

	// Use time.Now() for consistency, though DB default is used.
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

// GetPostByID retrieves a single post by ID.
func (s *PostgresStore) GetPostByID(id int) (*models.Post, error) {
	query := `
	SELECT id, title, content, category, tags, created_at, updated_at
	FROM posts
	WHERE id = $1`

	post := new(models.Post)
	err := s.db.QueryRow(context.Background(), query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Category,
		&post.Tags,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // Indicate not found
		}
		return nil, fmt.Errorf("error querying post by ID: %w", err)
	}
	return post, nil
}

// GetAllPosts retrieves all posts, optionally filtered by a search term.
func (s *PostgresStore) GetAllPosts(searchTerm string) ([]*models.Post, error) {
	var (
		posts   []*models.Post
		query   string
		args    []interface{}
	)

	query = `
	SELECT id, title, content, category, tags, created_at, updated_at
	FROM posts`
	
	if searchTerm != "" {
		// Wildcard search (case-insensitive ILIKE) on title, content, and category.
		query += ` WHERE title ILIKE $1 OR content ILIKE $1 OR category ILIKE $1`
		args = append(args, "%"+searchTerm+"%")
	}

	query += ` ORDER BY created_at DESC`

	rows, err := s.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying all posts: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		post := new(models.Post)
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Category,
			&post.Tags,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning post row: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// UpdatePost updates an existing post by ID.
func (s *PostgresStore) UpdatePost(id int, req *models.PostRequest) (*models.Post, error) {
	query := `
	UPDATE posts
	SET title = $1, content = $2, category = $3, tags = $4, updated_at = NOW()
	WHERE id = $5
	RETURNING id, title, content, category, tags, created_at, updated_at`

	updatedPost := new(models.Post)
	err := s.db.QueryRow(
		context.Background(),
		query,
		req.Title,
		req.Content,
		req.Category,
		req.Tags,
		id,
	).Scan(
		&updatedPost.ID,
		&updatedPost.Title,
		&updatedPost.Content,
		&updatedPost.Category,
		&updatedPost.Tags,
		&updatedPost.CreatedAt,
		&updatedPost.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // Not found
		}
		return nil, fmt.Errorf("error updating post: %w", err)
	}
	return updatedPost, nil
}

// DeletePost deletes a post by ID.
func (s *PostgresStore) DeletePost(id int) error {
	query := `DELETE FROM posts WHERE id = $1`
	cmdTag, err := s.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("error executing delete query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return pgx.ErrNoRows // Indicate that no rows were deleted (404)
	}
	return nil
}